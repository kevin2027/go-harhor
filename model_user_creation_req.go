/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package harbor
// UserCreationReq struct for UserCreationReq
type UserCreationReq struct {
	Email string `json:"email,omitempty"`
	Realname string `json:"realname,omitempty"`
	Comment string `json:"comment,omitempty"`
	Password string `json:"password,omitempty"`
	Username string `json:"username,omitempty"`
}
