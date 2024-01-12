// Copyright 2024 Gravitational, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package scim

import (
	"context"

	"github.com/gravitational/trace"
	"google.golang.org/grpc"

	scimpb "github.com/gravitational/teleport/api/gen/proto/go/teleport/scim/v1"
)

// Client wraps the underlying GRPC client with some more human-friendly tooling
type Client struct {
	grpcClient scimpb.ScimServiceClient
}

func NewClientFromConn(cc grpc.ClientConnInterface) *Client {
	return NewClient(scimpb.NewScimServiceClient(cc))
}

func NewClient(grpcClient scimpb.ScimServiceClient) *Client {
	return &Client{grpcClient: grpcClient}
}

func (c *Client) ListSCIMResources(ctx context.Context, req *scimpb.ListSCIMResourcesRequest) (*scimpb.ResourceList, error) {
	resp, err := c.grpcClient.ListSCIMResources(ctx, req)
	if err != nil {
		return nil, trace.Wrap(err, "handling SCIM list request")
	}
	return resp, nil
}

func (c *Client) GetSCIMResource(ctx context.Context, req *scimpb.GetSCIMResourceRequest) (*scimpb.Resource, error) {
	resp, err := c.grpcClient.GetSCIMResource(ctx, req)
	if err != nil {
		return nil, trace.Wrap(err, "handling SCIM get request")
	}
	return resp, nil
}

func (c *Client) CreateSCIMResource(ctx context.Context, req *scimpb.CreateSCIMResourceRequest) (*scimpb.Resource, error) {
	resp, err := c.grpcClient.CreateSCIMResource(ctx, req)
	if err != nil {
		return nil, trace.Wrap(err, "handling SCIM POST request")
	}
	return resp, nil
}

func (c *Client) UpdateSCIMResource(ctx context.Context, req *scimpb.UpdateSCIMResourceRequest) (*scimpb.Resource, error) {
	res, err := c.grpcClient.UpdateSCIMResource(ctx, req)
	if err != nil {
		return nil, trace.Wrap(err, "handling SCIM update request")
	}
	return res, nil
}
