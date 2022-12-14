# \LabelApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLabel**](LabelApi.md#CreateLabel) | **Post** /labels | Post creates a label
[**DeleteLabel**](LabelApi.md#DeleteLabel) | **Delete** /labels/{label_id} | Delete the label specified by ID.
[**GetLabelByID**](LabelApi.md#GetLabelByID) | **Get** /labels/{label_id} | Get the label specified by ID.
[**ListLabels**](LabelApi.md#ListLabels) | **Get** /labels | List labels according to the query strings.
[**UpdateLabel**](LabelApi.md#UpdateLabel) | **Put** /labels/{label_id} | Update the label properties.



## CreateLabel

> map[string]interface{} CreateLabel(ctx, optional)

Post creates a label

This endpoint let user creates a label. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateLabelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateLabelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 
 **label** | [**optional.Interface of Label**](Label.md)|  | 

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


## DeleteLabel

> map[string]interface{} DeleteLabel(ctx, labelId, optional)

Delete the label specified by ID.

Delete the label specified by ID. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelId** | **int32**| Label ID | 
 **optional** | ***DeleteLabelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteLabelOpts struct


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


## GetLabelByID

> Label GetLabelByID(ctx, labelId, optional)

Get the label specified by ID.

This endpoint let user get the label by specific ID. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelId** | **int32**| Label ID | 
 **optional** | ***GetLabelByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLabelByIDOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLabels

> []Label ListLabels(ctx, optional)

List labels according to the query strings.

This endpoint let user list labels by name, scope and project_id 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListLabelsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListLabelsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Query string to query resources. Supported query patterns are \&quot;exact match(k&#x3D;v)\&quot;, \&quot;fuzzy match(k&#x3D;~v)\&quot;, \&quot;range(k&#x3D;[min~max])\&quot;, \&quot;list with union releationship(k&#x3D;{v1 v2 v3})\&quot; and \&quot;list with intersetion relationship(k&#x3D;(v1 v2 v3))\&quot;. The value of range and list can be string(enclosed by \&quot; or &#39;), integer or time(in format \&quot;2020-04-09 02:36:00\&quot;). All of these query patterns should be put in the query string \&quot;q&#x3D;xxx\&quot; and splitted by \&quot;,\&quot;. e.g. q&#x3D;k1&#x3D;v1,k2&#x3D;~v2,k3&#x3D;[min~max] | 
 **sort** | **optional.String**| Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with \&quot;sort&#x3D;field1,-field2\&quot; | 
 **page** | **optional.Int32**| The page number | 
 **pageSize** | **optional.Int32**| The size of per page | 
 **name** | **optional.String**| The label name. | 
 **scope** | **optional.String**| The label scope. Valid values are g and p. g for global labels and p for project labels. | 
 **projectId** | **optional.Int32**| Relevant project ID, required when scope is p. | 
 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**[]Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLabel

> map[string]interface{} UpdateLabel(ctx, labelId, optional)

Update the label properties.

This endpoint let user update label properties. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelId** | **int32**| Label ID | 
 **optional** | ***UpdateLabelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateLabelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 
 **label** | [**optional.Interface of Label**](Label.md)|  | 

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

