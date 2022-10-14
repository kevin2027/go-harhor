# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** | Project ID | [optional] 
**OwnerId** | **int32** | The owner ID of the project always means the creator of the project. | [optional] 
**Name** | **string** | The name of the project. | [optional] 
**RegistryId** | **int64** | The ID of referenced registry when the project is a proxy cache project. | [optional] 
**CreationTime** | [**time.Time**](time.Time.md) | The creation time of the project. | [optional] 
**UpdateTime** | [**time.Time**](time.Time.md) | The update time of the project. | [optional] 
**Deleted** | **bool** | A deletion mark of the project. | [optional] 
**OwnerName** | **string** | The owner name of the project. | [optional] 
**Togglable** | **bool** | Correspond to the UI about whether the project&#39;s publicity is  updatable (for UI) | [optional] 
**CurrentUserRoleId** | **int32** | The role ID with highest permission of the current user who triggered the API (for UI).  This attribute is deprecated and will be removed in future versions. | [optional] 
**CurrentUserRoleIds** | **[]int32** | The list of role ID of the current user who triggered the API (for UI) | [optional] 
**RepoCount** | **int32** | The number of the repositories under this project. | [optional] 
**ChartCount** | **int32** | The total number of charts under this project. | [optional] 
**Metadata** | [**ProjectMetadata**](ProjectMetadata.md) |  | [optional] 
**CveAllowlist** | [**CveAllowlist**](CveAllowlist.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


