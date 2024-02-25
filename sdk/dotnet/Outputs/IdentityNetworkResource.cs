// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Cpln.Outputs
{

    [OutputType]
    public sealed class IdentityNetworkResource
    {
        public readonly string AgentLink;
        public readonly string? Fqdn;
        public readonly ImmutableArray<string> Ips;
        public readonly string Name;
        public readonly ImmutableArray<int> Ports;
        public readonly string? ResolverIp;

        [OutputConstructor]
        private IdentityNetworkResource(
            string agentLink,

            string? fqdn,

            ImmutableArray<string> ips,

            string name,

            ImmutableArray<int> ports,

            string? resolverIp)
        {
            AgentLink = agentLink;
            Fqdn = fqdn;
            Ips = ips;
            Name = name;
            Ports = ports;
            ResolverIp = resolverIp;
        }
    }
}