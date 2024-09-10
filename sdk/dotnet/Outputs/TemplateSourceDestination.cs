// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.PulumiService.Outputs
{

    [OutputType]
    public sealed class TemplateSourceDestination
    {
        /// <summary>
        /// Destination URL that gets filled in on new project creation.
        /// </summary>
        public readonly string? Url;

        [OutputConstructor]
        private TemplateSourceDestination(string? url)
        {
            Url = url;
        }
    }
}