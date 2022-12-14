/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package harbor
// RegistryProviderInfo The registry provider info contains the base info and capability declarations of the registry provider
type RegistryProviderInfo struct {
	EndpointPattern RegistryProviderEndpointPattern `json:"endpoint_pattern,omitempty"`
	CredentialPattern RegistryProviderCredentialPattern `json:"credential_pattern,omitempty"`
}
