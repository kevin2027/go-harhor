/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package harbor
// Platform struct for Platform
type Platform struct {
	// The architecture that the artifact applys to
	Architecture string `json:"architecture,omitempty"`
	// The OS that the artifact applys to
	Os string `json:"os,omitempty"`
	// The version of the OS that the artifact applys to
	OsVersion string `json:"&#39;os.version&#39;,omitempty"`
	// The features of the OS that the artifact applys to
	OsFeatures []string `json:"&#39;os.features&#39;,omitempty"`
	// The variant of the CPU
	Variant string `json:"variant,omitempty"`
}
