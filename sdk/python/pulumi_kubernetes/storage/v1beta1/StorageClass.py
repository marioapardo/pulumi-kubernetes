# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import utilities, tables


class StorageClass(pulumi.CustomResource):
    allow_volume_expansion: pulumi.Output[bool]
    """
    AllowVolumeExpansion shows whether the storage class allow volume expand
    """
    allowed_topologies: pulumi.Output[list]
    """
    Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
    """
    api_version: pulumi.Output[str]
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
    """
    kind: pulumi.Output[str]
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
    """
    metadata: pulumi.Output[dict]
    """
    Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
    """
    mount_options: pulumi.Output[list]
    """
    Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
    """
    parameters: pulumi.Output[dict]
    """
    Parameters holds the parameters for the provisioner that should create volumes of this storage class.
    """
    provisioner: pulumi.Output[str]
    """
    Provisioner indicates the type of the provisioner.
    """
    reclaim_policy: pulumi.Output[str]
    """
    Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
    """
    volume_binding_mode: pulumi.Output[str]
    """
    VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
    """
    def __init__(__self__, resource_name, opts=None, allow_volume_expansion=None, allowed_topologies=None, api_version=None, kind=None, metadata=None, mount_options=None, parameters=None, provisioner=None, reclaim_policy=None, volume_binding_mode=None, __props__=None, __name__=None, __opts__=None):
        """
        StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.

        StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_volume_expansion: AllowVolumeExpansion shows whether the storage class allow volume expand
        :param pulumi.Input[list] allowed_topologies: Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input[dict] metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input[list] mount_options: Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
        :param pulumi.Input[dict] parameters: Parameters holds the parameters for the provisioner that should create volumes of this storage class.
        :param pulumi.Input[str] provisioner: Provisioner indicates the type of the provisioner.
        :param pulumi.Input[str] reclaim_policy: Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
        :param pulumi.Input[str] volume_binding_mode: VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['allow_volume_expansion'] = allow_volume_expansion
            __props__['allowed_topologies'] = allowed_topologies
            __props__['api_version'] = 'storage.k8s.io/v1beta1'
            __props__['kind'] = 'StorageClass'
            __props__['metadata'] = metadata
            __props__['mount_options'] = mount_options
            __props__['parameters'] = parameters
            if provisioner is None:
                raise TypeError("Missing required property 'provisioner'")
            __props__['provisioner'] = provisioner
            __props__['reclaim_policy'] = reclaim_policy
            __props__['volume_binding_mode'] = volume_binding_mode
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="kubernetes:storage.k8s.io/v1:StorageClass")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(StorageClass, __self__).__init__(
            'kubernetes:storage.k8s.io/v1beta1:StorageClass',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing StorageClass resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return StorageClass(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
