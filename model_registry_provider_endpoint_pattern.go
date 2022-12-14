/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package harbor
// RegistryProviderEndpointPattern The registry endpoint pattern
type RegistryProviderEndpointPattern struct {
	// The endpoint type
	EndpointType string `json:"endpoint_type,omitempty"`
	// The endpoint list
	Endpoints []RegistryEndpoint `json:"endpoints,omitempty"`
}
