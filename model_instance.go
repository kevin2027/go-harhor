/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Instance struct for Instance
type Instance struct {
	// Unique ID
	Id int32 `json:"id,omitempty"`
	// Instance name
	Name string `json:"name,omitempty"`
	// Description of instance
	Description string `json:"description,omitempty"`
	// Based on which driver, identified by ID
	Vendor string `json:"vendor,omitempty"`
	// The service endpoint of this instance
	Endpoint string `json:"endpoint,omitempty"`
	// The authentication way supported
	AuthMode string `json:"auth_mode,omitempty"`
	// The auth credential data if exists
	AuthInfo map[string]string `json:"auth_info,omitempty"`
	// The health status
	Status string `json:"status,omitempty"`
	// Whether the instance is activated or not
	Enabled bool `json:"enabled,omitempty"`
	// Whether the instance is default or not
	Default bool `json:"default,omitempty"`
	// Whether the instance endpoint is insecure or not
	Insecure bool `json:"insecure,omitempty"`
	// The timestamp of instance setting up
	SetupTimestamp int64 `json:"setup_timestamp,omitempty"`
}
