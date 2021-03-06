// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package gaming_test

import (
	"context"

	gaming "cloud.google.com/go/gaming/apiv1"
	"google.golang.org/api/iterator"
	gamingpb "google.golang.org/genproto/googleapis/cloud/gaming/v1"
)

func ExampleNewGameServerDeploymentsClient() {
	ctx := context.Background()
	c, err := gaming.NewGameServerDeploymentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleGameServerDeploymentsClient_ListGameServerDeployments() {
	// import gamingpb "google.golang.org/genproto/googleapis/cloud/gaming/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := gaming.NewGameServerDeploymentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &gamingpb.ListGameServerDeploymentsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListGameServerDeployments(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleGameServerDeploymentsClient_GetGameServerDeployment() {
	// import gamingpb "google.golang.org/genproto/googleapis/cloud/gaming/v1"

	ctx := context.Background()
	c, err := gaming.NewGameServerDeploymentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &gamingpb.GetGameServerDeploymentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetGameServerDeployment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleGameServerDeploymentsClient_CreateGameServerDeployment() {
	// import gamingpb "google.golang.org/genproto/googleapis/cloud/gaming/v1"

	ctx := context.Background()
	c, err := gaming.NewGameServerDeploymentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &gamingpb.CreateGameServerDeploymentRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateGameServerDeployment(ctx, req)
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

func ExampleGameServerDeploymentsClient_DeleteGameServerDeployment() {
	// import gamingpb "google.golang.org/genproto/googleapis/cloud/gaming/v1"

	ctx := context.Background()
	c, err := gaming.NewGameServerDeploymentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &gamingpb.DeleteGameServerDeploymentRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteGameServerDeployment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleGameServerDeploymentsClient_UpdateGameServerDeployment() {
	// import gamingpb "google.golang.org/genproto/googleapis/cloud/gaming/v1"

	ctx := context.Background()
	c, err := gaming.NewGameServerDeploymentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &gamingpb.UpdateGameServerDeploymentRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpdateGameServerDeployment(ctx, req)
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

func ExampleGameServerDeploymentsClient_GetGameServerDeploymentRollout() {
	// import gamingpb "google.golang.org/genproto/googleapis/cloud/gaming/v1"

	ctx := context.Background()
	c, err := gaming.NewGameServerDeploymentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &gamingpb.GetGameServerDeploymentRolloutRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetGameServerDeploymentRollout(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleGameServerDeploymentsClient_UpdateGameServerDeploymentRollout() {
	// import gamingpb "google.golang.org/genproto/googleapis/cloud/gaming/v1"

	ctx := context.Background()
	c, err := gaming.NewGameServerDeploymentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &gamingpb.UpdateGameServerDeploymentRolloutRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpdateGameServerDeploymentRollout(ctx, req)
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

func ExampleGameServerDeploymentsClient_PreviewGameServerDeploymentRollout() {
	// import gamingpb "google.golang.org/genproto/googleapis/cloud/gaming/v1"

	ctx := context.Background()
	c, err := gaming.NewGameServerDeploymentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &gamingpb.PreviewGameServerDeploymentRolloutRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.PreviewGameServerDeploymentRollout(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleGameServerDeploymentsClient_FetchDeploymentState() {
	// import gamingpb "google.golang.org/genproto/googleapis/cloud/gaming/v1"

	ctx := context.Background()
	c, err := gaming.NewGameServerDeploymentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &gamingpb.FetchDeploymentStateRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.FetchDeploymentState(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
