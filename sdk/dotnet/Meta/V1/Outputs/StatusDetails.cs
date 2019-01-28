// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Meta.V1
{

    [OutputType]
    public sealed class StatusDetails
    {
        /// <summary>
        /// The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Meta.V1.StatusCause> Causes;
        /// <summary>
        /// The group attribute of the resource associated with the status StatusReason.
        /// </summary>
        public readonly string Group;
        /// <summary>
        /// The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.
        /// </summary>
        public readonly int RetryAfterSeconds;
        /// <summary>
        /// UID of the resource. (when there is a single resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids
        /// </summary>
        public readonly string Uid;

        [OutputConstructor]
        private StatusDetails(
            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Meta.V1.StatusCause> causes,

            string group,

            string kind,

            string name,

            int retryAfterSeconds,

            string uid)
        {
            Causes = causes;
            Group = group;
            Kind = kind;
            Name = name;
            RetryAfterSeconds = retryAfterSeconds;
            Uid = uid;
        }
    }
}