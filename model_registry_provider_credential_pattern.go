/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// RegistryProviderCredentialPattern The registry credential pattern
type RegistryProviderCredentialPattern struct {
	// The access key type
	AccessKeyType string `json:"access_key_type,omitempty"`
	// The access key data
	AccessKeyData string `json:"access_key_data,omitempty"`
	// The access secret type
	AccessSecretType string `json:"access_secret_type,omitempty"`
	// The access secret data
	AccessSecretData string `json:"access_secret_data,omitempty"`
}