// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cpln

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-cpln/sdk/go/cpln/internal"
)

type DomainRoute struct {
	pulumi.CustomResourceState

	DomainLink    pulumi.StringOutput    `pulumi:"domainLink"`
	DomainPort    pulumi.IntPtrOutput    `pulumi:"domainPort"`
	HostPrefix    pulumi.StringPtrOutput `pulumi:"hostPrefix"`
	Port          pulumi.IntPtrOutput    `pulumi:"port"`
	Prefix        pulumi.StringOutput    `pulumi:"prefix"`
	ReplacePrefix pulumi.StringPtrOutput `pulumi:"replacePrefix"`
	WorkloadLink  pulumi.StringOutput    `pulumi:"workloadLink"`
}

// NewDomainRoute registers a new resource with the given unique name, arguments, and options.
func NewDomainRoute(ctx *pulumi.Context,
	name string, args *DomainRouteArgs, opts ...pulumi.ResourceOption) (*DomainRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainLink == nil {
		return nil, errors.New("invalid value for required argument 'DomainLink'")
	}
	if args.Prefix == nil {
		return nil, errors.New("invalid value for required argument 'Prefix'")
	}
	if args.WorkloadLink == nil {
		return nil, errors.New("invalid value for required argument 'WorkloadLink'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DomainRoute
	err := ctx.RegisterResource("cpln:index/domainRoute:DomainRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainRoute gets an existing DomainRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainRouteState, opts ...pulumi.ResourceOption) (*DomainRoute, error) {
	var resource DomainRoute
	err := ctx.ReadResource("cpln:index/domainRoute:DomainRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainRoute resources.
type domainRouteState struct {
	DomainLink    *string `pulumi:"domainLink"`
	DomainPort    *int    `pulumi:"domainPort"`
	HostPrefix    *string `pulumi:"hostPrefix"`
	Port          *int    `pulumi:"port"`
	Prefix        *string `pulumi:"prefix"`
	ReplacePrefix *string `pulumi:"replacePrefix"`
	WorkloadLink  *string `pulumi:"workloadLink"`
}

type DomainRouteState struct {
	DomainLink    pulumi.StringPtrInput
	DomainPort    pulumi.IntPtrInput
	HostPrefix    pulumi.StringPtrInput
	Port          pulumi.IntPtrInput
	Prefix        pulumi.StringPtrInput
	ReplacePrefix pulumi.StringPtrInput
	WorkloadLink  pulumi.StringPtrInput
}

func (DomainRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainRouteState)(nil)).Elem()
}

type domainRouteArgs struct {
	DomainLink    string  `pulumi:"domainLink"`
	DomainPort    *int    `pulumi:"domainPort"`
	HostPrefix    *string `pulumi:"hostPrefix"`
	Port          *int    `pulumi:"port"`
	Prefix        string  `pulumi:"prefix"`
	ReplacePrefix *string `pulumi:"replacePrefix"`
	WorkloadLink  string  `pulumi:"workloadLink"`
}

// The set of arguments for constructing a DomainRoute resource.
type DomainRouteArgs struct {
	DomainLink    pulumi.StringInput
	DomainPort    pulumi.IntPtrInput
	HostPrefix    pulumi.StringPtrInput
	Port          pulumi.IntPtrInput
	Prefix        pulumi.StringInput
	ReplacePrefix pulumi.StringPtrInput
	WorkloadLink  pulumi.StringInput
}

func (DomainRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainRouteArgs)(nil)).Elem()
}

type DomainRouteInput interface {
	pulumi.Input

	ToDomainRouteOutput() DomainRouteOutput
	ToDomainRouteOutputWithContext(ctx context.Context) DomainRouteOutput
}

func (*DomainRoute) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainRoute)(nil)).Elem()
}

func (i *DomainRoute) ToDomainRouteOutput() DomainRouteOutput {
	return i.ToDomainRouteOutputWithContext(context.Background())
}

func (i *DomainRoute) ToDomainRouteOutputWithContext(ctx context.Context) DomainRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainRouteOutput)
}

func (i *DomainRoute) ToOutput(ctx context.Context) pulumix.Output[*DomainRoute] {
	return pulumix.Output[*DomainRoute]{
		OutputState: i.ToDomainRouteOutputWithContext(ctx).OutputState,
	}
}

// DomainRouteArrayInput is an input type that accepts DomainRouteArray and DomainRouteArrayOutput values.
// You can construct a concrete instance of `DomainRouteArrayInput` via:
//
//	DomainRouteArray{ DomainRouteArgs{...} }
type DomainRouteArrayInput interface {
	pulumi.Input

	ToDomainRouteArrayOutput() DomainRouteArrayOutput
	ToDomainRouteArrayOutputWithContext(context.Context) DomainRouteArrayOutput
}

type DomainRouteArray []DomainRouteInput

func (DomainRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainRoute)(nil)).Elem()
}

func (i DomainRouteArray) ToDomainRouteArrayOutput() DomainRouteArrayOutput {
	return i.ToDomainRouteArrayOutputWithContext(context.Background())
}

func (i DomainRouteArray) ToDomainRouteArrayOutputWithContext(ctx context.Context) DomainRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainRouteArrayOutput)
}

func (i DomainRouteArray) ToOutput(ctx context.Context) pulumix.Output[[]*DomainRoute] {
	return pulumix.Output[[]*DomainRoute]{
		OutputState: i.ToDomainRouteArrayOutputWithContext(ctx).OutputState,
	}
}

// DomainRouteMapInput is an input type that accepts DomainRouteMap and DomainRouteMapOutput values.
// You can construct a concrete instance of `DomainRouteMapInput` via:
//
//	DomainRouteMap{ "key": DomainRouteArgs{...} }
type DomainRouteMapInput interface {
	pulumi.Input

	ToDomainRouteMapOutput() DomainRouteMapOutput
	ToDomainRouteMapOutputWithContext(context.Context) DomainRouteMapOutput
}

type DomainRouteMap map[string]DomainRouteInput

func (DomainRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainRoute)(nil)).Elem()
}

func (i DomainRouteMap) ToDomainRouteMapOutput() DomainRouteMapOutput {
	return i.ToDomainRouteMapOutputWithContext(context.Background())
}

func (i DomainRouteMap) ToDomainRouteMapOutputWithContext(ctx context.Context) DomainRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainRouteMapOutput)
}

func (i DomainRouteMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DomainRoute] {
	return pulumix.Output[map[string]*DomainRoute]{
		OutputState: i.ToDomainRouteMapOutputWithContext(ctx).OutputState,
	}
}

type DomainRouteOutput struct{ *pulumi.OutputState }

func (DomainRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainRoute)(nil)).Elem()
}

func (o DomainRouteOutput) ToDomainRouteOutput() DomainRouteOutput {
	return o
}

func (o DomainRouteOutput) ToDomainRouteOutputWithContext(ctx context.Context) DomainRouteOutput {
	return o
}

func (o DomainRouteOutput) ToOutput(ctx context.Context) pulumix.Output[*DomainRoute] {
	return pulumix.Output[*DomainRoute]{
		OutputState: o.OutputState,
	}
}

func (o DomainRouteOutput) DomainLink() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainRoute) pulumi.StringOutput { return v.DomainLink }).(pulumi.StringOutput)
}

func (o DomainRouteOutput) DomainPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DomainRoute) pulumi.IntPtrOutput { return v.DomainPort }).(pulumi.IntPtrOutput)
}

func (o DomainRouteOutput) HostPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainRoute) pulumi.StringPtrOutput { return v.HostPrefix }).(pulumi.StringPtrOutput)
}

func (o DomainRouteOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DomainRoute) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

func (o DomainRouteOutput) Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainRoute) pulumi.StringOutput { return v.Prefix }).(pulumi.StringOutput)
}

func (o DomainRouteOutput) ReplacePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainRoute) pulumi.StringPtrOutput { return v.ReplacePrefix }).(pulumi.StringPtrOutput)
}

func (o DomainRouteOutput) WorkloadLink() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainRoute) pulumi.StringOutput { return v.WorkloadLink }).(pulumi.StringOutput)
}

type DomainRouteArrayOutput struct{ *pulumi.OutputState }

func (DomainRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainRoute)(nil)).Elem()
}

func (o DomainRouteArrayOutput) ToDomainRouteArrayOutput() DomainRouteArrayOutput {
	return o
}

func (o DomainRouteArrayOutput) ToDomainRouteArrayOutputWithContext(ctx context.Context) DomainRouteArrayOutput {
	return o
}

func (o DomainRouteArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DomainRoute] {
	return pulumix.Output[[]*DomainRoute]{
		OutputState: o.OutputState,
	}
}

func (o DomainRouteArrayOutput) Index(i pulumi.IntInput) DomainRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DomainRoute {
		return vs[0].([]*DomainRoute)[vs[1].(int)]
	}).(DomainRouteOutput)
}

type DomainRouteMapOutput struct{ *pulumi.OutputState }

func (DomainRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainRoute)(nil)).Elem()
}

func (o DomainRouteMapOutput) ToDomainRouteMapOutput() DomainRouteMapOutput {
	return o
}

func (o DomainRouteMapOutput) ToDomainRouteMapOutputWithContext(ctx context.Context) DomainRouteMapOutput {
	return o
}

func (o DomainRouteMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DomainRoute] {
	return pulumix.Output[map[string]*DomainRoute]{
		OutputState: o.OutputState,
	}
}

func (o DomainRouteMapOutput) MapIndex(k pulumi.StringInput) DomainRouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DomainRoute {
		return vs[0].(map[string]*DomainRoute)[vs[1].(string)]
	}).(DomainRouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainRouteInput)(nil)).Elem(), &DomainRoute{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainRouteArrayInput)(nil)).Elem(), DomainRouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainRouteMapInput)(nil)).Elem(), DomainRouteMap{})
	pulumi.RegisterOutputType(DomainRouteOutput{})
	pulumi.RegisterOutputType(DomainRouteArrayOutput{})
	pulumi.RegisterOutputType(DomainRouteMapOutput{})
}