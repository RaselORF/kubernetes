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

package etcd

import (
	"strings"
	"testing"
	"time"

	"fmt"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api/latest"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api/rest/resttest"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/fields"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/runtime"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/tools"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/tools/etcdtest"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/util"

	"github.com/coreos/go-etcd/etcd"
)

func newHelper(t *testing.T) (*tools.FakeEtcdClient, tools.EtcdHelper) {
	fakeEtcdClient := tools.NewFakeEtcdClient(t)
	fakeEtcdClient.TestIndex = true
	helper := tools.NewEtcdHelper(fakeEtcdClient, latest.Codec, etcdtest.PathPrefix())
	return fakeEtcdClient, helper
}

func newStorage(t *testing.T) (*REST, *StatusREST, *tools.FakeEtcdClient, tools.EtcdHelper) {
	fakeEtcdClient, h := newHelper(t)
	storage, statusStorage := NewStorage(h)
	return storage, statusStorage, fakeEtcdClient, h
}

func validNewAutoScaler(name string) *api.AutoScaler {
	return &api.AutoScaler{
		ObjectMeta: api.ObjectMeta{
			Name:      name,
			Namespace: api.NamespaceDefault,
		},
		Spec: api.AutoScalerSpec{
			TargetSelector:  map[string]string{"foo": "bar"},
			MonitoringSources: []string{"fizz", "buzz"},
		},
	}
}

func TestCreate(t *testing.T) {
	storage, _, fakeEtcdClient, _ := newStorage(t)
	test := resttest.New(t, storage, fakeEtcdClient.SetError)
	autoScaler := validNewAutoScaler("create-test")
	autoScaler.ObjectMeta = api.ObjectMeta{}
	test.TestCreate(
		//valid
		autoScaler,
		//invalid
		&api.AutoScaler{},
	)
}

func TestCreateRegistryError(t *testing.T) {
	fakeEtcdClient, helper := newHelper(t)
	fakeEtcdClient.Err = fmt.Errorf("test error")
	storage, _ := NewStorage(helper)

	autoScaler := validNewAutoScaler("foo")
	_, err := storage.Create(api.NewDefaultContext(), autoScaler)
	if err != fakeEtcdClient.Err {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCreateSetsFields(t *testing.T) {
	ctx := api.NewDefaultContext()
	fakeEtcdClient, helper := newHelper(t)
	storage, _ := NewStorage(helper)
	autoScaler := validNewAutoScaler("foo")
	_, err := storage.Create(ctx, autoScaler)
	if err != fakeEtcdClient.Err {
		t.Fatalf("unexpected error: %v", err)
	}

	actual := &api.AutoScaler{}
	key, err := storage.KeyFunc(ctx, "foo")
	if err != nil {
		t.Fatalf("unexpected key error: %v", err)
	}
	if err := helper.ExtractObj(key, actual, false); err != nil {
		t.Fatalf("unexpected extraction error: %v", err)
	}
	if actual.Name != autoScaler.Name {
		t.Errorf("unexpected autoScaler: %#v", actual)
	}
	if len(actual.UID) == 0 {
		t.Errorf("expected UID to be set: %#v", actual)
	}
}

func TestListError(t *testing.T) {
	fakeEtcdClient, helper := newHelper(t)
	fakeEtcdClient.Err = fmt.Errorf("test error")
	storage, _ := NewStorage(helper)
	scalers, err := storage.List(api.NewDefaultContext(), labels.Everything(), fields.Everything())
	if err != fakeEtcdClient.Err {
		t.Fatalf("Expected %#v, Got %#v", fakeEtcdClient.Err, err)
	}
	if scalers != nil {
		t.Errorf("Unexpected non-nil autoscaler list: %#v", scalers)
	}
}

func TestListEmptyAutoScalerList(t *testing.T) {
	ctx := api.NewDefaultContext()
	storage, _, fakeEtcdClient, _ := newStorage(t)
	fakeEtcdClient.ChangeIndex = 1
	key := storage.KeyRootFunc(ctx)
	key = etcdtest.AddPrefix(key)
	fakeEtcdClient.Data[key] = tools.EtcdResponseWithError{
		R: &etcd.Response{},
		E: fakeEtcdClient.NewError(tools.EtcdErrorCodeNotFound),
	}

	scalers, err := storage.List(ctx, labels.Everything(), fields.Everything())
	if err != nil {
		t.Fatalf("Unexpected error: %#v", err)
	}
	if len(scalers.(*api.AutoScalerList).Items) != 0 {
		t.Errorf("Unexpected non-zero list: %#v", scalers)
	}
	if scalers.(*api.AutoScalerList).ResourceVersion != "1" {
		t.Errorf("Unexpected resource version: %#v", scalers)
	}
}

func TestListAutoScalerList(t *testing.T) {
	ctx := api.NewDefaultContext()
	storage, _, fakeEtcdClient, _ := newStorage(t)
	fakeEtcdClient.ChangeIndex = 1

	autoScaler := validNewAutoScaler("foo")

	key := storage.KeyRootFunc(ctx)
	key = etcdtest.AddPrefix(key)
	fakeEtcdClient.Data[key] = tools.EtcdResponseWithError{
		R: &etcd.Response{
			Node: &etcd.Node{
				Nodes: []*etcd.Node{
					{
						Value: runtime.EncodeOrDie(latest.Codec, autoScaler),
					},
				},
			},
		},
	}

	scalers, err := storage.List(ctx, labels.Everything(), fields.Everything())
	scalerList := scalers.(*api.AutoScalerList)

	if err != nil {
		t.Fatalf("Unexpected error: %#v", err)
	}
	if len(scalerList.Items) != 1 {
		t.Errorf("Unexpected list: %#v", scalers)
	}
	if scalerList.Items[0].Name != autoScaler.Name {
		t.Errorf("Unexpected autoscaler: %#v", scalerList.Items[0])
	}
}

func TestListAutoScalerListSelection(t *testing.T) {
	ctx := api.NewDefaultContext()
	storage, _, fakeEtcdClient, _ := newStorage(t)
	fakeEtcdClient.ChangeIndex = 1

	validScaler1 := validNewAutoScaler("scaler1")
	validScaler2 := validNewAutoScaler("scaler2")
	validScaler3 := validNewAutoScaler("scaler3")
	validScaler3.ObjectMeta.Labels = map[string]string{"label": "test"}

	key := storage.KeyRootFunc(ctx)
	key = etcdtest.AddPrefix(key)
	fakeEtcdClient.Data[key] = tools.EtcdResponseWithError{
		R: &etcd.Response{
			Node: &etcd.Node{
				Nodes: []*etcd.Node{
					{
						Value: runtime.EncodeOrDie(latest.Codec, validScaler1),
					},
					{
						Value: runtime.EncodeOrDie(latest.Codec, validScaler2),
					},
					{
						Value: runtime.EncodeOrDie(latest.Codec, validScaler3),
					},
				},
			},
		},
	}

	testCases := []struct {
		name, label, field string
		expectedNames      util.StringSet
	}{
		{
			name:          "everything",
			expectedNames: util.NewStringSet(validScaler1.Name, validScaler2.Name, validScaler3.Name),
		},
		{
			name:          "field selector found",
			field:         "name=scaler1",
			expectedNames: util.NewStringSet(validScaler1.Name),
		},
		{
			name:          "field selector not found",
			field:         "foo=bar",
			expectedNames: util.NewStringSet(),
		},
		{
			name:          "label selector not found",
			label:         "biz=baz",
			expectedNames: util.NewStringSet(),
		},
		{
			name:          "label selector found",
			label:         "label=test",
			expectedNames: util.NewStringSet(validScaler3.Name),
		},
	}

	for _, tc := range testCases {
		label, err := labels.Parse(tc.label)
		if err != nil {
			t.Errorf("%s failed with error: %v", tc.name, err)
			continue
		}

		field, err := fields.ParseSelector(tc.field)
		if err != nil {
			t.Errorf("%s failed with error: %v", tc.name, err)
			continue
		}

		autoScalerObjs, err := storage.List(ctx, label, field)
		if err != nil {
			t.Errorf("%s failed with error: %v", tc.name, err)
			continue
		}
		autoScalerList := autoScalerObjs.(*api.AutoScalerList)

		if e, a := len(tc.expectedNames), len(autoScalerList.Items); e != a {
			t.Errorf("%s failed.  Unexpected number of objects were returned.  Expected: %d got: %d", tc.name, e, a)
			continue
		}

		for _, autoScaler := range autoScalerList.Items {
			if !tc.expectedNames.Has(autoScaler.Name) {
				t.Errorf("%s failed: unexpected name was returned: %s", tc.name, autoScaler.Name)
			}
		}
	}
}

func TestAutoScalerDecode(t *testing.T) {
	storage, _ := NewStorage(tools.EtcdHelper{})
	expected := validNewAutoScaler("foo")
	body, err := latest.Codec.Encode(expected)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	actual := storage.New()
	if err := latest.Codec.DecodeInto(body, actual); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !api.Semantic.DeepEqual(expected, actual) {
		t.Errorf("mismatch: %s", util.ObjectDiff(expected, actual))
	}
}

func TestGet(t *testing.T) {
	expect := validNewAutoScaler("foo")

	storage, _, fakeEtcdClient, _ := newStorage(t)
	ctx := api.NewDefaultContext()
	key, err := storage.KeyFunc(api.WithNamespace(ctx, "test"), "foo")
	if err != nil {
		t.Fatalf("unexpected key error: %v", err)
	}

	key = etcdtest.AddPrefix(key)
	fakeEtcdClient.Data[key] = tools.EtcdResponseWithError{
		R: &etcd.Response{
			Node: &etcd.Node{
				Value: runtime.EncodeOrDie(latest.Codec, expect),
			},
		},
	}

	obj, err := storage.Get(api.WithNamespace(ctx, "test"), "foo")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	autoScaler := obj.(*api.AutoScaler)
	if e, a := expect, autoScaler; !api.Semantic.DeepEqual(e, a) {
		t.Errorf("Unexpected autoscaler: %s", util.ObjectDiff(e, a))
	}
}

func TestUpdateWithConflictingNamespace(t *testing.T) {
	expect := validNewAutoScaler("foo")

	storage, _, fakeEtcdClient, _ := newStorage(t)
	ctx := api.NewDefaultContext()
	key, err := storage.KeyFunc(ctx, "foo")
	if err != nil {
		t.Fatalf("unexpected key error: %v", err)
	}

	key = etcdtest.AddPrefix(key)
	fakeEtcdClient.Data[key] = tools.EtcdResponseWithError{
		R: &etcd.Response{
			Node: &etcd.Node{
				Value:         runtime.EncodeOrDie(latest.Codec, expect),
				ModifiedIndex: 1,
			},
		},
	}
	expect.Namespace = "not-default"
	expect.ResourceVersion = "1"
	obj, created, err := storage.Update(ctx, expect)

	if obj != nil || created {
		t.Error("Expected a nil object, but we got a value or created")
	}
	if err == nil {
		t.Errorf("Expected an error, but we didn't get one")
	} else if strings.Index(err.Error(), "the namespace of the provided object does not match the namespace sent on the request") == -1 {
		t.Errorf("Expected 'AutoScaler.Namespace does not match the provided context' error, got '%v'", err.Error())
	}
}

func TestDelete(t *testing.T) {
	expect := validNewAutoScaler("foo")

	storage, _, fakeEtcdClient, _ := newStorage(t)
	ctx := api.NewDefaultContext()
	key, err := storage.KeyFunc(ctx, "foo")
	if err != nil {
		t.Fatalf("unexpected key error: %v", err)
	}

	key = etcdtest.AddPrefix(key)
	fakeEtcdClient.Data[key] = tools.EtcdResponseWithError{
		R: &etcd.Response{
			Node: &etcd.Node{
				Value:         runtime.EncodeOrDie(latest.Codec, expect),
				ModifiedIndex: 1,
			},
		},
	}
	result, err := storage.Delete(ctx, expect.Name, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatalf("got no deleted object where one was expected")
	}
}

func TestWatch(t *testing.T) {
	storage, _, fakeEtcdClient, _ := newStorage(t)
	watching, err := storage.Watch(api.NewDefaultContext(), labels.Everything(), fields.Everything(), "1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	fakeEtcdClient.WaitForWatchCompletion()
	select {
	case _, ok := <-watching.ResultChan():
		if !ok {
			t.Errorf("watching channel should be open")
		}
	default:
	}
	fakeEtcdClient.WatchInjectError <- nil
	if _, ok := <-watching.ResultChan(); ok {
		t.Errorf("watching channel should be closed")
	}
	watching.Stop()
}

func TestWatchWithSelectorMatch(t *testing.T) {
	storage, _, fakeEtcdClient, _ := newStorage(t)
	watching, err := storage.Watch(api.NewDefaultContext(),
		labels.SelectorFromSet(labels.Set{"name": "foo"}),
		fields.Everything(),
		"1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	fakeEtcdClient.WaitForWatchCompletion()

	autoScaler := validNewAutoScaler("foo")
	autoScaler.Labels = map[string]string{"name": "foo"}
	encoded, _ := latest.Codec.Encode(autoScaler)
	fakeEtcdClient.WatchResponse <- &etcd.Response{
		Action: "create",
		Node: &etcd.Node{
			Value: string(encoded),
		},
	}

	select {
	case _, ok := <-watching.ResultChan():
		if !ok {
			t.Errorf("watching channel should be open")
		}
	case <-time.After(time.Millisecond * 100):
		t.Error("unexpected timeout from result channel")
	}

	watching.Stop()
}

func TestWatchWithSelectorNoMatch(t *testing.T) {
	storage, _, fakeEtcdClient, _ := newStorage(t)
	watching, err := storage.Watch(api.NewDefaultContext(),
		labels.SelectorFromSet(labels.Set{"name": "foo"}),
		fields.Everything(),
		"1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	fakeEtcdClient.WaitForWatchCompletion()

	autoScaler := validNewAutoScaler("foo")
	encoded, _ := latest.Codec.Encode(autoScaler)
	fakeEtcdClient.WatchResponse <- &etcd.Response{
		Action: "create",
		Node: &etcd.Node{
			Value: string(encoded),
		},
	}

	select {
	case <-watching.ResultChan():
		t.Errorf("unexpected object returned during watch")
	case <-time.After(time.Millisecond * 100):
		//expected
	}

	watching.Stop()
}
