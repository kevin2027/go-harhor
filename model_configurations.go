/*
 * harbor
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Configurations struct for Configurations
type Configurations struct {
	// The auth mode of current system, such as \"db_auth\", \"ldap_auth\", \"oidc_auth\"
	AuthMode string `json:"auth_mode,omitempty"`
	// The sender name for Email notification.
	EmailFrom string `json:"email_from,omitempty"`
	// The hostname of SMTP server that sends Email notification.
	EmailHost string `json:"email_host,omitempty"`
	// By default it's empty so the email_username is picked
	EmailIdentity string `json:"email_identity,omitempty"`
	// Whether or not the certificate will be verified when Harbor tries to access the email server.
	EmailInsecure bool `json:"email_insecure,omitempty"`
	// Email password
	EmailPassword string `json:"email_password,omitempty"`
	// The port of SMTP server
	EmailPort int32 `json:"email_port,omitempty"`
	// When it''s set to true the system will access Email server via TLS by default.  If it''s set to false, it still will handle \"STARTTLS\" from server side.
	EmailSsl bool `json:"email_ssl,omitempty"`
	// The username for authenticate against SMTP server
	EmailUsername string `json:"email_username,omitempty"`
	// The Base DN for LDAP binding.
	LdapBaseDn string `json:"ldap_base_dn,omitempty"`
	// The filter for LDAP search
	LdapFilter string `json:"ldap_filter,omitempty"`
	// The base DN to search LDAP group.
	LdapGroupBaseDn string `json:"ldap_group_base_dn,omitempty"`
	// Specify the ldap group which have the same privilege with Harbor admin
	LdapGroupAdminDn string `json:"ldap_group_admin_dn,omitempty"`
	// The attribute which is used as identity of the LDAP group, default is cn.'
	LdapGroupAttributeName string `json:"ldap_group_attribute_name,omitempty"`
	// The filter to search the ldap group
	LdapGroupSearchFilter string `json:"ldap_group_search_filter,omitempty"`
	// The scope to search ldap group. ''0-LDAP_SCOPE_BASE, 1-LDAP_SCOPE_ONELEVEL, 2-LDAP_SCOPE_SUBTREE''
	LdapGroupSearchScope int32 `json:"ldap_group_search_scope,omitempty"`
	// The scope to search ldap users,'0-LDAP_SCOPE_BASE, 1-LDAP_SCOPE_ONELEVEL, 2-LDAP_SCOPE_SUBTREE'
	LdapScope int32 `json:"ldap_scope,omitempty"`
	// The DN of the user to do the search.
	LdapSearchDn string `json:"ldap_search_dn,omitempty"`
	// The password of the ldap search dn
	LdapSearchPassword string `json:"ldap_search_password,omitempty"`
	// Timeout in seconds for connection to LDAP server
	LdapTimeout int32 `json:"ldap_timeout,omitempty"`
	// The attribute which is used as identity for the LDAP binding, such as \"CN\" or \"SAMAccountname\"
	LdapUid string `json:"ldap_uid,omitempty"`
	// The URL of LDAP server
	LdapUrl string `json:"ldap_url,omitempty"`
	// Whether verify your OIDC server certificate, disable it if your OIDC server is hosted via self-hosted certificate.
	LdapVerifyCert bool `json:"ldap_verify_cert,omitempty"`
	// The user attribute to identify the group membership
	LdapGroupMembershipAttribute string `json:"ldap_group_membership_attribute,omitempty"`
	// Indicate who can create projects, it could be ''adminonly'' or ''everyone''.
	ProjectCreationRestriction string `json:"project_creation_restriction,omitempty"`
	// The flag to indicate whether Harbor is in readonly mode.
	ReadOnly bool `json:"read_only,omitempty"`
	// Whether the Harbor instance supports self-registration.  If it''s set to false, admin need to add user to the instance.
	SelfRegistration bool `json:"self_registration,omitempty"`
	// The expiration time of the token for internal Registry, in minutes.
	TokenExpiration int32 `json:"token_expiration,omitempty"`
	// The client id of UAA
	UaaClientId string `json:"uaa_client_id,omitempty"`
	// The client secret of the UAA
	UaaClientSecret string `json:"uaa_client_secret,omitempty"`
	// The endpoint of the UAA
	UaaEndpoint string `json:"uaa_endpoint,omitempty"`
	// Verify the certificate in UAA server
	UaaVerifyCert bool `json:"uaa_verify_cert,omitempty"`
	// The endpoint of the HTTP auth
	HttpAuthproxyEndpoint string `json:"http_authproxy_endpoint,omitempty"`
	// The token review endpoint
	HttpAuthproxyTokenreviewEndpoint string `json:"http_authproxy_tokenreview_endpoint,omitempty"`
	// The group which has the harbor admin privileges
	HttpAuthproxyAdminGroups string `json:"http_authproxy_admin_groups,omitempty"`
	// The username which has the harbor admin privileges
	HttpAuthproxyAdminUsernames string `json:"http_authproxy_admin_usernames,omitempty"`
	// Verify the HTTP auth provider's certificate
	HttpAuthproxyVerifyCert bool `json:"http_authproxy_verify_cert,omitempty"`
	// Search user before onboard
	HttpAuthproxySkipSearch bool `json:"http_authproxy_skip_search,omitempty"`
	// The certificate of the HTTP auth provider
	HttpAuthproxyServerCertificate string `json:"http_authproxy_server_certificate,omitempty"`
	// The OIDC provider name
	OidcName string `json:"oidc_name,omitempty"`
	// The endpoint of the OIDC provider
	OidcEndpoint string `json:"oidc_endpoint,omitempty"`
	// The client ID of the OIDC provider
	OidcClientId string `json:"oidc_client_id,omitempty"`
	// The OIDC provider secret
	OidcClientSecret string `json:"oidc_client_secret,omitempty"`
	// The attribute claims the group name
	OidcGroupsClaim string `json:"oidc_groups_claim,omitempty"`
	// The OIDC group which has the harbor admin privileges
	OidcAdminGroup string `json:"oidc_admin_group,omitempty"`
	// The scope of the OIDC provider
	OidcScope string `json:"oidc_scope,omitempty"`
	// The attribute claims the username
	OidcUserClaim string `json:"oidc_user_claim,omitempty"`
	// Verify the OIDC provider's certificate'
	OidcVerifyCert bool `json:"oidc_verify_cert,omitempty"`
	// Auto onboard the OIDC user
	OidcAutoOnboard bool `json:"oidc_auto_onboard,omitempty"`
	// Extra parameters to add when redirect request to OIDC provider
	OidcExtraRedirectParms string `json:"oidc_extra_redirect_parms,omitempty"`
	// The robot account token duration in days
	RobotTokenDuration int32 `json:"robot_token_duration,omitempty"`
	// The rebot account name prefix
	RobotNamePrefix string `json:"robot_name_prefix,omitempty"`
	// Enable notification
	NotificationEnable bool `json:"notification_enable,omitempty"`
	// Enable quota per project
	QuotaPerProjectEnable bool `json:"quota_per_project_enable,omitempty"`
	// The storage quota per project
	StoragePerProject int32 `json:"storage_per_project,omitempty"`
	// The audit log forward endpoint
	AuditLogForwardEndpoint string `json:"audit_log_forward_endpoint,omitempty"`
	// Skip audit log database
	SkipAuditLogDatabase bool `json:"skip_audit_log_database,omitempty"`
}
