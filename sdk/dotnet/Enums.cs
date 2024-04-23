// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.PulumiService
{
    [EnumType]
    public readonly struct PulumiOperation : IEquatable<PulumiOperation>
    {
        private readonly string _value;

        private PulumiOperation(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Analogous to `pulumi up` command.
        /// </summary>
        public static PulumiOperation Update { get; } = new PulumiOperation("update");
        /// <summary>
        /// Analogous to `pulumi preview` command.
        /// </summary>
        public static PulumiOperation Preview { get; } = new PulumiOperation("preview");
        /// <summary>
        /// Analogous to `pulumi refresh` command.
        /// </summary>
        public static PulumiOperation Refresh { get; } = new PulumiOperation("refresh");
        /// <summary>
        /// Analogous to `pulumi destroy` command.
        /// </summary>
        public static PulumiOperation Destroy { get; } = new PulumiOperation("destroy");

        public static bool operator ==(PulumiOperation left, PulumiOperation right) => left.Equals(right);
        public static bool operator !=(PulumiOperation left, PulumiOperation right) => !left.Equals(right);

        public static explicit operator string(PulumiOperation value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PulumiOperation other && Equals(other);
        public bool Equals(PulumiOperation other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct TeamStackPermissionScope : IEquatable<TeamStackPermissionScope>
    {
        private readonly double _value;

        private TeamStackPermissionScope(double value)
        {
            _value = value;
        }

        /// <summary>
        /// Grants read permissions to stack.
        /// </summary>
        public static TeamStackPermissionScope Read { get; } = new TeamStackPermissionScope(101);
        /// <summary>
        /// Grants edit permissions to stack.
        /// </summary>
        public static TeamStackPermissionScope Edit { get; } = new TeamStackPermissionScope(102);
        /// <summary>
        /// Grants admin permissions to stack.
        /// </summary>
        public static TeamStackPermissionScope Admin { get; } = new TeamStackPermissionScope(103);

        public static bool operator ==(TeamStackPermissionScope left, TeamStackPermissionScope right) => left.Equals(right);
        public static bool operator !=(TeamStackPermissionScope left, TeamStackPermissionScope right) => !left.Equals(right);

        public static explicit operator double(TeamStackPermissionScope value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TeamStackPermissionScope other && Equals(other);
        public bool Equals(TeamStackPermissionScope other) => _value == other._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value.GetHashCode();

        public override string ToString() => _value.ToString();
    }

    [EnumType]
    public readonly struct WebhookFilters : IEquatable<WebhookFilters>
    {
        private readonly string _value;

        private WebhookFilters(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Trigger a webhook when a stack is created. Only valid for org webhooks.
        /// </summary>
        public static WebhookFilters StackCreated { get; } = new WebhookFilters("stack_created");
        /// <summary>
        /// Trigger a webhook when a stack is deleted. Only valid for org webhooks.
        /// </summary>
        public static WebhookFilters StackDeleted { get; } = new WebhookFilters("stack_deleted");
        /// <summary>
        /// Trigger a webhook when a stack update succeeds.
        /// </summary>
        public static WebhookFilters UpdateSucceeded { get; } = new WebhookFilters("update_succeeded");
        /// <summary>
        /// Trigger a webhook when a stack update fails.
        /// </summary>
        public static WebhookFilters UpdateFailed { get; } = new WebhookFilters("update_failed");
        /// <summary>
        /// Trigger a webhook when a stack preview succeeds.
        /// </summary>
        public static WebhookFilters PreviewSucceeded { get; } = new WebhookFilters("preview_succeeded");
        /// <summary>
        /// Trigger a webhook when a stack preview fails.
        /// </summary>
        public static WebhookFilters PreviewFailed { get; } = new WebhookFilters("preview_failed");
        /// <summary>
        /// Trigger a webhook when a stack destroy succeeds.
        /// </summary>
        public static WebhookFilters DestroySucceeded { get; } = new WebhookFilters("destroy_succeeded");
        /// <summary>
        /// Trigger a webhook when a stack destroy fails.
        /// </summary>
        public static WebhookFilters DestroyFailed { get; } = new WebhookFilters("destroy_failed");
        /// <summary>
        /// Trigger a webhook when a stack refresh succeeds.
        /// </summary>
        public static WebhookFilters RefreshSucceeded { get; } = new WebhookFilters("refresh_succeeded");
        /// <summary>
        /// Trigger a webhook when a stack refresh fails.
        /// </summary>
        public static WebhookFilters RefreshFailed { get; } = new WebhookFilters("refresh_failed");
        /// <summary>
        /// Trigger a webhook when a deployment is queued.
        /// </summary>
        public static WebhookFilters DeploymentQueued { get; } = new WebhookFilters("deployment_queued");
        /// <summary>
        /// Trigger a webhook when a deployment starts running.
        /// </summary>
        public static WebhookFilters DeploymentStarted { get; } = new WebhookFilters("deployment_started");
        /// <summary>
        /// Trigger a webhook when a deployment succeeds.
        /// </summary>
        public static WebhookFilters DeploymentSucceeded { get; } = new WebhookFilters("deployment_succeeded");
        /// <summary>
        /// Trigger a webhook when a deployment fails.
        /// </summary>
        public static WebhookFilters DeploymentFailed { get; } = new WebhookFilters("deployment_failed");
        /// <summary>
        /// Trigger a webhook when drift is detected.
        /// </summary>
        public static WebhookFilters DriftDetected { get; } = new WebhookFilters("drift_detected");
        /// <summary>
        /// Trigger a webhook when a drift detection run succeeds, regardless of whether drift is detected.
        /// </summary>
        public static WebhookFilters DriftDetectionSucceeded { get; } = new WebhookFilters("drift_detection_succeeded");
        /// <summary>
        /// Trigger a webhook when a drift detection run fails.
        /// </summary>
        public static WebhookFilters DriftDetectionFailed { get; } = new WebhookFilters("drift_detection_failed");
        /// <summary>
        /// Trigger a webhook when a drift remediation run succeeds.
        /// </summary>
        public static WebhookFilters DriftRemediationSucceeded { get; } = new WebhookFilters("drift_remediation_succeeded");
        /// <summary>
        /// Trigger a webhook when a drift remediation run fails.
        /// </summary>
        public static WebhookFilters DriftRemediationFailed { get; } = new WebhookFilters("drift_remediation_failed");

        public static bool operator ==(WebhookFilters left, WebhookFilters right) => left.Equals(right);
        public static bool operator !=(WebhookFilters left, WebhookFilters right) => !left.Equals(right);

        public static explicit operator string(WebhookFilters value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WebhookFilters other && Equals(other);
        public bool Equals(WebhookFilters other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct WebhookFormat : IEquatable<WebhookFormat>
    {
        private readonly string _value;

        private WebhookFormat(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The default webhook format.
        /// </summary>
        public static WebhookFormat Raw { get; } = new WebhookFormat("raw");
        /// <summary>
        /// Messages formatted for consumption by Slack incoming webhooks.
        /// </summary>
        public static WebhookFormat Slack { get; } = new WebhookFormat("slack");
        /// <summary>
        /// Initiate deployments on a stack from a Pulumi Cloud webhook.
        /// </summary>
        public static WebhookFormat PulumiDeployments { get; } = new WebhookFormat("pulumi_deployments");
        /// <summary>
        /// Messages formatted for consumption by Microsoft Teams incoming webhooks.
        /// </summary>
        public static WebhookFormat MicrosoftTeams { get; } = new WebhookFormat("ms_teams");

        public static bool operator ==(WebhookFormat left, WebhookFormat right) => left.Equals(right);
        public static bool operator !=(WebhookFormat left, WebhookFormat right) => !left.Equals(right);

        public static explicit operator string(WebhookFormat value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WebhookFormat other && Equals(other);
        public bool Equals(WebhookFormat other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
