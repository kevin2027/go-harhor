/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package harbor
// Stats Stats provides the overall progress of the scan all process.
type Stats struct {
	// The total number of scan processes triggered by the scan all action
	Total int32 `json:"total,omitempty"`
	// The number of the finished scan processes triggered by the scan all action
	Completed int32 `json:"completed,omitempty"`
	// The metrics data for the each status
	Metrics map[string]int32 `json:"metrics,omitempty"`
	// A flag indicating job status of scan all.
	Ongoing bool `json:"ongoing,omitempty"`
	// The trigger of the scan all job.
	Trigger string `json:"trigger,omitempty"`
}
