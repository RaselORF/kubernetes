/*
Copyright 2017 The Kubernetes Authors.

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

package daemon

import (
	"fmt"

	"github.com/golang/glog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	intstrutil "k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/kubernetes/pkg/api/v1"
	extensions "k8s.io/kubernetes/pkg/apis/extensions/v1beta1"
	"k8s.io/kubernetes/pkg/controller/daemon/util"
	podutil "k8s.io/kubernetes/pkg/controller/deployment/util"
)

func (dsc *DaemonSetsController) rollingUpdate(ds *extensions.DaemonSet) error {
	newPods, oldPods, err := dsc.getAllDaemonSetPods(ds)
	allPods := append(oldPods, newPods...)

	maxUnavailable, numUnavailable, err := dsc.GetUnavailableNumbers(ds, allPods)
	if err != nil {
		glog.Errorf("Couldn't get unavailable numbers: %v", err)
		return nil
	}
	oldAvailablePods, oldUnavailablePods := util.SplitByAvailablePods(ds, oldPods)

	// for oldPods delete all not running pods
	var podsToDelete []string
	glog.V(4).Infof("Marking all unavailable old pods for deletion")
	for _, pod := range oldUnavailablePods {
		glog.V(4).Infof("Marking pod %s/%s for deletion", ds.Name, pod.Name)
		podsToDelete = append(podsToDelete, pod.Name)
	}

	glog.V(4).Infof("Marking old pods for deletion")
	for _, pod := range oldAvailablePods {
		if numUnavailable >= maxUnavailable {
			glog.V(4).Infof("Number of unavailable DaemonSet pods: %d, is equal to or exceeds allowed maximum: %d", numUnavailable, maxUnavailable)
			break
		}
		glog.V(4).Infof("Marking pod %s/%s for deletion", ds.Name, pod.Name)
		podsToDelete = append(podsToDelete, pod.Name)
		numUnavailable++
	}
	errors := dsc.syncNodes(ds, podsToDelete, []string{})
	return utilerrors.NewAggregate(errors)
}

func (dsc *DaemonSetsController) getAllDaemonSetPods(ds *extensions.DaemonSet) ([]*v1.Pod, []*v1.Pod, error) {
	var newPods []*v1.Pod
	var oldPods []*v1.Pod

	selector, err := metav1.LabelSelectorAsSelector(ds.Spec.Selector)
	if err != nil {
		return newPods, oldPods, err
	}
	daemonPods, err := dsc.podLister.Pods(ds.Namespace).List(selector)
	if err != nil {
		return newPods, oldPods, fmt.Errorf("Couldn't get list of pods for daemon set %#v: %v", ds, err)
	}
	dsPodTemplateSpecHash := podutil.GetPodTemplateSpecHashFnv(ds.Spec.Template)
	for _, pod := range daemonPods {
		if util.IsPodUpdated(dsPodTemplateSpecHash, pod) {
			newPods = append(newPods, pod)
		} else {
			oldPods = append(oldPods, pod)
		}
	}
	return newPods, oldPods, nil
}

func (dsc *DaemonSetsController) GetUnavailableNumbers(ds *extensions.DaemonSet, allPods []*v1.Pod) (int, int, error) {
	nodeList, err := dsc.nodeLister.List(labels.Everything())
	if err != nil {
		return -1, -1, fmt.Errorf("couldn't get list of nodes during rolling update of daemon set %#v: %v", ds, err)
	}

	var desiredNumberScheduled int
	for i := range nodeList {
		node := nodeList[i]
		wantToRun, _, _, err := dsc.nodeShouldRunDaemonPod(node, ds)
		if err != nil {
			return -1, -1, err
		}
		if wantToRun {
			desiredNumberScheduled++
		}
	}

	numUnavailable := desiredNumberScheduled - len(allPods)
	for _, pod := range allPods {
		if !v1.IsPodAvailable(pod, ds.Spec.MinReadySeconds, metav1.Now()) {
			numUnavailable++
		}
	}

	maxUnavailable, err := intstrutil.GetValueFromIntOrPercent(ds.Spec.UpdateStrategy.RollingUpdate.MaxUnavailable, desiredNumberScheduled, true)
	if err != nil {
		glog.Errorf("Invalid value for MaxUnavailable: %v", err)
		return -1, -1, nil
	}
	return maxUnavailable, numUnavailable, nil
}
