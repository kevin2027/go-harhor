/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package harbor
// ProjectMemberEntity struct for ProjectMemberEntity
type ProjectMemberEntity struct {
	// the project member id
	Id int32 `json:"id,omitempty"`
	// the project id
	ProjectId int32 `json:"project_id,omitempty"`
	// the name of the group member.
	EntityName string `json:"entity_name,omitempty"`
	// the name of the role
	RoleName string `json:"role_name,omitempty"`
	// the role id
	RoleId int32 `json:"role_id,omitempty"`
	// the id of entity, if the member is a user, it is user_id in user table. if the member is a user group, it is the user group's ID in user_group table.
	EntityId int32 `json:"entity_id,omitempty"`
	// the entity's type, u for user entity, g for group entity.
	EntityType string `json:"entity_type,omitempty"`
}
