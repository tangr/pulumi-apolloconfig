// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.pulumiservice;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.pulumiservice.inputs.DeploymentSettingsExecutorContextArgs;
import com.pulumi.pulumiservice.inputs.DeploymentSettingsGithubArgs;
import com.pulumi.pulumiservice.inputs.DeploymentSettingsOperationContextArgs;
import com.pulumi.pulumiservice.inputs.DeploymentSettingsSourceContextArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DeploymentSettingsArgs extends com.pulumi.resources.ResourceArgs {

    public static final DeploymentSettingsArgs Empty = new DeploymentSettingsArgs();

    /**
     * Settings related to the deployment executor.
     * 
     */
    @Import(name="executorContext")
    private @Nullable Output<DeploymentSettingsExecutorContextArgs> executorContext;

    /**
     * @return Settings related to the deployment executor.
     * 
     */
    public Optional<Output<DeploymentSettingsExecutorContextArgs>> executorContext() {
        return Optional.ofNullable(this.executorContext);
    }

    /**
     * GitHub settings for the deployment.
     * 
     */
    @Import(name="github")
    private @Nullable Output<DeploymentSettingsGithubArgs> github;

    /**
     * @return GitHub settings for the deployment.
     * 
     */
    public Optional<Output<DeploymentSettingsGithubArgs>> github() {
        return Optional.ofNullable(this.github);
    }

    /**
     * Settings related to the Pulumi operation environment during the deployment.
     * 
     */
    @Import(name="operationContext")
    private @Nullable Output<DeploymentSettingsOperationContextArgs> operationContext;

    /**
     * @return Settings related to the Pulumi operation environment during the deployment.
     * 
     */
    public Optional<Output<DeploymentSettingsOperationContextArgs>> operationContext() {
        return Optional.ofNullable(this.operationContext);
    }

    /**
     * Organization name.
     * 
     */
    @Import(name="organization", required=true)
    private Output<String> organization;

    /**
     * @return Organization name.
     * 
     */
    public Output<String> organization() {
        return this.organization;
    }

    /**
     * Project name.
     * 
     */
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return Project name.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    /**
     * Settings related to the source of the deployment.
     * 
     */
    @Import(name="sourceContext", required=true)
    private Output<DeploymentSettingsSourceContextArgs> sourceContext;

    /**
     * @return Settings related to the source of the deployment.
     * 
     */
    public Output<DeploymentSettingsSourceContextArgs> sourceContext() {
        return this.sourceContext;
    }

    /**
     * Stack name.
     * 
     */
    @Import(name="stack", required=true)
    private Output<String> stack;

    /**
     * @return Stack name.
     * 
     */
    public Output<String> stack() {
        return this.stack;
    }

    private DeploymentSettingsArgs() {}

    private DeploymentSettingsArgs(DeploymentSettingsArgs $) {
        this.executorContext = $.executorContext;
        this.github = $.github;
        this.operationContext = $.operationContext;
        this.organization = $.organization;
        this.project = $.project;
        this.sourceContext = $.sourceContext;
        this.stack = $.stack;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DeploymentSettingsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DeploymentSettingsArgs $;

        public Builder() {
            $ = new DeploymentSettingsArgs();
        }

        public Builder(DeploymentSettingsArgs defaults) {
            $ = new DeploymentSettingsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param executorContext Settings related to the deployment executor.
         * 
         * @return builder
         * 
         */
        public Builder executorContext(@Nullable Output<DeploymentSettingsExecutorContextArgs> executorContext) {
            $.executorContext = executorContext;
            return this;
        }

        /**
         * @param executorContext Settings related to the deployment executor.
         * 
         * @return builder
         * 
         */
        public Builder executorContext(DeploymentSettingsExecutorContextArgs executorContext) {
            return executorContext(Output.of(executorContext));
        }

        /**
         * @param github GitHub settings for the deployment.
         * 
         * @return builder
         * 
         */
        public Builder github(@Nullable Output<DeploymentSettingsGithubArgs> github) {
            $.github = github;
            return this;
        }

        /**
         * @param github GitHub settings for the deployment.
         * 
         * @return builder
         * 
         */
        public Builder github(DeploymentSettingsGithubArgs github) {
            return github(Output.of(github));
        }

        /**
         * @param operationContext Settings related to the Pulumi operation environment during the deployment.
         * 
         * @return builder
         * 
         */
        public Builder operationContext(@Nullable Output<DeploymentSettingsOperationContextArgs> operationContext) {
            $.operationContext = operationContext;
            return this;
        }

        /**
         * @param operationContext Settings related to the Pulumi operation environment during the deployment.
         * 
         * @return builder
         * 
         */
        public Builder operationContext(DeploymentSettingsOperationContextArgs operationContext) {
            return operationContext(Output.of(operationContext));
        }

        /**
         * @param organization Organization name.
         * 
         * @return builder
         * 
         */
        public Builder organization(Output<String> organization) {
            $.organization = organization;
            return this;
        }

        /**
         * @param organization Organization name.
         * 
         * @return builder
         * 
         */
        public Builder organization(String organization) {
            return organization(Output.of(organization));
        }

        /**
         * @param project Project name.
         * 
         * @return builder
         * 
         */
        public Builder project(Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project Project name.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param sourceContext Settings related to the source of the deployment.
         * 
         * @return builder
         * 
         */
        public Builder sourceContext(Output<DeploymentSettingsSourceContextArgs> sourceContext) {
            $.sourceContext = sourceContext;
            return this;
        }

        /**
         * @param sourceContext Settings related to the source of the deployment.
         * 
         * @return builder
         * 
         */
        public Builder sourceContext(DeploymentSettingsSourceContextArgs sourceContext) {
            return sourceContext(Output.of(sourceContext));
        }

        /**
         * @param stack Stack name.
         * 
         * @return builder
         * 
         */
        public Builder stack(Output<String> stack) {
            $.stack = stack;
            return this;
        }

        /**
         * @param stack Stack name.
         * 
         * @return builder
         * 
         */
        public Builder stack(String stack) {
            return stack(Output.of(stack));
        }

        public DeploymentSettingsArgs build() {
            $.organization = Objects.requireNonNull($.organization, "expected parameter 'organization' to be non-null");
            $.project = Objects.requireNonNull($.project, "expected parameter 'project' to be non-null");
            $.sourceContext = Objects.requireNonNull($.sourceContext, "expected parameter 'sourceContext' to be non-null");
            $.stack = Objects.requireNonNull($.stack, "expected parameter 'stack' to be non-null");
            return $;
        }
    }

}
