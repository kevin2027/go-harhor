/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package harbor
// ProviderUnderProject struct for ProviderUnderProject
type ProviderUnderProject struct {
	Id int32 `json:"id,omitempty"`
	Provider string `json:"provider,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Default bool `json:"default,omitempty"`
}
