// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about an existing [Global Virtual Cloud (GVC)](https://docs.controlplane.com/reference/gvc) within Control Plane.
 *
 * ## Required
 *
 * - **name** (String) Name of the GVC.
 *
 * ## Outputs
 *
 * The following attributes are exported:
 *
 * - **cpln_id** (String) The ID, in GUID format, of the GVC.
 * - **name** (String) Name of the GVC.
 * - **locations** (List of String) A list of [locations](https://docs.controlplane.com/reference/location#current) making up the Global Virtual Cloud.
 * - **description** (String) Description of the GVC.
 * - **domain** (String) Custom domain name used by associated workloads.
 * - **pull_secrets** (List of String) A list of [pull secret](https://docs.controlplane.com/reference/gvc#pull-secrets) names used to authenticate to any private image repository referenced by Workloads within the GVC.
 * - **tags** (Map of String) Key-value map of resource tags.
 * - **lightstep_tracing** (Block List, Max: 1) (see below).
 * - **self_link** (String) Full link to this resource. Can be referenced by other resources.
 *
 * <a id="nestedblock--lightstep_tracing"></a>
 * ### `lightstepTracing`
 *
 * - **sampling** (Int) Sampling percentage.
 * - **endpoint** (String) Tracing Endpoint Workload. Either the canonical endpoint or the internal endpoint.
 * - **credentials** (String) Full link to referenced Opaque Secret.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cpln from "@pulumi/cpln";
 *
 * const gvc = cpln.getGvc({});
 * export const gvcId = gvc.then(gvc => gvc.id);
 * export const gvcLocations = gvc.then(gvc => gvc.locations);
 * ```
 */
export function getGvc(args: GetGvcArgs, opts?: pulumi.InvokeOptions): Promise<GetGvcResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("cpln:index/getGvc:getGvc", {
        "controlplaneTracing": args.controlplaneTracing,
        "description": args.description,
        "domain": args.domain,
        "env": args.env,
        "lightstepTracing": args.lightstepTracing,
        "loadBalancer": args.loadBalancer,
        "locations": args.locations,
        "name": args.name,
        "otelTracing": args.otelTracing,
        "pullSecrets": args.pullSecrets,
        "sidecar": args.sidecar,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getGvc.
 */
export interface GetGvcArgs {
    controlplaneTracing?: inputs.GetGvcControlplaneTracing;
    description?: string;
    /**
     * @deprecated Selecting a domain on a GVC will be deprecated in the future. Use the 'cpln_domain resource' instead.
     */
    domain?: string;
    env?: {[key: string]: string};
    lightstepTracing?: inputs.GetGvcLightstepTracing;
    loadBalancer?: inputs.GetGvcLoadBalancer;
    locations?: string[];
    name: string;
    otelTracing?: inputs.GetGvcOtelTracing;
    pullSecrets?: string[];
    sidecar?: inputs.GetGvcSidecar;
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getGvc.
 */
export interface GetGvcResult {
    readonly alias: string;
    readonly controlplaneTracing?: outputs.GetGvcControlplaneTracing;
    readonly cplnId: string;
    readonly description?: string;
    /**
     * @deprecated Selecting a domain on a GVC will be deprecated in the future. Use the 'cpln_domain resource' instead.
     */
    readonly domain?: string;
    readonly env?: {[key: string]: string};
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lightstepTracing?: outputs.GetGvcLightstepTracing;
    readonly loadBalancer?: outputs.GetGvcLoadBalancer;
    readonly locations?: string[];
    readonly name: string;
    readonly otelTracing?: outputs.GetGvcOtelTracing;
    readonly pullSecrets?: string[];
    readonly selfLink: string;
    readonly sidecar?: outputs.GetGvcSidecar;
    readonly tags?: {[key: string]: string};
}
/**
 * Use this data source to access information about an existing [Global Virtual Cloud (GVC)](https://docs.controlplane.com/reference/gvc) within Control Plane.
 *
 * ## Required
 *
 * - **name** (String) Name of the GVC.
 *
 * ## Outputs
 *
 * The following attributes are exported:
 *
 * - **cpln_id** (String) The ID, in GUID format, of the GVC.
 * - **name** (String) Name of the GVC.
 * - **locations** (List of String) A list of [locations](https://docs.controlplane.com/reference/location#current) making up the Global Virtual Cloud.
 * - **description** (String) Description of the GVC.
 * - **domain** (String) Custom domain name used by associated workloads.
 * - **pull_secrets** (List of String) A list of [pull secret](https://docs.controlplane.com/reference/gvc#pull-secrets) names used to authenticate to any private image repository referenced by Workloads within the GVC.
 * - **tags** (Map of String) Key-value map of resource tags.
 * - **lightstep_tracing** (Block List, Max: 1) (see below).
 * - **self_link** (String) Full link to this resource. Can be referenced by other resources.
 *
 * <a id="nestedblock--lightstep_tracing"></a>
 * ### `lightstepTracing`
 *
 * - **sampling** (Int) Sampling percentage.
 * - **endpoint** (String) Tracing Endpoint Workload. Either the canonical endpoint or the internal endpoint.
 * - **credentials** (String) Full link to referenced Opaque Secret.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cpln from "@pulumi/cpln";
 *
 * const gvc = cpln.getGvc({});
 * export const gvcId = gvc.then(gvc => gvc.id);
 * export const gvcLocations = gvc.then(gvc => gvc.locations);
 * ```
 */
export function getGvcOutput(args: GetGvcOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGvcResult> {
    return pulumi.output(args).apply((a: any) => getGvc(a, opts))
}

/**
 * A collection of arguments for invoking getGvc.
 */
export interface GetGvcOutputArgs {
    controlplaneTracing?: pulumi.Input<inputs.GetGvcControlplaneTracingArgs>;
    description?: pulumi.Input<string>;
    /**
     * @deprecated Selecting a domain on a GVC will be deprecated in the future. Use the 'cpln_domain resource' instead.
     */
    domain?: pulumi.Input<string>;
    env?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    lightstepTracing?: pulumi.Input<inputs.GetGvcLightstepTracingArgs>;
    loadBalancer?: pulumi.Input<inputs.GetGvcLoadBalancerArgs>;
    locations?: pulumi.Input<pulumi.Input<string>[]>;
    name: pulumi.Input<string>;
    otelTracing?: pulumi.Input<inputs.GetGvcOtelTracingArgs>;
    pullSecrets?: pulumi.Input<pulumi.Input<string>[]>;
    sidecar?: pulumi.Input<inputs.GetGvcSidecarArgs>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}