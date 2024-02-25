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

type Workload struct {
	pulumi.CustomResourceState

	Containers         WorkloadContainerArrayOutput     `pulumi:"containers"`
	CplnId             pulumi.StringOutput              `pulumi:"cplnId"`
	Description        pulumi.StringPtrOutput           `pulumi:"description"`
	FirewallSpec       WorkloadFirewallSpecPtrOutput    `pulumi:"firewallSpec"`
	Gvc                pulumi.StringOutput              `pulumi:"gvc"`
	IdentityLink       pulumi.StringPtrOutput           `pulumi:"identityLink"`
	Job                WorkloadJobPtrOutput             `pulumi:"job"`
	LocalOptions       WorkloadLocalOptionArrayOutput   `pulumi:"localOptions"`
	Name               pulumi.StringOutput              `pulumi:"name"`
	Options            WorkloadOptionsOutput            `pulumi:"options"`
	RolloutOptions     WorkloadRolloutOptionsPtrOutput  `pulumi:"rolloutOptions"`
	SecurityOptions    WorkloadSecurityOptionsPtrOutput `pulumi:"securityOptions"`
	SelfLink           pulumi.StringOutput              `pulumi:"selfLink"`
	Sidecar            WorkloadSidecarPtrOutput         `pulumi:"sidecar"`
	Statuses           WorkloadStatusArrayOutput        `pulumi:"statuses"`
	SupportDynamicTags pulumi.BoolPtrOutput             `pulumi:"supportDynamicTags"`
	Tags               pulumi.StringMapOutput           `pulumi:"tags"`
	Type               pulumi.StringOutput              `pulumi:"type"`
}

// NewWorkload registers a new resource with the given unique name, arguments, and options.
func NewWorkload(ctx *pulumi.Context,
	name string, args *WorkloadArgs, opts ...pulumi.ResourceOption) (*Workload, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Containers == nil {
		return nil, errors.New("invalid value for required argument 'Containers'")
	}
	if args.Gvc == nil {
		return nil, errors.New("invalid value for required argument 'Gvc'")
	}
	if args.Options == nil {
		return nil, errors.New("invalid value for required argument 'Options'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Workload
	err := ctx.RegisterResource("cpln:index/workload:Workload", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkload gets an existing Workload resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkload(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadState, opts ...pulumi.ResourceOption) (*Workload, error) {
	var resource Workload
	err := ctx.ReadResource("cpln:index/workload:Workload", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workload resources.
type workloadState struct {
	Containers         []WorkloadContainer      `pulumi:"containers"`
	CplnId             *string                  `pulumi:"cplnId"`
	Description        *string                  `pulumi:"description"`
	FirewallSpec       *WorkloadFirewallSpec    `pulumi:"firewallSpec"`
	Gvc                *string                  `pulumi:"gvc"`
	IdentityLink       *string                  `pulumi:"identityLink"`
	Job                *WorkloadJob             `pulumi:"job"`
	LocalOptions       []WorkloadLocalOption    `pulumi:"localOptions"`
	Name               *string                  `pulumi:"name"`
	Options            *WorkloadOptions         `pulumi:"options"`
	RolloutOptions     *WorkloadRolloutOptions  `pulumi:"rolloutOptions"`
	SecurityOptions    *WorkloadSecurityOptions `pulumi:"securityOptions"`
	SelfLink           *string                  `pulumi:"selfLink"`
	Sidecar            *WorkloadSidecar         `pulumi:"sidecar"`
	Statuses           []WorkloadStatus         `pulumi:"statuses"`
	SupportDynamicTags *bool                    `pulumi:"supportDynamicTags"`
	Tags               map[string]string        `pulumi:"tags"`
	Type               *string                  `pulumi:"type"`
}

type WorkloadState struct {
	Containers         WorkloadContainerArrayInput
	CplnId             pulumi.StringPtrInput
	Description        pulumi.StringPtrInput
	FirewallSpec       WorkloadFirewallSpecPtrInput
	Gvc                pulumi.StringPtrInput
	IdentityLink       pulumi.StringPtrInput
	Job                WorkloadJobPtrInput
	LocalOptions       WorkloadLocalOptionArrayInput
	Name               pulumi.StringPtrInput
	Options            WorkloadOptionsPtrInput
	RolloutOptions     WorkloadRolloutOptionsPtrInput
	SecurityOptions    WorkloadSecurityOptionsPtrInput
	SelfLink           pulumi.StringPtrInput
	Sidecar            WorkloadSidecarPtrInput
	Statuses           WorkloadStatusArrayInput
	SupportDynamicTags pulumi.BoolPtrInput
	Tags               pulumi.StringMapInput
	Type               pulumi.StringPtrInput
}

func (WorkloadState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadState)(nil)).Elem()
}

type workloadArgs struct {
	Containers         []WorkloadContainer      `pulumi:"containers"`
	Description        *string                  `pulumi:"description"`
	FirewallSpec       *WorkloadFirewallSpec    `pulumi:"firewallSpec"`
	Gvc                string                   `pulumi:"gvc"`
	IdentityLink       *string                  `pulumi:"identityLink"`
	Job                *WorkloadJob             `pulumi:"job"`
	LocalOptions       []WorkloadLocalOption    `pulumi:"localOptions"`
	Name               *string                  `pulumi:"name"`
	Options            WorkloadOptions          `pulumi:"options"`
	RolloutOptions     *WorkloadRolloutOptions  `pulumi:"rolloutOptions"`
	SecurityOptions    *WorkloadSecurityOptions `pulumi:"securityOptions"`
	Sidecar            *WorkloadSidecar         `pulumi:"sidecar"`
	SupportDynamicTags *bool                    `pulumi:"supportDynamicTags"`
	Tags               map[string]string        `pulumi:"tags"`
	Type               string                   `pulumi:"type"`
}

// The set of arguments for constructing a Workload resource.
type WorkloadArgs struct {
	Containers         WorkloadContainerArrayInput
	Description        pulumi.StringPtrInput
	FirewallSpec       WorkloadFirewallSpecPtrInput
	Gvc                pulumi.StringInput
	IdentityLink       pulumi.StringPtrInput
	Job                WorkloadJobPtrInput
	LocalOptions       WorkloadLocalOptionArrayInput
	Name               pulumi.StringPtrInput
	Options            WorkloadOptionsInput
	RolloutOptions     WorkloadRolloutOptionsPtrInput
	SecurityOptions    WorkloadSecurityOptionsPtrInput
	Sidecar            WorkloadSidecarPtrInput
	SupportDynamicTags pulumi.BoolPtrInput
	Tags               pulumi.StringMapInput
	Type               pulumi.StringInput
}

func (WorkloadArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadArgs)(nil)).Elem()
}

type WorkloadInput interface {
	pulumi.Input

	ToWorkloadOutput() WorkloadOutput
	ToWorkloadOutputWithContext(ctx context.Context) WorkloadOutput
}

func (*Workload) ElementType() reflect.Type {
	return reflect.TypeOf((**Workload)(nil)).Elem()
}

func (i *Workload) ToWorkloadOutput() WorkloadOutput {
	return i.ToWorkloadOutputWithContext(context.Background())
}

func (i *Workload) ToWorkloadOutputWithContext(ctx context.Context) WorkloadOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadOutput)
}

func (i *Workload) ToOutput(ctx context.Context) pulumix.Output[*Workload] {
	return pulumix.Output[*Workload]{
		OutputState: i.ToWorkloadOutputWithContext(ctx).OutputState,
	}
}

// WorkloadArrayInput is an input type that accepts WorkloadArray and WorkloadArrayOutput values.
// You can construct a concrete instance of `WorkloadArrayInput` via:
//
//	WorkloadArray{ WorkloadArgs{...} }
type WorkloadArrayInput interface {
	pulumi.Input

	ToWorkloadArrayOutput() WorkloadArrayOutput
	ToWorkloadArrayOutputWithContext(context.Context) WorkloadArrayOutput
}

type WorkloadArray []WorkloadInput

func (WorkloadArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Workload)(nil)).Elem()
}

func (i WorkloadArray) ToWorkloadArrayOutput() WorkloadArrayOutput {
	return i.ToWorkloadArrayOutputWithContext(context.Background())
}

func (i WorkloadArray) ToWorkloadArrayOutputWithContext(ctx context.Context) WorkloadArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadArrayOutput)
}

func (i WorkloadArray) ToOutput(ctx context.Context) pulumix.Output[[]*Workload] {
	return pulumix.Output[[]*Workload]{
		OutputState: i.ToWorkloadArrayOutputWithContext(ctx).OutputState,
	}
}

// WorkloadMapInput is an input type that accepts WorkloadMap and WorkloadMapOutput values.
// You can construct a concrete instance of `WorkloadMapInput` via:
//
//	WorkloadMap{ "key": WorkloadArgs{...} }
type WorkloadMapInput interface {
	pulumi.Input

	ToWorkloadMapOutput() WorkloadMapOutput
	ToWorkloadMapOutputWithContext(context.Context) WorkloadMapOutput
}

type WorkloadMap map[string]WorkloadInput

func (WorkloadMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Workload)(nil)).Elem()
}

func (i WorkloadMap) ToWorkloadMapOutput() WorkloadMapOutput {
	return i.ToWorkloadMapOutputWithContext(context.Background())
}

func (i WorkloadMap) ToWorkloadMapOutputWithContext(ctx context.Context) WorkloadMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadMapOutput)
}

func (i WorkloadMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Workload] {
	return pulumix.Output[map[string]*Workload]{
		OutputState: i.ToWorkloadMapOutputWithContext(ctx).OutputState,
	}
}

type WorkloadOutput struct{ *pulumi.OutputState }

func (WorkloadOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workload)(nil)).Elem()
}

func (o WorkloadOutput) ToWorkloadOutput() WorkloadOutput {
	return o
}

func (o WorkloadOutput) ToWorkloadOutputWithContext(ctx context.Context) WorkloadOutput {
	return o
}

func (o WorkloadOutput) ToOutput(ctx context.Context) pulumix.Output[*Workload] {
	return pulumix.Output[*Workload]{
		OutputState: o.OutputState,
	}
}

func (o WorkloadOutput) Containers() WorkloadContainerArrayOutput {
	return o.ApplyT(func(v *Workload) WorkloadContainerArrayOutput { return v.Containers }).(WorkloadContainerArrayOutput)
}

func (o WorkloadOutput) CplnId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workload) pulumi.StringOutput { return v.CplnId }).(pulumi.StringOutput)
}

func (o WorkloadOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workload) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o WorkloadOutput) FirewallSpec() WorkloadFirewallSpecPtrOutput {
	return o.ApplyT(func(v *Workload) WorkloadFirewallSpecPtrOutput { return v.FirewallSpec }).(WorkloadFirewallSpecPtrOutput)
}

func (o WorkloadOutput) Gvc() pulumi.StringOutput {
	return o.ApplyT(func(v *Workload) pulumi.StringOutput { return v.Gvc }).(pulumi.StringOutput)
}

func (o WorkloadOutput) IdentityLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workload) pulumi.StringPtrOutput { return v.IdentityLink }).(pulumi.StringPtrOutput)
}

func (o WorkloadOutput) Job() WorkloadJobPtrOutput {
	return o.ApplyT(func(v *Workload) WorkloadJobPtrOutput { return v.Job }).(WorkloadJobPtrOutput)
}

func (o WorkloadOutput) LocalOptions() WorkloadLocalOptionArrayOutput {
	return o.ApplyT(func(v *Workload) WorkloadLocalOptionArrayOutput { return v.LocalOptions }).(WorkloadLocalOptionArrayOutput)
}

func (o WorkloadOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workload) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkloadOutput) Options() WorkloadOptionsOutput {
	return o.ApplyT(func(v *Workload) WorkloadOptionsOutput { return v.Options }).(WorkloadOptionsOutput)
}

func (o WorkloadOutput) RolloutOptions() WorkloadRolloutOptionsPtrOutput {
	return o.ApplyT(func(v *Workload) WorkloadRolloutOptionsPtrOutput { return v.RolloutOptions }).(WorkloadRolloutOptionsPtrOutput)
}

func (o WorkloadOutput) SecurityOptions() WorkloadSecurityOptionsPtrOutput {
	return o.ApplyT(func(v *Workload) WorkloadSecurityOptionsPtrOutput { return v.SecurityOptions }).(WorkloadSecurityOptionsPtrOutput)
}

func (o WorkloadOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Workload) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

func (o WorkloadOutput) Sidecar() WorkloadSidecarPtrOutput {
	return o.ApplyT(func(v *Workload) WorkloadSidecarPtrOutput { return v.Sidecar }).(WorkloadSidecarPtrOutput)
}

func (o WorkloadOutput) Statuses() WorkloadStatusArrayOutput {
	return o.ApplyT(func(v *Workload) WorkloadStatusArrayOutput { return v.Statuses }).(WorkloadStatusArrayOutput)
}

func (o WorkloadOutput) SupportDynamicTags() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Workload) pulumi.BoolPtrOutput { return v.SupportDynamicTags }).(pulumi.BoolPtrOutput)
}

func (o WorkloadOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workload) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o WorkloadOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Workload) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type WorkloadArrayOutput struct{ *pulumi.OutputState }

func (WorkloadArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Workload)(nil)).Elem()
}

func (o WorkloadArrayOutput) ToWorkloadArrayOutput() WorkloadArrayOutput {
	return o
}

func (o WorkloadArrayOutput) ToWorkloadArrayOutputWithContext(ctx context.Context) WorkloadArrayOutput {
	return o
}

func (o WorkloadArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Workload] {
	return pulumix.Output[[]*Workload]{
		OutputState: o.OutputState,
	}
}

func (o WorkloadArrayOutput) Index(i pulumi.IntInput) WorkloadOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Workload {
		return vs[0].([]*Workload)[vs[1].(int)]
	}).(WorkloadOutput)
}

type WorkloadMapOutput struct{ *pulumi.OutputState }

func (WorkloadMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Workload)(nil)).Elem()
}

func (o WorkloadMapOutput) ToWorkloadMapOutput() WorkloadMapOutput {
	return o
}

func (o WorkloadMapOutput) ToWorkloadMapOutputWithContext(ctx context.Context) WorkloadMapOutput {
	return o
}

func (o WorkloadMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Workload] {
	return pulumix.Output[map[string]*Workload]{
		OutputState: o.OutputState,
	}
}

func (o WorkloadMapOutput) MapIndex(k pulumi.StringInput) WorkloadOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Workload {
		return vs[0].(map[string]*Workload)[vs[1].(string)]
	}).(WorkloadOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadInput)(nil)).Elem(), &Workload{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadArrayInput)(nil)).Elem(), WorkloadArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadMapInput)(nil)).Elem(), WorkloadMap{})
	pulumi.RegisterOutputType(WorkloadOutput{})
	pulumi.RegisterOutputType(WorkloadArrayOutput{})
	pulumi.RegisterOutputType(WorkloadMapOutput{})
}