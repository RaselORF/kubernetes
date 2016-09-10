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

package system

import (
	"fmt"

	"github.com/docker/engine-api/client"
	"golang.org/x/net/context"
)

var _ Validator = &DockerValidator{}

// DockerValidator validates docker configuration.
type DockerValidator struct{}

func (d *DockerValidator) Name() string {
	return "docker"
}

const (
	dockerEndpoint     = "unix:///var/run/docker.sock"
	dockerConfigPrefix = "DOCKER_"
)

// TODO(random-liu): Add more validating items.
func (d *DockerValidator) Validate(spec SysSpec) error {
	if spec.RuntimeSpec.DockerSpec == nil {
		// If DockerSpec is not specified, assume current runtime is not
		// docker, skip the docker configuration validation.
		return nil
	}
	c, err := client.NewClient(dockerEndpoint, "", nil, nil)
	if err != nil {
		return fmt.Errorf("failed to create docker client: %v", err)
	}
	info, err := c.Info(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get docker info: %v", err)
	}
	// Valiate graph driver.
	item := dockerConfigPrefix + "GRAPH_DRIVER"
	for _, d := range spec.RuntimeSpec.GraphDriver {
		if info.Driver == d {
			report(item, info.Driver, good)
			return nil
		}
	}
	report(item, info.Driver, bad)
	return fmt.Errorf("unsupported graph driver: %s", info.Driver)
}
