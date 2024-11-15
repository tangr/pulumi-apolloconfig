package provider

import (
	"fmt"

	pbempty "google.golang.org/protobuf/types/known/emptypb"

	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

type ApollConfigUnknownResource struct{}
type ApollConfigUnknownFunction struct{}

func (u *ApollConfigUnknownResource) Name() string {
	return "apolloconfig:index:Unknown"
}

func (u *ApollConfigUnknownResource) Configure(config ApollConfig) {
}

func (u *ApollConfigUnknownResource) Diff(req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	return nil, createUnknownResourceErrorFromRequest(req)
}

func (u *ApollConfigUnknownResource) Delete(req *pulumirpc.DeleteRequest) (*pbempty.Empty, error) {
	return nil, createUnknownResourceErrorFromRequest(req)
}

func (u *ApollConfigUnknownResource) Create(req *pulumirpc.CreateRequest) (*pulumirpc.CreateResponse, error) {
	return nil, createUnknownResourceErrorFromRequest(req)
}

func (u *ApollConfigUnknownResource) Check(req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	return nil, createUnknownResourceErrorFromRequest(req)
}

func (u *ApollConfigUnknownResource) Update(req *pulumirpc.UpdateRequest) (*pulumirpc.UpdateResponse, error) {
	return nil, createUnknownResourceErrorFromRequest(req)
}

func (u *ApollConfigUnknownResource) Read(req *pulumirpc.ReadRequest) (*pulumirpc.ReadResponse, error) {
	return nil, createUnknownResourceErrorFromRequest(req)
}

func createUnknownResourceErrorFromRequest(req ResourceBase) error {
	rn := getResourceNameFromRequest(req)
	return fmt.Errorf("unknown resource type '%s'", rn)
}

func (u *ApollConfigUnknownResource) Invoke(s *apolloconfigProvider, req *pulumirpc.InvokeRequest) (*pulumirpc.InvokeResponse, error) {
	return &pulumirpc.InvokeResponse{Return: nil}, fmt.Errorf("unknown function '%s'", req.Tok)
}

func (f *ApollConfigUnknownFunction) Name() string {
	return "apolloconfig:index:Unknown"
}

func (f *ApollConfigUnknownFunction) Configure(config ApollConfig) {
}
