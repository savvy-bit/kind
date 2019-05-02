/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package logs

import (
	"archive/tar"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	osexec "os/exec"
	"path/filepath"
	"sigs.k8s.io/kind/pkg/cluster/nodes"
	"sigs.k8s.io/kind/pkg/exec"
	"sigs.k8s.io/kind/pkg/util"
	"sync"
)

type errFn func() error

// Collect collects logs related to / from the cluster nodes and the host
// system to the specified directory
func Collect(nodes []nodes.Node, dir string) error {
	prefixedPath := func(path string) string {
		return filepath.Join(dir, path)
	}
	// helper to run a cmd and write the output to path
	execToPath := func(cmd exec.Cmd, path string) error {
		realPath := prefixedPath(path)
		os.MkdirAll(filepath.Dir(realPath), os.ModePerm)
		f, err := os.Create(realPath)
		if err != nil {
			return err
		}
		defer f.Close()
		cmd.SetStdout(f)
		cmd.SetStderr(f)
		return cmd.Run()
	}
	execToPathFn := func(cmd exec.Cmd, path string) func() error {
		return func() error {
			return execToPath(cmd, path)
		}
	}
	// construct a slice of methods to collect logs
	fns := []errFn{
		// TODO(bentheelder): record the kind version here as well
		// record info about the host docker
		execToPathFn(
			exec.Command("docker", "info"),
			"docker-info.txt",
		),
	}
	// add a log collection method for each node
	for _, n := range nodes {
		node := n // https://golang.org/doc/faq#closures_and_goroutines
		name := node.String()
		// grab all logs under /var/log (pods and containers)
		cmd := osexec.Command("docker", "exec", name, "tar", "--hard-dereference", "-C", "/var/log", "-chf", "-", ".")
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}
		if err := cmd.Start(); err != nil {
			return err
		}
		log.Debugf("Running: %v %v", cmd.Path, cmd.Args)
		if err := untar(stdout, filepath.Join(dir, name)); err != nil {
			return errors.Wrapf(err, "Untarring %q: %v", name, err)
		}
		if err := cmd.Wait(); err != nil {
			return err
		}
		fns = append(fns, func() error {
			return coalesce(
				// record info about the node container
				execToPathFn(
					exec.Command("docker", "inspect", name),
					filepath.Join(name, "inspect.json"),
				),
				// grab all of the node logs
				execToPathFn(
					node.Command("cat", "/kind/version"),
					filepath.Join(name, "kubernetes-version.txt"),
				),
				execToPathFn(
					node.Command("journalctl", "--no-pager"),
					filepath.Join(name, "journal.log"),
				),
				execToPathFn(
					node.Command("journalctl", "--no-pager", "-u", "kubelet.service"),
					filepath.Join(name, "kubelet.log"),
				),
				execToPathFn(
					node.Command("journalctl", "--no-pager", "-u", "docker.service"),
					filepath.Join(name, "docker.log"),
				),
			)
		})
	}
	// run and collect up all errors
	return coalesce(fns...)
}

// colaese runs fns concurrently, returning an Errors if there are > 1 errors
func coalesce(fns ...errFn) error {
	// run all fns concurrently
	ch := make(chan error, len(fns))
	var wg sync.WaitGroup
	for _, fn := range fns {
		wg.Add(1)
		go func(f errFn) {
			defer wg.Done()
			ch <- f()
		}(fn)
	}
	wg.Wait()
	close(ch)
	// collect up and return errors
	errs := []error{}
	for err := range ch {
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 1 {
		return util.Flatten(errs)
	} else if len(errs) == 1 {
		return errs[0]
	}
	return nil
}

// untar reads the tar file from r and writes it into dir.
func untar(r io.Reader, dir string) (err error) {
	tr := tar.NewReader(r)
	for {
		f, err := tr.Next()

		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return errors.Wrapf(err, "tar reading error: %v", err)
		case f == nil:
			continue
		}

		rel := filepath.FromSlash(f.Name)
		abs := filepath.Join(dir, rel)

		switch f.Typeflag {
		case tar.TypeReg:
			wf, err := os.OpenFile(abs, os.O_CREATE|os.O_RDWR, os.FileMode(f.Mode))
			if err != nil {
				return err
			}
			n, err := io.Copy(wf, tr)
			if closeErr := wf.Close(); closeErr != nil && err == nil {
				err = closeErr
			}
			if err != nil {
				return errors.Errorf("error writing to %s: %v", abs, err)
			}
			if n != f.Size {
				return errors.Errorf("only wrote %d bytes to %s; expected %d", n, abs, f.Size)
			}
		case tar.TypeDir:
			if _, err := os.Stat(abs); err != nil {
				if err := os.MkdirAll(abs, 0755); err != nil {
					return err
				}
			}
		default:
			log.Warningf("tar file entry %s contained unsupported file type %v", f.Name, f.Typeflag)
		}
	}
}
