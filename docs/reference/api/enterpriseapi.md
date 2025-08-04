# EnterpriseApi


All URIs are relative to */api/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createGroupForOrganization**](EnterpriseApi.md#createGroupForOrganization) | **POST** /organizations/{organization}/groups | Create group for organization |
| [**createOauth2Application**](EnterpriseApi.md#createOauth2Application) | **POST** /oauth2-provider/apps | Create OAuth2 application. |
| [**createOauth2ApplicationSecret**](EnterpriseApi.md#createOauth2ApplicationSecret) | **POST** /oauth2-provider/apps/{app}/secrets | Create OAuth2 application secret. |
| [**createProvisionerKey**](EnterpriseApi.md#createProvisionerKey) | **POST** /organizations/{organization}/provisionerkeys | Create provisioner key |
| [**createWorkspaceProxy**](EnterpriseApi.md#createWorkspaceProxy) | **POST** /workspaceproxies | Create workspace proxy |
| [**deleteGroupByName**](EnterpriseApi.md#deleteGroupByName) | **DELETE** /groups/{group} | Delete group by name |
| [**deleteLicense**](EnterpriseApi.md#deleteLicense) | **DELETE** /licenses/{id} | Delete license |
| [**deleteOauth2Application**](EnterpriseApi.md#deleteOauth2Application) | **DELETE** /oauth2-provider/apps/{app} | Delete OAuth2 application. |
| [**deleteOauth2ApplicationSecret**](EnterpriseApi.md#deleteOauth2ApplicationSecret) | **DELETE** /oauth2-provider/apps/{app}/secrets/{secretID} | Delete OAuth2 application secret. |
| [**deleteOauth2ApplicationTokens**](EnterpriseApi.md#deleteOauth2ApplicationTokens) | **DELETE** /oauth2/tokens | Delete OAuth2 application tokens. |
| [**deleteOauth2ClientConfiguration**](EnterpriseApi.md#deleteOauth2ClientConfiguration) | **DELETE** /oauth2/clients/{client_id} | Delete OAuth2 client registration (RFC 7592) |
| [**deleteProvisionerKey**](EnterpriseApi.md#deleteProvisionerKey) | **DELETE** /organizations/{organization}/provisionerkeys/{provisionerkey} | Delete provisioner key |
| [**deleteWorkspaceProxy**](EnterpriseApi.md#deleteWorkspaceProxy) | **DELETE** /workspaceproxies/{workspaceproxy} | Delete workspace proxy |
| [**deregisterWorkspaceProxy**](EnterpriseApi.md#deregisterWorkspaceProxy) | **POST** /workspaceproxies/me/deregister | Deregister workspace proxy |
| [**fetchProvisionerKeyDetails**](EnterpriseApi.md#fetchProvisionerKeyDetails) | **GET** /provisionerkeys/{provisionerkey} | Fetch provisioner key details |
| [**getActiveReplicas**](EnterpriseApi.md#getActiveReplicas) | **GET** /replicas | Get active replicas |
| [**getAppearance**](EnterpriseApi.md#getAppearance) | **GET** /appearance | Get appearance |
| [**getConnectionLogs**](EnterpriseApi.md#getConnectionLogs) | **GET** /connectionlog | Get connection logs |
| [**getEntitlements**](EnterpriseApi.md#getEntitlements) | **GET** /entitlements | Get entitlements |
| [**getGroupById**](EnterpriseApi.md#getGroupById) | **GET** /groups/{group} | Get group by ID |
| [**getGroupByOrganizationAndGroupName**](EnterpriseApi.md#getGroupByOrganizationAndGroupName) | **GET** /organizations/{organization}/groups/{groupName} | Get group by organization and group name |
| [**getGroupIdpSyncSettingsByOrganization**](EnterpriseApi.md#getGroupIdpSyncSettingsByOrganization) | **GET** /organizations/{organization}/settings/idpsync/groups | Get group IdP Sync settings by organization |
| [**getGroups**](EnterpriseApi.md#getGroups) | **GET** /groups | Get groups |
| [**getGroupsByOrganization**](EnterpriseApi.md#getGroupsByOrganization) | **GET** /organizations/{organization}/groups | Get groups by organization |
| [**getLicenses**](EnterpriseApi.md#getLicenses) | **GET** /licenses | Get licenses |
| [**getOauth2Application**](EnterpriseApi.md#getOauth2Application) | **GET** /oauth2-provider/apps/{app} | Get OAuth2 application. |
| [**getOauth2ApplicationSecrets**](EnterpriseApi.md#getOauth2ApplicationSecrets) | **GET** /oauth2-provider/apps/{app}/secrets | Get OAuth2 application secrets. |
| [**getOauth2Applications**](EnterpriseApi.md#getOauth2Applications) | **GET** /oauth2-provider/apps | Get OAuth2 applications. |
| [**getOauth2ClientConfiguration**](EnterpriseApi.md#getOauth2ClientConfiguration) | **GET** /oauth2/clients/{client_id} | Get OAuth2 client configuration (RFC 7592) |
| [**getOrganizationIdpSyncSettings**](EnterpriseApi.md#getOrganizationIdpSyncSettings) | **GET** /settings/idpsync/organization | Get organization IdP Sync settings |
| [**getRoleIdpSyncSettingsByOrganization**](EnterpriseApi.md#getRoleIdpSyncSettingsByOrganization) | **GET** /organizations/{organization}/settings/idpsync/roles | Get role IdP Sync settings by organization |
| [**getTemplateAcls**](EnterpriseApi.md#getTemplateAcls) | **GET** /templates/{template}/acl | Get template ACLs |
| [**getTemplateAvailableAclUsersgroups**](EnterpriseApi.md#getTemplateAvailableAclUsersgroups) | **GET** /templates/{template}/acl/available | Get template available acl users/groups |
| [**getTheAvailableIdpSyncClaimFields**](EnterpriseApi.md#getTheAvailableIdpSyncClaimFields) | **GET** /settings/idpsync/available-fields | Get the available idp sync claim fields |
| [**getTheAvailableOrganizationIdpSyncClaimFields**](EnterpriseApi.md#getTheAvailableOrganizationIdpSyncClaimFields) | **GET** /organizations/{organization}/settings/idpsync/available-fields | Get the available organization idp sync claim fields |
| [**getTheIdpSyncClaimFieldValues**](EnterpriseApi.md#getTheIdpSyncClaimFieldValues) | **GET** /settings/idpsync/field-values | Get the idp sync claim field values |
| [**getTheOrganizationIdpSyncClaimFieldValues**](EnterpriseApi.md#getTheOrganizationIdpSyncClaimFieldValues) | **GET** /organizations/{organization}/settings/idpsync/field-values | Get the organization idp sync claim field values |
| [**getUserQuietHoursSchedule**](EnterpriseApi.md#getUserQuietHoursSchedule) | **GET** /users/{user}/quiet-hours | Get user quiet hours schedule |
| [**getWorkspaceProxies**](EnterpriseApi.md#getWorkspaceProxies) | **GET** /workspaceproxies | Get workspace proxies |
| [**getWorkspaceProxy**](EnterpriseApi.md#getWorkspaceProxy) | **GET** /workspaceproxies/{workspaceproxy} | Get workspace proxy |
| [**getWorkspaceProxyCryptoKeys**](EnterpriseApi.md#getWorkspaceProxyCryptoKeys) | **GET** /workspaceproxies/me/crypto-keys | Get workspace proxy crypto keys |
| [**getWorkspaceQuotaByUser**](EnterpriseApi.md#getWorkspaceQuotaByUser) | **GET** /organizations/{organization}/members/{user}/workspace-quota | Get workspace quota by user |
| [**getWorkspaceQuotaByUserDeprecated**](EnterpriseApi.md#getWorkspaceQuotaByUserDeprecated) | **GET** /workspace-quota/{user} | Get workspace quota by user deprecated |
| [**issueSignedAppTokenForReconnectingPty**](EnterpriseApi.md#issueSignedAppTokenForReconnectingPty) | **POST** /applications/reconnecting-pty-signed-token | Issue signed app token for reconnecting PTY |
| [**issueSignedWorkspaceAppToken**](EnterpriseApi.md#issueSignedWorkspaceAppToken) | **POST** /workspaceproxies/me/issue-signed-app-token | Issue signed workspace app token |
| [**listProvisionerKey**](EnterpriseApi.md#listProvisionerKey) | **GET** /organizations/{organization}/provisionerkeys | List provisioner key |
| [**listProvisionerKeyDaemons**](EnterpriseApi.md#listProvisionerKeyDaemons) | **GET** /organizations/{organization}/provisionerkeys/daemons | List provisioner key daemons |
| [**oauth2AuthorizationRequestGet**](EnterpriseApi.md#oauth2AuthorizationRequestGet) | **GET** /oauth2/authorize | OAuth2 authorization request (GET - show authorization page). |
| [**oauth2AuthorizationRequestPost**](EnterpriseApi.md#oauth2AuthorizationRequestPost) | **POST** /oauth2/authorize | OAuth2 authorization request (POST - process authorization). |
| [**oauth2AuthorizationServerMetadata**](EnterpriseApi.md#oauth2AuthorizationServerMetadata) | **GET** /.well-known/oauth-authorization-server | OAuth2 authorization server metadata. |
| [**oauth2DynamicClientRegistration**](EnterpriseApi.md#oauth2DynamicClientRegistration) | **POST** /oauth2/register | OAuth2 dynamic client registration (RFC 7591) |
| [**oauth2ProtectedResourceMetadata**](EnterpriseApi.md#oauth2ProtectedResourceMetadata) | **GET** /.well-known/oauth-protected-resource | OAuth2 protected resource metadata. |
| [**oauth2TokenExchange**](EnterpriseApi.md#oauth2TokenExchange) | **POST** /oauth2/tokens | OAuth2 token exchange. |
| [**putOauth2ClientConfiguration**](EnterpriseApi.md#putOauth2ClientConfiguration) | **PUT** /oauth2/clients/{client_id} | Update OAuth2 client configuration (RFC 7592) |
| [**registerWorkspaceProxy**](EnterpriseApi.md#registerWorkspaceProxy) | **POST** /workspaceproxies/me/register | Register workspace proxy |
| [**reportWorkspaceAppStats**](EnterpriseApi.md#reportWorkspaceAppStats) | **POST** /workspaceproxies/me/app-stats | Report workspace app stats |
| [**scimCreateNewUser**](EnterpriseApi.md#scimCreateNewUser) | **POST** /scim/v2/Users | SCIM 2.0: Create new user |
| [**scimGetServiceProviderConfig**](EnterpriseApi.md#scimGetServiceProviderConfig) | **GET** /scim/v2/ServiceProviderConfig | SCIM 2.0: Service Provider Config |
| [**scimGetUserById**](EnterpriseApi.md#scimGetUserById) | **GET** /scim/v2/Users/{id} | SCIM 2.0: Get user by ID |
| [**scimGetUsers**](EnterpriseApi.md#scimGetUsers) | **GET** /scim/v2/Users | SCIM 2.0: Get users |
| [**scimReplaceUserStatus**](EnterpriseApi.md#scimReplaceUserStatus) | **PUT** /scim/v2/Users/{id} | SCIM 2.0: Replace user account |
| [**scimUpdateUserStatus**](EnterpriseApi.md#scimUpdateUserStatus) | **PATCH** /scim/v2/Users/{id} | SCIM 2.0: Update user account |
| [**serveProvisionerDaemon**](EnterpriseApi.md#serveProvisionerDaemon) | **GET** /organizations/{organization}/provisionerdaemons/serve | Serve provisioner daemon |
| [**updateAppearance**](EnterpriseApi.md#updateAppearance) | **PUT** /appearance | Update appearance |
| [**updateGroupByName**](EnterpriseApi.md#updateGroupByName) | **PATCH** /groups/{group} | Update group by name |
| [**updateGroupIdpSyncConfig**](EnterpriseApi.md#updateGroupIdpSyncConfig) | **PATCH** /organizations/{organization}/settings/idpsync/groups/config | Update group IdP Sync config |
| [**updateGroupIdpSyncMapping**](EnterpriseApi.md#updateGroupIdpSyncMapping) | **PATCH** /organizations/{organization}/settings/idpsync/groups/mapping | Update group IdP Sync mapping |
| [**updateGroupIdpSyncSettingsByOrganization**](EnterpriseApi.md#updateGroupIdpSyncSettingsByOrganization) | **PATCH** /organizations/{organization}/settings/idpsync/groups | Update group IdP Sync settings by organization |
| [**updateNotificationTemplateDispatchMethod**](EnterpriseApi.md#updateNotificationTemplateDispatchMethod) | **PUT** /notifications/templates/{notification_template}/method | Update notification template dispatch method |
| [**updateOauth2Application**](EnterpriseApi.md#updateOauth2Application) | **PUT** /oauth2-provider/apps/{app} | Update OAuth2 application. |
| [**updateOrganizationIdpSyncConfig**](EnterpriseApi.md#updateOrganizationIdpSyncConfig) | **PATCH** /settings/idpsync/organization/config | Update organization IdP Sync config |
| [**updateOrganizationIdpSyncMapping**](EnterpriseApi.md#updateOrganizationIdpSyncMapping) | **PATCH** /settings/idpsync/organization/mapping | Update organization IdP Sync mapping |
| [**updateOrganizationIdpSyncSettings**](EnterpriseApi.md#updateOrganizationIdpSyncSettings) | **PATCH** /settings/idpsync/organization | Update organization IdP Sync settings |
| [**updateRoleIdpSyncConfig**](EnterpriseApi.md#updateRoleIdpSyncConfig) | **PATCH** /organizations/{organization}/settings/idpsync/roles/config | Update role IdP Sync config |
| [**updateRoleIdpSyncMapping**](EnterpriseApi.md#updateRoleIdpSyncMapping) | **PATCH** /organizations/{organization}/settings/idpsync/roles/mapping | Update role IdP Sync mapping |
| [**updateRoleIdpSyncSettingsByOrganization**](EnterpriseApi.md#updateRoleIdpSyncSettingsByOrganization) | **PATCH** /organizations/{organization}/settings/idpsync/roles | Update role IdP Sync settings by organization |
| [**updateTemplateAcl**](EnterpriseApi.md#updateTemplateAcl) | **PATCH** /templates/{template}/acl | Update template ACL |
| [**updateUserQuietHoursSchedule**](EnterpriseApi.md#updateUserQuietHoursSchedule) | **PUT** /users/{user}/quiet-hours | Update user quiet hours schedule |
| [**updateWorkspaceProxy**](EnterpriseApi.md#updateWorkspaceProxy) | **PATCH** /workspaceproxies/{workspaceproxy} | Update workspace proxy |
| [**workspaceProxyCoordinate**](EnterpriseApi.md#workspaceProxyCoordinate) | **GET** /workspaceproxies/me/coordinate | Workspace Proxy Coordinate |


<a name="createGroupForOrganization"></a>
# **createGroupForOrganization**
> codersdk.Group createGroupForOrganization(organization, request)

Create group for organization

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **String**| Organization ID | [default to null] |
| **request** | [**codersdk.CreateGroupRequest**](../Models/codersdk.CreateGroupRequest.md)| Create group request | |

### Return type

[**codersdk.Group**](../Models/codersdk.Group.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createOauth2Application"></a>
# **createOauth2Application**
> codersdk.OAuth2ProviderApp createOauth2Application(request)

Create OAuth2 application.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**codersdk.PostOAuth2ProviderAppRequest**](../Models/codersdk.PostOAuth2ProviderAppRequest.md)| The OAuth2 application to create. | |

### Return type

[**codersdk.OAuth2ProviderApp**](../Models/codersdk.OAuth2ProviderApp.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createOauth2ApplicationSecret"></a>
# **createOauth2ApplicationSecret**
> List createOauth2ApplicationSecret(app)

Create OAuth2 application secret.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **app** | **String**| App ID | [default to null] |

### Return type

[**List**](../Models/codersdk.OAuth2ProviderAppSecretFull.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="createProvisionerKey"></a>
# **createProvisionerKey**
> codersdk.CreateProvisionerKeyResponse createProvisionerKey(organization)

Create provisioner key

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **String**| Organization ID | [default to null] |

### Return type

[**codersdk.CreateProvisionerKeyResponse**](../Models/codersdk.CreateProvisionerKeyResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="createWorkspaceProxy"></a>
# **createWorkspaceProxy**
> codersdk.WorkspaceProxy createWorkspaceProxy(request)

Create workspace proxy

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**codersdk.CreateWorkspaceProxyRequest**](../Models/codersdk.CreateWorkspaceProxyRequest.md)| Create workspace proxy request | |

### Return type

[**codersdk.WorkspaceProxy**](../Models/codersdk.WorkspaceProxy.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteGroupByName"></a>
# **deleteGroupByName**
> codersdk.Group deleteGroupByName(group)

Delete group by name

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **group** | **String**| Group name | [default to null] |

### Return type

[**codersdk.Group**](../Models/codersdk.Group.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deleteLicense"></a>
# **deleteLicense**
> deleteLicense(id)

Delete license

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | **BigDecimal**| License ID | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="deleteOauth2Application"></a>
# **deleteOauth2Application**
> deleteOauth2Application(app)

Delete OAuth2 application.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **app** | **String**| App ID | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="deleteOauth2ApplicationSecret"></a>
# **deleteOauth2ApplicationSecret**
> deleteOauth2ApplicationSecret(app, secretID)

Delete OAuth2 application secret.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **app** | **String**| App ID | [default to null] |
| **secretID** | **String**| Secret ID | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="deleteOauth2ApplicationTokens"></a>
# **deleteOauth2ApplicationTokens**
> deleteOauth2ApplicationTokens(client\_id)

Delete OAuth2 application tokens.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **client\_id** | **String**| Client ID | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="deleteOauth2ClientConfiguration"></a>
# **deleteOauth2ClientConfiguration**
> deleteOauth2ClientConfiguration(client\_id)

Delete OAuth2 client registration (RFC 7592)

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **client\_id** | **String**| Client ID | [default to null] |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="deleteProvisionerKey"></a>
# **deleteProvisionerKey**
> deleteProvisionerKey(organization, provisionerkey)

Delete provisioner key

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **String**| Organization ID | [default to null] |
| **provisionerkey** | **String**| Provisioner key name | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="deleteWorkspaceProxy"></a>
# **deleteWorkspaceProxy**
> codersdk.Response deleteWorkspaceProxy(workspaceproxy)

Delete workspace proxy

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspaceproxy** | **UUID**| Proxy ID or name | [default to null] |

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deregisterWorkspaceProxy"></a>
# **deregisterWorkspaceProxy**
> deregisterWorkspaceProxy(request)

Deregister workspace proxy

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**wsproxysdk.DeregisterWorkspaceProxyRequest**](../Models/wsproxysdk.DeregisterWorkspaceProxyRequest.md)| Deregister workspace proxy request | |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

<a name="fetchProvisionerKeyDetails"></a>
# **fetchProvisionerKeyDetails**
> codersdk.ProvisionerKey fetchProvisionerKeyDetails(provisionerkey)

Fetch provisioner key details

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **provisionerkey** | **String**| Provisioner Key | [default to null] |

### Return type

[**codersdk.ProvisionerKey**](../Models/codersdk.ProvisionerKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getActiveReplicas"></a>
# **getActiveReplicas**
> List getActiveReplicas()

Get active replicas

### Parameters
This endpoint does not need any parameter.

### Return type

[**List**](../Models/codersdk.Replica.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAppearance"></a>
# **getAppearance**
> codersdk.AppearanceConfig getAppearance()

Get appearance

### Parameters
This endpoint does not need any parameter.

### Return type

[**codersdk.AppearanceConfig**](../Models/codersdk.AppearanceConfig.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getConnectionLogs"></a>
# **getConnectionLogs**
> codersdk.ConnectionLogResponse getConnectionLogs(limit, q, offset)

Get connection logs

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **limit** | **Integer**| Page limit | [default to null] |
| **q** | **String**| Search query | [optional] [default to null] |
| **offset** | **Integer**| Page offset | [optional] [default to null] |

### Return type

[**codersdk.ConnectionLogResponse**](../Models/codersdk.ConnectionLogResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getEntitlements"></a>
# **getEntitlements**
> codersdk.Entitlements getEntitlements()

Get entitlements

### Parameters
This endpoint does not need any parameter.

### Return type

[**codersdk.Entitlements**](../Models/codersdk.Entitlements.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getGroupById"></a>
# **getGroupById**
> codersdk.Group getGroupById(group)

Get group by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **group** | **String**| Group id | [default to null] |

### Return type

[**codersdk.Group**](../Models/codersdk.Group.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getGroupByOrganizationAndGroupName"></a>
# **getGroupByOrganizationAndGroupName**
> codersdk.Group getGroupByOrganizationAndGroupName(organization, groupName)

Get group by organization and group name

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |
| **groupName** | **String**| Group name | [default to null] |

### Return type

[**codersdk.Group**](../Models/codersdk.Group.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getGroupIdpSyncSettingsByOrganization"></a>
# **getGroupIdpSyncSettingsByOrganization**
> codersdk.GroupSyncSettings getGroupIdpSyncSettingsByOrganization(organization)

Get group IdP Sync settings by organization

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |

### Return type

[**codersdk.GroupSyncSettings**](../Models/codersdk.GroupSyncSettings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getGroups"></a>
# **getGroups**
> List getGroups(organization, has\_member, group\_ids)

Get groups

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **String**| Organization ID or name | [default to null] |
| **has\_member** | **String**| User ID or name | [default to null] |
| **group\_ids** | **String**| Comma separated list of group IDs | [default to null] |

### Return type

[**List**](../Models/codersdk.Group.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getGroupsByOrganization"></a>
# **getGroupsByOrganization**
> List getGroupsByOrganization(organization)

Get groups by organization

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |

### Return type

[**List**](../Models/codersdk.Group.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getLicenses"></a>
# **getLicenses**
> List getLicenses()

Get licenses

### Parameters
This endpoint does not need any parameter.

### Return type

[**List**](../Models/codersdk.License.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getOauth2Application"></a>
# **getOauth2Application**
> codersdk.OAuth2ProviderApp getOauth2Application(app)

Get OAuth2 application.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **app** | **String**| App ID | [default to null] |

### Return type

[**codersdk.OAuth2ProviderApp**](../Models/codersdk.OAuth2ProviderApp.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getOauth2ApplicationSecrets"></a>
# **getOauth2ApplicationSecrets**
> List getOauth2ApplicationSecrets(app)

Get OAuth2 application secrets.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **app** | **String**| App ID | [default to null] |

### Return type

[**List**](../Models/codersdk.OAuth2ProviderAppSecret.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getOauth2Applications"></a>
# **getOauth2Applications**
> List getOauth2Applications(user\_id)

Get OAuth2 applications.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user\_id** | **String**| Filter by applications authorized for a user | [optional] [default to null] |

### Return type

[**List**](../Models/codersdk.OAuth2ProviderApp.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getOauth2ClientConfiguration"></a>
# **getOauth2ClientConfiguration**
> codersdk.OAuth2ClientConfiguration getOauth2ClientConfiguration(client\_id)

Get OAuth2 client configuration (RFC 7592)

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **client\_id** | **String**| Client ID | [default to null] |

### Return type

[**codersdk.OAuth2ClientConfiguration**](../Models/codersdk.OAuth2ClientConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getOrganizationIdpSyncSettings"></a>
# **getOrganizationIdpSyncSettings**
> codersdk.OrganizationSyncSettings getOrganizationIdpSyncSettings()

Get organization IdP Sync settings

### Parameters
This endpoint does not need any parameter.

### Return type

[**codersdk.OrganizationSyncSettings**](../Models/codersdk.OrganizationSyncSettings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getRoleIdpSyncSettingsByOrganization"></a>
# **getRoleIdpSyncSettingsByOrganization**
> codersdk.RoleSyncSettings getRoleIdpSyncSettingsByOrganization(organization)

Get role IdP Sync settings by organization

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |

### Return type

[**codersdk.RoleSyncSettings**](../Models/codersdk.RoleSyncSettings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTemplateAcls"></a>
# **getTemplateAcls**
> codersdk.TemplateACL getTemplateAcls(template)

Get template ACLs

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **template** | **UUID**| Template ID | [default to null] |

### Return type

[**codersdk.TemplateACL**](../Models/codersdk.TemplateACL.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTemplateAvailableAclUsersgroups"></a>
# **getTemplateAvailableAclUsersgroups**
> List getTemplateAvailableAclUsersgroups(template)

Get template available acl users/groups

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **template** | **UUID**| Template ID | [default to null] |

### Return type

[**List**](../Models/codersdk.ACLAvailable.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTheAvailableIdpSyncClaimFields"></a>
# **getTheAvailableIdpSyncClaimFields**
> List getTheAvailableIdpSyncClaimFields(organization)

Get the available idp sync claim fields

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |

### Return type

**List**

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTheAvailableOrganizationIdpSyncClaimFields"></a>
# **getTheAvailableOrganizationIdpSyncClaimFields**
> List getTheAvailableOrganizationIdpSyncClaimFields(organization)

Get the available organization idp sync claim fields

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |

### Return type

**List**

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTheIdpSyncClaimFieldValues"></a>
# **getTheIdpSyncClaimFieldValues**
> List getTheIdpSyncClaimFieldValues(organization, claimField)

Get the idp sync claim field values

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |
| **claimField** | **String**| Claim Field | [default to null] |

### Return type

**List**

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTheOrganizationIdpSyncClaimFieldValues"></a>
# **getTheOrganizationIdpSyncClaimFieldValues**
> List getTheOrganizationIdpSyncClaimFieldValues(organization, claimField)

Get the organization idp sync claim field values

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |
| **claimField** | **String**| Claim Field | [default to null] |

### Return type

**List**

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserQuietHoursSchedule"></a>
# **getUserQuietHoursSchedule**
> List getUserQuietHoursSchedule(user)

Get user quiet hours schedule

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **UUID**| User ID | [default to null] |

### Return type

[**List**](../Models/codersdk.UserQuietHoursScheduleResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getWorkspaceProxies"></a>
# **getWorkspaceProxies**
> List getWorkspaceProxies()

Get workspace proxies

### Parameters
This endpoint does not need any parameter.

### Return type

[**List**](../Models/codersdk.RegionsResponse-codersdk_WorkspaceProxy.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getWorkspaceProxy"></a>
# **getWorkspaceProxy**
> codersdk.WorkspaceProxy getWorkspaceProxy(workspaceproxy)

Get workspace proxy

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspaceproxy** | **UUID**| Proxy ID or name | [default to null] |

### Return type

[**codersdk.WorkspaceProxy**](../Models/codersdk.WorkspaceProxy.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getWorkspaceProxyCryptoKeys"></a>
# **getWorkspaceProxyCryptoKeys**
> wsproxysdk.CryptoKeysResponse getWorkspaceProxyCryptoKeys(feature)

Get workspace proxy crypto keys

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **feature** | **String**| Feature key | [default to null] |

### Return type

[**wsproxysdk.CryptoKeysResponse**](../Models/wsproxysdk.CryptoKeysResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getWorkspaceQuotaByUser"></a>
# **getWorkspaceQuotaByUser**
> codersdk.WorkspaceQuota getWorkspaceQuotaByUser(user, organization)

Get workspace quota by user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |
| **organization** | **UUID**| Organization ID | [default to null] |

### Return type

[**codersdk.WorkspaceQuota**](../Models/codersdk.WorkspaceQuota.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getWorkspaceQuotaByUserDeprecated"></a>
# **getWorkspaceQuotaByUserDeprecated**
> codersdk.WorkspaceQuota getWorkspaceQuotaByUserDeprecated(user)

Get workspace quota by user deprecated

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |

### Return type

[**codersdk.WorkspaceQuota**](../Models/codersdk.WorkspaceQuota.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="issueSignedAppTokenForReconnectingPty"></a>
# **issueSignedAppTokenForReconnectingPty**
> codersdk.IssueReconnectingPTYSignedTokenResponse issueSignedAppTokenForReconnectingPty(request)

Issue signed app token for reconnecting PTY

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**codersdk.IssueReconnectingPTYSignedTokenRequest**](../Models/codersdk.IssueReconnectingPTYSignedTokenRequest.md)| Issue reconnecting PTY signed token request | |

### Return type

[**codersdk.IssueReconnectingPTYSignedTokenResponse**](../Models/codersdk.IssueReconnectingPTYSignedTokenResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="issueSignedWorkspaceAppToken"></a>
# **issueSignedWorkspaceAppToken**
> wsproxysdk.IssueSignedAppTokenResponse issueSignedWorkspaceAppToken(request)

Issue signed workspace app token

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**workspaceapps.IssueTokenRequest**](../Models/workspaceapps.IssueTokenRequest.md)| Issue signed app token request | |

### Return type

[**wsproxysdk.IssueSignedAppTokenResponse**](../Models/wsproxysdk.IssueSignedAppTokenResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="listProvisionerKey"></a>
# **listProvisionerKey**
> List listProvisionerKey(organization)

List provisioner key

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **String**| Organization ID | [default to null] |

### Return type

[**List**](../Models/codersdk.ProvisionerKey.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listProvisionerKeyDaemons"></a>
# **listProvisionerKeyDaemons**
> List listProvisionerKeyDaemons(organization)

List provisioner key daemons

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **String**| Organization ID | [default to null] |

### Return type

[**List**](../Models/codersdk.ProvisionerKeyDaemons.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="oauth2AuthorizationRequestGet"></a>
# **oauth2AuthorizationRequestGet**
> oauth2AuthorizationRequestGet(client\_id, state, response\_type, redirect\_uri, scope)

OAuth2 authorization request (GET - show authorization page).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **client\_id** | **String**| Client ID | [default to null] |
| **state** | **String**| A random unguessable string | [default to null] |
| **response\_type** | **String**| Response type | [default to null] [enum: code] |
| **redirect\_uri** | **String**| Redirect here after authorization | [optional] [default to null] |
| **scope** | **String**| Token scopes (currently ignored) | [optional] [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="oauth2AuthorizationRequestPost"></a>
# **oauth2AuthorizationRequestPost**
> oauth2AuthorizationRequestPost(client\_id, state, response\_type, redirect\_uri, scope)

OAuth2 authorization request (POST - process authorization).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **client\_id** | **String**| Client ID | [default to null] |
| **state** | **String**| A random unguessable string | [default to null] |
| **response\_type** | **String**| Response type | [default to null] [enum: code] |
| **redirect\_uri** | **String**| Redirect here after authorization | [optional] [default to null] |
| **scope** | **String**| Token scopes (currently ignored) | [optional] [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="oauth2AuthorizationServerMetadata"></a>
# **oauth2AuthorizationServerMetadata**
> codersdk.OAuth2AuthorizationServerMetadata oauth2AuthorizationServerMetadata()

OAuth2 authorization server metadata.

### Parameters
This endpoint does not need any parameter.

### Return type

[**codersdk.OAuth2AuthorizationServerMetadata**](../Models/codersdk.OAuth2AuthorizationServerMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="oauth2DynamicClientRegistration"></a>
# **oauth2DynamicClientRegistration**
> codersdk.OAuth2ClientRegistrationResponse oauth2DynamicClientRegistration(request)

OAuth2 dynamic client registration (RFC 7591)

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**codersdk.OAuth2ClientRegistrationRequest**](../Models/codersdk.OAuth2ClientRegistrationRequest.md)| Client registration request | |

### Return type

[**codersdk.OAuth2ClientRegistrationResponse**](../Models/codersdk.OAuth2ClientRegistrationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="oauth2ProtectedResourceMetadata"></a>
# **oauth2ProtectedResourceMetadata**
> codersdk.OAuth2ProtectedResourceMetadata oauth2ProtectedResourceMetadata()

OAuth2 protected resource metadata.

### Parameters
This endpoint does not need any parameter.

### Return type

[**codersdk.OAuth2ProtectedResourceMetadata**](../Models/codersdk.OAuth2ProtectedResourceMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="oauth2TokenExchange"></a>
# **oauth2TokenExchange**
> oauth2.Token oauth2TokenExchange(grant\_type, client\_id, client\_secret, code, refresh\_token)

OAuth2 token exchange.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **grant\_type** | **String**| Grant type | [default to null] [enum: authorization_code, refresh_token] |
| **client\_id** | **String**| Client ID, required if grant_type&#x3D;authorization_code | [optional] [default to null] |
| **client\_secret** | **String**| Client secret, required if grant_type&#x3D;authorization_code | [optional] [default to null] |
| **code** | **String**| Authorization code, required if grant_type&#x3D;authorization_code | [optional] [default to null] |
| **refresh\_token** | **String**| Refresh token, required if grant_type&#x3D;refresh_token | [optional] [default to null] |

### Return type

[**oauth2.Token**](../Models/oauth2.Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

<a name="putOauth2ClientConfiguration"></a>
# **putOauth2ClientConfiguration**
> codersdk.OAuth2ClientConfiguration putOauth2ClientConfiguration(client\_id, request)

Update OAuth2 client configuration (RFC 7592)

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **client\_id** | **String**| Client ID | [default to null] |
| **request** | [**codersdk.OAuth2ClientRegistrationRequest**](../Models/codersdk.OAuth2ClientRegistrationRequest.md)| Client update request | |

### Return type

[**codersdk.OAuth2ClientConfiguration**](../Models/codersdk.OAuth2ClientConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="registerWorkspaceProxy"></a>
# **registerWorkspaceProxy**
> wsproxysdk.RegisterWorkspaceProxyResponse registerWorkspaceProxy(request)

Register workspace proxy

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**wsproxysdk.RegisterWorkspaceProxyRequest**](../Models/wsproxysdk.RegisterWorkspaceProxyRequest.md)| Register workspace proxy request | |

### Return type

[**wsproxysdk.RegisterWorkspaceProxyResponse**](../Models/wsproxysdk.RegisterWorkspaceProxyResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="reportWorkspaceAppStats"></a>
# **reportWorkspaceAppStats**
> reportWorkspaceAppStats(request)

Report workspace app stats

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**wsproxysdk.ReportAppStatsRequest**](../Models/wsproxysdk.ReportAppStatsRequest.md)| Report app stats request | |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

<a name="scimCreateNewUser"></a>
# **scimCreateNewUser**
> coderd.SCIMUser scimCreateNewUser(request)

SCIM 2.0: Create new user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**coderd.SCIMUser**](../Models/coderd.SCIMUser.md)| New user | |

### Return type

[**coderd.SCIMUser**](../Models/coderd.SCIMUser.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="scimGetServiceProviderConfig"></a>
# **scimGetServiceProviderConfig**
> scimGetServiceProviderConfig()

SCIM 2.0: Service Provider Config

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="scimGetUserById"></a>
# **scimGetUserById**
> scimGetUserById(id)

SCIM 2.0: Get user by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | **UUID**| User ID | [default to null] |

### Return type

null (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="scimGetUsers"></a>
# **scimGetUsers**
> scimGetUsers()

SCIM 2.0: Get users

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="scimReplaceUserStatus"></a>
# **scimReplaceUserStatus**
> codersdk.User scimReplaceUserStatus(id, request)

SCIM 2.0: Replace user account

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | **UUID**| User ID | [default to null] |
| **request** | [**coderd.SCIMUser**](../Models/coderd.SCIMUser.md)| Replace user request | |

### Return type

[**codersdk.User**](../Models/codersdk.User.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/scim+json

<a name="scimUpdateUserStatus"></a>
# **scimUpdateUserStatus**
> codersdk.User scimUpdateUserStatus(id, request)

SCIM 2.0: Update user account

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | **UUID**| User ID | [default to null] |
| **request** | [**coderd.SCIMUser**](../Models/coderd.SCIMUser.md)| Update user request | |

### Return type

[**codersdk.User**](../Models/codersdk.User.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/scim+json

<a name="serveProvisionerDaemon"></a>
# **serveProvisionerDaemon**
> serveProvisionerDaemon(organization)

Serve provisioner daemon

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="updateAppearance"></a>
# **updateAppearance**
> codersdk.UpdateAppearanceConfig updateAppearance(request)

Update appearance

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**codersdk.UpdateAppearanceConfig**](../Models/codersdk.UpdateAppearanceConfig.md)| Update appearance request | |

### Return type

[**codersdk.UpdateAppearanceConfig**](../Models/codersdk.UpdateAppearanceConfig.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateGroupByName"></a>
# **updateGroupByName**
> codersdk.Group updateGroupByName(group, request)

Update group by name

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **group** | **String**| Group name | [default to null] |
| **request** | [**codersdk.PatchGroupRequest**](../Models/codersdk.PatchGroupRequest.md)| Patch group request | |

### Return type

[**codersdk.Group**](../Models/codersdk.Group.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateGroupIdpSyncConfig"></a>
# **updateGroupIdpSyncConfig**
> codersdk.GroupSyncSettings updateGroupIdpSyncConfig(organization, request)

Update group IdP Sync config

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID or name | [default to null] |
| **request** | [**codersdk.PatchGroupIDPSyncConfigRequest**](../Models/codersdk.PatchGroupIDPSyncConfigRequest.md)| New config values | |

### Return type

[**codersdk.GroupSyncSettings**](../Models/codersdk.GroupSyncSettings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateGroupIdpSyncMapping"></a>
# **updateGroupIdpSyncMapping**
> codersdk.GroupSyncSettings updateGroupIdpSyncMapping(organization, request)

Update group IdP Sync mapping

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID or name | [default to null] |
| **request** | [**codersdk.PatchGroupIDPSyncMappingRequest**](../Models/codersdk.PatchGroupIDPSyncMappingRequest.md)| Description of the mappings to add and remove | |

### Return type

[**codersdk.GroupSyncSettings**](../Models/codersdk.GroupSyncSettings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateGroupIdpSyncSettingsByOrganization"></a>
# **updateGroupIdpSyncSettingsByOrganization**
> codersdk.GroupSyncSettings updateGroupIdpSyncSettingsByOrganization(organization, request)

Update group IdP Sync settings by organization

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |
| **request** | [**codersdk.GroupSyncSettings**](../Models/codersdk.GroupSyncSettings.md)| New settings | |

### Return type

[**codersdk.GroupSyncSettings**](../Models/codersdk.GroupSyncSettings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateNotificationTemplateDispatchMethod"></a>
# **updateNotificationTemplateDispatchMethod**
> updateNotificationTemplateDispatchMethod(notification\_template)

Update notification template dispatch method

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **notification\_template** | **String**| Notification template UUID | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="updateOauth2Application"></a>
# **updateOauth2Application**
> codersdk.OAuth2ProviderApp updateOauth2Application(app, request)

Update OAuth2 application.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **app** | **String**| App ID | [default to null] |
| **request** | [**codersdk.PutOAuth2ProviderAppRequest**](../Models/codersdk.PutOAuth2ProviderAppRequest.md)| Update an OAuth2 application. | |

### Return type

[**codersdk.OAuth2ProviderApp**](../Models/codersdk.OAuth2ProviderApp.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateOrganizationIdpSyncConfig"></a>
# **updateOrganizationIdpSyncConfig**
> codersdk.OrganizationSyncSettings updateOrganizationIdpSyncConfig(request)

Update organization IdP Sync config

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**codersdk.PatchOrganizationIDPSyncConfigRequest**](../Models/codersdk.PatchOrganizationIDPSyncConfigRequest.md)| New config values | |

### Return type

[**codersdk.OrganizationSyncSettings**](../Models/codersdk.OrganizationSyncSettings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateOrganizationIdpSyncMapping"></a>
# **updateOrganizationIdpSyncMapping**
> codersdk.OrganizationSyncSettings updateOrganizationIdpSyncMapping(request)

Update organization IdP Sync mapping

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**codersdk.PatchOrganizationIDPSyncMappingRequest**](../Models/codersdk.PatchOrganizationIDPSyncMappingRequest.md)| Description of the mappings to add and remove | |

### Return type

[**codersdk.OrganizationSyncSettings**](../Models/codersdk.OrganizationSyncSettings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateOrganizationIdpSyncSettings"></a>
# **updateOrganizationIdpSyncSettings**
> codersdk.OrganizationSyncSettings updateOrganizationIdpSyncSettings(request)

Update organization IdP Sync settings

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**codersdk.OrganizationSyncSettings**](../Models/codersdk.OrganizationSyncSettings.md)| New settings | |

### Return type

[**codersdk.OrganizationSyncSettings**](../Models/codersdk.OrganizationSyncSettings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateRoleIdpSyncConfig"></a>
# **updateRoleIdpSyncConfig**
> codersdk.RoleSyncSettings updateRoleIdpSyncConfig(organization, request)

Update role IdP Sync config

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID or name | [default to null] |
| **request** | [**codersdk.PatchRoleIDPSyncConfigRequest**](../Models/codersdk.PatchRoleIDPSyncConfigRequest.md)| New config values | |

### Return type

[**codersdk.RoleSyncSettings**](../Models/codersdk.RoleSyncSettings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateRoleIdpSyncMapping"></a>
# **updateRoleIdpSyncMapping**
> codersdk.RoleSyncSettings updateRoleIdpSyncMapping(organization, request)

Update role IdP Sync mapping

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID or name | [default to null] |
| **request** | [**codersdk.PatchRoleIDPSyncMappingRequest**](../Models/codersdk.PatchRoleIDPSyncMappingRequest.md)| Description of the mappings to add and remove | |

### Return type

[**codersdk.RoleSyncSettings**](../Models/codersdk.RoleSyncSettings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateRoleIdpSyncSettingsByOrganization"></a>
# **updateRoleIdpSyncSettingsByOrganization**
> codersdk.RoleSyncSettings updateRoleIdpSyncSettingsByOrganization(organization, request)

Update role IdP Sync settings by organization

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |
| **request** | [**codersdk.RoleSyncSettings**](../Models/codersdk.RoleSyncSettings.md)| New settings | |

### Return type

[**codersdk.RoleSyncSettings**](../Models/codersdk.RoleSyncSettings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateTemplateAcl"></a>
# **updateTemplateAcl**
> codersdk.Response updateTemplateAcl(template, request)

Update template ACL

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **template** | **UUID**| Template ID | [default to null] |
| **request** | [**codersdk.UpdateTemplateACL**](../Models/codersdk.UpdateTemplateACL.md)| Update template ACL request | |

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateUserQuietHoursSchedule"></a>
# **updateUserQuietHoursSchedule**
> List updateUserQuietHoursSchedule(user, request)

Update user quiet hours schedule

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **UUID**| User ID | [default to null] |
| **request** | [**codersdk.UpdateUserQuietHoursScheduleRequest**](../Models/codersdk.UpdateUserQuietHoursScheduleRequest.md)| Update schedule request | |

### Return type

[**List**](../Models/codersdk.UserQuietHoursScheduleResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateWorkspaceProxy"></a>
# **updateWorkspaceProxy**
> codersdk.WorkspaceProxy updateWorkspaceProxy(workspaceproxy, request)

Update workspace proxy

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspaceproxy** | **UUID**| Proxy ID or name | [default to null] |
| **request** | [**codersdk.PatchWorkspaceProxy**](../Models/codersdk.PatchWorkspaceProxy.md)| Update workspace proxy request | |

### Return type

[**codersdk.WorkspaceProxy**](../Models/codersdk.WorkspaceProxy.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="workspaceProxyCoordinate"></a>
# **workspaceProxyCoordinate**
> workspaceProxyCoordinate()

Workspace Proxy Coordinate

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


