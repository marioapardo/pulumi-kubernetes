// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Core.V1
{

    [OutputType]
    public sealed class SecretReference
    {
        /// <summary>
        /// Name is unique within a namespace to reference a secret resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Namespace defines the space within which the secret name must be unique.
        /// </summary>
        public readonly string Namespace;

        [OutputConstructor]
        private SecretReference(
            string name,

            string @namespace)
        {
            Name = name;
            Namespace = @namespace;
        }
    }
}