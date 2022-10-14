/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package harbor
// ScannerAdapterMetadata The metadata info of the scanner adapter
type ScannerAdapterMetadata struct {
	Scanner Scanner `json:"scanner,omitempty"`
	Capabilities []ScannerCapability `json:"capabilities,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}
