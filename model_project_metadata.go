/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ProjectMetadata struct for ProjectMetadata
type ProjectMetadata struct {
	// The public status of the project. The valid values are \"true\", \"false\".
	Public string `json:"public,omitempty"`
	// Whether content trust is enabled or not. If it is enabled, user can't pull unsigned images from this project. The valid values are \"true\", \"false\".
	EnableContentTrust *string `json:"enable_content_trust,omitempty"`
	// Whether cosign content trust is enabled or not. If it is enabled, user can't pull images without cosign signature from this project. The valid values are \"true\", \"false\".
	EnableContentTrustCosign *string `json:"enable_content_trust_cosign,omitempty"`
	// Whether prevent the vulnerable images from running. The valid values are \"true\", \"false\".
	PreventVul *string `json:"prevent_vul,omitempty"`
	// If the vulnerability is high than severity defined here, the images can't be pulled. The valid values are \"none\", \"low\", \"medium\", \"high\", \"critical\".
	Severity *string `json:"severity,omitempty"`
	// Whether scan images automatically when pushing. The valid values are \"true\", \"false\".
	AutoScan *string `json:"auto_scan,omitempty"`
	// Whether this project reuse the system level CVE allowlist as the allowlist of its own.  The valid values are \"true\", \"false\". If it is set to \"true\" the actual allowlist associate with this project, if any, will be ignored.
	ReuseSysCveAllowlist *string `json:"reuse_sys_cve_allowlist,omitempty"`
	// The ID of the tag retention policy for the project
	RetentionId *string `json:"retention_id,omitempty"`
}
