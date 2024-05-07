// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.PulumiService
{
    /// <summary>
    /// A cron schedule to run drift detection.
    /// </summary>
    [PulumiServiceResourceType("pulumiservice:index:DriftSchedule")]
    public partial class DriftSchedule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether any drift detected should be remediated after a drift run.
        /// </summary>
        [Output("autoRemediate")]
        public Output<bool?> AutoRemediate { get; private set; } = null!;

        /// <summary>
        /// Organization name.
        /// </summary>
        [Output("organization")]
        public Output<string> Organization { get; private set; } = null!;

        /// <summary>
        /// Project name.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Cron expression for when to run drift detection.
        /// </summary>
        [Output("scheduleCron")]
        public Output<string> ScheduleCron { get; private set; } = null!;

        /// <summary>
        /// Schedule ID of the created schedule, assigned by Pulumi Cloud.
        /// </summary>
        [Output("scheduleId")]
        public Output<string> ScheduleId { get; private set; } = null!;

        /// <summary>
        /// Stack name.
        /// </summary>
        [Output("stack")]
        public Output<string> Stack { get; private set; } = null!;


        /// <summary>
        /// Create a DriftSchedule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DriftSchedule(string name, DriftScheduleArgs args, CustomResourceOptions? options = null)
            : base("pulumiservice:index:DriftSchedule", name, args ?? new DriftScheduleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DriftSchedule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("pulumiservice:index:DriftSchedule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DriftSchedule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DriftSchedule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DriftSchedule(name, id, options);
        }
    }

    public sealed class DriftScheduleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether any drift detected should be remediated after a drift run.
        /// </summary>
        [Input("autoRemediate")]
        public Input<bool>? AutoRemediate { get; set; }

        /// <summary>
        /// Organization name.
        /// </summary>
        [Input("organization", required: true)]
        public Input<string> Organization { get; set; } = null!;

        /// <summary>
        /// Project name.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Cron expression for when to run drift detection.
        /// </summary>
        [Input("scheduleCron", required: true)]
        public Input<string> ScheduleCron { get; set; } = null!;

        /// <summary>
        /// Stack name.
        /// </summary>
        [Input("stack", required: true)]
        public Input<string> Stack { get; set; } = null!;

        public DriftScheduleArgs()
        {
            AutoRemediate = false;
        }
        public static new DriftScheduleArgs Empty => new DriftScheduleArgs();
    }
}