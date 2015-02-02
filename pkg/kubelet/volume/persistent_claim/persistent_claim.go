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

package persistent_claim

import (
	"fmt"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/kubelet/volume"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/types"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/kubelet/volume/host_path"
	"github.com/golang/glog"
)

// This is the primary entrypoint for volume plugins.
func ProbeVolumePlugins() []volume.Plugin {
	return []volume.Plugin{&persistentClaimPlugin{nil}}
}

type persistentClaimPlugin struct {
	host       volume.Host
}

var _ volume.Plugin = &persistentClaimPlugin{}

const (
	persistentClaimPluginName = "kubernetes.io/persistent-claim"
)

func (plugin *persistentClaimPlugin) Init(host volume.Host) {
	plugin.host = host
}

func (plugin *persistentClaimPlugin) Name() string {
	return persistentClaimPluginName
}

func (plugin *persistentClaimPlugin) CanSupport(spec *api.Volume) bool {

	if spec.Source.PersistentVolumeClaimAttachment != nil ||
		spec.Source.HostPath != nil ||
		spec.Source.AWSElasticBlockStore != nil ||
		spec.Source.GCEPersistentDisk != nil ||
		spec.Source.NFSMount != nil {
			return true
	}

	return false
}

func (plugin *persistentClaimPlugin) NewBuilder(spec *api.Volume, podUID types.UID) (volume.Builder, error) {

	volPlugin := getVolumePlugin(spec)

	if volPlugin != nil {

		// get the real volume backing this claim

		claimName := spec.Source.PersistentVolumeClaimAttachment.PersistentVolumeClaimReference.Name
		volNamespace := spec.Source.PersistentVolumeClaimAttachment.PersistentVolumeClaimReference.Namespace
		claim, err := plugin.host.GetKubeClient().PersistentVolumeClaims(volNamespace).Get(claimName)

		if err != nil {
			glog.V(3).Infof("Error finding claim in namespace  PersistentVolume by ClaimRef: %+v\n", spec.Source.PersistentVolumeClaimAttachment.PersistentVolumeClaimReference)
			return nil, err
		}

		pv, err := plugin.host.GetKubeClient().PersistentVolumes().Get(claim.Status.VolumeRef.Name)

		if err != nil {
			glog.V(3).Infof("Error finding bound PersistentVolume by ClaimRef: %+v\n", spec.Source.PersistentVolumeClaimAttachment.PersistentVolumeClaimReference)
			return nil, err
		}

		wrapper := &api.Volume{
			Name: spec.Name,
			Source: pv.Spec.Source,
		}

		glog.V(3).Infof("WRAPPER -- %+v\n", wrapper)

		return volPlugin.NewBuilder(wrapper, podUID)
	}

	return nil, fmt.Errorf("Error creating builder for volume %+v", spec)
}

func (plugin *persistentClaimPlugin) NewCleaner(volName string, podUID types.UID) (volume.Cleaner, error) {
	return nil, fmt.Errorf("This should never be called.  There are no cleaners for persistent volume sources.  The cleaner for %s will come right from its own plugin.")
}


func getVolumePlugin(spec *api.Volume) volume.Plugin {

	if spec.Source.HostPath != nil {
		return host_path.ProbeVolumePlugins()[0]
	} else if spec.Source.GCEPersistentDisk != nil {
		return host_path.ProbeVolumePlugins()[0]
	} else if spec.Source.AWSElasticBlockStore != nil {
		return host_path.ProbeVolumePlugins()[0]
	} else if spec.Source.NFSMount != nil {
		return host_path.ProbeVolumePlugins()[0]
	}

	return nil
}
