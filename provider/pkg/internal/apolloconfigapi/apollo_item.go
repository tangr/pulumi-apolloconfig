// Copyright 2016-2022, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package apolloconfigapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

type ApolloItemClient interface {
	CreateApolloItem(ctx context.Context, params *CreateUpdateApollItemRequest) (*ApollItem, error)
	UpdateApolloItem(ctx context.Context, apolloItemId string, params *CreateUpdateApollItemRequest) error
	DeleteApolloItem(ctx context.Context, env, appId, clusterName, namespace, key, operator string) error
	// DeleteApolloItem(ctx context.Context, apolloItemId, orgName string, forceDestroy bool) error
	GetApolloItem(ctx context.Context, env, appId, clusterName, namespace, key string) (*ApollItem, error)
}

type ApollItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	TokenValue  string `json:"tokenValue"`

	Env                 string `json:"env"`
	AppId               string `json:"appId"`
	ClusterName         string `json:"clusterName"`
	Namespace           string `json:"namespace"`
	Key                 string `json:"key"`
	Value               string `json:"value"`
	Comment             string `json:"comment"`
	DataChangeCreatedBy string `json:"dataChangeCreatedBy"`

	DataChangeLastModifiedBy   string `json:"dataChangeLastModifiedBy"`
	DataChangeCreatedTime      string `json:"dataChangeCreatedTime"`
	DataChangeLastModifiedTime string `json:"dataChangeLastModifiedTime"`
}

type createApollItemResponse struct {
	ID         string `json:"id"`
	TokenValue string `json:"tokenValue"`

	DataChangeLastModifiedBy   string `json:"dataChangeLastModifiedBy"`
	DataChangeCreatedTime      string `json:"dataChangeCreatedTime"`
	DataChangeLastModifiedTime string `json:"dataChangeLastModifiedTime"`
}

type CreateUpdateApollItemRequest struct {
	AppID                    string `json:"appId"`
	Namespace                string `json:"namespace"`
	Env                      string `json:"env"`
	ClusterName              string `json:"clusterName"`
	DataChangeLastModifiedBy string `json:"dataChangeLastModifiedBy"`
	Key                      string `json:"key"`
	Value                    string `json:"value"`
	Comment                  string `json:"comment"`
	Operator                 string `json:"operator"`
	DataChangeCreatedBy      string `json:"dataChangeCreatedBy"`
}

type CreateUpdateApollItemPostData struct {
	Key                 string `json:"key"`
	Value               string `json:"value"`
	Comment             string `json:"comment"`
	DataChangeCreatedBy string `json:"dataChangeCreatedBy"`
}

type UpdateApollItemPostData struct {
	Key                      string `json:"key"`
	Value                    string `json:"value"`
	Comment                  string `json:"comment"`
	DataChangeLastModifiedBy string `json:"dataChangeLastModifiedBy"`
	DataChangeCreatedBy      string `json:"dataChangeCreatedBy"`
}

func (p *CreateUpdateApollItemRequest) Validate() error {
	if p.AppID == "" {
		return errors.New("empty appId")
	}
	if p.Namespace == "" {
		return errors.New("empty namespace")
	}
	if p.Env == "" {
		return errors.New("empty env")
	}
	if p.ClusterName == "" {
		return errors.New("empty clusterName")
	}
	return nil
}

func (c *Client) CreateApolloItem(ctx context.Context, params *CreateUpdateApollItemRequest) (*ApollItem, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	apiPath := path.Join("v1", "envs", params.Env, "apps", params.AppID, "clusters", params.ClusterName, "namespaces", params.Namespace, "items")

	createReq := CreateUpdateApollItemPostData{
		Key:                 params.Key,
		Value:               params.Value,
		Comment:             params.Comment,
		DataChangeCreatedBy: params.DataChangeCreatedBy,
	}

	var createRes createApollItemResponse

	fmt.Printf("apiPath createReq: %+v\n", createReq)
	fmt.Printf("apiPath createRes: %+v\n", createRes)
	fmt.Printf("apiPath apiPath: %+v\n", apiPath)
	fmt.Printf("apiPath http.MethodPost: %+v\n", http.MethodPost)

	_, err := c.do(ctx, http.MethodPost, apiPath, createReq, &createRes)

	if err != nil {
		return nil, fmt.Errorf("failed to create apollo item: %w", err)
	}

	return &ApollItem{
		ID:          createRes.ID,
		Name:        createReq.Key,
		Description: createReq.Comment,
		TokenValue:  createRes.TokenValue,

		Env:                 params.Env,
		AppId:               params.AppID,
		ClusterName:         params.ClusterName,
		Namespace:           params.Namespace,
		Key:                 params.Key,
		Value:               params.Value,
		Comment:             params.Comment,
		DataChangeCreatedBy: params.DataChangeCreatedBy,

		DataChangeLastModifiedBy:   createRes.DataChangeLastModifiedBy,
		DataChangeCreatedTime:      createRes.DataChangeCreatedTime,
		DataChangeLastModifiedTime: createRes.DataChangeLastModifiedTime,
	}, nil

}

func (c *Client) UpdateApolloItem(ctx context.Context, apolloItemId string, params *CreateUpdateApollItemRequest) error {
	if err := params.Validate(); err != nil {
		return err
	}

	if len(apolloItemId) == 0 {
		return errors.New("apolloItemId length must be greater than zero")
	}

	// apiPath := path.Join("orgs", orgName, "apollo-items", apolloItemId)
	apiPath := path.Join("v1", "envs", params.Env, "apps", params.AppID, "clusters", params.ClusterName, "namespaces", params.Namespace, "items", params.Key)

	updateReq := UpdateApollItemPostData{
		Key:                       params.Key,
		Value:                     params.Value,
		Comment:                   params.Comment,
		DataChangeCreatedBy:       params.DataChangeCreatedBy,
		DataChangeLastModifiedBy:  params.DataChangeLastModifiedBy,
	}

	_, err := c.do(ctx, http.MethodPut, apiPath, updateReq, nil)
	if err != nil {
		return fmt.Errorf("failed to update apollo item: %w", err)
	}
	return nil
}

// func (c *Client) DeleteApolloItem(ctx context.Context, apolloItemId, orgName string, forceDestroy bool) error {
// 	if len(apolloItemId) == 0 {
// 		return errors.New("apolloItemId length must be greater than zero")
// 	}

// 	if len(orgName) == 0 {
// 		return errors.New("orgName length must be greater than zero")
// 	}

// 	apiPath := path.Join("orgs", orgName, "apollo-items", apolloItemId)

// 	var err error
// 	if forceDestroy {
// 		_, err = c.doWithQuery(ctx, http.MethodDelete, apiPath, url.Values{"force": []string{"true"}}, nil, nil)
// 	} else {
// 		_, err = c.do(ctx, http.MethodDelete, apiPath, nil, nil)
// 	}
// 	if err != nil {
// 		return fmt.Errorf("failed to delete apollo item %q: %w", apolloItemId, err)
// 	}

// 	return nil
// }

func (c *Client) DeleteApolloItem(ctx context.Context, env, appId, clusterName, namespace, key, operator string) error {
	switch {
	case env == "":
		return errors.New("env length must be greater than zero")
	case appId == "":
		return errors.New("appId length must be greater than zero")
	case clusterName == "":
		return errors.New("clusterName length must be greater than zero")
	case namespace == "":
		return errors.New("namespace length must be greater than zero")
	case key == "":
		return errors.New("key length must be greater than zero")
	}

	// apiPath := path.Join("orgs", orgName, "apollo-items", apolloItemId)
	apiPath := path.Join("v1", "envs", env, "apps", appId, "clusters", clusterName, "namespaces", namespace, "items", key)

	var err error
	_, err = c.doWithQuery(ctx, http.MethodDelete, apiPath, url.Values{"operator": []string{operator}}, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete apollo item %q: %w", key, err)
	}

	return nil
}

func (c *Client) GetApolloItem(ctx context.Context, env, appId, clusterName, namespace, key string) (*ApollItem, error) {
	// apiPath := path.Join("orgs", orgName, "apollo-items", apolloItemId)
	apiPath := path.Join("v1", "envs", env, "apps", appId, "clusters", clusterName, "namespaces", namespace, "items", key)

	var pool ApollItem
	_, err := c.do(ctx, http.MethodGet, apiPath, nil, &pool)
	if err != nil {
		statusCode := GetErrorStatusCode(err)
		if statusCode == http.StatusNotFound {
			// Important: we return nil here to hint it was not found
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get apollo item: %w", err)
	}

	return &pool, nil
}
