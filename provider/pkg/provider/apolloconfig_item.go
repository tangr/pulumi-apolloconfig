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

	AppId                    string `json:"appId"`
	Namespace                string `json:"namespace"`
	Env                      string `json:"env"`
	ClusterName              string `json:"clusterName"`
	Key                      string `json:"key"`
	Value                    string `json:"value"`
	Comment                  string `json:"comment"`
	Operator                 string `json:"operator"`
	DataChangeCreatedBy      string `json:"dataChangeCreatedBy"`
	DataChangeLastModifiedBy string `json:"dataChangeLastModifiedBy"`
}

func GenerateApolloItemProperties(input ApolloConfigItemInput, apolloItem apolloconfigapi.ApollItem) (outputs *structpb.Struct, inputs *structpb.Struct, err error) {
	inputMap := resource.PropertyMap{}

	inputMap["env"] = resource.NewPropertyValue(input.Env)
	inputMap["appId"] = resource.NewPropertyValue(input.AppId)
	inputMap["clusterName"] = resource.NewPropertyValue(input.ClusterName)
	inputMap["namespace"] = resource.NewPropertyValue(input.Namespace)
	inputMap["key"] = resource.NewPropertyValue(input.Key)
	inputMap["value"] = resource.NewPropertyValue(input.Value)
	inputMap["dataChangeCreatedBy"] = resource.NewPropertyValue(input.DataChangeCreatedBy)
	inputMap["comment"] = resource.NewPropertyValue(input.Comment)
	inputMap["dataChangeLastModifiedBy"] = resource.NewPropertyValue(input.DataChangeLastModifiedBy)
	if input.Operator != "" {
		inputMap["operator"] = resource.NewPropertyValue(input.Operator)
	}

	outputMap := resource.PropertyMap{}
	outputMap["env"] = resource.NewPropertyValue(input.Env)
	outputMap["appId"] = resource.NewPropertyValue(input.AppId)
	outputMap["clusterName"] = resource.NewPropertyValue(input.ClusterName)
	outputMap["namespace"] = resource.NewPropertyValue(input.Namespace)
	outputMap["key"] = resource.NewPropertyValue(input.Key)
	outputMap["value"] = resource.NewPropertyValue(input.Value)
	outputMap["comment"] = resource.NewPropertyValue(input.Comment)
	outputMap["operator"] = resource.NewPropertyValue(input.Operator)
	outputMap["dataChangeCreatedBy"] = resource.NewPropertyValue(input.DataChangeCreatedBy)

	outputMap["dataChangeLastModifiedBy"] = resource.NewPropertyValue(apolloItem.DataChangeLastModifiedBy)
	outputMap["dataChangeCreatedTime"] = resource.NewPropertyValue(apolloItem.DataChangeCreatedTime)
	outputMap["dataChangeLastModifiedTime"] = resource.NewPropertyValue(apolloItem.DataChangeLastModifiedTime)

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

	getStringValue := func(key string) string {
		if v, ok := inputMap[resource.PropertyKey(key)]; ok && v.HasValue() && v.IsString() {
			return v.StringValue()
		}
		return ""
	}

	input.AppId = getStringValue("appId")
	input.Namespace = getStringValue("namespace")
	input.Env = getStringValue("env")
	input.ClusterName = getStringValue("clusterName")
	input.Key = getStringValue("key")
	input.Value = getStringValue("value")
	input.Comment = getStringValue("comment")
	input.Operator = getStringValue("operator")

	input.DataChangeCreatedBy = getStringValue("dataChangeCreatedBy")
	input.DataChangeLastModifiedBy = getStringValue("dataChangeLastModifiedBy")

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
		"env": true,
		"appId": true,
		"clusterName": true,
		"namespace": true,
		// "key": true,
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

	item := aci.ToApolloConfigItemInput(inputs)

	operator := item.Operator
	if operator == "" {
		operator = item.DataChangeCreatedBy
	}

	err = aci.deleteApolloItem(ctx, req.Id, operator)

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
	itemId := fmt.Sprintf("%s/%s/%s/%s/%s", input.Env, input.AppId, input.ClusterName, input.Namespace, input.Key)
	apolloItem, err := aci.createApolloItem(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("error creating apollo item: '%s': %s", itemId, err.Error())
	}

	outputProperties, _, err := GenerateApolloItemProperties(input, *apolloItem)
	if err != nil {
		return nil, err
	}

	return &pulumirpc.CreateResponse{
		Id:         itemId,
		Properties: outputProperties,
	}, nil

}

func (aci *ApolloConfigItemResource) Check(req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	return &pulumirpc.CheckResponse{Inputs: req.News, Failures: nil}, nil
}

func (aci *ApolloConfigItemResource) Update(req *pulumirpc.UpdateRequest) (*pulumirpc.UpdateResponse, error) {
	ctx := context.Background()

	apolloItemId := req.GetId()

	olds, err := plugin.UnmarshalProperties(req.GetOldInputs(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	changedInputs := olds

	changedInputs["appId"] = news["appId"]
	changedInputs["namespace"] = news["namespace"]
	changedInputs["env"] = news["env"]
	changedInputs["clusterName"] = news["clusterName"]
	changedInputs["dataChangeLastModifiedBy"] = news["dataChangeLastModifiedBy"]
	changedInputs["env"] = news["env"]
	changedInputs["value"] = news["value"]
	changedInputs["comment"] = news["comment"]
	changedInputs["operator"] = news["operator"]
	changedInputs["dataChangeCreatedBy"] = news["dataChangeCreatedBy"]

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
	itemId := req.GetId()

	env, appId, clusterName, namespace, key, err := splitApolloItemId(itemId)
	if err != nil {
		return nil, err
	}

	// the item id is immutable; if we get nil it got deleted, otherwise all data is the same
	apolloItem, err := aci.client.GetApolloItem(ctx, env, appId, clusterName, namespace, key)
	if err != nil {
		return nil, err
	}
	if apolloItem == nil {
		return &pulumirpc.ReadResponse{}, nil
	}

	input := ApolloConfigItemInput{
		AppId:                    appId,
		Namespace:                namespace,
		Env:                      env,
		ClusterName:              clusterName,
		Key:                      key,
		Value:                    apolloItem.Value,
		Comment:                  apolloItem.Comment,
		// Operator:                 apolloItem.
		DataChangeCreatedBy:      apolloItem.DataChangeCreatedBy,
		DataChangeLastModifiedBy: apolloItem.DataChangeLastModifiedBy,
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

func (aci *ApolloConfigItemResource) Configure(_ ApollConfig) {
}

func (aci *ApolloConfigItemResource) createApolloItem(ctx context.Context, input ApolloConfigItemInput) (*apolloconfigapi.ApollItem, error) {
	params := &apolloconfigapi.CreateUpdateApollItemRequest{
		AppID:                    input.AppId,
		Namespace:                input.Namespace,
		Env:                      input.Env,
		ClusterName:              input.ClusterName,
		Key:                      input.Key,
		Value:                    input.Value,
		Comment:                  input.Comment,
		Operator:                 input.Operator,
		DataChangeCreatedBy:      input.DataChangeCreatedBy,
		DataChangeLastModifiedBy: input.DataChangeLastModifiedBy,
	}

	apolloItem, err := aci.client.CreateApolloItem(ctx, params)
	if err != nil {
		return nil, err
	}

	return apolloItem, nil
}

func (aci *ApolloConfigItemResource) updateApolloItem(ctx context.Context, apolloItemId string, input ApolloConfigItemInput) error {
	env, appId, clusterName, namespace, key, err := splitApolloItemId(apolloItemId)
	if err != nil {
		return err
	}

	params := &apolloconfigapi.CreateUpdateApollItemRequest{
		AppID:                    appId,
		Namespace:                namespace,
		Env:                      env,
		ClusterName:              clusterName,
		Key:                      key,
		Value:                    input.Value,
		Comment:                  input.Comment,
		Operator:                 input.Operator,
		DataChangeCreatedBy:      input.DataChangeCreatedBy,
		DataChangeLastModifiedBy: input.DataChangeLastModifiedBy,
	}
	return aci.client.UpdateApolloItem(ctx, apolloItemId, params)
}

func (aci *ApolloConfigItemResource) deleteApolloItem(ctx context.Context, id string, operator string) error {
	// we don't need the token name when we delete
	env, appId, clusterName, namespace, key, err := splitApolloItemId(id)
	if err != nil {
		return err
	}
	return aci.client.DeleteApolloItem(ctx, env, appId, clusterName, namespace, key, operator)

}

func splitApolloItemId(id string) (string, string, string, string, string, error) {
	s := strings.Split(id, "/")
	if len(s) != 5 {
		return "", "", "", "", "", fmt.Errorf("%q is invalid, must be in the format: env/appId/clusterName/namespace/key", id)
	}
	return s[0], s[1], s[2], s[3], s[4], nil
}
