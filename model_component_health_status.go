/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package harbor
// ComponentHealthStatus The health status of component
type ComponentHealthStatus struct {
	// The component name
	Name string `json:"name,omitempty"`
	// The health status of component
	Status string `json:"status,omitempty"`
	// (optional) The error message when the status is \"unhealthy\"
	Error string `json:"error,omitempty"`
}
