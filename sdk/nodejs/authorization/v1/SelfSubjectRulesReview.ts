// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import { getVersion } from "../../version";

    /**
     * SelfSubjectRulesReview enumerates the set of actions the current user can perform within a
     * namespace. The returned list of actions may be incomplete depending on the server's
     * authorization mode, and any errors experienced during the evaluation. SelfSubjectRulesReview
     * should be used by UIs to show/hide actions, or to quickly let an end user reason about their
     * permissions. It should NOT Be used by external systems to drive authorization decisions as
     * this raises confused deputy, cache lifetime/revocation, and correctness concerns.
     * SubjectAccessReview, and LocalAccessReview are the correct way to defer authorization
     * decisions to the API server.
     */
    export class SelfSubjectRulesReview extends pulumi.CustomResource {
      /**
       * APIVersion defines the versioned schema of this representation of an object. Servers should
       * convert recognized schemas to the latest internal value, and may reject unrecognized
       * values. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
       */
      public readonly apiVersion: pulumi.Output<"authorization.k8s.io/v1">;

      /**
       * Kind is a string value representing the REST resource this object represents. Servers may
       * infer this from the endpoint the client submits requests to. Cannot be updated. In
       * CamelCase. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
       */
      public readonly kind: pulumi.Output<"SelfSubjectRulesReview">;

      
      public readonly metadata: pulumi.Output<outputs.meta.v1.ObjectMeta>;

      /**
       * Spec holds information about the request being evaluated.
       */
      public readonly spec: pulumi.Output<outputs.authorization.v1.SelfSubjectRulesReviewSpec>;

      /**
       * Status is filled in by the server and indicates the set of actions a user can perform.
       */
      public readonly status: pulumi.Output<outputs.authorization.v1.SubjectRulesReviewStatus>;

      /**
       * Get the state of an existing `SelfSubjectRulesReview` resource, as identified by `id`.
       * Typically this ID  is of the form <namespace>/<name>; if <namespace> is omitted, then (per
       * Kubernetes convention) the ID becomes default/<name>.
       *
       * Pulumi will keep track of this resource using `name` as the Pulumi ID.
       *
       * @param name _Unique_ name used to register this resource with Pulumi.
       * @param id An ID for the Kubernetes resource to retrieve. Takes the form
       *  <namespace>/<name> or <name>.
       * @param opts Uniquely specifies a CustomResource to select.
       */
      public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SelfSubjectRulesReview {
          return new SelfSubjectRulesReview(name, undefined, { ...opts, id: id });
      }

      /** @internal */
      private static readonly __pulumiType = "kubernetes:authorization.k8s.io/v1:SelfSubjectRulesReview";

      /**
       * Returns true if the given object is an instance of SelfSubjectRulesReview.  This is designed to work even
       * when multiple copies of the Pulumi SDK have been loaded into the same process.
       */
      public static isInstance(obj: any): obj is SelfSubjectRulesReview {
          if (obj === undefined || obj === null) {
              return false;
          }

          return obj["__pulumiType"] === SelfSubjectRulesReview.__pulumiType;
      }

      /**
       * Create a authorization.v1.SelfSubjectRulesReview resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputs.authorization.v1.SelfSubjectRulesReview, opts?: pulumi.CustomResourceOptions) {
          const props: pulumi.Inputs = {};
          props["apiVersion"] = "authorization.k8s.io/v1";
          props["kind"] = "SelfSubjectRulesReview";
          props["metadata"] = args && args.metadata || undefined;
          props["spec"] = args && args.spec || undefined;
          props["status"] = args && args.status || undefined;

          if (!opts) {
              opts = {};
          }

          if (!opts.version) {
              opts.version = getVersion();
          }
          super(SelfSubjectRulesReview.__pulumiType, name, props, opts);
      }
    }
