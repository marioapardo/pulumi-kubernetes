// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package v1beta1

import (
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// APIService represents a server for a particular GroupVersion. Name must be "version.group".
type APIService struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec contains information for locating and communicating with a server
	Spec APIServiceSpecPtrOutput `pulumi:"spec"`
	// Status contains derived information about an API server
	Status APIServiceStatusPtrOutput `pulumi:"status"`
}

// NewAPIService registers a new resource with the given unique name, arguments, and options.
func NewAPIService(ctx *pulumi.Context,
	name string, args *APIServiceArgs, opts ...pulumi.ResourceOption) (*APIService, error) {
	if args == nil {
		args = &APIServiceArgs{}
	}
	if args.ApiVersion == nil {
		args.ApiVersion = pulumi.StringPtr("apiregistration.k8s.io/v1beta1")
	}
	if args.Kind == nil {
		args.Kind = pulumi.StringPtr("APIService")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:apiregistration.k8s.io/v1:APIService"),
		},
		{
			Type: pulumi.String("kubernetes:apiregistration/v1:APIService"),
		},
		{
			Type: pulumi.String("kubernetes:apiregistration/v1beta1:APIService"),
		},
	})
	opts = append(opts, aliases)
	var resource APIService
	err := ctx.RegisterResource("kubernetes:apiregistration.k8s.io/v1beta1:APIService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAPIService gets an existing APIService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAPIService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *APIServiceState, opts ...pulumi.ResourceOption) (*APIService, error) {
	var resource APIService
	err := ctx.ReadResource("kubernetes:apiregistration.k8s.io/v1beta1:APIService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering APIService resources.
type apiserviceState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec contains information for locating and communicating with a server
	Spec *APIServiceSpec `pulumi:"spec"`
	// Status contains derived information about an API server
	Status *APIServiceStatus `pulumi:"status"`
}

type APIServiceState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Spec contains information for locating and communicating with a server
	Spec APIServiceSpecPtrInput
	// Status contains derived information about an API server
	Status APIServiceStatusPtrInput
}

func (APIServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiserviceState)(nil)).Elem()
}

type apiserviceArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec contains information for locating and communicating with a server
	Spec *APIServiceSpec `pulumi:"spec"`
}

// The set of arguments for constructing a APIService resource.
type APIServiceArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Spec contains information for locating and communicating with a server
	Spec APIServiceSpecPtrInput
}

func (APIServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiserviceArgs)(nil)).Elem()
}
