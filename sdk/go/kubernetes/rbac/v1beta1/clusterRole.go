// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package v1beta1

import (
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ClusterRole is a cluster level, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding or ClusterRoleBinding. Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 ClusterRole, and will no longer be served in v1.20.
type ClusterRole struct {
	pulumi.CustomResourceState

	// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
	AggregationRule AggregationRulePtrOutput `pulumi:"aggregationRule"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Rules holds all the PolicyRules for this ClusterRole
	Rules PolicyRuleArrayOutput `pulumi:"rules"`
}

// NewClusterRole registers a new resource with the given unique name, arguments, and options.
func NewClusterRole(ctx *pulumi.Context,
	name string, args *ClusterRoleArgs, opts ...pulumi.ResourceOption) (*ClusterRole, error) {
	if args == nil {
		args = &ClusterRoleArgs{}
	}
	if args.ApiVersion == nil {
		args.ApiVersion = pulumi.StringPtr("rbac.authorization.k8s.io/v1beta1")
	}
	if args.Kind == nil {
		args.Kind = pulumi.StringPtr("ClusterRole")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:rbac.authorization.k8s.io/v1:ClusterRole"),
		},
		{
			Type: pulumi.String("kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRole"),
		},
	})
	opts = append(opts, aliases)
	var resource ClusterRole
	err := ctx.RegisterResource("kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterRole gets an existing ClusterRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterRoleState, opts ...pulumi.ResourceOption) (*ClusterRole, error) {
	var resource ClusterRole
	err := ctx.ReadResource("kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterRole resources.
type clusterRoleState struct {
	// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
	AggregationRule *AggregationRule `pulumi:"aggregationRule"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata.
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Rules holds all the PolicyRules for this ClusterRole
	Rules []PolicyRule `pulumi:"rules"`
}

type ClusterRoleState struct {
	// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
	AggregationRule AggregationRulePtrInput
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPtrInput
	// Rules holds all the PolicyRules for this ClusterRole
	Rules PolicyRuleArrayInput
}

func (ClusterRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterRoleState)(nil)).Elem()
}

type clusterRoleArgs struct {
	// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
	AggregationRule *AggregationRule `pulumi:"aggregationRule"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata.
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Rules holds all the PolicyRules for this ClusterRole
	Rules []PolicyRule `pulumi:"rules"`
}

// The set of arguments for constructing a ClusterRole resource.
type ClusterRoleArgs struct {
	// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
	AggregationRule AggregationRulePtrInput
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPtrInput
	// Rules holds all the PolicyRules for this ClusterRole
	Rules PolicyRuleArrayInput
}

func (ClusterRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterRoleArgs)(nil)).Elem()
}
