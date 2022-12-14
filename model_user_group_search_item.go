/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package harbor
// UserGroupSearchItem struct for UserGroupSearchItem
type UserGroupSearchItem struct {
	// The ID of the user group
	Id int32 `json:"id,omitempty"`
	// The name of the user group
	GroupName string `json:"group_name,omitempty"`
	// The group type, 1 for LDAP group, 2 for HTTP group, 3 for OIDC group.
	GroupType int32 `json:"group_type,omitempty"`
}
