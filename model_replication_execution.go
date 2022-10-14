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
// ReplicationExecution The replication execution
type ReplicationExecution struct {
	// The ID of the execution
	Id int32 `json:"id,omitempty"`
	// The ID if the policy that the execution belongs to
	PolicyId int32 `json:"policy_id,omitempty"`
	// The status of the execution
	Status string `json:"status,omitempty"`
	// The trigger mode
	Trigger string `json:"trigger,omitempty"`
	// The start time
	StartTime time.Time `json:"start_time,omitempty"`
	// The end time
	EndTime time.Time `json:"end_time,omitempty"`
	// The status text
	StatusText string `json:"status_text,omitempty"`
	// The total count of all executions
	Total int32 `json:"total,omitempty"`
	// The count of failed executions
	Failed int32 `json:"failed,omitempty"`
	// The count of succeed executions
	Succeed int32 `json:"succeed,omitempty"`
	// The count of in_progress executions
	InProgress int32 `json:"in_progress,omitempty"`
	// The count of stopped executions
	Stopped int32 `json:"stopped,omitempty"`
}
