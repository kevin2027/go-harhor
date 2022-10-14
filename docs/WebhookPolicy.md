# WebhookPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The webhook policy ID. | [optional] 
**Name** | **string** | The name of webhook policy. | [optional] 
**Description** | **string** | The description of webhook policy. | [optional] 
**ProjectId** | **int32** | The project ID of webhook policy. | [optional] 
**Targets** | [**[]WebhookTargetObject**](WebhookTargetObject.md) |  | [optional] 
**EventTypes** | **[]string** |  | [optional] 
**Creator** | **string** | The creator of the webhook policy. | [optional] 
**CreationTime** | [**time.Time**](time.Time.md) | The create time of the webhook policy. | [optional] 
**UpdateTime** | [**time.Time**](time.Time.md) | The update time of the webhook policy. | [optional] 
**Enabled** | **bool** | Whether the webhook policy is enabled or not. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


