package provider

import (
	"context"
	"fmt"
	"strings"

	pbempty "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/tangr/pulumi-apolloconfig/provider/pkg/internal/apolloconfigapi"
)

type ApolloConfigItemResource struct {
	client apolloconfigapi.ApolloItemClient
}

type ApolloConfigItemInput struct {
	OrgName      string
	Description  string
	Name         string
	ForceDestroy bool
}

func GenerateApolloItemProperties(input ApolloConfigItemInput, apolloItem apolloconfigapi.ApollItem) (outputs *structpb.Struct, inputs *structpb.Struct, err error) {
	inputMap := resource.PropertyMap{}
	inputMap["name"] = resource.NewPropertyValue(input.Name)
	inputMap["organizationName"] = resource.NewPropertyValue(input.OrgName)
	if input.Description != "" {
		inputMap["description"] = resource.NewPropertyValue(input.Description)
	}
	if input.ForceDestroy {
		inputMap["forceDestroy"] = resource.NewPropertyValue(input.ForceDestroy)
	}

	outputMap := resource.PropertyMap{}
	outputMap["apolloItemId"] = resource.NewPropertyValue(apolloItem.ID)
	outputMap["name"] = inputMap["name"]
	outputMap["organizationName"] = inputMap["organizationName"]
	outputMap["tokenValue"] = resource.NewPropertyValue(apolloItem.TokenValue)
	if input.Description != "" {
		outputMap["description"] = inputMap["description"]
	}
	if input.ForceDestroy {
		outputMap["forceDestroy"] = inputMap["forceDestroy"]
	}

	inputs, err = plugin.MarshalProperties(inputMap, plugin.MarshalOptions{})
	if err != nil {
		return nil, nil, err
	}

	outputs, err = plugin.MarshalProperties(outputMap, plugin.MarshalOptions{})
	if err != nil {
		return nil, nil, err
	}

	return outputs, inputs, err
}

func (aci *ApolloConfigItemResource) ToApolloConfigItemInput(inputMap resource.PropertyMap) ApolloConfigItemInput {
	input := ApolloConfigItemInput{}

	if inputMap["name"].HasValue() && inputMap["name"].IsString() {
		input.Name = inputMap["name"].StringValue()
	}

	if inputMap["description"].HasValue() && inputMap["description"].IsString() {
		input.Description = inputMap["description"].StringValue()
	}

	if inputMap["organizationName"].HasValue() && inputMap["organizationName"].IsString() {
		input.OrgName = inputMap["organizationName"].StringValue()
	}

	if inputMap["forceDestroy"].HasValue() && inputMap["forceDestroy"].IsBool() {
		input.ForceDestroy = inputMap["forceDestroy"].BoolValue()
	}

	return input
}

func (aci *ApolloConfigItemResource) Name() string {
	return "apolloconfig:index:ApolloItem"
}

func (aci *ApolloConfigItemResource) Diff(req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	olds, err := plugin.UnmarshalProperties(req.GetOldInputs(), plugin.MarshalOptions{KeepUnknowns: false, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	diffs := olds.Diff(news)
	if diffs == nil {
		return &pulumirpc.DiffResponse{
			Changes: pulumirpc.DiffResponse_DIFF_NONE,
		}, nil
	}

	dd := plugin.NewDetailedDiffFromObjectDiff(diffs, false)

	detailedDiffs := map[string]*pulumirpc.PropertyDiff{}
	replaceProperties := map[string]bool{
		"organizationName": true,
	}
	for k, v := range dd {
		if _, ok := replaceProperties[k]; ok {
			v.Kind = v.Kind.AsReplace()
		}
		detailedDiffs[k] = &pulumirpc.PropertyDiff{
			Kind:      pulumirpc.PropertyDiff_Kind(v.Kind),
			InputDiff: v.InputDiff,
		}
	}

	changes := pulumirpc.DiffResponse_DIFF_NONE
	if len(detailedDiffs) > 0 {
		changes = pulumirpc.DiffResponse_DIFF_SOME
	}
	return &pulumirpc.DiffResponse{
		Changes:         changes,
		DetailedDiff:    detailedDiffs,
		HasDetailedDiff: true,
	}, nil
}

func (aci *ApolloConfigItemResource) Delete(req *pulumirpc.DeleteRequest) (*pbempty.Empty, error) {
	ctx := context.Background()
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	pool := aci.ToApolloConfigItemInput(inputs)
	if err != nil {
		return nil, err
	}

	err = aci.deleteApolloItem(ctx, req.Id, pool.ForceDestroy)

	if err != nil {
		return &pbempty.Empty{}, err
	}

	return &pbempty.Empty{}, nil
}

func (aci *ApolloConfigItemResource) Create(req *pulumirpc.CreateRequest) (*pulumirpc.CreateResponse, error) {
	ctx := context.Background()
	inputMap, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	input := aci.ToApolloConfigItemInput(inputMap)
	apolloItem, err := aci.createApolloItem(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("error creating apollo item '%s': %s", input.Name, err.Error())
	}

	outputProperties, _, err := GenerateApolloItemProperties(input, *apolloItem)
	if err != nil {
		return nil, err
	}

	return &pulumirpc.CreateResponse{
		Id:         input.OrgName + "/" + input.Name + "/" + apolloItem.ID,
		Properties: outputProperties,
	}, nil

}

func (aci *ApolloConfigItemResource) Check(req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	return &pulumirpc.CheckResponse{Inputs: req.News, Failures: nil}, nil
}

func (aci *ApolloConfigItemResource) Update(req *pulumirpc.UpdateRequest) (*pulumirpc.UpdateResponse, error) {
	ctx := context.Background()

	// ignore orgName because if that changed, we would have done a replace, so update would never have been called
	_, _, apolloItemId, err := splitApolloItemId(req.GetId())
	if err != nil {
		return nil, fmt.Errorf("invalid resource id: %v", err)
	}

	olds, err := plugin.UnmarshalProperties(req.GetOldInputs(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	changedInputs := olds
	changedInputs["name"] = news["name"]
	changedInputs["description"] = news["description"]
	changedInputs["forceDestroy"] = news["forceDestroy"]

	inputsApolloItem := aci.ToApolloConfigItemInput(changedInputs)
	err = aci.updateApolloItem(ctx, apolloItemId, inputsApolloItem)
	if err != nil {
		return nil, fmt.Errorf("error updating apollo item '%s': %s", inputsApolloItem.Name, err.Error())
	}

	outputProperties, err := plugin.MarshalProperties(
		changedInputs,
		plugin.MarshalOptions{},
	)
	if err != nil {
		return nil, err
	}

	return &pulumirpc.UpdateResponse{
		Properties: outputProperties,
	}, nil
}

func (aci *ApolloConfigItemResource) Read(req *pulumirpc.ReadRequest) (*pulumirpc.ReadResponse, error) {
	ctx := context.Background()
	urn := req.GetId()

	orgName, _, apolloItemId, err := splitApolloItemId(urn)
	if err != nil {
		return nil, err
	}

	// the item id is immutable; if we get nil it got deleted, otherwise all data is the same
	apolloItem, err := aci.client.GetApolloItem(ctx, apolloItemId, orgName)
	if err != nil {
		return nil, err
	}
	if apolloItem == nil {
		return &pulumirpc.ReadResponse{}, nil
	}

	propertyMap, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{KeepSecrets: true})
	if err != nil {
		return nil, err
	}
	if propertyMap["tokenValue"].HasValue() {
		apolloItem.TokenValue = getSecretOrStringValue(propertyMap["tokenValue"])
	}

	input := ApolloConfigItemInput{
		OrgName:     orgName,
		Description: apolloItem.Description,
		Name:        apolloItem.Name,
	}
	outputProperties, inputs, err := GenerateApolloItemProperties(input, *apolloItem)
	if err != nil {
		return nil, err
	}

	return &pulumirpc.ReadResponse{
		Id:         req.GetId(),
		Properties: outputProperties,
		Inputs:     inputs,
	}, nil
}

func (aci *ApolloConfigItemResource) Invoke(_ *apolloconfigProvider, req *pulumirpc.InvokeRequest) (*pulumirpc.InvokeResponse, error) {
	return &pulumirpc.InvokeResponse{Return: nil}, fmt.Errorf("unknown function '%s'", req.Tok)
}

func (aci *ApolloConfigItemResource) Configure(_ PulumiServiceConfig) {
}

func (aci *ApolloConfigItemResource) createApolloItem(ctx context.Context, input ApolloConfigItemInput) (*apolloconfigapi.ApollItem, error) {
	apolloItem, err := aci.client.CreateApolloItem(ctx, input.OrgName, input.Name, input.Description)
	if err != nil {
		return nil, err
	}

	return apolloItem, nil
}

func (aci *ApolloConfigItemResource) updateApolloItem(ctx context.Context, apolloItemId string, input ApolloConfigItemInput) error {
	return aci.client.UpdateApolloItem(ctx, apolloItemId, input.OrgName, input.Name, input.Description)
}

func (aci *ApolloConfigItemResource) deleteApolloItem(ctx context.Context, id string, forceDestroy bool) error {
	// we don't need the token name when we delete
	orgName, _, apolloItemId, err := splitApolloItemId(id)
	if err != nil {
		return err
	}
	return aci.client.DeleteApolloItem(ctx, apolloItemId, orgName, forceDestroy)

}

func splitApolloItemId(id string) (string, string, string, error) {
	// format: organization/name/apolloItemId
	s := strings.Split(id, "/")
	if len(s) != 3 {
		return "", "", "", fmt.Errorf("%q is invalid, must be in the format: organization/name/apolloItemId", id)
	}
	return s[0], s[1], s[2], nil
}
