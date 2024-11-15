package provider

import (
	"fmt"
	"os"
)

const (
	EnvVarApolloconfigAuthorizationToken = "APOLLOCONFIG_AUTH_TOKEN"
	EnvVarApolloconfigBackendUrl = "APOLLOCONFIG_APIURL"
)

var ErrAuthTokenNotFound = fmt.Errorf("apolloconfig auth token not found")
var ErrApiUrlNotFound = fmt.Errorf("apolloconfig api url not found")

type ApollConfig struct {
	Config map[string]string
}

func (ac *ApollConfig) getConfig(configName, envName string) string {
	if val, ok := ac.Config[configName]; ok {
		return val
	}

	return os.Getenv(envName)
}

func (ac *ApollConfig) getApolloConfigAuthToken() (*string, error) {
	token := ac.getConfig("authToken", EnvVarApolloconfigAuthorizationToken)

	if len(token) > 0 {
		// found the token
		return &token, nil
	}

	return nil, ErrAuthTokenNotFound
}

func (ac *ApollConfig) getApolloConfigUrl() (*string, error) {
	url := ac.getConfig("apiUrl", EnvVarApolloconfigBackendUrl)

	if len(url) > 0 {
		return &url, nil
	}

	return nil, ErrApiUrlNotFound
}
