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
	CreateApolloItem(ctx context.Context, orgName, name, description string) (*ApollItem, error)
	UpdateApolloItem(ctx context.Context, apolloItemId, orgName, name, description string) error
	DeleteApolloItem(ctx context.Context, apolloItemId, orgName string, forceDestroy bool) error
	GetApolloItem(ctx context.Context, apolloItemId, orgName string) (*ApollItem, error)
}

type ApollItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	TokenValue  string `json:"tokenValue"`
}

type createApollItemResponse struct {
	ID         string `json:"id"`
	TokenValue string `json:"tokenValue"`
}

type createUpdateApollItemRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func (c *Client) CreateApolloItem(ctx context.Context, orgName, name, description string) (*ApollItem, error) {

	if len(orgName) == 0 {
		return nil, errors.New("empty orgName")
	}

	if len(name) == 0 {
		return nil, errors.New("empty name")
	}

	apiPath := path.Join("orgs", orgName, "apollo-items")

	createReq := createUpdateApollItemRequest{
		Name:        name,
		Description: description,
	}

	var createRes createApollItemResponse

	_, err := c.do(ctx, http.MethodPost, apiPath, createReq, &createRes)

	if err != nil {
		return nil, fmt.Errorf("failed to create apollo item: %w", err)
	}

	return &ApollItem{
		ID:          createRes.ID,
		Name:        createReq.Name,
		Description: createReq.Description,
		TokenValue:  createRes.TokenValue,
	}, nil

}

func (c *Client) UpdateApolloItem(ctx context.Context, apolloItemId, orgName, name, description string) error {
	if len(apolloItemId) == 0 {
		return errors.New("apolloItemId length must be greater than zero")
	}

	if len(orgName) == 0 {
		return errors.New("empty orgName")
	}

	if len(name) == 0 {
		return errors.New("empty name")
	}

	apiPath := path.Join("orgs", orgName, "apollo-items", apolloItemId)

	updateReq := createUpdateApollItemRequest{
		Name:        name,
		Description: description,
	}

	_, err := c.do(ctx, http.MethodPatch, apiPath, updateReq, nil)
	if err != nil {
		return fmt.Errorf("failed to update apollo item: %w", err)
	}
	return nil
}

func (c *Client) DeleteApolloItem(ctx context.Context, apolloItemId, orgName string, forceDestroy bool) error {
	if len(apolloItemId) == 0 {
		return errors.New("apolloItemId length must be greater than zero")
	}

	if len(orgName) == 0 {
		return errors.New("orgName length must be greater than zero")
	}

	apiPath := path.Join("orgs", orgName, "apollo-items", apolloItemId)

	var err error
	if forceDestroy {
		_, err = c.doWithQuery(ctx, http.MethodDelete, apiPath, url.Values{"force": []string{"true"}}, nil, nil)
	} else {
		_, err = c.do(ctx, http.MethodDelete, apiPath, nil, nil)
	}
	if err != nil {
		return fmt.Errorf("failed to delete apollo item %q: %w", apolloItemId, err)
	}

	return nil
}

func (c *Client) GetApolloItem(ctx context.Context, apolloItemId, orgName string) (*ApollItem, error) {
	apiPath := path.Join("orgs", orgName, "apollo-items", apolloItemId)

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
