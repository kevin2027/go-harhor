# Configurations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | **string** | The auth mode of current system, such as \&quot;db_auth\&quot;, \&quot;ldap_auth\&quot;, \&quot;oidc_auth\&quot; | [optional] 
**EmailFrom** | **string** | The sender name for Email notification. | [optional] 
**EmailHost** | **string** | The hostname of SMTP server that sends Email notification. | [optional] 
**EmailIdentity** | **string** | By default it&#39;s empty so the email_username is picked | [optional] 
**EmailInsecure** | **bool** | Whether or not the certificate will be verified when Harbor tries to access the email server. | [optional] 
**EmailPassword** | **string** | Email password | [optional] 
**EmailPort** | **int32** | The port of SMTP server | [optional] 
**EmailSsl** | **bool** | When it&#39;&#39;s set to true the system will access Email server via TLS by default.  If it&#39;&#39;s set to false, it still will handle \&quot;STARTTLS\&quot; from server side. | [optional] 
**EmailUsername** | **string** | The username for authenticate against SMTP server | [optional] 
**LdapBaseDn** | **string** | The Base DN for LDAP binding. | [optional] 
**LdapFilter** | **string** | The filter for LDAP search | [optional] 
**LdapGroupBaseDn** | **string** | The base DN to search LDAP group. | [optional] 
**LdapGroupAdminDn** | **string** | Specify the ldap group which have the same privilege with Harbor admin | [optional] 
**LdapGroupAttributeName** | **string** | The attribute which is used as identity of the LDAP group, default is cn.&#39; | [optional] 
**LdapGroupSearchFilter** | **string** | The filter to search the ldap group | [optional] 
**LdapGroupSearchScope** | **int32** | The scope to search ldap group. &#39;&#39;0-LDAP_SCOPE_BASE, 1-LDAP_SCOPE_ONELEVEL, 2-LDAP_SCOPE_SUBTREE&#39;&#39; | [optional] 
**LdapScope** | **int32** | The scope to search ldap users,&#39;0-LDAP_SCOPE_BASE, 1-LDAP_SCOPE_ONELEVEL, 2-LDAP_SCOPE_SUBTREE&#39; | [optional] 
**LdapSearchDn** | **string** | The DN of the user to do the search. | [optional] 
**LdapSearchPassword** | **string** | The password of the ldap search dn | [optional] 
**LdapTimeout** | **int32** | Timeout in seconds for connection to LDAP server | [optional] 
**LdapUid** | **string** | The attribute which is used as identity for the LDAP binding, such as \&quot;CN\&quot; or \&quot;SAMAccountname\&quot; | [optional] 
**LdapUrl** | **string** | The URL of LDAP server | [optional] 
**LdapVerifyCert** | **bool** | Whether verify your OIDC server certificate, disable it if your OIDC server is hosted via self-hosted certificate. | [optional] 
**LdapGroupMembershipAttribute** | **string** | The user attribute to identify the group membership | [optional] 
**ProjectCreationRestriction** | **string** | Indicate who can create projects, it could be &#39;&#39;adminonly&#39;&#39; or &#39;&#39;everyone&#39;&#39;. | [optional] 
**ReadOnly** | **bool** | The flag to indicate whether Harbor is in readonly mode. | [optional] 
**SelfRegistration** | **bool** | Whether the Harbor instance supports self-registration.  If it&#39;&#39;s set to false, admin need to add user to the instance. | [optional] 
**TokenExpiration** | **int32** | The expiration time of the token for internal Registry, in minutes. | [optional] 
**UaaClientId** | **string** | The client id of UAA | [optional] 
**UaaClientSecret** | **string** | The client secret of the UAA | [optional] 
**UaaEndpoint** | **string** | The endpoint of the UAA | [optional] 
**UaaVerifyCert** | **bool** | Verify the certificate in UAA server | [optional] 
**HttpAuthproxyEndpoint** | **string** | The endpoint of the HTTP auth | [optional] 
**HttpAuthproxyTokenreviewEndpoint** | **string** | The token review endpoint | [optional] 
**HttpAuthproxyAdminGroups** | **string** | The group which has the harbor admin privileges | [optional] 
**HttpAuthproxyAdminUsernames** | **string** | The username which has the harbor admin privileges | [optional] 
**HttpAuthproxyVerifyCert** | **bool** | Verify the HTTP auth provider&#39;s certificate | [optional] 
**HttpAuthproxySkipSearch** | **bool** | Search user before onboard | [optional] 
**HttpAuthproxyServerCertificate** | **string** | The certificate of the HTTP auth provider | [optional] 
**OidcName** | **string** | The OIDC provider name | [optional] 
**OidcEndpoint** | **string** | The endpoint of the OIDC provider | [optional] 
**OidcClientId** | **string** | The client ID of the OIDC provider | [optional] 
**OidcClientSecret** | **string** | The OIDC provider secret | [optional] 
**OidcGroupsClaim** | **string** | The attribute claims the group name | [optional] 
**OidcAdminGroup** | **string** | The OIDC group which has the harbor admin privileges | [optional] 
**OidcScope** | **string** | The scope of the OIDC provider | [optional] 
**OidcUserClaim** | **string** | The attribute claims the username | [optional] 
**OidcVerifyCert** | **bool** | Verify the OIDC provider&#39;s certificate&#39; | [optional] 
**OidcAutoOnboard** | **bool** | Auto onboard the OIDC user | [optional] 
**OidcExtraRedirectParms** | **string** | Extra parameters to add when redirect request to OIDC provider | [optional] 
**RobotTokenDuration** | **int32** | The robot account token duration in days | [optional] 
**RobotNamePrefix** | **string** | The rebot account name prefix | [optional] 
**NotificationEnable** | **bool** | Enable notification | [optional] 
**QuotaPerProjectEnable** | **bool** | Enable quota per project | [optional] 
**StoragePerProject** | **int32** | The storage quota per project | [optional] 
**AuditLogForwardEndpoint** | **string** | The audit log forward endpoint | [optional] 
**SkipAuditLogDatabase** | **bool** | Skip audit log database | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


