# ProjectMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Public** | **string** | The public status of the project. The valid values are \&quot;true\&quot;, \&quot;false\&quot;. | [optional] 
**EnableContentTrust** | Pointer to **string** | Whether content trust is enabled or not. If it is enabled, user can&#39;t pull unsigned images from this project. The valid values are \&quot;true\&quot;, \&quot;false\&quot;. | [optional] 
**EnableContentTrustCosign** | Pointer to **string** | Whether cosign content trust is enabled or not. If it is enabled, user can&#39;t pull images without cosign signature from this project. The valid values are \&quot;true\&quot;, \&quot;false\&quot;. | [optional] 
**PreventVul** | Pointer to **string** | Whether prevent the vulnerable images from running. The valid values are \&quot;true\&quot;, \&quot;false\&quot;. | [optional] 
**Severity** | Pointer to **string** | If the vulnerability is high than severity defined here, the images can&#39;t be pulled. The valid values are \&quot;none\&quot;, \&quot;low\&quot;, \&quot;medium\&quot;, \&quot;high\&quot;, \&quot;critical\&quot;. | [optional] 
**AutoScan** | Pointer to **string** | Whether scan images automatically when pushing. The valid values are \&quot;true\&quot;, \&quot;false\&quot;. | [optional] 
**ReuseSysCveAllowlist** | Pointer to **string** | Whether this project reuse the system level CVE allowlist as the allowlist of its own.  The valid values are \&quot;true\&quot;, \&quot;false\&quot;. If it is set to \&quot;true\&quot; the actual allowlist associate with this project, if any, will be ignored. | [optional] 
**RetentionId** | Pointer to **string** | The ID of the tag retention policy for the project | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


