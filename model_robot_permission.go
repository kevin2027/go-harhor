/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// RobotPermission struct for RobotPermission
type RobotPermission struct {
	// The kind of the permission
	Kind string `json:"kind,omitempty"`
	// The namespace of the permission
	Namespace string `json:"namespace,omitempty"`
	Access []Access `json:"access,omitempty"`
}
