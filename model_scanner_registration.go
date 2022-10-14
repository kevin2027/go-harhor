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
// ScannerRegistration Registration represents a named configuration for invoking a scanner via its adapter. 
type ScannerRegistration struct {
	// The unique identifier of this registration.
	Uuid string `json:"uuid,omitempty"`
	// The name of this registration.
	Name string `json:"name,omitempty"`
	// An optional description of this registration.
	Description string `json:"description,omitempty"`
	// A base URL of the scanner adapter
	Url string `json:"url,omitempty"`
	// Indicate whether the registration is enabled or not
	Disabled bool `json:"disabled,omitempty"`
	// Indicate if the registration is set as the system default one
	IsDefault bool `json:"is_default,omitempty"`
	// Specify what authentication approach is adopted for the HTTP communications. Supported types Basic\", \"Bearer\" and api key header \"X-ScannerAdapter-API-Key\" 
	Auth string `json:"auth,omitempty"`
	// An optional value of the HTTP Authorization header sent with each request to the Scanner Adapter API. 
	AccessCredential string `json:"access_credential,omitempty"`
	// Indicate if skip the certificate verification when sending HTTP requests
	SkipCertVerify bool `json:"skip_certVerify,omitempty"`
	// Indicate whether use internal registry addr for the scanner to pull content or not
	UseInternalAddr bool `json:"use_internal_addr,omitempty"`
	// The creation time of this registration
	CreateTime time.Time `json:"create_time,omitempty"`
	// The update time of this registration
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Optional property to describe the name of the scanner registration
	Adapter string `json:"adapter,omitempty"`
	// Optional property to describe the vendor of the scanner registration
	Vendor string `json:"vendor,omitempty"`
	// Optional property to describe the version of the scanner registration
	Version string `json:"version,omitempty"`
	// Indicate the healthy of the registration
	Health string `json:"health,omitempty"`
}
