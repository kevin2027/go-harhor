/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Permission struct for Permission
type Permission struct {
	// The permission resoruce
	Resource string `json:"resource,omitempty"`
	// The permission action
	Action string `json:"action,omitempty"`
}
