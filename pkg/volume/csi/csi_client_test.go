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

package csi

import (
	"context"
	"errors"
	"reflect"
	"testing"

	csipb "github.com/container-storage-interface/spec/lib/go/csi/v0"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	api "k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/volume/csi/fake"
)

type fakeCsiDriverClient struct {
	t          *testing.T
	nodeClient *fake.NodeClient
}

func newFakeCsiDriverClient(t *testing.T, stagingCapable bool) *fakeCsiDriverClient {
	return &fakeCsiDriverClient{
		t:          t,
		nodeClient: fake.NewNodeClient(stagingCapable),
	}
}

func (c *fakeCsiDriverClient) NodeGetInfo(ctx context.Context) (
	nodeID string,
	maxVolumePerNode int64,
	accessibleTopology *csipb.Topology,
	err error) {
	resp, err := c.nodeClient.NodeGetInfo(ctx, &csipb.NodeGetInfoRequest{})
	return resp.GetNodeId(), resp.GetMaxVolumesPerNode(), resp.GetAccessibleTopology(), err
}

// for pre-v1.0
func (c *fakeCsiDriverClient) NodeGetId(ctx context.Context) (nodeID string, err error) {
	resp, err := c.nodeClient.NodeGetId(ctx, &csipb.NodeGetIdRequest{})
	return resp.GetNodeId(), err
}

func (c *fakeCsiDriverClient) NodePublishVolume(
	ctx context.Context,
	volID string,
	readOnly bool,
	stagingTargetPath string,
	targetPath string,
	accessMode api.PersistentVolumeAccessMode,
	volumeInfo map[string]string,
	volumeAttribs map[string]string,
	nodePublishSecrets map[string]string,
	fsType string,
) error {
	c.t.Log("calling fake.NodePublishVolume...")
	req := &csipb.NodePublishVolumeRequest{
		VolumeId:           volID,
		TargetPath:         targetPath,
		Readonly:           readOnly,
		PublishInfo:        volumeInfo,
		VolumeAttributes:   volumeAttribs,
		NodePublishSecrets: nodePublishSecrets,
		VolumeCapability: &csipb.VolumeCapability{
			AccessMode: &csipb.VolumeCapability_AccessMode{
				Mode: asCSIAccessMode(accessMode),
			},
			AccessType: &csipb.VolumeCapability_Mount{
				Mount: &csipb.VolumeCapability_MountVolume{
					FsType: fsType,
				},
			},
		},
	}

	_, err := c.nodeClient.NodePublishVolume(ctx, req)
	return err
}

func (c *fakeCsiDriverClient) NodeUnpublishVolume(ctx context.Context, volID string, targetPath string) error {
	c.t.Log("calling fake.NodeUnpublishVolume...")
	req := &csipb.NodeUnpublishVolumeRequest{
		VolumeId:   volID,
		TargetPath: targetPath,
	}

	_, err := c.nodeClient.NodeUnpublishVolume(ctx, req)
	return err
}

func (c *fakeCsiDriverClient) NodeStageVolume(ctx context.Context,
	volID string,
	publishInfo map[string]string,
	stagingTargetPath string,
	fsType string,
	accessMode api.PersistentVolumeAccessMode,
	nodeStageSecrets map[string]string,
	volumeAttribs map[string]string,
) error {
	c.t.Log("calling fake.NodeStageVolume...")
	req := &csipb.NodeStageVolumeRequest{
		VolumeId:          volID,
		PublishInfo:       publishInfo,
		StagingTargetPath: stagingTargetPath,
		VolumeCapability: &csipb.VolumeCapability{
			AccessMode: &csipb.VolumeCapability_AccessMode{
				Mode: asCSIAccessMode(accessMode),
			},
			AccessType: &csipb.VolumeCapability_Mount{
				Mount: &csipb.VolumeCapability_MountVolume{
					FsType: fsType,
				},
			},
		},
		NodeStageSecrets: nodeStageSecrets,
		VolumeAttributes: volumeAttribs,
	}

	_, err := c.nodeClient.NodeStageVolume(ctx, req)
	return err
}

func (c *fakeCsiDriverClient) NodeUnstageVolume(ctx context.Context, volID, stagingTargetPath string) error {
	c.t.Log("calling fake.NodeUnstageVolume...")
	req := &csipb.NodeUnstageVolumeRequest{
		VolumeId:          volID,
		StagingTargetPath: stagingTargetPath,
	}
	_, err := c.nodeClient.NodeUnstageVolume(ctx, req)
	return err
}

func (c *fakeCsiDriverClient) NodeGetCapabilities(ctx context.Context) ([]*csipb.NodeServiceCapability, error) {
	c.t.Log("calling fake.NodeGetCapabilities...")
	req := &csipb.NodeGetCapabilitiesRequest{}
	resp, err := c.nodeClient.NodeGetCapabilities(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.GetCapabilities(), nil
}

func setupClient(t *testing.T, stageUnstageSet bool) csiClient {
	return newFakeCsiDriverClient(t, stageUnstageSet)
}

func TestClientNodeGetInfo(t *testing.T) {
	testCases := []struct {
		name                       string
		expectedNodeID             string
		expectedMaxVolumePerNode   int64
		expectedAccessibleTopology *csipb.Topology
		mustFail                   bool
		err                        error
	}{
		{
			name:                     "test ok",
			expectedNodeID:           "node1",
			expectedMaxVolumePerNode: 16,
			expectedAccessibleTopology: &csipb.Topology{
				Segments: map[string]string{"com.example.csi-topology/zone": "zone1"},
			},
		},
		{name: "grpc error", mustFail: true, err: errors.New("grpc error")},
	}

	client := setupClient(t, false /* stageUnstageSet */)

	for _, tc := range testCases {
		t.Logf("test case: %s", tc.name)
		client.(*fakeCsiDriverClient).nodeClient.SetNextError(tc.err)
		client.(*fakeCsiDriverClient).nodeClient.SetNodeGetInfoResp(&csipb.NodeGetInfoResponse{
			NodeId:             tc.expectedNodeID,
			MaxVolumesPerNode:  tc.expectedMaxVolumePerNode,
			AccessibleTopology: tc.expectedAccessibleTopology,
		})
		nodeID, maxVolumePerNode, accessibleTopology, err := client.NodeGetInfo(context.Background())

		if tc.mustFail && err == nil {
			t.Error("expected an error but got none")
		}

		if !tc.mustFail && err != nil {
			t.Errorf("expected no errors but got: %v", err)
		}

		if nodeID != tc.expectedNodeID {
			t.Errorf("expected nodeID: %v; got: %v", tc.expectedNodeID, nodeID)
		}

		if maxVolumePerNode != tc.expectedMaxVolumePerNode {
			t.Errorf("expected maxVolumePerNode: %v; got: %v", tc.expectedMaxVolumePerNode, maxVolumePerNode)
		}

		if !reflect.DeepEqual(accessibleTopology, tc.expectedAccessibleTopology) {
			t.Errorf("expected accessibleTopology: %v; got: %v", *tc.expectedAccessibleTopology, *accessibleTopology)
		}
	}
}

func TestClientNodeGetId(t *testing.T) {
	testCases := []struct {
		name           string
		expectedNodeID string
		mustFail       bool
		err            error
	}{
		{
			name:           "test ok",
			expectedNodeID: "node1",
		},
		{
			name:     "grpc error",
			mustFail: true,
			err:      errors.New("grpc error"),
		},
	}

	client := setupClient(t, false /* stageUnstageSet */)

	for _, tc := range testCases {
		t.Logf("test case: %s", tc.name)
		client.(*fakeCsiDriverClient).nodeClient.SetNextError(tc.err)
		client.(*fakeCsiDriverClient).nodeClient.SetNodeGetIdResp(&csipb.NodeGetIdResponse{
			NodeId: tc.expectedNodeID,
		})
		nodeID, err := client.NodeGetId(context.Background())

		if tc.mustFail && err == nil {
			t.Error("expected an error but got none")
		}

		if !tc.mustFail && err != nil {
			t.Errorf("expected no errors but got: %v", err)
		}

		if nodeID != tc.expectedNodeID {
			t.Errorf("expected nodeID: %v; got: %v", tc.expectedNodeID, nodeID)
		}
	}
}

func TestClientNodePublishVolume(t *testing.T) {
	testCases := []struct {
		name       string
		volID      string
		targetPath string
		fsType     string
		mustFail   bool
		err        error
	}{
		{name: "test ok", volID: "vol-test", targetPath: "/test/path"},
		{name: "missing volID", targetPath: "/test/path", mustFail: true},
		{name: "missing target path", volID: "vol-test", mustFail: true},
		{name: "bad fs", volID: "vol-test", targetPath: "/test/path", fsType: "badfs", mustFail: true},
		{name: "grpc error", volID: "vol-test", targetPath: "/test/path", mustFail: true, err: errors.New("grpc error")},
	}

	client := setupClient(t, false)

	for _, tc := range testCases {
		t.Logf("test case: %s", tc.name)
		client.(*fakeCsiDriverClient).nodeClient.SetNextError(tc.err)
		err := client.NodePublishVolume(
			context.Background(),
			tc.volID,
			false,
			"",
			tc.targetPath,
			api.ReadWriteOnce,
			map[string]string{"device": "/dev/null"},
			map[string]string{"attr0": "val0"},
			map[string]string{},
			tc.fsType,
		)

		if tc.mustFail && err == nil {
			t.Error("test must fail, but err is nil")
		}
	}
}

func TestClientNodeUnpublishVolume(t *testing.T) {
	testCases := []struct {
		name       string
		volID      string
		targetPath string
		mustFail   bool
		err        error
	}{
		{name: "test ok", volID: "vol-test", targetPath: "/test/path"},
		{name: "missing volID", targetPath: "/test/path", mustFail: true},
		{name: "missing target path", volID: "vol-test", mustFail: true},
		{name: "grpc error", volID: "vol-test", targetPath: "/test/path", mustFail: true, err: errors.New("grpc error")},
	}

	client := setupClient(t, false)

	for _, tc := range testCases {
		t.Logf("test case: %s", tc.name)
		client.(*fakeCsiDriverClient).nodeClient.SetNextError(tc.err)
		err := client.NodeUnpublishVolume(context.Background(), tc.volID, tc.targetPath)
		if tc.mustFail && err == nil {
			t.Error("test must fail, but err is nil")
		}
	}
}

func TestClientNodeStageVolume(t *testing.T) {
	testCases := []struct {
		name              string
		volID             string
		stagingTargetPath string
		fsType            string
		secret            map[string]string
		mustFail          bool
		err               error
	}{
		{name: "test ok", volID: "vol-test", stagingTargetPath: "/test/path", fsType: "ext4"},
		{name: "missing volID", stagingTargetPath: "/test/path", mustFail: true},
		{name: "missing target path", volID: "vol-test", mustFail: true},
		{name: "bad fs", volID: "vol-test", stagingTargetPath: "/test/path", fsType: "badfs", mustFail: true},
		{name: "grpc error", volID: "vol-test", stagingTargetPath: "/test/path", mustFail: true, err: errors.New("grpc error")},
	}

	client := setupClient(t, false)

	for _, tc := range testCases {
		t.Logf("Running test case: %s", tc.name)
		client.(*fakeCsiDriverClient).nodeClient.SetNextError(tc.err)
		err := client.NodeStageVolume(
			context.Background(),
			tc.volID,
			map[string]string{"device": "/dev/null"},
			tc.stagingTargetPath,
			tc.fsType,
			api.ReadWriteOnce,
			tc.secret,
			map[string]string{"attr0": "val0"},
		)

		if tc.mustFail && err == nil {
			t.Error("test must fail, but err is nil")
		}
	}
}

func TestClientNodeUnstageVolume(t *testing.T) {
	testCases := []struct {
		name              string
		volID             string
		stagingTargetPath string
		mustFail          bool
		err               error
	}{
		{name: "test ok", volID: "vol-test", stagingTargetPath: "/test/path"},
		{name: "missing volID", stagingTargetPath: "/test/path", mustFail: true},
		{name: "missing target path", volID: "vol-test", mustFail: true},
		{name: "grpc error", volID: "vol-test", stagingTargetPath: "/test/path", mustFail: true, err: errors.New("grpc error")},
	}

	client := setupClient(t, false)

	for _, tc := range testCases {
		t.Logf("Running test case: %s", tc.name)
		client.(*fakeCsiDriverClient).nodeClient.SetNextError(tc.err)
		err := client.NodeUnstageVolume(
			context.Background(),
			tc.volID, tc.stagingTargetPath,
		)
		if tc.mustFail && err == nil {
			t.Error("test must fail, but err is nil")
		}
	}
}

func TestGetNodeInfoWithFallback(t *testing.T) {
	var errRandom = errors.New("some other error")
	var errUnimplemented = status.Error(codes.Unimplemented, "unknown method someRandomMethod")

	testCases := map[string]struct {
		nodeGetInfoErr          error
		nodeGetIdErr            error
		expectedErr             error
		expectedGetNodeInfoCall bool
		expectedGetNodeIdCall   bool
	}{
		"no error on neither method": {
			nodeGetInfoErr:          nil,
			nodeGetIdErr:            nil,
			expectedErr:             nil,
			expectedGetNodeInfoCall: true,
			expectedGetNodeIdCall:   false,
		},
		"unimplemented NodeGetInfo": {
			nodeGetInfoErr:          errUnimplemented,
			nodeGetIdErr:            nil,
			expectedErr:             nil,
			expectedGetNodeInfoCall: true,
			expectedGetNodeIdCall:   true,
		},
		"other error for NodeGetInfo": {
			nodeGetInfoErr:          errRandom,
			nodeGetIdErr:            nil,
			expectedErr:             errRandom,
			expectedGetNodeInfoCall: true,
			expectedGetNodeIdCall:   false,
		},
		"unimplemented NodeGetInfo and error on NodeGetId": {
			nodeGetInfoErr:          errUnimplemented,
			nodeGetIdErr:            errRandom,
			expectedErr:             errRandom,
			expectedGetNodeInfoCall: true,
			expectedGetNodeIdCall:   true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			ig := &nodeInfoGetter{
				nodeGetInfoErr: tc.nodeGetInfoErr,
				nodeGetIdErr:   tc.nodeGetIdErr,
			}
			_, _, _, err := getNodeInfoWithFallback(ctx, ig)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got: %v", tc.expectedErr, err)
			}
			checkCall(t, ig.nodeGetInfoCallCount, tc.expectedGetNodeInfoCall, "NodeGetInfo")
			checkCall(t, ig.nodeGetIdCallCount, tc.expectedGetNodeIdCall, "NodeGetId")
		})
	}
}

func checkCall(t *testing.T, callCount int, expectedCall bool, funcName string) {
	t.Helper()

	if expectedCall {
		if callCount < 1 {
			t.Errorf("Expected %s to be called, was not called", funcName)
		}
		return
	}

	if callCount > 0 {
		t.Errorf("Expected %s not to be calle, was called %d times", funcName, callCount)
	}
}

type nodeInfoGetter struct {
	nodeGetIdErr   error
	nodeGetInfoErr error

	nodeGetIdCallCount   int
	nodeGetInfoCallCount int
}

func (g *nodeInfoGetter) NodeGetId(ctx context.Context) (string, error) {
	g.nodeGetIdCallCount += 1
	if err := g.nodeGetIdErr; err != nil {
		return "", err
	}
	return "some node id", nil
}

func (g *nodeInfoGetter) NodeGetInfo(ctx context.Context) (string, int64, *csipb.Topology, error) {
	g.nodeGetInfoCallCount += 1
	if err := g.nodeGetInfoErr; err != nil {
		return "", 0, nil, err
	}
	return "some node id", 0, nil, nil
}
