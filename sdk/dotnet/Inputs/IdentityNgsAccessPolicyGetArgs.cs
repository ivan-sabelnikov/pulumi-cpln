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

    public sealed class IdentityNgsAccessPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("cloudAccountLink", required: true)]
        public Input<string> CloudAccountLink { get; set; } = null!;

        [Input("data")]
        public Input<int>? Data { get; set; }

        [Input("payload")]
        public Input<int>? Payload { get; set; }

        [Input("pub")]
        public Input<Inputs.IdentityNgsAccessPolicyPubGetArgs>? Pub { get; set; }

        [Input("resp")]
        public Input<Inputs.IdentityNgsAccessPolicyRespGetArgs>? Resp { get; set; }

        [Input("sub")]
        public Input<Inputs.IdentityNgsAccessPolicySubGetArgs>? Sub { get; set; }

        [Input("subs")]
        public Input<int>? Subs { get; set; }

        public IdentityNgsAccessPolicyGetArgs()
        {
        }
        public static new IdentityNgsAccessPolicyGetArgs Empty => new IdentityNgsAccessPolicyGetArgs();
    }
}