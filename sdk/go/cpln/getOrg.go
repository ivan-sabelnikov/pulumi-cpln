// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cpln

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-cpln/sdk/go/cpln/internal"
)

// Output the ID and name of the current [org](https://docs.controlplane.com/reference/org).
//
// ## Outputs
//
// The following attributes are exported:
//
// - **cpln_id** (String) The ID, in GUID format, of the Org.
// - **name** (String) The name of Org.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-cpln/sdk/go/cpln"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			org, err := cpln.LookupOrg(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("orgId", org.Id)
//			ctx.Export("orgName", org.Name)
//			return nil
//		})
//	}
//
// ```
func LookupOrg(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*LookupOrgResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOrgResult
	err := ctx.Invoke("cpln:index/getOrg:getOrg", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getOrg.
type LookupOrgResult struct {
	CplnId string `pulumi:"cplnId"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}

func LookupOrgOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) LookupOrgResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (LookupOrgResult, error) {
		r, err := LookupOrg(ctx, opts...)
		var s LookupOrgResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(LookupOrgResultOutput)
}

// A collection of values returned by getOrg.
type LookupOrgResultOutput struct{ *pulumi.OutputState }

func (LookupOrgResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgResult)(nil)).Elem()
}

func (o LookupOrgResultOutput) ToLookupOrgResultOutput() LookupOrgResultOutput {
	return o
}

func (o LookupOrgResultOutput) ToLookupOrgResultOutputWithContext(ctx context.Context) LookupOrgResultOutput {
	return o
}

func (o LookupOrgResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupOrgResult] {
	return pulumix.Output[LookupOrgResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupOrgResultOutput) CplnId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgResult) string { return v.CplnId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupOrgResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOrgResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrgResultOutput{})
}