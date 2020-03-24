// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package v1

import (
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// SubjectAccessReview checks whether or not a user or group can perform an action.
type SubjectAccessReview struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec holds information about the request being evaluated
	Spec SubjectAccessReviewSpecPtrOutput `pulumi:"spec"`
	// Status is filled in by the server and indicates whether the request is allowed or not
	Status SubjectAccessReviewStatusPtrOutput `pulumi:"status"`
}

// NewSubjectAccessReview registers a new resource with the given unique name, arguments, and options.
func NewSubjectAccessReview(ctx *pulumi.Context,
	name string, args *SubjectAccessReviewArgs, opts ...pulumi.ResourceOption) (*SubjectAccessReview, error) {
	if args == nil || args.Spec == nil {
		return nil, errors.New("missing required argument 'Spec'")
	}
	if args == nil {
		args = &SubjectAccessReviewArgs{}
	}
	if args.ApiVersion == nil {
		args.ApiVersion = pulumi.StringPtr("authorization.k8s.io/v1")
	}
	if args.Kind == nil {
		args.Kind = pulumi.StringPtr("SubjectAccessReview")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:authorization.k8s.io/v1beta1:SubjectAccessReview"),
		},
	})
	opts = append(opts, aliases)
	var resource SubjectAccessReview
	err := ctx.RegisterResource("kubernetes:authorization.k8s.io/v1:SubjectAccessReview", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubjectAccessReview gets an existing SubjectAccessReview resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubjectAccessReview(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubjectAccessReviewState, opts ...pulumi.ResourceOption) (*SubjectAccessReview, error) {
	var resource SubjectAccessReview
	err := ctx.ReadResource("kubernetes:authorization.k8s.io/v1:SubjectAccessReview", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubjectAccessReview resources.
type subjectAccessReviewState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec holds information about the request being evaluated
	Spec *SubjectAccessReviewSpec `pulumi:"spec"`
	// Status is filled in by the server and indicates whether the request is allowed or not
	Status *SubjectAccessReviewStatus `pulumi:"status"`
}

type SubjectAccessReviewState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Spec holds information about the request being evaluated
	Spec SubjectAccessReviewSpecPtrInput
	// Status is filled in by the server and indicates whether the request is allowed or not
	Status SubjectAccessReviewStatusPtrInput
}

func (SubjectAccessReviewState) ElementType() reflect.Type {
	return reflect.TypeOf((*subjectAccessReviewState)(nil)).Elem()
}

type subjectAccessReviewArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec holds information about the request being evaluated
	Spec SubjectAccessReviewSpec `pulumi:"spec"`
}

// The set of arguments for constructing a SubjectAccessReview resource.
type SubjectAccessReviewArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Spec holds information about the request being evaluated
	Spec SubjectAccessReviewSpecInput
}

func (SubjectAccessReviewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subjectAccessReviewArgs)(nil)).Elem()
}
