// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tests

import (
	"github.com/onosproject/helmit/pkg/helm"
	"github.com/onosproject/helmit/pkg/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ONOSRICSuite is the onos-ric chart test suite
type ONOSRICSuite struct {
	test.Suite
}

// TestInstall tests installing the onos-ric chart
func (s *ONOSRICSuite) TestInstall(t *testing.T) {
	atomix := helm.Chart("kubernetes-controller", "https://charts.atomix.io").
		Release("onos-ric-atomix").
		Set("scope", "Namespace")
	assert.NoError(t, atomix.Install(true))

	raft := helm.Chart("raft-storage-controller", "https://charts.atomix.io").
		Release("onos-ric-raft").
		Set("scope", "Namespace")
	assert.NoError(t, raft.Install(true))

	cache := helm.Chart("cache-storage-controller", "https://charts.atomix.io").
		Release("onos-ric-cache").
		Set("scope", "Namespace")
	assert.NoError(t, cache.Install(true))

	topo := helm.Chart("onos-topo").
		Release("onos-topo").
		Set("global.store.controller", "onos-ric-atomix-kubernetes-controller:5679")
	assert.NoError(t, topo.Install(false))

	ric := helm.Chart("onos-ric").
		Release("onos-ric").
		Set("global.store.controller", "onos-ric-atomix-kubernetes-controller:5679")
	assert.NoError(t, ric.Install(true))
}
