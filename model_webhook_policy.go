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
// WebhookPolicy The webhook policy object
type WebhookPolicy struct {
	// The webhook policy ID.
	Id int64 `json:"id,omitempty"`
	// The name of webhook policy.
	Name string `json:"name,omitempty"`
	// The description of webhook policy.
	Description string `json:"description,omitempty"`
	// The project ID of webhook policy.
	ProjectId int32 `json:"project_id,omitempty"`
	Targets []WebhookTargetObject `json:"targets,omitempty"`
	EventTypes []string `json:"event_types,omitempty"`
	// The creator of the webhook policy.
	Creator string `json:"creator,omitempty"`
	// The create time of the webhook policy.
	CreationTime time.Time `json:"creation_time,omitempty"`
	// The update time of the webhook policy.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Whether the webhook policy is enabled or not.
	Enabled bool `json:"enabled,omitempty"`
}
