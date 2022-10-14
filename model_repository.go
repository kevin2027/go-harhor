/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// Repository struct for Repository
type Repository struct {
	// The ID of the repository
	Id int64 `json:"id,omitempty"`
	// The ID of the project that the repository belongs to
	ProjectId int64 `json:"project_id,omitempty"`
	// The name of the repository
	Name string `json:"name,omitempty"`
	// The description of the repository
	Description string `json:"description,omitempty"`
	// The count of the artifacts inside the repository
	ArtifactCount int64 `json:"artifact_count,omitempty"`
	// The count that the artifact inside the repository pulled
	PullCount int64 `json:"pull_count,omitempty"`
	// The creation time of the repository
	CreationTime time.Time `json:"creation_time,omitempty"`
	// The update time of the repository
	UpdateTime time.Time `json:"update_time,omitempty"`
}
