// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Cpln.Inputs
{

    public sealed class OrgObservabilityArgs : global::Pulumi.ResourceArgs
    {
        [Input("logsRetentionDays")]
        public Input<int>? LogsRetentionDays { get; set; }

        [Input("metricsRetentionDays")]
        public Input<int>? MetricsRetentionDays { get; set; }

        [Input("tracesRetentionDays")]
        public Input<int>? TracesRetentionDays { get; set; }

        public OrgObservabilityArgs()
        {
        }
        public static new OrgObservabilityArgs Empty => new OrgObservabilityArgs();
    }
}