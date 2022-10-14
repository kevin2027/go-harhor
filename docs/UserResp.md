# UserResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | [optional] 
**Realname** | **string** |  | [optional] 
**Comment** | **string** |  | [optional] 
**UserId** | **int32** |  | [optional] 
**Username** | **string** |  | [optional] 
**SysadminFlag** | **bool** |  | [optional] 
**AdminRoleInAuth** | **bool** | indicate the admin privilege is grant by authenticator (LDAP), is always false unless it is the current login user | [optional] 
**OidcUserMeta** | [**OidcUserInfo**](OidcUserInfo.md) |  | [optional] 
**CreationTime** | [**time.Time**](time.Time.md) | The creation time of the user. | [optional] 
**UpdateTime** | [**time.Time**](time.Time.md) | The update time of the user. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


