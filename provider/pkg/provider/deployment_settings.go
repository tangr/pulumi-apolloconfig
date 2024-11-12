package provider

import (
	"context"
	"fmt"
	"path"
	"time"

	pbempty "google.golang.org/protobuf/types/known/emptypb"

	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/tangr/pulumi-apolloconfig/provider/pkg/internal/pulumiapi"
)

// This is a value for imported secrets, to hint that value needs to be replaced
// in generated code
const replaceMe = "<REPLACE WITH ACTUAL SECRET VALUE>"

type PulumiServiceDeploymentSettingsInput struct {
	pulumiapi.DeploymentSettings
	Stack pulumiapi.StackIdentifier
}

// plaintextInputSettings are the latest inputs of the resource, containing plaintext values wrapped in Secrets
// currentStateCipherSettings are the latest outputs/properties of the resource, containing ciphertext strings of secret values
// isInput is a flag that selects whether to generating an input PropertyMap that contains plaintext (true) or an output PropertyMap that contains ciphertext (false)
func (ds *PulumiServiceDeploymentSettingsInput) ToPropertyMap(plaintextInputSettings *pulumiapi.DeploymentSettings, currentStateCipherSettings *pulumiapi.DeploymentSettings, isInput bool) resource.PropertyMap {
	// Below flags are used throughout this method and direct the serialization of twin value secrets
	// Twin value secrets are values whose plaintext cannot be retrieved from the API, thus forcing the development of this fairly complex system
	// When plaintextInputSettings is passed in, but currentStateCipherSettings is not, that means the resource is being created or updated
	createMode := plaintextInputSettings != nil && currentStateCipherSettings == nil
	// When both plaintextInputSettings and currentStateCipherSettings are passed in, that means an existing resource is being refreshed, and it's necessary to merge values
	// In case we are merging, but some of the properties don't previously exist, we will merge with empty value, setting plaintext to be empty string
	mergeMode := plaintextInputSettings != nil && currentStateCipherSettings != nil
	// If neither one is passed in, we are importing an existing resource into the state

	pm := resource.PropertyMap{}
	pm["organization"] = resource.NewPropertyValue(ds.Stack.OrgName)
	pm["project"] = resource.NewPropertyValue(ds.Stack.ProjectName)
	pm["stack"] = resource.NewPropertyValue(ds.Stack.StackName)

	if ds.AgentPoolId != "" {
		pm["agentPoolId"] = resource.NewPropertyValue(ds.AgentPoolId)
	}

	if ds.SourceContext != nil {
		scMap := resource.PropertyMap{}
		if ds.SourceContext.Git != nil {
			gitPropertyMap := resource.PropertyMap{}
			if ds.SourceContext.Git.RepoURL != "" {
				gitPropertyMap["repoUrl"] = resource.NewPropertyValue(ds.SourceContext.Git.RepoURL)
			}
			if ds.SourceContext.Git.Commit != "" {
				gitPropertyMap["commit"] = resource.NewPropertyValue(ds.SourceContext.Git.Commit)
			}
			if ds.SourceContext.Git.Branch != "" {
				gitPropertyMap["branch"] = resource.NewPropertyValue(ds.SourceContext.Git.Branch)
			}
			if ds.SourceContext.Git.RepoDir != "" {
				gitPropertyMap["repoDir"] = resource.NewPropertyValue(ds.SourceContext.Git.RepoDir)
			}
			if ds.SourceContext.Git.GitAuth != nil {
				gitAuthPropertyMap := resource.PropertyMap{}
				if ds.SourceContext.Git.GitAuth.SSHAuth != nil {
					sshAuthPropertyMap := resource.PropertyMap{}
					if ds.SourceContext.Git.GitAuth.SSHAuth.SSHPrivateKey.Value != "" {
						if mergeMode {
							var plaintextValue *pulumiapi.SecretValue
							var currentCipherValue *pulumiapi.SecretValue
							if currentStateCipherSettings.SourceContext != nil &&
								currentStateCipherSettings.SourceContext.Git != nil &&
								currentStateCipherSettings.SourceContext.Git.GitAuth != nil &&
								currentStateCipherSettings.SourceContext.Git.GitAuth.SSHAuth != nil {
								plaintextValue = &plaintextInputSettings.SourceContext.Git.GitAuth.SSHAuth.SSHPrivateKey
								currentCipherValue = &currentStateCipherSettings.SourceContext.Git.GitAuth.SSHAuth.SSHPrivateKey
							}
							mergeSecretValue(sshAuthPropertyMap, "sshPrivateKey", ds.SourceContext.Git.GitAuth.SSHAuth.SSHPrivateKey, plaintextValue, currentCipherValue, isInput)
						} else if createMode {
							createSecretValue(sshAuthPropertyMap, "sshPrivateKey", ds.SourceContext.Git.GitAuth.SSHAuth.SSHPrivateKey,
								plaintextInputSettings.SourceContext.Git.GitAuth.SSHAuth.SSHPrivateKey, isInput)
						} else {
							importSecretValue(sshAuthPropertyMap, "sshPrivateKey", ds.SourceContext.Git.GitAuth.SSHAuth.SSHPrivateKey, isInput)
						}
					}
					if ds.SourceContext.Git.GitAuth.SSHAuth.Password != nil && ds.SourceContext.Git.GitAuth.SSHAuth.Password.Value != "" {
						if mergeMode {
							var plaintextValue *pulumiapi.SecretValue
							var currentCipherValue *pulumiapi.SecretValue
							if currentStateCipherSettings.SourceContext != nil &&
								currentStateCipherSettings.SourceContext.Git != nil &&
								currentStateCipherSettings.SourceContext.Git.GitAuth != nil &&
								currentStateCipherSettings.SourceContext.Git.GitAuth.SSHAuth != nil &&
								currentStateCipherSettings.SourceContext.Git.GitAuth.SSHAuth.Password != nil {
								plaintextValue = plaintextInputSettings.SourceContext.Git.GitAuth.SSHAuth.Password
								currentCipherValue = currentStateCipherSettings.SourceContext.Git.GitAuth.SSHAuth.Password
							}
							mergeSecretValue(sshAuthPropertyMap, "password", *ds.SourceContext.Git.GitAuth.SSHAuth.Password, plaintextValue, currentCipherValue, isInput)
						} else if createMode {
							createSecretValue(sshAuthPropertyMap, "password", *ds.SourceContext.Git.GitAuth.SSHAuth.Password,
								*plaintextInputSettings.SourceContext.Git.GitAuth.SSHAuth.Password, isInput)
						} else {
							importSecretValue(sshAuthPropertyMap, "password", *ds.SourceContext.Git.GitAuth.SSHAuth.Password, isInput)
						}
					}
					gitAuthPropertyMap["sshAuth"] = resource.PropertyValue{V: sshAuthPropertyMap}
				}
				if ds.SourceContext.Git.GitAuth.BasicAuth != nil {
					basicAuthPropertyMap := resource.PropertyMap{}
					if ds.SourceContext.Git.GitAuth.BasicAuth.UserName.Value != "" {
						basicAuthPropertyMap["username"] = resource.NewPropertyValue(ds.SourceContext.Git.GitAuth.BasicAuth.UserName.Value)
					}
					if ds.SourceContext.Git.GitAuth.BasicAuth.Password.Value != "" {
						if mergeMode {
							var plaintextValue *pulumiapi.SecretValue
							var currentCipherValue *pulumiapi.SecretValue
							if currentStateCipherSettings.SourceContext != nil &&
								currentStateCipherSettings.SourceContext.Git != nil &&
								currentStateCipherSettings.SourceContext.Git.GitAuth != nil &&
								currentStateCipherSettings.SourceContext.Git.GitAuth.BasicAuth != nil {
								plaintextValue = &plaintextInputSettings.SourceContext.Git.GitAuth.BasicAuth.Password
								currentCipherValue = &currentStateCipherSettings.SourceContext.Git.GitAuth.BasicAuth.Password
							}
							mergeSecretValue(basicAuthPropertyMap, "password", ds.SourceContext.Git.GitAuth.BasicAuth.Password, plaintextValue, currentCipherValue, isInput)
						} else if createMode {
							createSecretValue(basicAuthPropertyMap, "password", ds.SourceContext.Git.GitAuth.BasicAuth.Password,
								plaintextInputSettings.SourceContext.Git.GitAuth.BasicAuth.Password, isInput)
						} else {
							importSecretValue(basicAuthPropertyMap, "password", ds.SourceContext.Git.GitAuth.BasicAuth.Password, isInput)
						}
					}
					gitAuthPropertyMap["basicAuth"] = resource.PropertyValue{V: basicAuthPropertyMap}
				}
				gitPropertyMap["gitAuth"] = resource.PropertyValue{V: gitAuthPropertyMap}
			}
			scMap["git"] = resource.PropertyValue{V: gitPropertyMap}
		}
		pm["sourceContext"] = resource.PropertyValue{V: scMap}
	}

	if ds.OperationContext != nil {
		ocMap := resource.PropertyMap{}
		if ds.OperationContext.PreRunCommands != nil {
			ocMap["preRunCommands"] = resource.NewPropertyValue(ds.OperationContext.PreRunCommands)
		}
		if ds.OperationContext.EnvironmentVariables != nil {
			evMap := resource.PropertyMap{}
			for k, v := range ds.OperationContext.EnvironmentVariables {
				if v.Secret {
					if mergeMode {
						var plaintextValue pulumiapi.SecretValue
						var currentCipherValue pulumiapi.SecretValue
						if currentStateCipherSettings.OperationContext != nil {
							plaintextValue = plaintextInputSettings.OperationContext.EnvironmentVariables[k]
							currentCipherValue = currentStateCipherSettings.OperationContext.EnvironmentVariables[k]
						}
						mergeSecretValue(evMap, k, v, &plaintextValue, &currentCipherValue, isInput)
					} else if createMode {
						createSecretValue(evMap, k, v,
							plaintextInputSettings.OperationContext.EnvironmentVariables[k], isInput)
					} else {
						importSecretValue(evMap, k, v, isInput)
					}
				} else {
					evMap[resource.PropertyKey(k)] = resource.NewPropertyValue(v.Value)
				}
			}
			ocMap["environmentVariables"] = resource.PropertyValue{V: evMap}
		}
		if ds.OperationContext.Options != nil {
			optionsMap := resource.PropertyMap{}
			if ds.OperationContext.Options.Shell != "" {
				optionsMap["shell"] = resource.NewPropertyValue(ds.OperationContext.Options.Shell)
			}
			if ds.OperationContext.Options.SkipInstallDependencies {
				optionsMap["skipInstallDependencies"] = resource.NewPropertyValue(true)
			}
			if ds.OperationContext.Options.SkipIntermediateDeployments {
				optionsMap["skipIntermediateDeployments"] = resource.NewPropertyValue(true)
			}
			if ds.OperationContext.Options.DeleteAfterDestroy {
				optionsMap["deleteAfterDestroy"] = resource.NewPropertyValue(true)
			}
			ocMap["options"] = resource.PropertyValue{V: optionsMap}
		}
		if ds.OperationContext.OIDC != nil {
			if ds.OperationContext.OIDC.AWS != nil || ds.OperationContext.OIDC.GCP != nil || ds.OperationContext.OIDC.Azure != nil {
				oidcMap := resource.PropertyMap{}
				if ds.OperationContext.OIDC.AWS != nil {
					awsMap := resource.PropertyMap{}
					if ds.OperationContext.OIDC.AWS.RoleARN != "" {
						awsMap["roleARN"] = resource.NewPropertyValue(ds.OperationContext.OIDC.AWS.RoleARN)
					}
					if ds.OperationContext.OIDC.AWS.SessionName != "" {
						awsMap["sessionName"] = resource.NewPropertyValue(ds.OperationContext.OIDC.AWS.SessionName)
					}
					if ds.OperationContext.OIDC.AWS.PolicyARNs != nil {
						awsMap["policyARNs"] = resource.NewPropertyValue(ds.OperationContext.OIDC.AWS.PolicyARNs)
					}
					if ds.OperationContext.OIDC.AWS.Duration != "" {
						awsMap["duration"] = resource.NewPropertyValue(ds.OperationContext.OIDC.AWS.Duration)
					}
					oidcMap["aws"] = resource.PropertyValue{V: awsMap}
				}
				if ds.OperationContext.OIDC.GCP != nil {
					gcpMap := resource.PropertyMap{}
					if ds.OperationContext.OIDC.GCP.ProviderID != "" {
						gcpMap["providerId"] = resource.NewPropertyValue(ds.OperationContext.OIDC.GCP.ProviderID)
					}
					if ds.OperationContext.OIDC.GCP.ServiceAccount != "" {
						gcpMap["serviceAccount"] = resource.NewPropertyValue(ds.OperationContext.OIDC.GCP.ServiceAccount)
					}
					if ds.OperationContext.OIDC.GCP.Region != "" {
						gcpMap["region"] = resource.NewPropertyValue(ds.OperationContext.OIDC.GCP.Region)
					}
					if ds.OperationContext.OIDC.GCP.WorkloadPoolID != "" {
						gcpMap["workloadPoolId"] = resource.NewPropertyValue(ds.OperationContext.OIDC.GCP.WorkloadPoolID)
					}
					if ds.OperationContext.OIDC.GCP.ProjectID != "" {
						gcpMap["projectId"] = resource.NewPropertyValue(ds.OperationContext.OIDC.GCP.ProjectID)
					}
					if ds.OperationContext.OIDC.GCP.TokenLifetime != "" {
						gcpMap["tokenLifetime"] = resource.NewPropertyValue(ds.OperationContext.OIDC.GCP.TokenLifetime)
					}
					oidcMap["gcp"] = resource.PropertyValue{V: gcpMap}
				}
				if ds.OperationContext.OIDC.Azure != nil {
					azureMap := resource.PropertyMap{}
					if ds.OperationContext.OIDC.Azure.TenantID != "" {
						azureMap["tenantId"] = resource.NewPropertyValue(ds.OperationContext.OIDC.Azure.TenantID)
					}
					if ds.OperationContext.OIDC.Azure.ClientID != "" {
						azureMap["clientId"] = resource.NewPropertyValue(ds.OperationContext.OIDC.Azure.ClientID)
					}
					if ds.OperationContext.OIDC.Azure.SubscriptionID != "" {
						azureMap["subscriptionId"] = resource.NewPropertyValue(ds.OperationContext.OIDC.Azure.SubscriptionID)
					}
					oidcMap["azure"] = resource.PropertyValue{V: azureMap}
				}
				ocMap["oidc"] = resource.PropertyValue{V: oidcMap}
			}
		}
		pm["operationContext"] = resource.PropertyValue{V: ocMap}
	}

	if ds.GitHub != nil {
		githubMap := resource.PropertyMap{}
		githubMap["previewPullRequests"] = resource.NewPropertyValue(ds.GitHub.PreviewPullRequests)
		githubMap["deployCommits"] = resource.NewPropertyValue(ds.GitHub.DeployCommits)
		githubMap["pullRequestTemplate"] = resource.NewPropertyValue(ds.GitHub.PullRequestTemplate)
		if ds.GitHub.Repository != "" {
			githubMap["repository"] = resource.NewPropertyValue(ds.GitHub.Repository)
		}
		if len(ds.GitHub.Paths) > 0 {
			githubMap["paths"] = resource.NewPropertyValue(ds.GitHub.Paths)
		}
		pm["github"] = resource.PropertyValue{V: githubMap}
	}

	if ds.ExecutorContext != nil && ds.ExecutorContext.ExecutorImage != nil && ds.ExecutorContext.ExecutorImage.Reference != "" {
		ecMap := resource.PropertyMap{}
		ecMap["executorImage"] = resource.NewPropertyValue(ds.ExecutorContext.ExecutorImage.Reference)
		pm["executorContext"] = resource.PropertyValue{V: ecMap}
	}

	if ds.CacheOptions != nil {
		coMap := resource.PropertyMap{}
		coMap["enable"] = resource.NewPropertyValue(ds.CacheOptions.Enable)
		pm["cacheOptions"] = resource.PropertyValue{coMap}
	}

	return pm
}

// All imported inputs will have a dummy value, asking to be replaced in real code
// All imported properties are just set to ciphertext read from Pulumi Service
func importSecretValue(propertyMap resource.PropertyMap, propertyName string, cipherValue pulumiapi.SecretValue, isInput bool) {
	if isInput {
		propertyMap[resource.PropertyKey(propertyName)] = resource.MakeSecret(resource.NewPropertyValue(replaceMe))
	} else {
		propertyMap[resource.PropertyKey(propertyName)] = resource.NewPropertyValue(cipherValue.Value)
	}
}

// On Create or Update, inputs already have a plaintext value, just set it
// Properties are just set to ciphertext returned from Pulumi Service
func createSecretValue(propertyMap resource.PropertyMap, propertyName string, cipherValue pulumiapi.SecretValue, plaintextValue pulumiapi.SecretValue, isInput bool) {
	if isInput {
		propertyMap[resource.PropertyKey(propertyName)] = resource.MakeSecret(resource.NewPropertyValue(plaintextValue.Value))
	} else {
		propertyMap[resource.PropertyKey(propertyName)] = resource.NewPropertyValue(cipherValue.Value)
	}
}

// Merge happens when existing resource is refreshed from Pulumi Service
// Output properties are just replaced with ciphertext retrieved from Pulumi Service
// Inputs are more complicated :
// If ciphertext never changed, keep existing plaintext value
// If ciphertext is different, set plaintext to empty string
// If retrieved state has a value that current state does not have, pass in nil, which will fill plaintext with empty string
func mergeSecretValue(propertyMap resource.PropertyMap, propertyName string, cipherValue pulumiapi.SecretValue, plaintextValue *pulumiapi.SecretValue, oldCipherValue *pulumiapi.SecretValue, isInput bool) {
	if isInput {
		if oldCipherValue != nil && cipherValue.Value == oldCipherValue.Value {
			propertyMap[resource.PropertyKey(propertyName)] = resource.MakeSecret(resource.NewPropertyValue(plaintextValue.Value))
		} else {
			propertyMap[resource.PropertyKey(propertyName)] = resource.MakeSecret(resource.NewPropertyValue(""))
		}
	} else {
		propertyMap[resource.PropertyKey(propertyName)] = resource.NewPropertyValue(cipherValue.Value)
	}
}

type PulumiServiceDeploymentSettingsResource struct {
	client pulumiapi.DeploymentSettingsClient
}

func (ds *PulumiServiceDeploymentSettingsResource) ToPulumiServiceDeploymentSettingsInput(inputMap resource.PropertyMap) PulumiServiceDeploymentSettingsInput {
	input := PulumiServiceDeploymentSettingsInput{}

	input.Stack.OrgName = getSecretOrStringValue(inputMap["organization"])
	input.Stack.ProjectName = getSecretOrStringValue(inputMap["project"])
	input.Stack.StackName = getSecretOrStringValue(inputMap["stack"])

	if inputMap["agentPoolId"].HasValue() && inputMap["agentPoolId"].IsString() {
		input.AgentPoolId = inputMap["agentPoolId"].StringValue()
	}

	input.ExecutorContext = toExecutorContext(inputMap)
	input.GitHub = toGitHubConfig(inputMap)
	input.SourceContext = toSourceContext(inputMap)
	input.OperationContext = toOperationContext(inputMap)
	input.CacheOptions = toCacheOptions(inputMap)

	return input
}

func toExecutorContext(inputMap resource.PropertyMap) *apitype.ExecutorContext {
	if !inputMap["executorContext"].HasValue() || !inputMap["executorContext"].IsObject() {
		return nil
	}

	ecInput := inputMap["executorContext"].ObjectValue()
	var ec apitype.ExecutorContext

	if ecInput["executorImage"].HasValue() {
		ec.ExecutorImage = &apitype.DockerImage{
			Reference: getSecretOrStringValue(ecInput["executorImage"]),
		}
	}

	return &ec
}

func toGitHubConfig(inputMap resource.PropertyMap) *pulumiapi.GitHubConfiguration {
	if !inputMap["github"].HasValue() || !inputMap["github"].IsObject() {
		return nil
	}

	githubInput := inputMap["github"].ObjectValue()
	var github pulumiapi.GitHubConfiguration

	if githubInput["repository"].HasValue() {
		github.Repository = getSecretOrStringValue(githubInput["repository"])
	}

	if githubInput["deployCommits"].HasValue() && githubInput["deployCommits"].IsBool() {
		github.DeployCommits = githubInput["deployCommits"].BoolValue()
	}
	if githubInput["previewPullRequests"].HasValue() && githubInput["previewPullRequests"].IsBool() {
		github.PreviewPullRequests = githubInput["previewPullRequests"].BoolValue()
	}
	if githubInput["pullRequestTemplate"].HasValue() && githubInput["pullRequestTemplate"].IsBool() {
		github.PullRequestTemplate = githubInput["pullRequestTemplate"].BoolValue()
	}
	if githubInput["paths"].HasValue() && githubInput["paths"].IsArray() {
		pathsInput := githubInput["paths"].ArrayValue()
		paths := make([]string, len(pathsInput))

		for i, v := range pathsInput {
			paths[i] = getSecretOrStringValue(v)
		}

		github.Paths = paths
	}

	return &github
}

func toSourceContext(inputMap resource.PropertyMap) *pulumiapi.SourceContext {
	if !inputMap["sourceContext"].HasValue() || !inputMap["sourceContext"].IsObject() {
		return nil
	}

	scInput := inputMap["sourceContext"].ObjectValue()
	var sc pulumiapi.SourceContext

	if scInput["git"].HasValue() && scInput["git"].IsObject() {
		gitInput := scInput["git"].ObjectValue()
		var g pulumiapi.SourceContextGit

		if gitInput["repoUrl"].HasValue() {
			g.RepoURL = getSecretOrStringValue(gitInput["repoUrl"])
		}
		if gitInput["branch"].HasValue() {
			g.Branch = getSecretOrStringValue(gitInput["branch"])
		}
		if gitInput["repoDir"].HasValue() {
			g.RepoDir = getSecretOrStringValue(gitInput["repoDir"])
		}

		if gitInput["gitAuth"].HasValue() && gitInput["gitAuth"].IsObject() {
			authInput := gitInput["gitAuth"].ObjectValue()
			var a pulumiapi.GitAuthConfig

			if authInput["sshAuth"].HasValue() && authInput["sshAuth"].IsObject() {
				sshInput := authInput["sshAuth"].ObjectValue()
				var s pulumiapi.SSHAuth

				if sshInput["sshPrivateKey"].HasValue() || sshInput["sshPrivateKeyCipher"].HasValue() {
					s.SSHPrivateKey = pulumiapi.SecretValue{
						Secret: true,
						Value:  getSecretOrStringValue(sshInput["sshPrivateKey"]),
					}
				}
				if sshInput["password"].HasValue() || sshInput["passwordCipher"].HasValue() {
					s.Password = &pulumiapi.SecretValue{
						Secret: true,
						Value:  getSecretOrStringValue(sshInput["password"]),
					}
				}

				a.SSHAuth = &s
			}

			if authInput["basicAuth"].HasValue() && authInput["basicAuth"].IsObject() {
				basicInput := authInput["basicAuth"].ObjectValue()
				var b pulumiapi.BasicAuth

				if basicInput["username"].HasValue() {
					b.UserName = pulumiapi.SecretValue{
						Value:  getSecretOrStringValue(basicInput["username"]),
						Secret: false,
					}
				}
				if basicInput["password"].HasValue() || basicInput["passwordCipher"].HasValue() {
					b.Password = pulumiapi.SecretValue{
						Value:  getSecretOrStringValue(basicInput["password"]),
						Secret: true,
					}
				}

				a.BasicAuth = &b
			}

			g.GitAuth = &a
		}

		sc.Git = &g
	}

	return &sc
}

func toOperationContext(inputMap resource.PropertyMap) *pulumiapi.OperationContext {
	if !inputMap["operationContext"].HasValue() || !inputMap["operationContext"].IsObject() {
		return nil
	}

	ocInput := inputMap["operationContext"].ObjectValue()
	var oc pulumiapi.OperationContext

	if ocInput["environmentVariables"].HasValue() && ocInput["environmentVariables"].IsObject() {
		ev := map[string]pulumiapi.SecretValue{}
		evInput := ocInput["environmentVariables"].ObjectValue()

		for k, v := range evInput {
			value := getSecretOrStringValue(v)
			ev[string(k)] = pulumiapi.SecretValue{Secret: v.IsSecret(), Value: value}
		}

		oc.EnvironmentVariables = ev
	}

	if ocInput["preRunCommands"].HasValue() && ocInput["preRunCommands"].IsArray() {
		pcInput := ocInput["preRunCommands"].ArrayValue()
		pc := make([]string, len(pcInput))

		for i, v := range pcInput {
			if v.IsString() {
				pc[i] = v.StringValue()
			}
		}

		oc.PreRunCommands = pc
	}

	if ocInput["options"].HasValue() && ocInput["options"].IsObject() {
		oInput := ocInput["options"].ObjectValue()
		var o pulumiapi.OperationContextOptions

		if oInput["skipInstallDependencies"].HasValue() && oInput["skipInstallDependencies"].IsBool() {
			o.SkipInstallDependencies = oInput["skipInstallDependencies"].BoolValue()
		}

		if oInput["skipIntermediateDeployments"].HasValue() && oInput["skipIntermediateDeployments"].IsBool() {
			o.SkipIntermediateDeployments = oInput["skipIntermediateDeployments"].BoolValue()
		}

		if oInput["Shell"].HasValue() && oInput["Shell"].IsString() {
			o.Shell = oInput["Shell"].StringValue()
		}

		if oInput["deleteAfterDestroy"].HasValue() && oInput["deleteAfterDestroy"].IsBool() {
			o.DeleteAfterDestroy = oInput["deleteAfterDestroy"].BoolValue()
		}

		oc.Options = &o
	}

	if ocInput["oidc"].HasValue() && ocInput["oidc"].IsObject() {
		oidcInput := ocInput["oidc"].ObjectValue()
		var oidc pulumiapi.OIDCConfiguration

		if oidcInput["aws"].HasValue() && oidcInput["aws"].IsObject() {
			awsInput := oidcInput["aws"].ObjectValue()
			var aws pulumiapi.AWSOIDCConfiguration

			if awsInput["roleARN"].HasValue() {
				aws.RoleARN = getSecretOrStringValue(awsInput["roleARN"])
			}
			if awsInput["duration"].HasValue() {
				aws.Duration = getSecretOrStringValue(awsInput["duration"])
			}
			if awsInput["sessionName"].HasValue() {
				aws.SessionName = getSecretOrStringValue(awsInput["sessionName"])
			}
			if awsInput["policyARNs"].HasValue() && awsInput["policyARNs"].IsArray() {
				policyARNsInput := awsInput["policyARNs"].ArrayValue()
				policyARNs := make([]string, len(policyARNsInput))

				for i, v := range policyARNsInput {
					policyARNs[i] = getSecretOrStringValue(v)
				}

				aws.PolicyARNs = policyARNs
			}

			oidc.AWS = &aws
		}

		if oidcInput["gcp"].HasValue() && oidcInput["gcp"].IsObject() {
			gcpInput := oidcInput["gcp"].ObjectValue()
			var gcp pulumiapi.GCPOIDCConfiguration

			if gcpInput["projectId"].HasValue() {
				gcp.ProjectID = getSecretOrStringValue(gcpInput["projectId"])
			}
			if gcpInput["region"].HasValue() {
				gcp.Region = getSecretOrStringValue(gcpInput["region"])
			}
			if gcpInput["workloadPoolId"].HasValue() {
				gcp.WorkloadPoolID = getSecretOrStringValue(gcpInput["workloadPoolId"])
			}
			if gcpInput["providerId"].HasValue() {
				gcp.ProviderID = getSecretOrStringValue(gcpInput["providerId"])
			}
			if gcpInput["serviceAccount"].HasValue() {
				gcp.ServiceAccount = getSecretOrStringValue(gcpInput["serviceAccount"])
			}
			if gcpInput["tokenLifetime"].HasValue() {
				gcp.TokenLifetime = getSecretOrStringValue(gcpInput["tokenLifetime"])
			}

			oidc.GCP = &gcp
		}

		if oidcInput["azure"].HasValue() && oidcInput["azure"].IsObject() {
			azureInput := oidcInput["azure"].ObjectValue()
			var azure pulumiapi.AzureOIDCConfiguration

			if azureInput["tenantId"].HasValue() {
				azure.TenantID = getSecretOrStringValue(azureInput["tenantId"])
			}
			if azureInput["clientId"].HasValue() {
				azure.ClientID = getSecretOrStringValue(azureInput["clientId"])
			}
			if azureInput["subscriptionId"].HasValue() {
				azure.SubscriptionID = getSecretOrStringValue(azureInput["subscriptionId"])
			}

			oidc.Azure = &azure
		}

		oc.OIDC = &oidc
	}

	return &oc
}

func toCacheOptions(inputMap resource.PropertyMap) *pulumiapi.CacheOptions {
	if !inputMap["cacheOptions"].HasValue() || !inputMap["cacheOptions"].IsObject() {
		return nil
	}

	coInput := inputMap["cacheOptions"].ObjectValue()
	var co pulumiapi.CacheOptions

	if coInput["enable"].HasValue() && coInput["enable"].IsBool() {
		co.Enable = coInput["enable"].BoolValue()
	}

	return &co
}

func getSecretOrStringValue(prop resource.PropertyValue) string {
	switch prop.V.(type) {
	case *resource.Secret:
		return prop.SecretValue().Element.StringValue()
	case nil:
		return ""
	default:
		return prop.StringValue()
	}
}

func (ds *PulumiServiceDeploymentSettingsResource) Diff(req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	olds, err := plugin.UnmarshalProperties(req.GetOldInputs(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
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
	replaces := []string(nil)
	replaceProperties := map[string]bool{
		"organization": true,
		"project":      true,
		"stack":        true,
	}
	for k, v := range dd {
		if _, ok := replaceProperties[k]; ok {
			v.Kind = v.Kind.AsReplace()
			replaces = append(replaces, k)
		}
		detailedDiffs[k] = &pulumirpc.PropertyDiff{
			Kind:      pulumirpc.PropertyDiff_Kind(v.Kind),
			InputDiff: v.InputDiff,
		}
	}

	return &pulumirpc.DiffResponse{
		Changes:             pulumirpc.DiffResponse_DIFF_SOME,
		Replaces:            replaces,
		DetailedDiff:        detailedDiffs,
		HasDetailedDiff:     true,
		DeleteBeforeReplace: true,
	}, nil
}

func (ds *PulumiServiceDeploymentSettingsResource) Check(req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true, KeepSecrets: true})
	if err != nil {
		return nil, err
	}

	var failures []*pulumirpc.CheckFailure
	for _, p := range []resource.PropertyKey{"organization", "project", "stack"} {
		if !news[(p)].HasValue() {
			failures = append(failures, &pulumirpc.CheckFailure{
				Reason:   fmt.Sprintf("missing required property '%s'", p),
				Property: string(p),
			})
		}
	}

	// Normalizing duration input
	if news["operationContext"].HasValue() && news["operationContext"].IsObject() {
		operationContext := news["operationContext"].ObjectValue()
		if operationContext["oidc"].HasValue() && operationContext["oidc"].IsObject() {
			oidc := operationContext["oidc"].ObjectValue()
			if oidc["aws"].HasValue() && oidc["aws"].IsObject() {
				aws := oidc["aws"].ObjectValue()
				if aws["duration"].HasValue() && aws["duration"].IsString() {
					durationString := aws["duration"].StringValue()
					normalized, err := normalizeDurationString(durationString)
					if err != nil {
						failures = append(failures, &pulumirpc.CheckFailure{
							Reason:   fmt.Sprintf("Failed to normalize duration string due to error: %s", err.Error()),
							Property: string("operationContext.oidc.aws.duration"),
						})
					} else {
						aws["duration"] = resource.NewStringProperty(*normalized)
					}
				}
			}
		}
	}

	checkedNews, err := plugin.MarshalProperties(news, plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true, KeepSecrets: true})
	if err != nil {
		return nil, err
	}

	return &pulumirpc.CheckResponse{Inputs: checkedNews, Failures: failures}, nil
}

func (ds *PulumiServiceDeploymentSettingsResource) Configure(_ PulumiServiceConfig) {}

func (ds *PulumiServiceDeploymentSettingsResource) Read(req *pulumirpc.ReadRequest) (*pulumirpc.ReadResponse, error) {
	ctx := context.Background()

	stack, err := pulumiapi.NewStackIdentifier(req.GetId())
	if err != nil {
		return nil, err
	}
	settings, err := ds.client.GetDeploymentSettings(ctx, stack)
	if err != nil {
		return nil, err
	}

	if settings == nil {
		// Empty response causes the resource to be deleted from the state.
		return &pulumirpc.ReadResponse{Id: "", Properties: nil}, nil
	}

	dsInput := PulumiServiceDeploymentSettingsInput{
		Stack:              stack,
		DeploymentSettings: *settings,
	}

	var plaintextSettings *pulumiapi.DeploymentSettings
	var ciphertextSettings *pulumiapi.DeploymentSettings
	propertyMap, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true, KeepSecrets: true})
	if err != nil {
		return nil, err
	}
	inputMap, err := plugin.UnmarshalProperties(req.GetInputs(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true, KeepSecrets: true})
	if err != nil {
		return nil, err
	}
	if propertyMap["stack"].HasValue() {
		tempPlain := ds.ToPulumiServiceDeploymentSettingsInput(inputMap)
		plaintextSettings = &tempPlain.DeploymentSettings
		tempCipher := ds.ToPulumiServiceDeploymentSettingsInput(propertyMap)
		ciphertextSettings = &tempCipher.DeploymentSettings
	}

	properties, err := plugin.MarshalProperties(
		dsInput.ToPropertyMap(plaintextSettings, ciphertextSettings, false),
		plugin.MarshalOptions{
			KeepUnknowns: true,
			SkipNulls:    true,
			KeepSecrets:  true,
		},
	)
	if err != nil {
		return nil, err
	}

	inputs, err := plugin.MarshalProperties(
		dsInput.ToPropertyMap(plaintextSettings, ciphertextSettings, true),
		plugin.MarshalOptions{
			KeepUnknowns: true,
			SkipNulls:    true,
			KeepSecrets:  true,
		},
	)
	if err != nil {
		return nil, err
	}

	return &pulumirpc.ReadResponse{
		Id:         req.Id,
		Properties: properties,
		Inputs:     inputs,
	}, nil
}

func (ds *PulumiServiceDeploymentSettingsResource) Delete(req *pulumirpc.DeleteRequest) (*pbempty.Empty, error) {
	ctx := context.Background()
	stack, err := pulumiapi.NewStackIdentifier(req.GetId())
	if err != nil {
		return nil, err
	}

	err = ds.client.DeleteDeploymentSettings(ctx, stack)
	if err != nil {
		return nil, err
	}
	return &pbempty.Empty{}, nil
}

func (ds *PulumiServiceDeploymentSettingsResource) Create(req *pulumirpc.CreateRequest) (*pulumirpc.CreateResponse, error) {
	ctx := context.Background()
	inputsMap, err := plugin.UnmarshalProperties(req.GetProperties(),
		plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true, KeepSecrets: true})
	if err != nil {
		return nil, err
	}

	input := ds.ToPulumiServiceDeploymentSettingsInput(inputsMap)
	settings := input.DeploymentSettings
	response, err := ds.client.CreateDeploymentSettings(ctx, input.Stack, settings)
	if err != nil {
		return nil, err
	}

	responseInput := PulumiServiceDeploymentSettingsInput{
		DeploymentSettings: *response,
		Stack:              input.Stack,
	}

	outputProperties, err := plugin.MarshalProperties(
		responseInput.ToPropertyMap(&input.DeploymentSettings, nil, false),
		plugin.MarshalOptions{
			KeepUnknowns: true,
			SkipNulls:    true,
			KeepSecrets:  true,
		},
	)
	if err != nil {
		return nil, err
	}

	return &pulumirpc.CreateResponse{
		Id:         path.Join(input.Stack.OrgName, input.Stack.ProjectName, input.Stack.StackName),
		Properties: outputProperties,
	}, nil
}

func (ds *PulumiServiceDeploymentSettingsResource) Update(req *pulumirpc.UpdateRequest) (*pulumirpc.UpdateResponse, error) {
	ctx := context.Background()
	inputsMap, err := plugin.UnmarshalProperties(req.GetNews(),
		plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true, KeepSecrets: true})
	if err != nil {
		return nil, err
	}

	input := ds.ToPulumiServiceDeploymentSettingsInput(inputsMap)
	settings := input.DeploymentSettings
	response, err := ds.client.UpdateDeploymentSettings(ctx, input.Stack, settings)
	if err != nil {
		return nil, err
	}

	responseInput := PulumiServiceDeploymentSettingsInput{
		DeploymentSettings: *response,
		Stack:              input.Stack,
	}

	outputProperties, err := plugin.MarshalProperties(
		responseInput.ToPropertyMap(&input.DeploymentSettings, nil, false),
		plugin.MarshalOptions{
			KeepUnknowns: true,
			SkipNulls:    true,
			KeepSecrets:  true,
		},
	)
	if err != nil {
		return nil, err
	}

	return &pulumirpc.UpdateResponse{
		Properties: outputProperties,
	}, nil
}

func (ds *PulumiServiceDeploymentSettingsResource) Name() string {
	return "apolloconfig:index:DeploymentSettings"
}

func normalizeDurationString(input string) (*string, error) {
	if input == "" {
		return nil, fmt.Errorf("empty value provided for duration string")
	}
	duration, err := time.ParseDuration(input)
	if err != nil {
		return nil, fmt.Errorf("failed to parse duration string `%s` due to error: %w", input, err)
	}
	result := duration.String()

	return &result, nil
}
