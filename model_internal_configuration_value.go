/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InternalConfigurationValue struct for InternalConfigurationValue
type InternalConfigurationValue struct {
	// The value of current config item
	Value map[string]interface{} `json:"value,omitempty"`
	// The configure item can be updated or not
	Editable bool `json:"editable,omitempty"`
}
