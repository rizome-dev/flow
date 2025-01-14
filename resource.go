/*
Copyright 2025 rizome labs llc, hi@rizome.dev

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
package flow

import (
	"context"

	alloydb "cloud.google.com/go/alloydb/apiv1"
	"cloud.google.com/go/alloydb/apiv1/alloydbpb"
)

// TODO
type (
	Resource struct{}
)

func CreateAlloyDBVectorStore(parent, instanceId string, instance *alloydbpb.Instance) {
	ctx := context.Background()
	client, err := alloydb.NewAlloyDBAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer client.Close()
	req := &alloydbpb.BatchCreateInstancesRequest{
		Parent: parent,
		Requests: &alloydbpb.CreateInstanceRequests{
			CreateInstanceRequests: []*alloydbpb.CreateInstanceRequest{{
				Parent:     parent,
				InstanceId: instanceId,
				Instance:   instance,
			}},
		},
	}
	op, err := client.BatchCreateInstances(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
