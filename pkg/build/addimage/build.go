/*
Copyright 2020 The Kubernetes Authors.

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

package addimage

import (
	"runtime"

	"sigs.k8s.io/kind/pkg/apis/config/defaults"
	"sigs.k8s.io/kind/pkg/log"
)

// Build builds a node image using the supplied options
func Build(options ...Option) error {
	// default options
	ctx := &buildContext{
		image:     DefaultImage,
		baseImage: defaults.Image,
		logger:    log.NoopLogger{},
		arch:      runtime.GOARCH,
	}

	// apply user options
	for _, option := range options {
		if err := option.apply(ctx); err != nil {
			return err
		}
	}

	// verify that we're using a supported arch
	if !supportedArch(ctx.arch) {
		ctx.logger.Warnf("unsupported architecture %q", ctx.arch)
	}
	return ctx.Build()
}

func supportedArch(arch string) bool {
	switch arch {
	default:
		return false
	// currently we nominally support building node images for these
	case "amd64":
	case "arm64":
	case "ppc64le":
	}
	return true
}