/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package harbor
// LdapFailedImportUser struct for LdapFailedImportUser
type LdapFailedImportUser struct {
	// the uid can't add to system.
	Uid string `json:"uid,omitempty"`
	// fail reason.
	Error string `json:"error,omitempty"`
}
