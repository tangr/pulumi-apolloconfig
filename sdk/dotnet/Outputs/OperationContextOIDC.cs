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
    public sealed class OperationContextOIDC
    {
        /// <summary>
        /// AWS-specific OIDC configuration.
        /// </summary>
        public readonly Outputs.AWSOIDCConfiguration? Aws;
        /// <summary>
        /// Azure-specific OIDC configuration.
        /// </summary>
        public readonly Outputs.AzureOIDCConfiguration? Azure;
        /// <summary>
        /// GCP-specific OIDC configuration.
        /// </summary>
        public readonly Outputs.GCPOIDCConfiguration? Gcp;

        [OutputConstructor]
        private OperationContextOIDC(
            Outputs.AWSOIDCConfiguration? aws,

            Outputs.AzureOIDCConfiguration? azure,

            Outputs.GCPOIDCConfiguration? gcp)
        {
            Aws = aws;
            Azure = azure;
            Gcp = gcp;
        }
    }
}
