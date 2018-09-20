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

// Package docker contains helpers for working with docker
package docker

import (
	"strings"
)

// JoinNameAndTag combines a docker image name and tag
// The tag may be empty, a tag as in name:tag, or a a sha as in `@sha256:asdf`
func JoinNameAndTag(name, tag string) string {
	// Join the name and tag with a colon IFF the tag is not any of:
	// - empty, in which case we can just the name
	// - starts with @, this is either a digest, or an invalid value
	// - starts with :, we can gracefully not double up the :
	if tag != "" && !strings.HasPrefix(tag, "@") && !strings.HasPrefix(tag, ":") {
		return name + ":" + tag
	}
	return name + tag
}
