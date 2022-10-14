# ScanDataExportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobName** | **string** | Name of the scan data export job | [optional] 
**Projects** | **[]int64** | A list of one or more projects for which to export the scan data, currently only one project is supported due to performance concerns, but define as array for extension in the future. | [optional] 
**Labels** | **[]int64** | A list of one or more labels for which to export the scan data, defaults to all if empty | [optional] 
**Repositories** | **string** | A list of repositories for which to export the scan data, defaults to all if empty | [optional] 
**CveIds** | **string** | CVE-IDs for which to export data. Multiple CVE-IDs can be specified by separating using &#39;,&#39; and enclosed between &#39;{}&#39;. Defaults to all if empty | [optional] 
**Tags** | **string** | A list of tags enclosed within &#39;{}&#39;. Defaults to all if empty | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


