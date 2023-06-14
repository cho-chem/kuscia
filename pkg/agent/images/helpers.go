/*
Copyright 2016 The Kubernetes Authors.

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

// Modified by Ant Group in 2023.

package images

import (
	"fmt"

	"k8s.io/client-go/util/flowcontrol"
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1"
	"k8s.io/kubernetes/pkg/credentialprovider"

	pkgcontainer "github.com/secretflow/kuscia/pkg/agent/container"
)

// throttleImagePulling wraps pkgcontainer.ImageService to throttle image
// pulling based on the given QPS and burst limits. If QPS is zero, defaults
// to no throttling.
func throttleImagePulling(imageService pkgcontainer.ImageService, qps float32, burst int) pkgcontainer.ImageService {
	if qps == 0.0 {
		return imageService
	}
	return &throttledImageService{
		ImageService: imageService,
		limiter:      flowcontrol.NewTokenBucketRateLimiter(qps, burst),
	}
}

type throttledImageService struct {
	pkgcontainer.ImageService
	limiter flowcontrol.RateLimiter
}

func (ts throttledImageService) PullImage(image pkgcontainer.ImageSpec, auth *credentialprovider.AuthConfig, podSandboxConfig *runtimeapi.PodSandboxConfig) (string, error) {
	if ts.limiter.TryAccept() {
		return ts.ImageService.PullImage(image, auth, podSandboxConfig)
	}
	return "", fmt.Errorf("pull QPS exceeded")
}
