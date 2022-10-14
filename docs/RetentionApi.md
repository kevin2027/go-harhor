# \RetentionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRetention**](RetentionApi.md#CreateRetention) | **Post** /retentions | Create Retention Policy
[**DeleteRetention**](RetentionApi.md#DeleteRetention) | **Delete** /retentions/{id} | Delete Retention Policy
[**GetRentenitionMetadata**](RetentionApi.md#GetRentenitionMetadata) | **Get** /retentions/metadatas | Get Retention Metadatas
[**GetRetention**](RetentionApi.md#GetRetention) | **Get** /retentions/{id} | Get Retention Policy
[**GetRetentionTaskLog**](RetentionApi.md#GetRetentionTaskLog) | **Get** /retentions/{id}/executions/{eid}/tasks/{tid} | Get Retention job task log
[**ListRetentionExecutions**](RetentionApi.md#ListRetentionExecutions) | **Get** /retentions/{id}/executions | Get Retention executions
[**ListRetentionTasks**](RetentionApi.md#ListRetentionTasks) | **Get** /retentions/{id}/executions/{eid}/tasks | Get Retention tasks
[**OperateRetentionExecution**](RetentionApi.md#OperateRetentionExecution) | **Patch** /retentions/{id}/executions/{eid} | Stop a Retention execution
[**TriggerRetentionExecution**](RetentionApi.md#TriggerRetentionExecution) | **Post** /retentions/{id}/executions | Trigger a Retention Execution
[**UpdateRetention**](RetentionApi.md#UpdateRetention) | **Put** /retentions/{id} | Update Retention Policy



## CreateRetention

> map[string]interface{} CreateRetention(ctx, optional)

Create Retention Policy

Create Retention Policy, you can reference metadatas API for the policy model. You can check project metadatas to find whether a retention policy is already binded. This method should only be called when no retention policy binded to project yet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateRetentionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateRetentionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 
 **retentionPolicy** | [**optional.Interface of RetentionPolicy**](RetentionPolicy.md)|  | 

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


## DeleteRetention

> map[string]interface{} DeleteRetention(ctx, id, optional)

Delete Retention Policy

Delete Retention Policy, you can reference metadatas API for the policy model. You can check project metadatas to find whether a retention policy is already binded. This method should only be called when retention policy has already binded to project.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| Retention ID. | 
 **optional** | ***DeleteRetentionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteRetentionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

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


## GetRentenitionMetadata

> RetentionMetadata GetRentenitionMetadata(ctx, optional)

Get Retention Metadatas

Get Retention Metadatas.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetRentenitionMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRentenitionMetadataOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**RetentionMetadata**](RetentionMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRetention

> RetentionPolicy GetRetention(ctx, id, optional)

Get Retention Policy

Get Retention Policy.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| Retention ID. | 
 **optional** | ***GetRetentionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRetentionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**RetentionPolicy**](RetentionPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRetentionTaskLog

> map[string]interface{} GetRetentionTaskLog(ctx, id, eid, tid, optional)

Get Retention job task log

Get Retention job task log, tags ratain or deletion detail will be shown in a table.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| Retention ID. | 
**eid** | **int32**| Retention execution ID. | 
**tid** | **int32**| Retention execution ID. | 
 **optional** | ***GetRetentionTaskLogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRetentionTaskLogOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRequestId** | **optional.String**| An unique ID for the request | 

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


## ListRetentionExecutions

> []RetentionExecution ListRetentionExecutions(ctx, id, optional)

Get Retention executions

Get Retention executions, execution status may be delayed before job service schedule it up.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| Retention ID. | 
 **optional** | ***ListRetentionExecutionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRetentionExecutionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| The page number. | 
 **pageSize** | **optional.Int32**| The size of per page. | 
 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**[]RetentionExecution**](RetentionExecution.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRetentionTasks

> []RetentionExecutionTask ListRetentionTasks(ctx, id, eid, optional)

Get Retention tasks

Get Retention tasks, each repository as a task.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| Retention ID. | 
**eid** | **int32**| Retention execution ID. | 
 **optional** | ***ListRetentionTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRetentionTasksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| The page number. | 
 **pageSize** | **optional.Int32**| The size of per page. | 
 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**[]RetentionExecutionTask**](RetentionExecutionTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperateRetentionExecution

> map[string]interface{} OperateRetentionExecution(ctx, id, eid, optional)

Stop a Retention execution

Stop a Retention execution, only support \"stop\" action now.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| Retention ID. | 
**eid** | **int32**| Retention execution ID. | 
 **optional** | ***OperateRetentionExecutionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OperateRetentionExecutionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| An unique ID for the request | 
 **inlineObject2** | [**optional.Interface of InlineObject2**](InlineObject2.md)|  | 

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


## TriggerRetentionExecution

> map[string]interface{} TriggerRetentionExecution(ctx, id, optional)

Trigger a Retention Execution

Trigger a Retention Execution, if dry_run is True, nothing would be deleted actually.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| Retention ID. | 
 **optional** | ***TriggerRetentionExecutionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TriggerRetentionExecutionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 
 **inlineObject1** | [**optional.Interface of InlineObject1**](InlineObject1.md)|  | 

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


## UpdateRetention

> map[string]interface{} UpdateRetention(ctx, id, optional)

Update Retention Policy

Update Retention Policy, you can reference metadatas API for the policy model. You can check project metadatas to find whether a retention policy is already binded. This method should only be called when retention policy has already binded to project.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| Retention ID. | 
 **optional** | ***UpdateRetentionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateRetentionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 
 **retentionPolicy** | [**optional.Interface of RetentionPolicy**](RetentionPolicy.md)|  | 

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

