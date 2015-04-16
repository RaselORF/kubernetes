/*
Copyright 2015 Google Inc. All rights reserved.

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

package testclient

import (
	"testing"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api/errors"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api/latest"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/runtime"
)

func TestNewClient(t *testing.T) {
	o := NewObjects(api.Scheme)
	if err := AddObjectsFromPath("../../../examples/guestbook/frontend-service.json", o); err != nil {
		t.Fatal(err)
	}
	client := &Fake{ReactFn: ObjectReaction(o, latest.RESTMapper)}
	list, err := client.Services("test").List(labels.Everything())
	if err != nil {
		t.Fatal(err)
	}
	if len(list.Items) != 1 {
		t.Fatalf("unexpected list %#v", list)
	}

	// When list is invoked a second time, the same results are returned.
	list, err = client.Services("test").List(labels.Everything())
	if err != nil {
		t.Fatal(err)
	}
	if len(list.Items) != 1 {
		t.Fatalf("unexpected list %#v", list)
	}
	t.Logf("list: %#v", list)
}

func TestErrors(t *testing.T) {
	o := NewObjects(api.Scheme)
	o.Add(&api.List{
		Items: []runtime.Object{
			// This first call to List will return this error
			&(errors.NewNotFound("ServiceList", "").(*errors.StatusError).ErrStatus),
			// The second call to List will return this error
			&(errors.NewForbidden("ServiceList", "", nil).(*errors.StatusError).ErrStatus),
		},
	})
	client := &Fake{ReactFn: ObjectReaction(o, latest.RESTMapper)}
	_, err := client.Services("test").List(labels.Everything())
	if !errors.IsNotFound(err) {
		t.Fatalf("unexpected error: %v", err)
	}
	t.Logf("error: %#v", err.(*errors.StatusError).Status())
	_, err = client.Services("test").List(labels.Everything())
	if !errors.IsForbidden(err) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDeleteFromClient(t *testing.T) {
	podA := &api.Pod{ObjectMeta: api.ObjectMeta{Name: "foo", Namespace: "ns"}}
	podB := &api.Pod{ObjectMeta: api.ObjectMeta{Name: "bar", Namespace: "ns"}}

	o := NewObjects(api.Scheme)
	o.Add(podA)
	o.Add(podB)

	client := &Fake{ReactFn: ObjectReaction(o, latest.RESTMapper)}
	pod, err := client.Pods(api.NamespaceAll).Get("any")
	if err != nil {
		t.Fatal(err)
	}
	if pod.Name != podA.Name {
		t.Errorf("Expected %+v\n\n but got %+v\n", podA, pod)
	}

	o.Delete(podA)
	pod, err = client.Pods(api.NamespaceAll).Get("any")
	if err != nil {
		t.Fatal(err)
	}
	if pod.Name != podB.Name {
		t.Errorf("Expected %+v\n\n but got %+v\n", podB, pod)
	}
}
