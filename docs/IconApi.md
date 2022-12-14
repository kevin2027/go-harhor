# \IconApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIcon**](IconApi.md#GetIcon) | **Get** /icons/{digest} | Get artifact icon



## GetIcon

> Icon GetIcon(ctx, digest, optional)

Get artifact icon

Get the artifact icon with the specified digest. As the original icon image is resized and encoded before returning, the parameter \"digest\" in the path doesn't match the hash of the returned content

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**digest** | **string**| The digest of the resource | 
 **optional** | ***GetIconOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetIconOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**Icon**](Icon.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

