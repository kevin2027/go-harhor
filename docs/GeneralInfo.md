# GeneralInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentTime** | Pointer to [**time.Time**](time.Time.md) | The current time of the server. | [optional] 
**WithNotary** | Pointer to **bool** | If the Harbor instance is deployed with nested notary. | [optional] 
**WithChartmuseum** | Pointer to **bool** | If the Harbor instance is deployed with nested chartmuseum. | [optional] 
**RegistryUrl** | Pointer to **string** | The url of registry against which the docker command should be issued. | [optional] 
**ExternalUrl** | Pointer to **string** | The external URL of Harbor, with protocol. | [optional] 
**AuthMode** | Pointer to **string** | The auth mode of current Harbor instance. | [optional] 
**ProjectCreationRestriction** | Pointer to **string** | Indicate who can create projects, it could be &#39;adminonly&#39; or &#39;everyone&#39;. | [optional] 
**SelfRegistration** | Pointer to **bool** | Indicate whether the Harbor instance enable user to register himself. | [optional] 
**HasCaRoot** | Pointer to **bool** | Indicate whether there is a ca root cert file ready for download in the file system. | [optional] 
**HarborVersion** | Pointer to **string** | The build version of Harbor. | [optional] 
**RegistryStorageProviderName** | Pointer to **string** | The storage provider&#39;s name of Harbor registry | [optional] 
**ReadOnly** | Pointer to **bool** | The flag to indicate whether Harbor is in readonly mode. | [optional] 
**NotificationEnable** | Pointer to **bool** | The flag to indicate whether notification mechanism is enabled on Harbor instance. | [optional] 
**AuthproxySettings** | [**AuthproxySetting**](AuthproxySetting.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


