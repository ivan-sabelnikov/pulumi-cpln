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

    public sealed class GroupIdentityMatcherGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("language")]
        public Input<string>? Language { get; set; }

        public GroupIdentityMatcherGetArgs()
        {
        }
        public static new GroupIdentityMatcherGetArgs Empty => new GroupIdentityMatcherGetArgs();
    }
}