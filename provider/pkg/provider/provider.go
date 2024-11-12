// Copyright 2016-2020, Pulumi Corporation.
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

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pbempty "google.golang.org/protobuf/types/known/emptypb"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/tangr/pulumi-apolloconfig/provider/pkg/internal/pulumiapi"
)

type ApolloconfigResource interface {
	Configure(config PulumiServiceConfig)
	Diff(req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error)
	Create(req *pulumirpc.CreateRequest) (*pulumirpc.CreateResponse, error)
	Delete(req *pulumirpc.DeleteRequest) (*pbempty.Empty, error)
	Check(req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error)
	Update(req *pulumirpc.UpdateRequest) (*pulumirpc.UpdateResponse, error)
	Read(req *pulumirpc.ReadRequest) (*pulumirpc.ReadResponse, error)
	Name() string
}

type apolloconfigProvider struct {
	pulumirpc.UnimplementedResourceProviderServer

	host            *provider.HostClient
	name            string
	version         string
	schema          string
	pulumiResources []ApolloconfigResource
	AccessToken     string
}

func makeProvider(host *provider.HostClient, name, version, schema string) (pulumirpc.ResourceProviderServer, error) {
	// inject version into schema
	versionedSchema := mustSetSchemaVersion(schema, version)
	// Return the new provider
	return &apolloconfigProvider{
		host:        host,
		name:        name,
		schema:      versionedSchema,
		version:     version,
		AccessToken: "",
	}, nil
}

// Call dynamically executes a method in the provider associated with a component resource.
func (k *apolloconfigProvider) Call(ctx context.Context, req *pulumirpc.CallRequest) (*pulumirpc.CallResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Call is not yet implemented")
}

// Attach implements pulumirpc.ResourceProviderServer
func (k *apolloconfigProvider) Attach(_ context.Context, req *pulumirpc.PluginAttach) (*pbempty.Empty, error) {
	host, err := provider.NewHostClient(req.Address)
	if err != nil {
		return nil, err
	}
	k.host = host
	return &pbempty.Empty{}, nil
}

// Construct creates a new component resource.
func (k *apolloconfigProvider) Construct(ctx context.Context, req *pulumirpc.ConstructRequest) (*pulumirpc.ConstructResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Construct is not yet implemented")
}

// CheckConfig validates the configuration for this provider.
func (k *apolloconfigProvider) CheckConfig(ctx context.Context, req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	return &pulumirpc.CheckResponse{Inputs: req.GetNews()}, nil
}

// DiffConfig diffs the configuration for this provider.
func (k *apolloconfigProvider) DiffConfig(ctx context.Context, req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	return &pulumirpc.DiffResponse{}, nil
}

// Configure configures the resource provider with "globals" that control its behavior.
func (k *apolloconfigProvider) Configure(_ context.Context, req *pulumirpc.ConfigureRequest) (*pulumirpc.ConfigureResponse, error) {

	sc := PulumiServiceConfig{}
	sc.Config = make(map[string]string)
	for key, val := range req.GetVariables() {
		sc.Config[strings.TrimPrefix(key, "apolloconfig:config:")] = val
	}

	httpClient := http.Client{
		Timeout: 60 * time.Second,
	}
	token, err := sc.getPulumiAccessToken()
	if err != nil {
		return nil, err
	}
	url, err := sc.getPulumiServiceUrl()
	if err != nil {
		return nil, err
	}
	client, err := pulumiapi.NewClient(&httpClient, *token, *url)


	if err != nil {
		return nil, err
	}

	k.pulumiResources = []ApolloconfigResource{
		&PulumiServiceTeamResource{
			client: client,
		},
		&PulumiServiceAccessTokenResource{
			client: client,
		},
		&PulumiServiceWebhookResource{
			client: client,
		},
		&PulumiServiceStackTagResource{
			client: client,
		},
		&PulumiServiceAgentPoolResource{
			client: client,
		},
	}

	for _, sr := range k.pulumiResources {
		sr.Configure(sc)
	}

	return &pulumirpc.ConfigureResponse{
		AcceptSecrets: true,
	}, nil
}

// Invoke dynamically executes a built-in function in the provider.
func (k *apolloconfigProvider) Invoke(_ context.Context, req *pulumirpc.InvokeRequest) (*pulumirpc.InvokeResponse, error) {
	tok := req.GetTok()
	return nil, fmt.Errorf("unknown Invoke token '%s'", tok)
}

// StreamInvoke dynamically executes a built-in function in the provider. The result is streamed
// back as a series of messages.
func (k *apolloconfigProvider) StreamInvoke(req *pulumirpc.InvokeRequest, server pulumirpc.ResourceProvider_StreamInvokeServer) error {
	tok := req.GetTok()
	return fmt.Errorf("unknown StreamInvoke token '%s'", tok)
}

// Check validates that the given property bag is valid for a resource of the given type and returns
// the inputs that should be passed to successive calls to Diff, Create, or Update for this
// resource. As a rule, the provider inputs returned by a call to Check should preserve the original
// representation of the properties as present in the program inputs. Though this rule is not
// required for correctness, violations thereof can negatively impact the end-user experience, as
// the provider inputs are using for detecting and rendering diffs.
func (k *apolloconfigProvider) Check(ctx context.Context, req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	rn := getResourceNameFromRequest(req)
	res := k.getApolloconfigResource(rn)
	return res.Check(req)
}

// Diff checks what impacts a hypothetical update will have on the resource's properties.
func (k *apolloconfigProvider) Diff(ctx context.Context, req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	rn := getResourceNameFromRequest(req)
	res := k.getApolloconfigResource(rn)
	return res.Diff(req)
}

// Create allocates a new instance of the provided resource and returns its unique ID afterwards.
func (k *apolloconfigProvider) Create(ctx context.Context, req *pulumirpc.CreateRequest) (*pulumirpc.CreateResponse, error) {
	rn := getResourceNameFromRequest(req)
	res := k.getApolloconfigResource(rn)
	return res.Create(req)
}

// Read the current live state associated with a resource.
func (k *apolloconfigProvider) Read(ctx context.Context, req *pulumirpc.ReadRequest) (*pulumirpc.ReadResponse, error) {
	rn := getResourceNameFromRequest(req)
	res := k.getApolloconfigResource(rn)
	return res.Read(req)
}

// Update updates an existing resource with new values.
func (k *apolloconfigProvider) Update(ctx context.Context, req *pulumirpc.UpdateRequest) (*pulumirpc.UpdateResponse, error) {
	rn := getResourceNameFromRequest(req)
	res := k.getApolloconfigResource(rn)
	return res.Update(req)
}

// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed
// to still exist.
func (k *apolloconfigProvider) Delete(ctx context.Context, req *pulumirpc.DeleteRequest) (*pbempty.Empty, error) {
	rn := getResourceNameFromRequest(req)
	res := k.getApolloconfigResource(rn)
	return res.Delete(req)
}

// GetPluginInfo returns generic information about this plugin, like its version.
func (k *apolloconfigProvider) GetPluginInfo(context.Context, *pbempty.Empty) (*pulumirpc.PluginInfo, error) {
	return &pulumirpc.PluginInfo{
		Version: k.version,
	}, nil
}

// GetSchema returns the JSON-serialized schema for the provider.
func (k *apolloconfigProvider) GetSchema(ctx context.Context, req *pulumirpc.GetSchemaRequest) (*pulumirpc.GetSchemaResponse, error) {
	return &pulumirpc.GetSchemaResponse{
		Schema: k.schema,
	}, nil
}

// Cancel signals the provider to gracefully shut down and abort any ongoing resource operations.
// Operations aborted in this way will return an error (e.g., `Update` and `Create` will either a
// creation error or an initialization error). Since Cancel is advisory and non-blocking, it is up
// to the host to decide how long to wait after Cancel is called before (e.g.)
// hard-closing any gRPC connection.
func (k *apolloconfigProvider) Cancel(context.Context, *pbempty.Empty) (*pbempty.Empty, error) {
	// TODO
	return &pbempty.Empty{}, nil
}

func (k *apolloconfigProvider) getApolloconfigResource(name string) ApolloconfigResource {
	for _, r := range k.pulumiResources {
		if r.Name() == name {
			return r
		}
	}

	return &PulumiServiceUnknownResource{}
}

func getResourceNameFromRequest(req ResourceBase) string {
	urn := resource.URN(req.GetUrn())
	return urn.Type().String()
}

// mustSetSchemaVersion deserializes schemaStr from json, sets Version field
// then serializes back to json string
func mustSetSchemaVersion(schemaStr string, version string) string {
	var spec schema.PackageSpec
	if err := json.Unmarshal([]byte(schemaStr), &spec); err != nil {
		panic(fmt.Errorf("failed to parse schema: %w", err))
	}
	spec.Version = version
	bytes, err := json.Marshal(spec)
	if err != nil {
		panic(fmt.Errorf("failed to serialize versioned schema: %w", err))
	}
	return string(bytes)
}

type ResourceBase interface {
	GetUrn() string
}
