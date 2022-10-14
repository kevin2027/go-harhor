/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ConfigurationsResponseScanAllPolicyParameter The parameters of the policy, the values are dependent on the type of the policy.
type ConfigurationsResponseScanAllPolicyParameter struct {
	// The offset in seconds of UTC 0 o'clock, only valid when the policy type is \"daily\"
	DailyTime int32 `json:"daily_time,omitempty"`
}
