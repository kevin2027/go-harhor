/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SupportedWebhookEventTypes Supportted webhook event types and notify types.
type SupportedWebhookEventTypes struct {
	EventType []string `json:"event_type,omitempty"`
	NotifyType []string `json:"notify_type,omitempty"`
}