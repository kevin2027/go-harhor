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
// NativeReportSummary The summary for the native report
type NativeReportSummary struct {
	// id of the native scan report
	ReportId string `json:"report_id,omitempty"`
	// The status of the report generating process
	ScanStatus string `json:"scan_status,omitempty"`
	// The overall severity
	Severity string `json:"severity,omitempty"`
	// The seconds spent for generating the report
	Duration int64 `json:"duration,omitempty"`
	Summary VulnerabilitySummary `json:"summary,omitempty"`
	// The start time of the scan process that generating report
	StartTime time.Time `json:"start_time,omitempty"`
	// The end time of the scan process that generating report
	EndTime time.Time `json:"end_time,omitempty"`
	// The complete percent of the scanning which value is between 0 and 100
	CompletePercent int32 `json:"complete_percent,omitempty"`
	Scanner Scanner `json:"scanner,omitempty"`
}
