# \MemberApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectMember**](MemberApi.md#CreateProjectMember) | **Post** /projects/{project_name_or_id}/members | Create project member
[**DeleteProjectMember**](MemberApi.md#DeleteProjectMember) | **Delete** /projects/{project_name_or_id}/members/{mid} | Delete project member
[**GetProjectMember**](MemberApi.md#GetProjectMember) | **Get** /projects/{project_name_or_id}/members/{mid} | Get the project member information
[**ListProjectMembers**](MemberApi.md#ListProjectMembers) | **Get** /projects/{project_name_or_id}/members | Get all project member information
[**UpdateProjectMember**](MemberApi.md#UpdateProjectMember) | **Put** /projects/{project_name_or_id}/members/{mid} | Update project member



## CreateProjectMember

> map[string]interface{} CreateProjectMember(ctx, projectNameOrId, optional)

Create project member

Create project member relationship, the member can be one of the user_member and group_member,  The user_member need to specify user_id or username. If the user already exist in harbor DB, specify the user_id,  If does not exist in harbor DB, it will SearchAndOnBoard the user. The group_member need to specify id or ldap_group_dn. If the group already exist in harbor DB. specify the user group's id,  If does not exist, it will SearchAndOnBoard the group. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectNameOrId** | **string**| The name or id of the project | 
 **optional** | ***CreateProjectMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateProjectMemberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.String**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | 
 **projectMember** | [**optional.Interface of ProjectMember**](ProjectMember.md)|  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectMember

> map[string]interface{} DeleteProjectMember(ctx, projectNameOrId, mid, optional)

Delete project member

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectNameOrId** | **string**| The name or id of the project | 
**mid** | **int32**| Member ID. | 
 **optional** | ***DeleteProjectMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteProjectMemberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.String**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectMember

> ProjectMemberEntity GetProjectMember(ctx, projectNameOrId, mid, optional)

Get the project member information

Get the project member information

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectNameOrId** | **string**| The name or id of the project | 
**mid** | **int32**| The member ID | 
 **optional** | ***GetProjectMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProjectMemberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.String**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | 

### Return type

[**ProjectMemberEntity**](ProjectMemberEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjectMembers

> []ProjectMemberEntity ListProjectMembers(ctx, projectNameOrId, optional)

Get all project member information

Get all project member information

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectNameOrId** | **string**| The name or id of the project | 
 **optional** | ***ListProjectMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListProjectMembersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| The page number | 
 **pageSize** | **optional.Int32**| The size of per page | 
 **entityname** | **optional.String**| The entity name to search. | 
 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.String**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | 

### Return type

[**[]ProjectMemberEntity**](ProjectMemberEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProjectMember

> map[string]interface{} UpdateProjectMember(ctx, projectNameOrId, mid, optional)

Update project member

Update project member relationship

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectNameOrId** | **string**| The name or id of the project | 
**mid** | **int32**| Member ID. | 
 **optional** | ***UpdateProjectMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateProjectMemberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.String**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | 
 **roleRequest** | [**optional.Interface of RoleRequest**](RoleRequest.md)|  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

