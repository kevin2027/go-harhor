# NativeReportSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportId** | **string** | id of the native scan report | [optional] 
**ScanStatus** | **string** | The status of the report generating process | [optional] 
**Severity** | **string** | The overall severity | [optional] 
**Duration** | **int64** | The seconds spent for generating the report | [optional] 
**Summary** | [**VulnerabilitySummary**](VulnerabilitySummary.md) |  | [optional] 
**StartTime** | [**time.Time**](time.Time.md) | The start time of the scan process that generating report | [optional] 
**EndTime** | [**time.Time**](time.Time.md) | The end time of the scan process that generating report | [optional] 
**CompletePercent** | **int32** | The complete percent of the scanning which value is between 0 and 100 | [optional] 
**Scanner** | [**Scanner**](Scanner.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


