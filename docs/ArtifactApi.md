# \ArtifactApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLabel**](ArtifactApi.md#AddLabel) | **Post** /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels | Add label to artifact
[**CopyArtifact**](ArtifactApi.md#CopyArtifact) | **Post** /projects/{project_name}/repositories/{repository_name}/artifacts | Copy artifact
[**CreateTag**](ArtifactApi.md#CreateTag) | **Post** /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags | Create tag
[**DeleteArtifact**](ArtifactApi.md#DeleteArtifact) | **Delete** /projects/{project_name}/repositories/{repository_name}/artifacts/{reference} | Delete the specific artifact
[**DeleteTag**](ArtifactApi.md#DeleteTag) | **Delete** /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags/{tag_name} | Delete tag
[**GetAddition**](ArtifactApi.md#GetAddition) | **Get** /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/additions/{addition} | Get the addition of the specific artifact
[**GetArtifact**](ArtifactApi.md#GetArtifact) | **Get** /projects/{project_name}/repositories/{repository_name}/artifacts/{reference} | Get the specific artifact
[**GetVulnerabilitiesAddition**](ArtifactApi.md#GetVulnerabilitiesAddition) | **Get** /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/additions/vulnerabilities | Get the vulnerabilities addition of the specific artifact
[**ListAccessories**](ArtifactApi.md#ListAccessories) | **Get** /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/accessories | List accessories
[**ListArtifacts**](ArtifactApi.md#ListArtifacts) | **Get** /projects/{project_name}/repositories/{repository_name}/artifacts | List artifacts
[**ListTags**](ArtifactApi.md#ListTags) | **Get** /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags | List tags
[**RemoveLabel**](ArtifactApi.md#RemoveLabel) | **Delete** /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels/{label_id} | Remove label from artifact



## AddLabel

> map[string]interface{} AddLabel(ctx, projectName, repositoryName, reference, optional)

Add label to artifact

Add label to the specified artiact.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| The name of the project | 
**repositoryName** | **string**| The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -&gt; a%252Fb | 
**reference** | **string**| The reference of the artifact, can be digest or tag | 
 **optional** | ***AddLabelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddLabelOpts struct


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


## CopyArtifact

> map[string]interface{} CopyArtifact(ctx, projectName, repositoryName, from, optional)

Copy artifact

Copy the artifact specified in the \"from\" parameter to the repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| The name of the project | 
**repositoryName** | **string**| The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -&gt; a%252Fb | 
**from** | **string**| The artifact from which the new artifact is copied from, the format should be \&quot;project/repository:tag\&quot; or \&quot;project/repository@digest\&quot;. | 
 **optional** | ***CopyArtifactOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CopyArtifactOpts struct


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


## CreateTag

> map[string]interface{} CreateTag(ctx, projectName, repositoryName, reference, optional)

Create tag

Create a tag for the specified artifact

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| The name of the project | 
**repositoryName** | **string**| The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -&gt; a%252Fb | 
**reference** | **string**| The reference of the artifact, can be digest or tag | 
 **optional** | ***CreateTagOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTagOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRequestId** | **optional.String**| An unique ID for the request | 
 **tag** | [**optional.Interface of Tag**](Tag.md)|  | 

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


## DeleteArtifact

> map[string]interface{} DeleteArtifact(ctx, projectName, repositoryName, reference, optional)

Delete the specific artifact

Delete the artifact specified by the reference under the project and repository. The reference can be digest or tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| The name of the project | 
**repositoryName** | **string**| The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -&gt; a%252Fb | 
**reference** | **string**| The reference of the artifact, can be digest or tag | 
 **optional** | ***DeleteArtifactOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteArtifactOpts struct


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


## DeleteTag

> map[string]interface{} DeleteTag(ctx, projectName, repositoryName, reference, tagName, optional)

Delete tag

Delete the tag of the specified artifact

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| The name of the project | 
**repositoryName** | **string**| The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -&gt; a%252Fb | 
**reference** | **string**| The reference of the artifact, can be digest or tag | 
**tagName** | **string**| The name of the tag | 
 **optional** | ***DeleteTagOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteTagOpts struct


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


## GetAddition

> string GetAddition(ctx, projectName, repositoryName, reference, addition, optional)

Get the addition of the specific artifact

Get the addition of the artifact specified by the reference under the project and repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| The name of the project | 
**repositoryName** | **string**| The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -&gt; a%252Fb | 
**reference** | **string**| The reference of the artifact, can be digest or tag | 
**addition** | **string**| The type of addition. | 
 **optional** | ***GetAdditionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAdditionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifact

> Artifact GetArtifact(ctx, projectName, repositoryName, reference, optional)

Get the specific artifact

Get the artifact specified by the reference under the project and repository. The reference can be digest or tag.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| The name of the project | 
**repositoryName** | **string**| The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -&gt; a%252Fb | 
**reference** | **string**| The reference of the artifact, can be digest or tag | 
 **optional** | ***GetArtifactOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetArtifactOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **optional.Int32**| The page number | 
 **pageSize** | **optional.Int32**| The size of per page | 
 **withTag** | **optional.String**| Specify whether the tags are inclued inside the returning artifacts | 
 **withLabel** | **optional.String**| Specify whether the labels are inclued inside the returning artifacts | 
 **withScanOverview** | **optional.String**| Specify whether the scan overview is inclued inside the returning artifacts | 
 **withAccessory** | **optional.String**| Specify whether the accessories are included of the returning artifacts. | 
 **withSignature** | **optional.String**| Specify whether the signature is inclued inside the returning artifacts | 
 **withImmutableStatus** | **optional.String**| Specify whether the immutable status is inclued inside the tags of the returning artifacts. | 
 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xAcceptVulnerabilities** | **optional.String**| A comma-separated lists of MIME types for the scan report or scan summary. The first mime type will be used when the report found for it. Currently the mime type supports &#39;application/vnd.scanner.adapter.vuln.report.harbor+json; version&#x3D;1.0&#39; and &#39;application/vnd.security.vulnerability.report; version&#x3D;1.1&#39; | 

### Return type

[**Artifact**](Artifact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVulnerabilitiesAddition

> string GetVulnerabilitiesAddition(ctx, projectName, repositoryName, reference, optional)

Get the vulnerabilities addition of the specific artifact

Get the vulnerabilities addition of the artifact specified by the reference under the project and repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| The name of the project | 
**repositoryName** | **string**| The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -&gt; a%252Fb | 
**reference** | **string**| The reference of the artifact, can be digest or tag | 
 **optional** | ***GetVulnerabilitiesAdditionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetVulnerabilitiesAdditionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xAcceptVulnerabilities** | **optional.String**| A comma-separated lists of MIME types for the scan report or scan summary. The first mime type will be used when the report found for it. Currently the mime type supports &#39;application/vnd.scanner.adapter.vuln.report.harbor+json; version&#x3D;1.0&#39; and &#39;application/vnd.security.vulnerability.report; version&#x3D;1.1&#39; | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccessories

> []Accessory ListAccessories(ctx, projectName, repositoryName, reference, optional)

List accessories

List accessories of the specific artifact

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| The name of the project | 
**repositoryName** | **string**| The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -&gt; a%252Fb | 
**reference** | **string**| The reference of the artifact, can be digest or tag | 
 **optional** | ***ListAccessoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAccessoriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **q** | **optional.String**| Query string to query resources. Supported query patterns are \&quot;exact match(k&#x3D;v)\&quot;, \&quot;fuzzy match(k&#x3D;~v)\&quot;, \&quot;range(k&#x3D;[min~max])\&quot;, \&quot;list with union releationship(k&#x3D;{v1 v2 v3})\&quot; and \&quot;list with intersetion relationship(k&#x3D;(v1 v2 v3))\&quot;. The value of range and list can be string(enclosed by \&quot; or &#39;), integer or time(in format \&quot;2020-04-09 02:36:00\&quot;). All of these query patterns should be put in the query string \&quot;q&#x3D;xxx\&quot; and splitted by \&quot;,\&quot;. e.g. q&#x3D;k1&#x3D;v1,k2&#x3D;~v2,k3&#x3D;[min~max] | 
 **sort** | **optional.String**| Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with \&quot;sort&#x3D;field1,-field2\&quot; | 
 **page** | **optional.Int32**| The page number | 
 **pageSize** | **optional.Int32**| The size of per page | 
 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**[]Accessory**](Accessory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArtifacts

> []Artifact ListArtifacts(ctx, projectName, repositoryName, optional)

List artifacts

List artifacts under the specific project and repository. Except the basic properties, the other supported queries in \"q\" includes \"tags=*\" to list only tagged artifacts, \"tags=nil\" to list only untagged artifacts, \"tags=~v\" to list artifacts whose tag fuzzy matches \"v\", \"tags=v\" to list artifact whose tag exactly matches \"v\", \"labels=(id1, id2)\" to list artifacts that both labels with id1 and id2 are added to

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| The name of the project | 
**repositoryName** | **string**| The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -&gt; a%252Fb | 
 **optional** | ***ListArtifactsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListArtifactsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **q** | **optional.String**| Query string to query resources. Supported query patterns are \&quot;exact match(k&#x3D;v)\&quot;, \&quot;fuzzy match(k&#x3D;~v)\&quot;, \&quot;range(k&#x3D;[min~max])\&quot;, \&quot;list with union releationship(k&#x3D;{v1 v2 v3})\&quot; and \&quot;list with intersetion relationship(k&#x3D;(v1 v2 v3))\&quot;. The value of range and list can be string(enclosed by \&quot; or &#39;), integer or time(in format \&quot;2020-04-09 02:36:00\&quot;). All of these query patterns should be put in the query string \&quot;q&#x3D;xxx\&quot; and splitted by \&quot;,\&quot;. e.g. q&#x3D;k1&#x3D;v1,k2&#x3D;~v2,k3&#x3D;[min~max] | 
 **sort** | **optional.String**| Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with \&quot;sort&#x3D;field1,-field2\&quot; | 
 **page** | **optional.Int32**| The page number | 
 **pageSize** | **optional.Int32**| The size of per page | 
 **withTag** | **optional.String**| Specify whether the tags are included inside the returning artifacts | 
 **withLabel** | **optional.String**| Specify whether the labels are included inside the returning artifacts | 
 **withScanOverview** | **optional.String**| Specify whether the scan overview is included inside the returning artifacts | 
 **withSignature** | **optional.String**| Specify whether the signature is included inside the tags of the returning artifacts. Only works when setting \&quot;with_tag&#x3D;true\&quot; | 
 **withImmutableStatus** | **optional.String**| Specify whether the immutable status is included inside the tags of the returning artifacts. Only works when setting \&quot;with_immutable_status&#x3D;true\&quot; | 
 **withAccessory** | **optional.String**| Specify whether the accessories are included of the returning artifacts. Only works when setting \&quot;with_accessory&#x3D;true\&quot; | 
 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xAcceptVulnerabilities** | **optional.String**| A comma-separated lists of MIME types for the scan report or scan summary. The first mime type will be used when the report found for it. Currently the mime type supports &#39;application/vnd.scanner.adapter.vuln.report.harbor+json; version&#x3D;1.0&#39; and &#39;application/vnd.security.vulnerability.report; version&#x3D;1.1&#39; | 

### Return type

[**[]Artifact**](Artifact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTags

> []Tag ListTags(ctx, projectName, repositoryName, reference, optional)

List tags

List tags of the specific artifact

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| The name of the project | 
**repositoryName** | **string**| The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -&gt; a%252Fb | 
**reference** | **string**| The reference of the artifact, can be digest or tag | 
 **optional** | ***ListTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **q** | **optional.String**| Query string to query resources. Supported query patterns are \&quot;exact match(k&#x3D;v)\&quot;, \&quot;fuzzy match(k&#x3D;~v)\&quot;, \&quot;range(k&#x3D;[min~max])\&quot;, \&quot;list with union releationship(k&#x3D;{v1 v2 v3})\&quot; and \&quot;list with intersetion relationship(k&#x3D;(v1 v2 v3))\&quot;. The value of range and list can be string(enclosed by \&quot; or &#39;), integer or time(in format \&quot;2020-04-09 02:36:00\&quot;). All of these query patterns should be put in the query string \&quot;q&#x3D;xxx\&quot; and splitted by \&quot;,\&quot;. e.g. q&#x3D;k1&#x3D;v1,k2&#x3D;~v2,k3&#x3D;[min~max] | 
 **sort** | **optional.String**| Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with \&quot;sort&#x3D;field1,-field2\&quot; | 
 **page** | **optional.Int32**| The page number | 
 **pageSize** | **optional.Int32**| The size of per page | 
 **withSignature** | **optional.String**| Specify whether the signature is included inside the returning tags | 
 **withImmutableStatus** | **optional.String**| Specify whether the immutable status is included inside the returning tags | 
 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**[]Tag**](Tag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLabel

> map[string]interface{} RemoveLabel(ctx, projectName, repositoryName, reference, labelId, optional)

Remove label from artifact

Remove the label from the specified artiact.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| The name of the project | 
**repositoryName** | **string**| The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -&gt; a%252Fb | 
**reference** | **string**| The reference of the artifact, can be digest or tag | 
**labelId** | **int32**| The ID of the label that removed from the artifact. | 
 **optional** | ***RemoveLabelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveLabelOpts struct


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

