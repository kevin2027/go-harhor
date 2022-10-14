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
// WebhookLastTrigger The webhook policy and last trigger time group by event type.
type WebhookLastTrigger struct {
	// The webhook policy name.
	PolicyName string `json:"policy_name,omitempty"`
	// The webhook event type.
	EventType string `json:"event_type,omitempty"`
	// Whether or not the webhook policy enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The creation time of webhook policy.
	CreationTime time.Time `json:"creation_time,omitempty"`
	// The last trigger time of webhook policy.
	LastTriggerTime time.Time `json:"last_trigger_time,omitempty"`
}
