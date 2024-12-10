/*
Copyright 2024 The Kubermatic Kubernetes Platform contributors.

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

package sync

import (
	"context"
	"testing"

	"github.com/kcp-dev/logicalcluster/v3"

	"sigs.k8s.io/controller-runtime/pkg/kontext"
)

func TestNewContext(t *testing.T) {
	clusterName := logicalcluster.Name("foo")
	ctx := kontext.WithCluster(context.Background(), clusterName)

	combinedCtx := NewContext(context.Background(), ctx)

	if combinedCtx.clusterName != clusterName.String() {
		t.Fatalf("Expected function to recognize the cluster name in the context, but got %q", combinedCtx.clusterName)
	}
}