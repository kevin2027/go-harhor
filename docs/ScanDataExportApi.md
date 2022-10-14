# \ScanDataExportApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadScanData**](ScanDataExportApi.md#DownloadScanData) | **Get** /export/cve/download/{execution_id} | Download the scan data export file
[**ExportScanData**](ScanDataExportApi.md#ExportScanData) | **Post** /export/cve | Export scan data for selected projects
[**GetScanDataExportExecution**](ScanDataExportApi.md#GetScanDataExportExecution) | **Get** /export/cve/execution/{execution_id} | Get the specific scan data export execution
[**GetScanDataExportExecutionList**](ScanDataExportApi.md#GetScanDataExportExecutionList) | **Get** /export/cve/executions | Get a list of specific scan data export execution jobs for a specified user



## DownloadScanData

> map[string]interface{} DownloadScanData(ctx, executionId, optional)

Download the scan data export file

Download the scan data report. Default format is CSV

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **int32**| Execution ID | 
 **optional** | ***DownloadScanDataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DownloadScanDataOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**| The format of the data to be exported. e.g. CSV or PDF | 
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


## ExportScanData

> ScanDataExportJob ExportScanData(ctx, xScanDataType, optional)

Export scan data for selected projects

Export scan data for selected projects

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xScanDataType** | **string**| The type of scan data to export | 
 **optional** | ***ExportScanDataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportScanDataOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 
 **scanDataExportRequest** | [**optional.Interface of ScanDataExportRequest**](ScanDataExportRequest.md)|  | 

### Return type

[**ScanDataExportJob**](ScanDataExportJob.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScanDataExportExecution

> ScanDataExportExecution GetScanDataExportExecution(ctx, executionId, optional)

Get the specific scan data export execution

Get the scan data export execution specified by ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **int32**| Execution ID | 
 **optional** | ***GetScanDataExportExecutionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetScanDataExportExecutionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**ScanDataExportExecution**](ScanDataExportExecution.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScanDataExportExecutionList

> ScanDataExportExecutionList GetScanDataExportExecutionList(ctx, optional)

Get a list of specific scan data export execution jobs for a specified user

Get a list of specific scan data export execution jobs for a specified user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetScanDataExportExecutionListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetScanDataExportExecutionListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**ScanDataExportExecutionList**](ScanDataExportExecutionList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

