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
// GcHistory struct for GcHistory
type GcHistory struct {
	// the id of gc job.
	Id int32 `json:"id,omitempty"`
	// the job name of gc job.
	JobName string `json:"job_name,omitempty"`
	// the job kind of gc job.
	JobKind string `json:"job_kind,omitempty"`
	// the job parameters of gc job.
	JobParameters string `json:"job_parameters,omitempty"`
	Schedule ScheduleObj `json:"schedule,omitempty"`
	// the status of gc job.
	JobStatus string `json:"job_status,omitempty"`
	// if gc job was deleted.
	Deleted bool `json:"deleted,omitempty"`
	// the creation time of gc job.
	CreationTime time.Time `json:"creation_time,omitempty"`
	// the update time of gc job.
	UpdateTime time.Time `json:"update_time,omitempty"`
}