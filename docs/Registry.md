# Registry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The registry ID. | [optional] 
**Url** | **string** | The registry URL string. | [optional] 
**Name** | **string** | The registry name. | [optional] 
**Credential** | [**RegistryCredential**](RegistryCredential.md) |  | [optional] 
**Type** | **string** | Type of the registry, e.g. &#39;harbor&#39;. | [optional] 
**Insecure** | **bool** | Whether or not the certificate will be verified when Harbor tries to access the server. | [optional] 
**Description** | **string** | Description of the registry. | [optional] 
**Status** | **string** | Health status of the registry. | [optional] 
**CreationTime** | [**time.Time**](time.Time.md) | The create time of the policy. | [optional] 
**UpdateTime** | [**time.Time**](time.Time.md) | The update time of the policy. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


