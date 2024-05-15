// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pulumiservice

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-pulumiservice/sdk/go/pulumiservice/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Agent Pool for customer managed deployments
type AgentPool struct {
	pulumi.CustomResourceState

	// The agent pool identifier.
	AgentPoolId pulumi.StringOutput `pulumi:"agentPoolId"`
	// Description of the agent pool.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the agent pool.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization's name.
	OrganizationName pulumi.StringOutput `pulumi:"organizationName"`
	// The agent pool's token's value.
	TokenValue pulumi.StringOutput `pulumi:"tokenValue"`
}

// NewAgentPool registers a new resource with the given unique name, arguments, and options.
func NewAgentPool(ctx *pulumi.Context,
	name string, args *AgentPoolArgs, opts ...pulumi.ResourceOption) (*AgentPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.OrganizationName == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationName'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tokenValue",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AgentPool
	err := ctx.RegisterResource("pulumiservice:index:AgentPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgentPool gets an existing AgentPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgentPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentPoolState, opts ...pulumi.ResourceOption) (*AgentPool, error) {
	var resource AgentPool
	err := ctx.ReadResource("pulumiservice:index:AgentPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AgentPool resources.
type agentPoolState struct {
}

type AgentPoolState struct {
}

func (AgentPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentPoolState)(nil)).Elem()
}

type agentPoolArgs struct {
	// Description of the agent pool.
	Description *string `pulumi:"description"`
	// Name of the agent pool.
	Name string `pulumi:"name"`
	// The organization's name.
	OrganizationName string `pulumi:"organizationName"`
}

// The set of arguments for constructing a AgentPool resource.
type AgentPoolArgs struct {
	// Description of the agent pool.
	Description pulumi.StringPtrInput
	// Name of the agent pool.
	Name pulumi.StringInput
	// The organization's name.
	OrganizationName pulumi.StringInput
}

func (AgentPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentPoolArgs)(nil)).Elem()
}

type AgentPoolInput interface {
	pulumi.Input

	ToAgentPoolOutput() AgentPoolOutput
	ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput
}

func (*AgentPool) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPool)(nil)).Elem()
}

func (i *AgentPool) ToAgentPoolOutput() AgentPoolOutput {
	return i.ToAgentPoolOutputWithContext(context.Background())
}

func (i *AgentPool) ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolOutput)
}

// AgentPoolArrayInput is an input type that accepts AgentPoolArray and AgentPoolArrayOutput values.
// You can construct a concrete instance of `AgentPoolArrayInput` via:
//
//	AgentPoolArray{ AgentPoolArgs{...} }
type AgentPoolArrayInput interface {
	pulumi.Input

	ToAgentPoolArrayOutput() AgentPoolArrayOutput
	ToAgentPoolArrayOutputWithContext(context.Context) AgentPoolArrayOutput
}

type AgentPoolArray []AgentPoolInput

func (AgentPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AgentPool)(nil)).Elem()
}

func (i AgentPoolArray) ToAgentPoolArrayOutput() AgentPoolArrayOutput {
	return i.ToAgentPoolArrayOutputWithContext(context.Background())
}

func (i AgentPoolArray) ToAgentPoolArrayOutputWithContext(ctx context.Context) AgentPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolArrayOutput)
}

// AgentPoolMapInput is an input type that accepts AgentPoolMap and AgentPoolMapOutput values.
// You can construct a concrete instance of `AgentPoolMapInput` via:
//
//	AgentPoolMap{ "key": AgentPoolArgs{...} }
type AgentPoolMapInput interface {
	pulumi.Input

	ToAgentPoolMapOutput() AgentPoolMapOutput
	ToAgentPoolMapOutputWithContext(context.Context) AgentPoolMapOutput
}

type AgentPoolMap map[string]AgentPoolInput

func (AgentPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AgentPool)(nil)).Elem()
}

func (i AgentPoolMap) ToAgentPoolMapOutput() AgentPoolMapOutput {
	return i.ToAgentPoolMapOutputWithContext(context.Background())
}

func (i AgentPoolMap) ToAgentPoolMapOutputWithContext(ctx context.Context) AgentPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolMapOutput)
}

type AgentPoolOutput struct{ *pulumi.OutputState }

func (AgentPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPool)(nil)).Elem()
}

func (o AgentPoolOutput) ToAgentPoolOutput() AgentPoolOutput {
	return o
}

func (o AgentPoolOutput) ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput {
	return o
}

// The agent pool identifier.
func (o AgentPoolOutput) AgentPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.AgentPoolId }).(pulumi.StringOutput)
}

// Description of the agent pool.
func (o AgentPoolOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the agent pool.
func (o AgentPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization's name.
func (o AgentPoolOutput) OrganizationName() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.OrganizationName }).(pulumi.StringOutput)
}

// The agent pool's token's value.
func (o AgentPoolOutput) TokenValue() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.TokenValue }).(pulumi.StringOutput)
}

type AgentPoolArrayOutput struct{ *pulumi.OutputState }

func (AgentPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AgentPool)(nil)).Elem()
}

func (o AgentPoolArrayOutput) ToAgentPoolArrayOutput() AgentPoolArrayOutput {
	return o
}

func (o AgentPoolArrayOutput) ToAgentPoolArrayOutputWithContext(ctx context.Context) AgentPoolArrayOutput {
	return o
}

func (o AgentPoolArrayOutput) Index(i pulumi.IntInput) AgentPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AgentPool {
		return vs[0].([]*AgentPool)[vs[1].(int)]
	}).(AgentPoolOutput)
}

type AgentPoolMapOutput struct{ *pulumi.OutputState }

func (AgentPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AgentPool)(nil)).Elem()
}

func (o AgentPoolMapOutput) ToAgentPoolMapOutput() AgentPoolMapOutput {
	return o
}

func (o AgentPoolMapOutput) ToAgentPoolMapOutputWithContext(ctx context.Context) AgentPoolMapOutput {
	return o
}

func (o AgentPoolMapOutput) MapIndex(k pulumi.StringInput) AgentPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AgentPool {
		return vs[0].(map[string]*AgentPool)[vs[1].(string)]
	}).(AgentPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AgentPoolInput)(nil)).Elem(), &AgentPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*AgentPoolArrayInput)(nil)).Elem(), AgentPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AgentPoolMapInput)(nil)).Elem(), AgentPoolMap{})
	pulumi.RegisterOutputType(AgentPoolOutput{})
	pulumi.RegisterOutputType(AgentPoolArrayOutput{})
	pulumi.RegisterOutputType(AgentPoolMapOutput{})
}
