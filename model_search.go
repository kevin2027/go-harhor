/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package harbor
// Search struct for Search
type Search struct {
	// Search results of the projects that matched the filter keywords.
	Project []Project `json:"project,omitempty"`
	// Search results of the repositories that matched the filter keywords.
	Repository []SearchRepository `json:"repository,omitempty"`
	// Search results of the charts that macthed the filter keywords.
	Chart []SearchResult `json:"chart,omitempty"`
}
