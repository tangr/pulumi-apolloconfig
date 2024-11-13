package provider

import (
	"fmt"
	"os"
)

const (
	EnvVarApolloconfigAuthorizationToken = "APOLLOCONFIG_AUTH_TOKEN"
	EnvVarApolloconfigBackendUrl = "APOLLOCONFIG_URL"
)

var ErrAuthTokenNotFound = fmt.Errorf("pulumi access token not found")

type ApollConfig struct {
	Config map[string]string
}

func (ac *ApollConfig) getConfig(configName, envName string) string {
	if val, ok := ac.Config[configName]; ok {
		return val
	}

	return os.Getenv(envName)
}

func (ac *ApollConfig) getPulumiAccessToken() (*string, error) {
	token := ac.getConfig("accessToken", EnvVarPulumiAccessToken)

	if len(token) > 0 {
		// found the token
		return &token, nil
	}

	return nil, ErrAuthTokenNotFound
}

func (ac *ApollConfig) getPulumiServiceUrl() (*string, error) {
	url := ac.getConfig("apiUrl", EnvVarPulumiBackendUrl)
	baseurl := "https://api.pulumi.com"

	if len(url) == 0 {
		url = baseurl
	}

	return &url, nil
}
