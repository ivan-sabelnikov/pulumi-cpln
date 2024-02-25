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

    public sealed class PolicyTargetQuerySpecArgs : global::Pulumi.ResourceArgs
    {
        [Input("match")]
        public Input<string>? Match { get; set; }

        [Input("terms")]
        private InputList<Inputs.PolicyTargetQuerySpecTermArgs>? _terms;
        public InputList<Inputs.PolicyTargetQuerySpecTermArgs> Terms
        {
            get => _terms ?? (_terms = new InputList<Inputs.PolicyTargetQuerySpecTermArgs>());
            set => _terms = value;
        }

        public PolicyTargetQuerySpecArgs()
        {
        }
        public static new PolicyTargetQuerySpecArgs Empty => new PolicyTargetQuerySpecArgs();
    }
}