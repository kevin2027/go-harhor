# Artifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the artifact | [optional] 
**Type** | **string** | The type of the artifact, e.g. image, chart, etc | [optional] 
**MediaType** | **string** | The media type of the artifact | [optional] 
**ManifestMediaType** | **string** | The manifest media type of the artifact | [optional] 
**ProjectId** | **int64** | The ID of the project that the artifact belongs to | [optional] 
**RepositoryId** | **int64** | The ID of the repository that the artifact belongs to | [optional] 
**Digest** | **string** | The digest of the artifact | [optional] 
**Size** | **int64** | The size of the artifact | [optional] 
**Icon** | **string** | The digest of the icon | [optional] 
**PushTime** | [**time.Time**](time.Time.md) | The push time of the artifact | [optional] 
**PullTime** | [**time.Time**](time.Time.md) | The latest pull time of the artifact | [optional] 
**ExtraAttrs** | **map[string]map[string]interface{}** |  | [optional] 
**Annotations** | **map[string]string** |  | [optional] 
**References** | [**[]Reference**](Reference.md) |  | [optional] 
**Tags** | [**[]Tag**](Tag.md) |  | [optional] 
**AdditionLinks** | [**map[string]AdditionLink**](AdditionLink.md) |  | [optional] 
**Labels** | [**[]Label**](Label.md) |  | [optional] 
**ScanOverview** | [**map[string]NativeReportSummary**](NativeReportSummary.md) | The scan overview attached in the metadata of tag | [optional] 
**Accessories** | [**[]Accessory**](Accessory.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


