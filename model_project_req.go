/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ProjectReq struct for ProjectReq
type ProjectReq struct {
	// The name of the project.
	ProjectName string `json:"project_name,omitempty"`
	// deprecated, reserved for project creation in replication
	Public *bool `json:"public,omitempty"`
	Metadata ProjectMetadata `json:"metadata,omitempty"`
	CveAllowlist CveAllowlist `json:"cve_allowlist,omitempty"`
	// The storage quota of the project.
	StorageLimit *int64 `json:"storage_limit,omitempty"`
	// The ID of referenced registry when creating the proxy cache project
	RegistryId *int64 `json:"registry_id,omitempty"`
}