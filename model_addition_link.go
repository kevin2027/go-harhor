/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// AdditionLink struct for AdditionLink
type AdditionLink struct {
	// The link of the addition
	Href string `json:"href,omitempty"`
	// Determine whether the link is an absolute URL or not
	Absolute bool `json:"absolute,omitempty"`
}