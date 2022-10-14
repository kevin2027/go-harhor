/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// AuthproxySetting struct for AuthproxySetting
type AuthproxySetting struct {
	// The fully qualified URI of login endpoint of authproxy, such as 'https://192.168.1.2:8443/login'
	Endpoint string `json:"endpoint,omitempty"`
	// The fully qualified URI of token review endpoint of authproxy, such as 'https://192.168.1.2:8443/tokenreview'
	TokenreivewEndpoint string `json:"tokenreivew_endpoint,omitempty"`
	// The flag to determine whether Harbor can skip search the user/group when adding him as a member.
	SkipSearch bool `json:"skip_search,omitempty"`
	// The flag to determine whether Harbor should verify the certificate when connecting to the auth proxy.
	VerifyCert bool `json:"verify_cert,omitempty"`
	// The certificate to be pinned when connecting auth proxy.
	ServerCertificate string `json:"server_certificate,omitempty"`
}
