/*
Copyright 2014 Google Inc. All rights reserved.

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

package kubelet

// LifecycleOutput is the output from a lifecycle event
type LifecycleOutput struct {
	Details string
}

// Lifecycle is an interface implemented by things that handle lifecycle events
type Lifecycle interface {
	// OnStart is called when a container is started.  The management system blocks until it returns
	OnStart(containerID string) (LifecycleOutput, error)
	// OnStop is called when a container is stopped.  The management system blocks until it returns
	OnStop(containerID string) (LifecycleOutput, error)
}

func newCommandLineLifecycle(r ContainerCommandRunner) Lifecycle {
	return &commandLineLifecycle{
		runner: r,
	}
}

type commandLineLifecycle struct {
	runner ContainerCommandRunner
}

func (c *commandLineLifecycle) runLifecycleCommand(containerID string, cmd []string) (LifecycleOutput, error) {
	data, err := c.runner.RunInContainer(containerID, cmd)
	return LifecycleOutput{string(data)}, err
}

func (c *commandLineLifecycle) OnStart(containerID string) (LifecycleOutput, error) {
	return c.runLifecycleCommand(containerID, []string{"start.sh"})
}

func (c *commandLineLifecycle) OnStop(containerID string) (LifecycleOutput, error) {
	return c.runLifecycleCommand(containerID, []string{"stop.sh"})
}
