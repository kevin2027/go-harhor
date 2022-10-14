# ReplicationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The policy ID. | [optional] 
**Name** | **string** | The policy name. | [optional] 
**Description** | **string** | The description of the policy. | [optional] 
**SrcRegistry** | [**Registry**](Registry.md) |  | [optional] 
**DestRegistry** | [**Registry**](Registry.md) |  | [optional] 
**DestNamespace** | **string** | The destination namespace. | [optional] 
**DestNamespaceReplaceCount** | **int32** | Specify how many path components will be replaced by the provided destination namespace. The default value is -1 in which case the legacy mode will be applied. | [optional] 
**Trigger** | [**ReplicationTrigger**](ReplicationTrigger.md) |  | [optional] 
**Filters** | [**[]ReplicationFilter**](ReplicationFilter.md) | The replication policy filter array. | [optional] 
**ReplicateDeletion** | **bool** | Whether to replicate the deletion operation. | [optional] 
**Deletion** | **bool** | Deprecated, use \&quot;replicate_deletion\&quot; instead. Whether to replicate the deletion operation. | [optional] 
**Override** | **bool** | Whether to override the resources on the destination registry. | [optional] 
**Enabled** | **bool** | Whether the policy is enabled or not. | [optional] 
**CreationTime** | [**time.Time**](time.Time.md) | The create time of the policy. | [optional] 
**UpdateTime** | [**time.Time**](time.Time.md) | The update time of the policy. | [optional] 
**Speed** | **int32** | speed limit for each task | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


