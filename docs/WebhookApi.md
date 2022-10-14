# \WebhookApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhookPolicyOfProject**](WebhookApi.md#CreateWebhookPolicyOfProject) | **Post** /projects/{project_name_or_id}/webhook/policies | Create project webhook policy.
[**DeleteWebhookPolicyOfProject**](WebhookApi.md#DeleteWebhookPolicyOfProject) | **Delete** /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id} | Delete webhook policy of a project
[**GetSupportedEventTypes**](WebhookApi.md#GetSupportedEventTypes) | **Get** /projects/{project_name_or_id}/webhook/events | Get supported event types and notify types.
[**GetWebhookPolicyOfProject**](WebhookApi.md#GetWebhookPolicyOfProject) | **Get** /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id} | Get project webhook policy
[**LastTrigger**](WebhookApi.md#LastTrigger) | **Get** /projects/{project_name_or_id}/webhook/lasttrigger | Get project webhook policy last trigger info
[**ListWebhookPoliciesOfProject**](WebhookApi.md#ListWebhookPoliciesOfProject) | **Get** /projects/{project_name_or_id}/webhook/policies | List project webhook policies.
[**UpdateWebhookPolicyOfProject**](WebhookApi.md#UpdateWebhookPolicyOfProject) | **Put** /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id} | Update webhook policy of a project.



## CreateWebhookPolicyOfProject

> map[string]interface{} CreateWebhookPolicyOfProject(ctx, projectNameOrId, optional)

Create project webhook policy.

This endpoint create a webhook policy if the project does not have one. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectNameOrId** | **string**| The name or id of the project | 
 **optional** | ***CreateWebhookPolicyOfProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateWebhookPolicyOfProjectOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.String**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | 
 **webhookPolicy** | [**optional.Interface of WebhookPolicy**](WebhookPolicy.md)|  | 

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


## DeleteWebhookPolicyOfProject

> map[string]interface{} DeleteWebhookPolicyOfProject(ctx, projectNameOrId, webhookPolicyId, optional)

Delete webhook policy of a project

This endpoint is aimed to delete webhookpolicy of a project. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectNameOrId** | **string**| The name or id of the project | 
**webhookPolicyId** | **int32**| The ID of the webhook policy | 
 **optional** | ***DeleteWebhookPolicyOfProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteWebhookPolicyOfProjectOpts struct


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


## GetSupportedEventTypes

> SupportedWebhookEventTypes GetSupportedEventTypes(ctx, projectNameOrId, optional)

Get supported event types and notify types.

Get supportted event types and notify types.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectNameOrId** | **string**| The name or id of the project | 
 **optional** | ***GetSupportedEventTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSupportedEventTypesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.String**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | 

### Return type

[**SupportedWebhookEventTypes**](SupportedWebhookEventTypes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookPolicyOfProject

> WebhookPolicy GetWebhookPolicyOfProject(ctx, projectNameOrId, webhookPolicyId, optional)

Get project webhook policy

This endpoint returns specified webhook policy of a project. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectNameOrId** | **string**| The name or id of the project | 
**webhookPolicyId** | **int32**| The ID of the webhook policy | 
 **optional** | ***GetWebhookPolicyOfProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWebhookPolicyOfProjectOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.String**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | 

### Return type

[**WebhookPolicy**](WebhookPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LastTrigger

> []WebhookLastTrigger LastTrigger(ctx, projectNameOrId, optional)

Get project webhook policy last trigger info

This endpoint returns last trigger information of project webhook policy. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectNameOrId** | **string**| The name or id of the project | 
 **optional** | ***LastTriggerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LastTriggerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.String**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | 

### Return type

[**[]WebhookLastTrigger**](WebhookLastTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhookPoliciesOfProject

> []WebhookPolicy ListWebhookPoliciesOfProject(ctx, projectNameOrId, optional)

List project webhook policies.

This endpoint returns webhook policies of a project. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectNameOrId** | **string**| The name or id of the project | 
 **optional** | ***ListWebhookPoliciesOfProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListWebhookPoliciesOfProjectOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **optional.String**| Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with \&quot;sort&#x3D;field1,-field2\&quot; | 
 **q** | **optional.String**| Query string to query resources. Supported query patterns are \&quot;exact match(k&#x3D;v)\&quot;, \&quot;fuzzy match(k&#x3D;~v)\&quot;, \&quot;range(k&#x3D;[min~max])\&quot;, \&quot;list with union releationship(k&#x3D;{v1 v2 v3})\&quot; and \&quot;list with intersetion relationship(k&#x3D;(v1 v2 v3))\&quot;. The value of range and list can be string(enclosed by \&quot; or &#39;), integer or time(in format \&quot;2020-04-09 02:36:00\&quot;). All of these query patterns should be put in the query string \&quot;q&#x3D;xxx\&quot; and splitted by \&quot;,\&quot;. e.g. q&#x3D;k1&#x3D;v1,k2&#x3D;~v2,k3&#x3D;[min~max] | 
 **page** | **optional.Int32**| The page number | 
 **pageSize** | **optional.Int32**| The size of per page | 
 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.String**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | 

### Return type

[**[]WebhookPolicy**](WebhookPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhookPolicyOfProject

> map[string]interface{} UpdateWebhookPolicyOfProject(ctx, projectNameOrId, webhookPolicyId, optional)

Update webhook policy of a project.

This endpoint is aimed to update the webhook policy of a project. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectNameOrId** | **string**| The name or id of the project | 
**webhookPolicyId** | **int32**| The ID of the webhook policy | 
 **optional** | ***UpdateWebhookPolicyOfProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateWebhookPolicyOfProjectOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.String**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | 
 **webhookPolicy** | [**optional.Interface of WebhookPolicy**](WebhookPolicy.md)|  | 

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

