/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Access struct for Access
type Access struct {
	// The resource of the access
	Resource string `json:"resource,omitempty"`
	// The action of the access
	Action string `json:"action,omitempty"`
	// The effect of the access
	Effect string `json:"effect,omitempty"`
}