/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package harbor
import (
	"time"
)
// WebhookJob The webhook job.
type WebhookJob struct {
	// The webhook job ID.
	Id int64 `json:"id,omitempty"`
	// The webhook policy ID.
	PolicyId int64 `json:"policy_id,omitempty"`
	// The webhook job event type.
	EventType string `json:"event_type,omitempty"`
	// The webhook job notify type.
	NotifyType string `json:"notify_type,omitempty"`
	// The webhook job status.
	Status string `json:"status,omitempty"`
	// The webhook job notify detailed data.
	JobDetail string `json:"job_detail,omitempty"`
	// The webhook job creation time.
	CreationTime time.Time `json:"creation_time,omitempty"`
	// The webhook job update time.
	UpdateTime time.Time `json:"update_time,omitempty"`
}
