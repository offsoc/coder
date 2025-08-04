# Schemas

## Models

### agentsdk.AWSInstanceIdentityToken

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **document** | **String** |  | [default to null] |
| **signature** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### agentsdk.AuthenticateResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **session\_token** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### agentsdk.AzureInstanceIdentityToken

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **encoding** | **String** |  | [default to null] |
| **signature** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### agentsdk.ExternalAuthResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **access\_token** | **String** |  | [optional] [default to null] |
| **password** | **String** |  | [optional] [default to null] |
| **token\_extra** | [**Map**](AnyType.md) |  | [optional] [default to null] |
| **type** | **String** |  | [optional] [default to null] |
| **url** | **String** |  | [optional] [default to null] |
| **username** | **String** | Deprecated: Only supported on &#x60;/workspaceagents/me/gitauth&#x60; for backwards compatibility. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### agentsdk.GitSSHKey

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **private\_key** | **String** |  | [optional] [default to null] |
| **public\_key** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### agentsdk.GoogleInstanceIdentityToken

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **json\_web\_token** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### agentsdk.Log

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **created\_at** | **String** |  | [optional] [default to null] |
| **level** | [**codersdk.LogLevel**](codersdk.LogLevel.md) |  | [optional] [default to null] |
| **output** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### agentsdk.PatchAppStatus

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **app\_slug** | **String** |  | [optional] [default to null] |
| **icon** | **String** | Deprecated: this field is unused and will be removed in a future version. | [optional] [default to null] |
| **message** | **String** |  | [optional] [default to null] |
| **needs\_user\_attention** | **Boolean** | Deprecated: this field is unused and will be removed in a future version. | [optional] [default to null] |
| **state** | [**codersdk.WorkspaceAppStatusState**](codersdk.WorkspaceAppStatusState.md) |  | [optional] [default to null] |
| **uri** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### agentsdk.PatchLogs

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **log\_source\_id** | **String** |  | [optional] [default to null] |
| **logs** | [**List**](agentsdk.Log.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### agentsdk.PostLogSourceRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **display\_name** | **String** |  | [optional] [default to null] |
| **icon** | **String** |  | [optional] [default to null] |
| **id** | **String** | ID is a unique identifier for the log source. It is scoped to a workspace agent, and can be statically defined inside code to prevent duplicate sources from being created for the same agent. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### agentsdk.ReinitializationEvent

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **reason** | [**agentsdk.ReinitializationReason**](agentsdk.ReinitializationReason.md) |  | [optional] [default to null] |
| **workspaceID** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### agentsdk.ReinitializationReason

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### coderd.SCIMUser

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **active** | **Boolean** | Active is a ptr to prevent the empty value from being interpreted as false. | [optional] [default to null] |
| **emails** | [**List**](coderd_SCIMUser_emails_inner.md) |  | [optional] [default to null] |
| **groups** | **List** |  | [optional] [default to null] |
| **id** | **String** |  | [optional] [default to null] |
| **meta** | [**coderd_SCIMUser_meta**](coderd_SCIMUser_meta.md) |  | [optional] [default to null] |
| **name** | [**coderd_SCIMUser_name**](coderd_SCIMUser_name.md) |  | [optional] [default to null] |
| **schemas** | **List** |  | [optional] [default to null] |
| **userName** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### coderd.cspViolation

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **csp-report** | [**Map**](AnyType.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### coderd_SCIMUser_emails_inner

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **display** | **String** |  | [optional] [default to null] |
| **primary** | **Boolean** |  | [optional] [default to null] |
| **type** | **String** |  | [optional] [default to null] |
| **value** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### coderd_SCIMUser_meta

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **resourceType** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### coderd_SCIMUser_name

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **familyName** | **String** |  | [optional] [default to null] |
| **givenName** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ACLAvailable

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **groups** | [**List**](codersdk.Group.md) |  | [optional] [default to null] |
| **users** | [**List**](codersdk.ReducedUser.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.APIKey

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **created\_at** | **Date** |  | [default to null] |
| **expires\_at** | **Date** |  | [default to null] |
| **id** | **String** |  | [default to null] |
| **last\_used** | **Date** |  | [default to null] |
| **lifetime\_seconds** | **Integer** |  | [default to null] |
| **login\_type** | [**codersdk.LoginType**](codersdk.LoginType.md) |  | [default to null] |
| **scope** | [**codersdk.APIKeyScope**](codersdk.APIKeyScope.md) |  | [default to null] |
| **token\_name** | **String** |  | [default to null] |
| **updated\_at** | **Date** |  | [default to null] |
| **user\_id** | **UUID** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.APIKeyScope

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.AddLicenseRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **license** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.AgentConnectionTiming

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **ended\_at** | **Date** |  | [optional] [default to null] |
| **stage** | [**codersdk.TimingStage**](codersdk.TimingStage.md) |  | [optional] [default to null] |
| **started\_at** | **Date** |  | [optional] [default to null] |
| **workspace\_agent\_id** | **String** |  | [optional] [default to null] |
| **workspace\_agent\_name** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.AgentScriptTiming

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **display\_name** | **String** |  | [optional] [default to null] |
| **ended\_at** | **Date** |  | [optional] [default to null] |
| **exit\_code** | **Integer** |  | [optional] [default to null] |
| **stage** | [**codersdk.TimingStage**](codersdk.TimingStage.md) |  | [optional] [default to null] |
| **started\_at** | **Date** |  | [optional] [default to null] |
| **status** | **String** |  | [optional] [default to null] |
| **workspace\_agent\_id** | **String** |  | [optional] [default to null] |
| **workspace\_agent\_name** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.AgentSubsystem

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.AppHostResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **host** | **String** | Host is the externally accessible URL for the Coder instance. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.AppearanceConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **announcement\_banners** | [**List**](codersdk.BannerConfig.md) |  | [optional] [default to null] |
| **application\_name** | **String** |  | [optional] [default to null] |
| **docs\_url** | **String** |  | [optional] [default to null] |
| **logo\_url** | **String** |  | [optional] [default to null] |
| **service\_banner** | [**codersdk.BannerConfig**](codersdk.BannerConfig.md) | Deprecated: ServiceBanner has been replaced by AnnouncementBanners. | [optional] [default to null] |
| **support\_links** | [**List**](codersdk.LinkConfig.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ArchiveTemplateVersionsRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **all** | **Boolean** | By default, only failed versions are archived. Set this to true to archive all unused versions regardless of job status. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.AssignableRoles

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **assignable** | **Boolean** |  | [optional] [default to null] |
| **built\_in** | **Boolean** | BuiltIn roles are immutable | [optional] [default to null] |
| **display\_name** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **organization\_id** | **UUID** |  | [optional] [default to null] |
| **organization\_permissions** | [**List**](codersdk.Permission.md) | OrganizationPermissions are specific for the organization in the field &#39;OrganizationID&#39; above. | [optional] [default to null] |
| **site\_permissions** | [**List**](codersdk.Permission.md) |  | [optional] [default to null] |
| **user\_permissions** | [**List**](codersdk.Permission.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.AuditAction

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.AuditDiffField

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **new** | [**Object**](.md) |  | [optional] [default to null] |
| **old** | [**Object**](.md) |  | [optional] [default to null] |
| **secret** | **Boolean** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.AuditLog

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **action** | [**codersdk.AuditAction**](codersdk.AuditAction.md) |  | [optional] [default to null] |
| **additional\_fields** | [**Object**](.md) |  | [optional] [default to null] |
| **description** | **String** |  | [optional] [default to null] |
| **diff** | [**Map**](codersdk.AuditDiffField.md) |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **ip** | **String** |  | [optional] [default to null] |
| **is\_deleted** | **Boolean** |  | [optional] [default to null] |
| **organization** | [**codersdk.MinimalOrganization**](codersdk.MinimalOrganization.md) |  | [optional] [default to null] |
| **organization\_id** | **UUID** | Deprecated: Use &#39;organization.id&#39; instead. | [optional] [default to null] |
| **request\_id** | **UUID** |  | [optional] [default to null] |
| **resource\_icon** | **String** |  | [optional] [default to null] |
| **resource\_id** | **UUID** |  | [optional] [default to null] |
| **resource\_link** | **String** |  | [optional] [default to null] |
| **resource\_target** | **String** | ResourceTarget is the name of the resource. | [optional] [default to null] |
| **resource\_type** | [**codersdk.ResourceType**](codersdk.ResourceType.md) |  | [optional] [default to null] |
| **status\_code** | **Integer** |  | [optional] [default to null] |
| **time** | **Date** |  | [optional] [default to null] |
| **user** | [**codersdk.User**](codersdk.User.md) |  | [optional] [default to null] |
| **user\_agent** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.AuditLogResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **audit\_logs** | [**List**](codersdk.AuditLog.md) |  | [optional] [default to null] |
| **count** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.AuthMethod

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **enabled** | **Boolean** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.AuthMethods

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **github** | [**codersdk.GithubAuthMethod**](codersdk.GithubAuthMethod.md) |  | [optional] [default to null] |
| **oidc** | [**codersdk.OIDCAuthMethod**](codersdk.OIDCAuthMethod.md) |  | [optional] [default to null] |
| **password** | [**codersdk.AuthMethod**](codersdk.AuthMethod.md) |  | [optional] [default to null] |
| **terms\_of\_service\_url** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.AuthorizationCheck

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **action** | [**codersdk.RBACAction**](codersdk.RBACAction.md) |  | [optional] [default to null] |
| **object** | [**codersdk.AuthorizationObject**](codersdk.AuthorizationObject.md) | Object can represent a \&quot;set\&quot; of objects, such as: all workspaces in an organization, all workspaces owned by me, and all workspaces across the entire product. When defining an object, use the most specific language when possible to produce the smallest set. Meaning to set as many fields on &#39;Object&#39; as you can. Example, if you want to check if you can update all workspaces owned by &#39;me&#39;, try to also add an &#39;OrganizationID&#39; to the settings. Omitting the &#39;OrganizationID&#39; could produce the incorrect value, as workspaces have both &#x60;user&#x60; and &#x60;organization&#x60; owners. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.AuthorizationObject

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **any\_org** | **Boolean** | AnyOrgOwner (optional) will disregard the org_owner when checking for permissions. This cannot be set to true if the OrganizationID is set. | [optional] [default to null] |
| **organization\_id** | **String** | OrganizationID (optional) adds the set constraint to all resources owned by a given organization. | [optional] [default to null] |
| **owner\_id** | **String** | OwnerID (optional) adds the set constraint to all resources owned by a given user. | [optional] [default to null] |
| **resource\_id** | **String** | ResourceID (optional) reduces the set to a singular resource. This assigns a resource ID to the resource type, eg: a single workspace. The rbac library will not fetch the resource from the database, so if you are using this option, you should also set the owner ID and organization ID if possible. Be as specific as possible using all the fields relevant. | [optional] [default to null] |
| **resource\_type** | [**codersdk.RBACResource**](codersdk.RBACResource.md) | ResourceType is the name of the resource. &#x60;./coderd/rbac/object.go&#x60; has the list of valid resource types. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.AuthorizationRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **checks** | [**Map**](codersdk.AuthorizationCheck.md) | Checks is a map keyed with an arbitrary string to a permission check. The key can be any string that is helpful to the caller, and allows multiple permission checks to be run in a single request. The key ensures that each permission check has the same key in the response. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.AutomaticUpdates

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.BannerConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **background\_color** | **String** |  | [optional] [default to null] |
| **enabled** | **Boolean** |  | [optional] [default to null] |
| **message** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.BuildInfoResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **agent\_api\_version** | **String** | AgentAPIVersion is the current version of the Agent API (back versions MAY still be supported). | [optional] [default to null] |
| **dashboard\_url** | **String** | DashboardURL is the URL to hit the deployment&#39;s dashboard. For external workspace proxies, this is the coderd they are connected to. | [optional] [default to null] |
| **deployment\_id** | **String** | DeploymentID is the unique identifier for this deployment. | [optional] [default to null] |
| **external\_url** | **String** | ExternalURL references the current Coder version. For production builds, this will link directly to a release. For development builds, this will link to a commit. | [optional] [default to null] |
| **provisioner\_api\_version** | **String** | ProvisionerAPIVersion is the current version of the Provisioner API | [optional] [default to null] |
| **telemetry** | **Boolean** | Telemetry is a boolean that indicates whether telemetry is enabled. | [optional] [default to null] |
| **upgrade\_message** | **String** | UpgradeMessage is the message displayed to users when an outdated client is detected. | [optional] [default to null] |
| **version** | **String** | Version returns the semantic version of the build. | [optional] [default to null] |
| **webpush\_public\_key** | **String** | WebPushPublicKey is the public key for push notifications via Web Push. | [optional] [default to null] |
| **workspace\_proxy** | **Boolean** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.BuildReason

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.CORSBehavior

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ChangePasswordWithOneTimePasscodeRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **email** | **String** |  | [default to null] |
| **one\_time\_passcode** | **String** |  | [default to null] |
| **password** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ConnectionLatency

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **p50** | **BigDecimal** |  | [optional] [default to null] |
| **p95** | **BigDecimal** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ConnectionLog

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **agent\_name** | **String** |  | [optional] [default to null] |
| **connect\_time** | **Date** |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **ip** | **String** |  | [optional] [default to null] |
| **organization** | [**codersdk.MinimalOrganization**](codersdk.MinimalOrganization.md) |  | [optional] [default to null] |
| **ssh\_info** | [**codersdk.ConnectionLogSSHInfo**](codersdk.ConnectionLogSSHInfo.md) | SSHInfo is only set when &#x60;type&#x60; is one of: - &#x60;ConnectionTypeSSH&#x60; - &#x60;ConnectionTypeReconnectingPTY&#x60; - &#x60;ConnectionTypeVSCode&#x60; - &#x60;ConnectionTypeJetBrains&#x60; | [optional] [default to null] |
| **type** | [**codersdk.ConnectionType**](codersdk.ConnectionType.md) |  | [optional] [default to null] |
| **web\_info** | [**codersdk.ConnectionLogWebInfo**](codersdk.ConnectionLogWebInfo.md) | WebInfo is only set when &#x60;type&#x60; is one of: - &#x60;ConnectionTypePortForwarding&#x60; - &#x60;ConnectionTypeWorkspaceApp&#x60; | [optional] [default to null] |
| **workspace\_id** | **UUID** |  | [optional] [default to null] |
| **workspace\_name** | **String** |  | [optional] [default to null] |
| **workspace\_owner\_id** | **UUID** |  | [optional] [default to null] |
| **workspace\_owner\_username** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ConnectionLogResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **connection\_logs** | [**List**](codersdk.ConnectionLog.md) |  | [optional] [default to null] |
| **count** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ConnectionLogSSHInfo

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **connection\_id** | **UUID** |  | [optional] [default to null] |
| **disconnect\_reason** | **String** | DisconnectReason is omitted if a disconnect event with the same connection ID has not yet been seen. | [optional] [default to null] |
| **disconnect\_time** | **Date** | DisconnectTime is omitted if a disconnect event with the same connection ID has not yet been seen. | [optional] [default to null] |
| **exit\_code** | **Integer** | ExitCode is the exit code of the SSH session. It is omitted if a disconnect event with the same connection ID has not yet been seen. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ConnectionLogWebInfo

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **slug\_or\_port** | **String** |  | [optional] [default to null] |
| **status\_code** | **Integer** | StatusCode is the HTTP status code of the request. | [optional] [default to null] |
| **user** | [**codersdk.User**](codersdk.User.md) | User is omitted if the connection event was from an unauthenticated user. | [optional] [default to null] |
| **user\_agent** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ConnectionType

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ConvertLoginRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **password** | **String** |  | [default to null] |
| **to\_type** | [**codersdk.LoginType**](codersdk.LoginType.md) | ToType is the login type to convert to. | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.CreateFirstUserRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **email** | **String** |  | [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **password** | **String** |  | [default to null] |
| **trial** | **Boolean** |  | [optional] [default to null] |
| **trial\_info** | [**codersdk.CreateFirstUserTrialInfo**](codersdk.CreateFirstUserTrialInfo.md) |  | [optional] [default to null] |
| **username** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.CreateFirstUserResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **organization\_id** | **UUID** |  | [optional] [default to null] |
| **user\_id** | **UUID** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.CreateFirstUserTrialInfo

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **company\_name** | **String** |  | [optional] [default to null] |
| **country** | **String** |  | [optional] [default to null] |
| **developers** | **String** |  | [optional] [default to null] |
| **first\_name** | **String** |  | [optional] [default to null] |
| **job\_title** | **String** |  | [optional] [default to null] |
| **last\_name** | **String** |  | [optional] [default to null] |
| **phone\_number** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.CreateGroupRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **avatar\_url** | **String** |  | [optional] [default to null] |
| **display\_name** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [default to null] |
| **quota\_allowance** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.CreateOrganizationRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **description** | **String** |  | [optional] [default to null] |
| **display\_name** | **String** | DisplayName will default to the same value as &#x60;Name&#x60; if not provided. | [optional] [default to null] |
| **icon** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.CreateProvisionerKeyResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **key** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.CreateTemplateRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **activity\_bump\_ms** | **Integer** | ActivityBumpMillis allows optionally specifying the activity bump duration for all workspaces created from this template. Defaults to 1h but can be set to 0 to disable activity bumping. | [optional] [default to null] |
| **allow\_user\_autostart** | **Boolean** | AllowUserAutostart allows users to set a schedule for autostarting their workspace. By default this is true. This can only be disabled when using an enterprise license. | [optional] [default to null] |
| **allow\_user\_autostop** | **Boolean** | AllowUserAutostop allows users to set a custom workspace TTL to use in place of the template&#39;s DefaultTTL field. By default this is true. If false, the DefaultTTL will always be used. This can only be disabled when using an enterprise license. | [optional] [default to null] |
| **allow\_user\_cancel\_workspace\_jobs** | **Boolean** | Allow users to cancel in-progress workspace jobs. *bool as the default value is \&quot;true\&quot;. | [optional] [default to null] |
| **autostart\_requirement** | [**codersdk.TemplateAutostartRequirement**](codersdk.TemplateAutostartRequirement.md) | AutostartRequirement allows optionally specifying the autostart allowed days for workspaces created from this template. This is an enterprise feature. | [optional] [default to null] |
| **autostop\_requirement** | [**codersdk.TemplateAutostopRequirement**](codersdk.TemplateAutostopRequirement.md) | AutostopRequirement allows optionally specifying the autostop requirement for workspaces created from this template. This is an enterprise feature. | [optional] [default to null] |
| **cors\_behavior** | [**codersdk.CORSBehavior**](codersdk.CORSBehavior.md) | CORSBehavior allows optionally specifying the CORS behavior for all shared ports. | [optional] [default to null] |
| **default\_ttl\_ms** | **Integer** | DefaultTTLMillis allows optionally specifying the default TTL for all workspaces created from this template. | [optional] [default to null] |
| **delete\_ttl\_ms** | **Integer** | TimeTilDormantAutoDeleteMillis allows optionally specifying the max lifetime before Coder permanently deletes dormant workspaces created from this template. | [optional] [default to null] |
| **description** | **String** | Description is a description of what the template contains. It must be less than 128 bytes. | [optional] [default to null] |
| **disable\_everyone\_group\_access** | **Boolean** | DisableEveryoneGroupAccess allows optionally disabling the default behavior of granting the &#39;everyone&#39; group access to use the template. If this is set to true, the template will not be available to all users, and must be explicitly granted to users or groups in the permissions settings of the template. | [optional] [default to null] |
| **display\_name** | **String** | DisplayName is the displayed name of the template. | [optional] [default to null] |
| **dormant\_ttl\_ms** | **Integer** | TimeTilDormantMillis allows optionally specifying the max lifetime before Coder locks inactive workspaces created from this template. | [optional] [default to null] |
| **failure\_ttl\_ms** | **Integer** | FailureTTLMillis allows optionally specifying the max lifetime before Coder stops all resources for failed workspaces created from this template. | [optional] [default to null] |
| **icon** | **String** | Icon is a relative path or external URL that specifies an icon to be displayed in the dashboard. | [optional] [default to null] |
| **max\_port\_share\_level** | [**codersdk.WorkspaceAgentPortShareLevel**](codersdk.WorkspaceAgentPortShareLevel.md) | MaxPortShareLevel allows optionally specifying the maximum port share level for workspaces created from the template. | [optional] [default to null] |
| **name** | **String** | Name is the name of the template. | [default to null] |
| **require\_active\_version** | **Boolean** | RequireActiveVersion mandates that workspaces are built with the active template version. | [optional] [default to null] |
| **template\_use\_classic\_parameter\_flow** | **Boolean** | UseClassicParameterFlow allows optionally specifying whether the template should use the classic parameter flow. The default if unset is true, and is why &#x60;*bool&#x60; is used here. When dynamic parameters becomes the default, this will default to false. | [optional] [default to null] |
| **template\_version\_id** | **UUID** | VersionID is an in-progress or completed job to use as an initial version of the template.  This is required on creation to enable a user-flow of validating a template works. There is no reason the data-model cannot support empty templates, but it doesn&#39;t make sense for users. | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.CreateTemplateVersionDryRunRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **rich\_parameter\_values** | [**List**](codersdk.WorkspaceBuildParameter.md) |  | [optional] [default to null] |
| **user\_variable\_values** | [**List**](codersdk.VariableValue.md) |  | [optional] [default to null] |
| **workspace\_name** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.CreateTemplateVersionRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **example\_id** | **String** |  | [optional] [default to null] |
| **file\_id** | **UUID** |  | [optional] [default to null] |
| **message** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **provisioner** | **String** |  | [default to null] |
| **storage\_method** | [**codersdk.ProvisionerStorageMethod**](codersdk.ProvisionerStorageMethod.md) |  | [default to null] |
| **tags** | **Map** |  | [optional] [default to null] |
| **template\_id** | **UUID** | TemplateID optionally associates a version with a template. | [optional] [default to null] |
| **user\_variable\_values** | [**List**](codersdk.VariableValue.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.CreateTestAuditLogRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **action** | [**codersdk.AuditAction**](codersdk.AuditAction.md) |  | [optional] [default to null] |
| **additional\_fields** | **List** |  | [optional] [default to null] |
| **build\_reason** | [**codersdk.BuildReason**](codersdk.BuildReason.md) |  | [optional] [default to null] |
| **organization\_id** | **UUID** |  | [optional] [default to null] |
| **request\_id** | **UUID** |  | [optional] [default to null] |
| **resource\_id** | **UUID** |  | [optional] [default to null] |
| **resource\_type** | [**codersdk.ResourceType**](codersdk.ResourceType.md) |  | [optional] [default to null] |
| **time** | **Date** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.CreateTokenRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **lifetime** | **Integer** |  | [optional] [default to null] |
| **scope** | [**codersdk.APIKeyScope**](codersdk.APIKeyScope.md) |  | [optional] [default to null] |
| **token\_name** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.CreateUserRequestWithOrgs

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **email** | **String** |  | [default to null] |
| **login\_type** | [**codersdk.LoginType**](codersdk.LoginType.md) | UserLoginType defaults to LoginTypePassword. | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **organization\_ids** | **List** | OrganizationIDs is a list of organization IDs that the user should be a member of. | [optional] [default to null] |
| **password** | **String** |  | [optional] [default to null] |
| **user\_status** | [**codersdk.UserStatus**](codersdk.UserStatus.md) | UserStatus defaults to UserStatusDormant. | [optional] [default to null] |
| **username** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.CreateWorkspaceBuildReason

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.CreateWorkspaceBuildRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **dry\_run** | **Boolean** |  | [optional] [default to null] |
| **log\_level** | [**codersdk.ProvisionerLogLevel**](codersdk.ProvisionerLogLevel.md) | Log level changes the default logging verbosity of a provider (\&quot;info\&quot; if empty). | [optional] [default to null] |
| **orphan** | **Boolean** | Orphan may be set for the Destroy transition. | [optional] [default to null] |
| **reason** | [**codersdk.CreateWorkspaceBuildReason**](codersdk.CreateWorkspaceBuildReason.md) | Reason sets the reason for the workspace build. | [optional] [default to null] |
| **rich\_parameter\_values** | [**List**](codersdk.WorkspaceBuildParameter.md) | ParameterValues are optional. It will write params to the &#39;workspace&#39; scope. This will overwrite any existing parameters with the same name. This will not delete old params not included in this list. | [optional] [default to null] |
| **state** | **List** |  | [optional] [default to null] |
| **template\_version\_id** | **UUID** |  | [optional] [default to null] |
| **template\_version\_preset\_id** | **UUID** | TemplateVersionPresetID is the ID of the template version preset to use for the build. | [optional] [default to null] |
| **transition** | [**codersdk.WorkspaceTransition**](codersdk.WorkspaceTransition.md) |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.CreateWorkspaceProxyRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **display\_name** | **String** |  | [optional] [default to null] |
| **icon** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.CreateWorkspaceRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **automatic\_updates** | [**codersdk.AutomaticUpdates**](codersdk.AutomaticUpdates.md) |  | [optional] [default to null] |
| **autostart\_schedule** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [default to null] |
| **rich\_parameter\_values** | [**List**](codersdk.WorkspaceBuildParameter.md) | RichParameterValues allows for additional parameters to be provided during the initial provision. | [optional] [default to null] |
| **template\_id** | **UUID** | TemplateID specifies which template should be used for creating the workspace. | [optional] [default to null] |
| **template\_version\_id** | **UUID** | TemplateVersionID can be used to specify a specific version of a template for creating the workspace. | [optional] [default to null] |
| **template\_version\_preset\_id** | **UUID** |  | [optional] [default to null] |
| **ttl\_ms** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.CryptoKey

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **deletes\_at** | **Date** |  | [optional] [default to null] |
| **feature** | [**codersdk.CryptoKeyFeature**](codersdk.CryptoKeyFeature.md) |  | [optional] [default to null] |
| **secret** | **String** |  | [optional] [default to null] |
| **sequence** | **Integer** |  | [optional] [default to null] |
| **starts\_at** | **Date** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.CryptoKeyFeature

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.CustomRoleRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **display\_name** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **organization\_permissions** | [**List**](codersdk.Permission.md) | OrganizationPermissions are specific to the organization the role belongs to. | [optional] [default to null] |
| **site\_permissions** | [**List**](codersdk.Permission.md) |  | [optional] [default to null] |
| **user\_permissions** | [**List**](codersdk.Permission.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.DAUEntry

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **amount** | **Integer** |  | [optional] [default to null] |
| **date** | **String** | Date is a string formatted as 2024-01-31. Timezone and time information is not included. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.DAUsResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **entries** | [**List**](codersdk.DAUEntry.md) |  | [optional] [default to null] |
| **tz\_hour\_offset** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.DERP

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **config** | [**codersdk.DERPConfig**](codersdk.DERPConfig.md) |  | [optional] [default to null] |
| **server** | [**codersdk.DERPServerConfig**](codersdk.DERPServerConfig.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.DERPConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **block\_direct** | **Boolean** |  | [optional] [default to null] |
| **force\_websockets** | **Boolean** |  | [optional] [default to null] |
| **path** | **String** |  | [optional] [default to null] |
| **url** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.DERPRegion

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **latency\_ms** | **BigDecimal** |  | [optional] [default to null] |
| **preferred** | **Boolean** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.DERPServerConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **enable** | **Boolean** |  | [optional] [default to null] |
| **region\_code** | **String** |  | [optional] [default to null] |
| **region\_id** | **Integer** |  | [optional] [default to null] |
| **region\_name** | **String** |  | [optional] [default to null] |
| **relay\_url** | [**serpent.URL**](serpent.URL.md) |  | [optional] [default to null] |
| **stun\_addresses** | **List** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.DangerousConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **allow\_all\_cors** | **Boolean** |  | [optional] [default to null] |
| **allow\_path\_app\_sharing** | **Boolean** |  | [optional] [default to null] |
| **allow\_path\_app\_site\_owner\_access** | **Boolean** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.DeleteWebpushSubscription

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **endpoint** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.DeleteWorkspaceAgentPortShareRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **agent\_name** | **String** |  | [optional] [default to null] |
| **port** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.DeploymentConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **config** | [**codersdk.DeploymentValues**](codersdk.DeploymentValues.md) |  | [optional] [default to null] |
| **options** | [**List**](serpent.Option.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.DeploymentStats

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **aggregated\_from** | **Date** | AggregatedFrom is the time in which stats are aggregated from. This might be back in time a specific duration or interval. | [optional] [default to null] |
| **collected\_at** | **Date** | CollectedAt is the time in which stats are collected at. | [optional] [default to null] |
| **next\_update\_at** | **Date** | NextUpdateAt is the time when the next batch of stats will be updated. | [optional] [default to null] |
| **session\_count** | [**codersdk.SessionCountDeploymentStats**](codersdk.SessionCountDeploymentStats.md) |  | [optional] [default to null] |
| **workspaces** | [**codersdk.WorkspaceDeploymentStats**](codersdk.WorkspaceDeploymentStats.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.DeploymentValues

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **access\_url** | [**serpent.URL**](serpent.URL.md) |  | [optional] [default to null] |
| **additional\_csp\_policy** | **List** |  | [optional] [default to null] |
| **address** | [**serpent.HostPort**](serpent.HostPort.md) | Deprecated: Use HTTPAddress or TLS.Address instead. | [optional] [default to null] |
| **agent\_fallback\_troubleshooting\_url** | [**serpent.URL**](serpent.URL.md) |  | [optional] [default to null] |
| **agent\_stat\_refresh\_interval** | **Integer** |  | [optional] [default to null] |
| **allow\_workspace\_renames** | **Boolean** |  | [optional] [default to null] |
| **autobuild\_poll\_interval** | **Integer** |  | [optional] [default to null] |
| **browser\_only** | **Boolean** |  | [optional] [default to null] |
| **cache\_directory** | **String** |  | [optional] [default to null] |
| **cli\_upgrade\_message** | **String** |  | [optional] [default to null] |
| **config** | **String** |  | [optional] [default to null] |
| **config\_ssh** | [**codersdk.SSHConfig**](codersdk.SSHConfig.md) |  | [optional] [default to null] |
| **dangerous** | [**codersdk.DangerousConfig**](codersdk.DangerousConfig.md) |  | [optional] [default to null] |
| **derp** | [**codersdk.DERP**](codersdk.DERP.md) |  | [optional] [default to null] |
| **disable\_owner\_workspace\_exec** | **Boolean** |  | [optional] [default to null] |
| **disable\_password\_auth** | **Boolean** |  | [optional] [default to null] |
| **disable\_path\_apps** | **Boolean** |  | [optional] [default to null] |
| **docs\_url** | [**serpent.URL**](serpent.URL.md) |  | [optional] [default to null] |
| **enable\_terraform\_debug\_mode** | **Boolean** |  | [optional] [default to null] |
| **ephemeral\_deployment** | **Boolean** |  | [optional] [default to null] |
| **experiments** | **List** |  | [optional] [default to null] |
| **external\_auth** | [**serpent.Struct-array_codersdk_ExternalAuthConfig**](serpent.Struct-array_codersdk_ExternalAuthConfig.md) |  | [optional] [default to null] |
| **external\_token\_encryption\_keys** | **List** |  | [optional] [default to null] |
| **healthcheck** | [**codersdk.HealthcheckConfig**](codersdk.HealthcheckConfig.md) |  | [optional] [default to null] |
| **hide\_ai\_tasks** | **Boolean** |  | [optional] [default to null] |
| **http\_address** | **String** | HTTPAddress is a string because it may be set to zero to disable. | [optional] [default to null] |
| **http\_cookies** | [**codersdk.HTTPCookieConfig**](codersdk.HTTPCookieConfig.md) |  | [optional] [default to null] |
| **job\_hang\_detector\_interval** | **Integer** |  | [optional] [default to null] |
| **logging** | [**codersdk.LoggingConfig**](codersdk.LoggingConfig.md) |  | [optional] [default to null] |
| **metrics\_cache\_refresh\_interval** | **Integer** |  | [optional] [default to null] |
| **notifications** | [**codersdk.NotificationsConfig**](codersdk.NotificationsConfig.md) |  | [optional] [default to null] |
| **oauth2** | [**codersdk.OAuth2Config**](codersdk.OAuth2Config.md) |  | [optional] [default to null] |
| **oidc** | [**codersdk.OIDCConfig**](codersdk.OIDCConfig.md) |  | [optional] [default to null] |
| **pg\_auth** | **String** |  | [optional] [default to null] |
| **pg\_connection\_url** | **String** |  | [optional] [default to null] |
| **pprof** | [**codersdk.PprofConfig**](codersdk.PprofConfig.md) |  | [optional] [default to null] |
| **prometheus** | [**codersdk.PrometheusConfig**](codersdk.PrometheusConfig.md) |  | [optional] [default to null] |
| **provisioner** | [**codersdk.ProvisionerConfig**](codersdk.ProvisionerConfig.md) |  | [optional] [default to null] |
| **proxy\_health\_status\_interval** | **Integer** |  | [optional] [default to null] |
| **proxy\_trusted\_headers** | **List** |  | [optional] [default to null] |
| **proxy\_trusted\_origins** | **List** |  | [optional] [default to null] |
| **rate\_limit** | [**codersdk.RateLimitConfig**](codersdk.RateLimitConfig.md) |  | [optional] [default to null] |
| **redirect\_to\_access\_url** | **Boolean** |  | [optional] [default to null] |
| **scim\_api\_key** | **String** |  | [optional] [default to null] |
| **session\_lifetime** | [**codersdk.SessionLifetime**](codersdk.SessionLifetime.md) |  | [optional] [default to null] |
| **ssh\_keygen\_algorithm** | **String** |  | [optional] [default to null] |
| **strict\_transport\_security** | **Integer** |  | [optional] [default to null] |
| **strict\_transport\_security\_options** | **List** |  | [optional] [default to null] |
| **support** | [**codersdk.SupportConfig**](codersdk.SupportConfig.md) |  | [optional] [default to null] |
| **swagger** | [**codersdk.SwaggerConfig**](codersdk.SwaggerConfig.md) |  | [optional] [default to null] |
| **telemetry** | [**codersdk.TelemetryConfig**](codersdk.TelemetryConfig.md) |  | [optional] [default to null] |
| **terms\_of\_service\_url** | **String** |  | [optional] [default to null] |
| **tls** | [**codersdk.TLSConfig**](codersdk.TLSConfig.md) |  | [optional] [default to null] |
| **trace** | [**codersdk.TraceConfig**](codersdk.TraceConfig.md) |  | [optional] [default to null] |
| **update\_check** | **Boolean** |  | [optional] [default to null] |
| **user\_quiet\_hours\_schedule** | [**codersdk.UserQuietHoursScheduleConfig**](codersdk.UserQuietHoursScheduleConfig.md) |  | [optional] [default to null] |
| **verbose** | **Boolean** |  | [optional] [default to null] |
| **web\_terminal\_renderer** | **String** |  | [optional] [default to null] |
| **wgtunnel\_host** | **String** |  | [optional] [default to null] |
| **wildcard\_access\_url** | **String** |  | [optional] [default to null] |
| **workspace\_hostname\_suffix** | **String** |  | [optional] [default to null] |
| **workspace\_prebuilds** | [**codersdk.PrebuildsConfig**](codersdk.PrebuildsConfig.md) |  | [optional] [default to null] |
| **write\_config** | **Boolean** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.DiagnosticExtra

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **code** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.DiagnosticSeverityString

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.DisplayApp

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.DynamicParametersRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **Integer** | ID identifies the request. The response contains the same ID so that the client can match it to the request. | [optional] [default to null] |
| **inputs** | **Map** |  | [optional] [default to null] |
| **owner\_id** | **UUID** | OwnerID if uuid.Nil, it defaults to &#x60;codersdk.Me&#x60; | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.DynamicParametersResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **diagnostics** | [**List**](codersdk.FriendlyDiagnostic.md) |  | [optional] [default to null] |
| **id** | **Integer** |  | [optional] [default to null] |
| **parameters** | [**List**](codersdk.PreviewParameter.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.Entitlement

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.Entitlements

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **errors** | **List** |  | [optional] [default to null] |
| **features** | [**Map**](codersdk.Feature.md) |  | [optional] [default to null] |
| **has\_license** | **Boolean** |  | [optional] [default to null] |
| **refreshed\_at** | **Date** |  | [optional] [default to null] |
| **require\_telemetry** | **Boolean** |  | [optional] [default to null] |
| **trial** | **Boolean** |  | [optional] [default to null] |
| **warnings** | **List** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.Experiment

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ExternalAuth

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **app\_install\_url** | **String** | AppInstallURL is the URL to install the app. | [optional] [default to null] |
| **app\_installable** | **Boolean** | AppInstallable is true if the request for app installs was successful. | [optional] [default to null] |
| **authenticated** | **Boolean** |  | [optional] [default to null] |
| **device** | **Boolean** |  | [optional] [default to null] |
| **display\_name** | **String** |  | [optional] [default to null] |
| **installations** | [**List**](codersdk.ExternalAuthAppInstallation.md) | AppInstallations are the installations that the user has access to. | [optional] [default to null] |
| **user** | [**codersdk.ExternalAuthUser**](codersdk.ExternalAuthUser.md) | User is the user that authenticated with the provider. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ExternalAuthAppInstallation

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **account** | [**codersdk.ExternalAuthUser**](codersdk.ExternalAuthUser.md) |  | [optional] [default to null] |
| **configure\_url** | **String** |  | [optional] [default to null] |
| **id** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ExternalAuthConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **app\_install\_url** | **String** |  | [optional] [default to null] |
| **app\_installations\_url** | **String** |  | [optional] [default to null] |
| **auth\_url** | **String** |  | [optional] [default to null] |
| **client\_id** | **String** |  | [optional] [default to null] |
| **device\_code\_url** | **String** |  | [optional] [default to null] |
| **device\_flow** | **Boolean** |  | [optional] [default to null] |
| **display\_icon** | **String** | DisplayIcon is a URL to an icon to display in the UI. | [optional] [default to null] |
| **display\_name** | **String** | DisplayName is shown in the UI to identify the auth config. | [optional] [default to null] |
| **id** | **String** | ID is a unique identifier for the auth config. It defaults to &#x60;type&#x60; when not provided. | [optional] [default to null] |
| **no\_refresh** | **Boolean** |  | [optional] [default to null] |
| **regex** | **String** | Regex allows API requesters to match an auth config by a string (e.g. coder.com) instead of by it&#39;s type.  Git clone makes use of this by parsing the URL from: &#39;Username for \&quot;https://github.com\&quot;:&#39; And sending it to the Coder server to match against the Regex. | [optional] [default to null] |
| **scopes** | **List** |  | [optional] [default to null] |
| **token\_url** | **String** |  | [optional] [default to null] |
| **type** | **String** | Type is the type of external auth config. | [optional] [default to null] |
| **validate\_url** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ExternalAuthDevice

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **device\_code** | **String** |  | [optional] [default to null] |
| **expires\_in** | **Integer** |  | [optional] [default to null] |
| **interval** | **Integer** |  | [optional] [default to null] |
| **user\_code** | **String** |  | [optional] [default to null] |
| **verification\_uri** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ExternalAuthLink

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **authenticated** | **Boolean** |  | [optional] [default to null] |
| **created\_at** | **Date** |  | [optional] [default to null] |
| **expires** | **Date** |  | [optional] [default to null] |
| **has\_refresh\_token** | **Boolean** |  | [optional] [default to null] |
| **provider\_id** | **String** |  | [optional] [default to null] |
| **updated\_at** | **Date** |  | [optional] [default to null] |
| **validate\_error** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ExternalAuthUser

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **avatar\_url** | **String** |  | [optional] [default to null] |
| **id** | **Integer** |  | [optional] [default to null] |
| **login** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **profile\_url** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.Feature

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **actual** | **Integer** |  | [optional] [default to null] |
| **enabled** | **Boolean** |  | [optional] [default to null] |
| **entitlement** | [**codersdk.Entitlement**](codersdk.Entitlement.md) |  | [optional] [default to null] |
| **limit** | **Integer** |  | [optional] [default to null] |
| **soft\_limit** | **Integer** | SoftLimit is the soft limit of the feature, and is only used for showing included limits in the dashboard. No license validation or warnings are generated from this value. | [optional] [default to null] |
| **usage\_period** | [**codersdk.UsagePeriod**](codersdk.UsagePeriod.md) | UsagePeriod denotes that the usage is a counter that accumulates over this period (and most likely resets with the issuance of the next license).  These dates are determined from the license that this entitlement comes from, see enterprise/coderd/license/license.go.  Only certain features set these fields: - FeatureManagedAgentLimit | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.FriendlyDiagnostic

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **detail** | **String** |  | [optional] [default to null] |
| **extra** | [**codersdk.DiagnosticExtra**](codersdk.DiagnosticExtra.md) |  | [optional] [default to null] |
| **severity** | [**codersdk.DiagnosticSeverityString**](codersdk.DiagnosticSeverityString.md) |  | [optional] [default to null] |
| **summary** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.GenerateAPIKeyResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **key** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.GetInboxNotificationResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **notification** | [**codersdk.InboxNotification**](codersdk.InboxNotification.md) |  | [optional] [default to null] |
| **unread\_count** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.GetUserStatusCountsResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **status\_counts** | [**Map**](array.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.GetUsersResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **count** | **Integer** |  | [optional] [default to null] |
| **users** | [**List**](codersdk.User.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.GitSSHKey

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **created\_at** | **Date** |  | [optional] [default to null] |
| **public\_key** | **String** | PublicKey is the SSH public key in OpenSSH format. Example: \&quot;ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAID3OmYJvT7q1cF1azbybYy0OZ9yrXfA+M6Lr4vzX5zlp\\n\&quot; Note: The key includes a trailing newline (\\n). | [optional] [default to null] |
| **updated\_at** | **Date** |  | [optional] [default to null] |
| **user\_id** | **UUID** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.GithubAuthMethod

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **default\_provider\_configured** | **Boolean** |  | [optional] [default to null] |
| **enabled** | **Boolean** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.Group

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **avatar\_url** | **URI** |  | [optional] [default to null] |
| **display\_name** | **String** |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **members** | [**List**](codersdk.ReducedUser.md) |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **organization\_display\_name** | **String** |  | [optional] [default to null] |
| **organization\_id** | **UUID** |  | [optional] [default to null] |
| **organization\_name** | **String** |  | [optional] [default to null] |
| **quota\_allowance** | **Integer** |  | [optional] [default to null] |
| **source** | [**codersdk.GroupSource**](codersdk.GroupSource.md) |  | [optional] [default to null] |
| **total\_member\_count** | **Integer** | How many members are in this group. Shows the total count, even if the user is not authorized to read group member details. May be greater than &#x60;len(Group.Members)&#x60;. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.GroupSource

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.GroupSyncSettings

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **auto\_create\_missing\_groups** | **Boolean** | AutoCreateMissing controls whether groups returned by the OIDC provider are automatically created in Coder if they are missing. | [optional] [default to null] |
| **field** | **String** | Field is the name of the claim field that specifies what groups a user should be in. If empty, no groups will be synced. | [optional] [default to null] |
| **legacy\_group\_name\_mapping** | **Map** | LegacyNameMapping is deprecated. It remaps an IDP group name to a Coder group name. Since configuration is now done at runtime, group IDs are used to account for group renames. For legacy configurations, this config option has to remain. Deprecated: Use Mapping instead. | [optional] [default to null] |
| **mapping** | [**Map**](array.md) | Mapping is a map from OIDC groups to Coder group IDs | [optional] [default to null] |
| **regex\_filter** | [**Object**](.md) | RegexFilter is a regular expression that filters the groups returned by the OIDC provider. Any group not matched by this regex will be ignored. If the group filter is nil, then no group filtering will occur. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.HTTPCookieConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **same\_site** | **String** |  | [optional] [default to null] |
| **secure\_auth\_cookie** | **Boolean** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.Healthcheck

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **interval** | **Integer** | Interval specifies the seconds between each health check. | [optional] [default to null] |
| **threshold** | **Integer** | Threshold specifies the number of consecutive failed health checks before returning \&quot;unhealthy\&quot;. | [optional] [default to null] |
| **url** | **String** | URL specifies the endpoint to check for the app health. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.HealthcheckConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **refresh** | **Integer** |  | [optional] [default to null] |
| **threshold\_database** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.InboxNotification

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **actions** | [**List**](codersdk.InboxNotificationAction.md) |  | [optional] [default to null] |
| **content** | **String** |  | [optional] [default to null] |
| **created\_at** | **Date** |  | [optional] [default to null] |
| **icon** | **String** |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **read\_at** | **String** |  | [optional] [default to null] |
| **targets** | **List** |  | [optional] [default to null] |
| **template\_id** | **UUID** |  | [optional] [default to null] |
| **title** | **String** |  | [optional] [default to null] |
| **user\_id** | **UUID** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.InboxNotificationAction

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **label** | **String** |  | [optional] [default to null] |
| **url** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.InsightsReportInterval

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.IssueReconnectingPTYSignedTokenRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **agentID** | **UUID** |  | [default to null] |
| **url** | **String** | URL is the URL of the reconnecting-pty endpoint you are connecting to. | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.IssueReconnectingPTYSignedTokenResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **signed\_token** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.JobErrorCode

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.License

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **claims** | [**Map**](AnyType.md) | Claims are the JWT claims asserted by the license.  Here we use a generic string map to ensure that all data from the server is parsed verbatim, not just the fields this version of Coder understands. | [optional] [default to null] |
| **id** | **Integer** |  | [optional] [default to null] |
| **uploaded\_at** | **Date** |  | [optional] [default to null] |
| **uuid** | **UUID** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.LinkConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **icon** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **target** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ListInboxNotificationsResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **notifications** | [**List**](codersdk.InboxNotification.md) |  | [optional] [default to null] |
| **unread\_count** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.LogLevel

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.LogSource

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.LoggingConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **human** | **String** |  | [optional] [default to null] |
| **json** | **String** |  | [optional] [default to null] |
| **log\_filter** | **List** |  | [optional] [default to null] |
| **stackdriver** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.LoginType

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.LoginWithPasswordRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **email** | **String** |  | [default to null] |
| **password** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.LoginWithPasswordResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **session\_token** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.MatchedProvisioners

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **available** | **Integer** | Available is the number of provisioner daemons that are available to take jobs. This may be less than the count if some provisioners are busy or have been stopped. | [optional] [default to null] |
| **count** | **Integer** | Count is the number of provisioner daemons that matched the given tags. If the count is 0, it means no provisioner daemons matched the requested tags. | [optional] [default to null] |
| **most\_recently\_seen** | **Date** | MostRecentlySeen is the most recently seen time of the set of matched provisioners. If no provisioners matched, this field will be null. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.MinimalOrganization

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **display\_name** | **String** |  | [optional] [default to null] |
| **icon** | **String** |  | [optional] [default to null] |
| **id** | **UUID** |  | [default to null] |
| **name** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.MinimalUser

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **avatar\_url** | **URI** |  | [optional] [default to null] |
| **id** | **UUID** |  | [default to null] |
| **username** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.NotificationMethodsResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **available** | **List** |  | [optional] [default to null] |
| **default** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.NotificationPreference

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **disabled** | **Boolean** |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **updated\_at** | **Date** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.NotificationTemplate

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **actions** | **String** |  | [optional] [default to null] |
| **body\_template** | **String** |  | [optional] [default to null] |
| **enabled\_by\_default** | **Boolean** |  | [optional] [default to null] |
| **group** | **String** |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **kind** | **String** |  | [optional] [default to null] |
| **method** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **title\_template** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.NotificationsConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **dispatch\_timeout** | **Integer** | How long to wait while a notification is being sent before giving up. | [optional] [default to null] |
| **email** | [**codersdk.NotificationsEmailConfig**](codersdk.NotificationsEmailConfig.md) | SMTP settings. | [optional] [default to null] |
| **fetch\_interval** | **Integer** | How often to query the database for queued notifications. | [optional] [default to null] |
| **inbox** | [**codersdk.NotificationsInboxConfig**](codersdk.NotificationsInboxConfig.md) | Inbox settings. | [optional] [default to null] |
| **lease\_count** | **Integer** | How many notifications a notifier should lease per fetch interval. | [optional] [default to null] |
| **lease\_period** | **Integer** | How long a notifier should lease a message. This is effectively how long a notification is &#39;owned&#39; by a notifier, and once this period expires it will be available for lease by another notifier. Leasing is important in order for multiple running notifiers to not pick the same messages to deliver concurrently. This lease period will only expire if a notifier shuts down ungracefully; a dispatch of the notification releases the lease. | [optional] [default to null] |
| **max\_send\_attempts** | **Integer** | The upper limit of attempts to send a notification. | [optional] [default to null] |
| **method** | **String** | Which delivery method to use (available options: &#39;smtp&#39;, &#39;webhook&#39;). | [optional] [default to null] |
| **retry\_interval** | **Integer** | The minimum time between retries. | [optional] [default to null] |
| **sync\_buffer\_size** | **Integer** | The notifications system buffers message updates in memory to ease pressure on the database. This option controls how many updates are kept in memory. The lower this value the lower the change of state inconsistency in a non-graceful shutdown - but it also increases load on the database. It is recommended to keep this option at its default value. | [optional] [default to null] |
| **sync\_interval** | **Integer** | The notifications system buffers message updates in memory to ease pressure on the database. This option controls how often it synchronizes its state with the database. The shorter this value the lower the change of state inconsistency in a non-graceful shutdown - but it also increases load on the database. It is recommended to keep this option at its default value. | [optional] [default to null] |
| **webhook** | [**codersdk.NotificationsWebhookConfig**](codersdk.NotificationsWebhookConfig.md) | Webhook settings. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.NotificationsEmailAuthConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **identity** | **String** | Identity for PLAIN auth. | [optional] [default to null] |
| **password** | **String** | Password for LOGIN/PLAIN auth. | [optional] [default to null] |
| **password\_file** | **String** | File from which to load the password for LOGIN/PLAIN auth. | [optional] [default to null] |
| **username** | **String** | Username for LOGIN/PLAIN auth. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.NotificationsEmailConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **auth** | [**codersdk.NotificationsEmailAuthConfig**](codersdk.NotificationsEmailAuthConfig.md) | Authentication details. | [optional] [default to null] |
| **force\_tls** | **Boolean** | ForceTLS causes a TLS connection to be attempted. | [optional] [default to null] |
| **from** | **String** | The sender&#39;s address. | [optional] [default to null] |
| **hello** | **String** | The hostname identifying the SMTP server. | [optional] [default to null] |
| **smarthost** | **String** | The intermediary SMTP host through which emails are sent (host:port). | [optional] [default to null] |
| **tls** | [**codersdk.NotificationsEmailTLSConfig**](codersdk.NotificationsEmailTLSConfig.md) | TLS details. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.NotificationsEmailTLSConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **ca\_file** | **String** | CAFile specifies the location of the CA certificate to use. | [optional] [default to null] |
| **cert\_file** | **String** | CertFile specifies the location of the certificate to use. | [optional] [default to null] |
| **insecure\_skip\_verify** | **Boolean** | InsecureSkipVerify skips target certificate validation. | [optional] [default to null] |
| **key\_file** | **String** | KeyFile specifies the location of the key to use. | [optional] [default to null] |
| **server\_name** | **String** | ServerName to verify the hostname for the targets. | [optional] [default to null] |
| **start\_tls** | **Boolean** | StartTLS attempts to upgrade plain connections to TLS. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.NotificationsInboxConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **enabled** | **Boolean** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.NotificationsSettings

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **notifier\_paused** | **Boolean** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.NotificationsWebhookConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **endpoint** | [**serpent.URL**](serpent.URL.md) | The URL to which the payload will be sent with an HTTP POST request. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.NullHCLString

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **valid** | **Boolean** |  | [optional] [default to null] |
| **value** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.OAuth2AppEndpoints

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **authorization** | **String** |  | [optional] [default to null] |
| **device\_authorization** | **String** | DeviceAuth is optional. | [optional] [default to null] |
| **token** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.OAuth2AuthorizationServerMetadata

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **authorization\_endpoint** | **String** |  | [optional] [default to null] |
| **code\_challenge\_methods\_supported** | **List** |  | [optional] [default to null] |
| **grant\_types\_supported** | **List** |  | [optional] [default to null] |
| **issuer** | **String** |  | [optional] [default to null] |
| **registration\_endpoint** | **String** |  | [optional] [default to null] |
| **response\_types\_supported** | **List** |  | [optional] [default to null] |
| **scopes\_supported** | **List** |  | [optional] [default to null] |
| **token\_endpoint** | **String** |  | [optional] [default to null] |
| **token\_endpoint\_auth\_methods\_supported** | **List** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.OAuth2ClientConfiguration

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **client\_id** | **String** |  | [optional] [default to null] |
| **client\_id\_issued\_at** | **Integer** |  | [optional] [default to null] |
| **client\_name** | **String** |  | [optional] [default to null] |
| **client\_secret\_expires\_at** | **Integer** |  | [optional] [default to null] |
| **client\_uri** | **String** |  | [optional] [default to null] |
| **contacts** | **List** |  | [optional] [default to null] |
| **grant\_types** | **List** |  | [optional] [default to null] |
| **jwks** | [**Object**](.md) |  | [optional] [default to null] |
| **jwks\_uri** | **String** |  | [optional] [default to null] |
| **logo\_uri** | **String** |  | [optional] [default to null] |
| **policy\_uri** | **String** |  | [optional] [default to null] |
| **redirect\_uris** | **List** |  | [optional] [default to null] |
| **registration\_access\_token** | **String** |  | [optional] [default to null] |
| **registration\_client\_uri** | **String** |  | [optional] [default to null] |
| **response\_types** | **List** |  | [optional] [default to null] |
| **scope** | **String** |  | [optional] [default to null] |
| **software\_id** | **String** |  | [optional] [default to null] |
| **software\_version** | **String** |  | [optional] [default to null] |
| **token\_endpoint\_auth\_method** | **String** |  | [optional] [default to null] |
| **tos\_uri** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.OAuth2ClientRegistrationRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **client\_name** | **String** |  | [optional] [default to null] |
| **client\_uri** | **String** |  | [optional] [default to null] |
| **contacts** | **List** |  | [optional] [default to null] |
| **grant\_types** | **List** |  | [optional] [default to null] |
| **jwks** | [**Object**](.md) |  | [optional] [default to null] |
| **jwks\_uri** | **String** |  | [optional] [default to null] |
| **logo\_uri** | **String** |  | [optional] [default to null] |
| **policy\_uri** | **String** |  | [optional] [default to null] |
| **redirect\_uris** | **List** |  | [optional] [default to null] |
| **response\_types** | **List** |  | [optional] [default to null] |
| **scope** | **String** |  | [optional] [default to null] |
| **software\_id** | **String** |  | [optional] [default to null] |
| **software\_statement** | **String** |  | [optional] [default to null] |
| **software\_version** | **String** |  | [optional] [default to null] |
| **token\_endpoint\_auth\_method** | **String** |  | [optional] [default to null] |
| **tos\_uri** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.OAuth2ClientRegistrationResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **client\_id** | **String** |  | [optional] [default to null] |
| **client\_id\_issued\_at** | **Integer** |  | [optional] [default to null] |
| **client\_name** | **String** |  | [optional] [default to null] |
| **client\_secret** | **String** |  | [optional] [default to null] |
| **client\_secret\_expires\_at** | **Integer** |  | [optional] [default to null] |
| **client\_uri** | **String** |  | [optional] [default to null] |
| **contacts** | **List** |  | [optional] [default to null] |
| **grant\_types** | **List** |  | [optional] [default to null] |
| **jwks** | [**Object**](.md) |  | [optional] [default to null] |
| **jwks\_uri** | **String** |  | [optional] [default to null] |
| **logo\_uri** | **String** |  | [optional] [default to null] |
| **policy\_uri** | **String** |  | [optional] [default to null] |
| **redirect\_uris** | **List** |  | [optional] [default to null] |
| **registration\_access\_token** | **String** |  | [optional] [default to null] |
| **registration\_client\_uri** | **String** |  | [optional] [default to null] |
| **response\_types** | **List** |  | [optional] [default to null] |
| **scope** | **String** |  | [optional] [default to null] |
| **software\_id** | **String** |  | [optional] [default to null] |
| **software\_version** | **String** |  | [optional] [default to null] |
| **token\_endpoint\_auth\_method** | **String** |  | [optional] [default to null] |
| **tos\_uri** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.OAuth2Config

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **github** | [**codersdk.OAuth2GithubConfig**](codersdk.OAuth2GithubConfig.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.OAuth2GithubConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **allow\_everyone** | **Boolean** |  | [optional] [default to null] |
| **allow\_signups** | **Boolean** |  | [optional] [default to null] |
| **allowed\_orgs** | **List** |  | [optional] [default to null] |
| **allowed\_teams** | **List** |  | [optional] [default to null] |
| **client\_id** | **String** |  | [optional] [default to null] |
| **client\_secret** | **String** |  | [optional] [default to null] |
| **default\_provider\_enable** | **Boolean** |  | [optional] [default to null] |
| **device\_flow** | **Boolean** |  | [optional] [default to null] |
| **enterprise\_base\_url** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.OAuth2ProtectedResourceMetadata

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **authorization\_servers** | **List** |  | [optional] [default to null] |
| **bearer\_methods\_supported** | **List** |  | [optional] [default to null] |
| **resource** | **String** |  | [optional] [default to null] |
| **scopes\_supported** | **List** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.OAuth2ProviderApp

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **callback\_url** | **String** |  | [optional] [default to null] |
| **endpoints** | [**codersdk.OAuth2AppEndpoints**](codersdk.OAuth2AppEndpoints.md) | Endpoints are included in the app response for easier discovery. The OAuth2 spec does not have a defined place to find these (for comparison, OIDC has a &#39;/.well-known/openid-configuration&#39; endpoint). | [optional] [default to null] |
| **icon** | **String** |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.OAuth2ProviderAppSecret

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **client\_secret\_truncated** | **String** |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **last\_used\_at** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.OAuth2ProviderAppSecretFull

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **client\_secret\_full** | **String** |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.OAuthConversionResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **expires\_at** | **Date** |  | [optional] [default to null] |
| **state\_string** | **String** |  | [optional] [default to null] |
| **to\_type** | [**codersdk.LoginType**](codersdk.LoginType.md) |  | [optional] [default to null] |
| **user\_id** | **UUID** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.OIDCAuthMethod

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **enabled** | **Boolean** |  | [optional] [default to null] |
| **iconUrl** | **String** |  | [optional] [default to null] |
| **signInText** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.OIDCConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **allow\_signups** | **Boolean** |  | [optional] [default to null] |
| **auth\_url\_params** | [**Object**](.md) |  | [optional] [default to null] |
| **client\_cert\_file** | **String** |  | [optional] [default to null] |
| **client\_id** | **String** |  | [optional] [default to null] |
| **client\_key\_file** | **String** | ClientKeyFile &amp; ClientCertFile are used in place of ClientSecret for PKI auth. | [optional] [default to null] |
| **client\_secret** | **String** |  | [optional] [default to null] |
| **email\_domain** | **List** |  | [optional] [default to null] |
| **email\_field** | **String** |  | [optional] [default to null] |
| **group\_allow\_list** | **List** |  | [optional] [default to null] |
| **group\_auto\_create** | **Boolean** |  | [optional] [default to null] |
| **group\_mapping** | [**Object**](.md) |  | [optional] [default to null] |
| **group\_regex\_filter** | [**Object**](.md) |  | [optional] [default to null] |
| **groups\_field** | **String** |  | [optional] [default to null] |
| **icon\_url** | [**serpent.URL**](serpent.URL.md) |  | [optional] [default to null] |
| **ignore\_email\_verified** | **Boolean** |  | [optional] [default to null] |
| **ignore\_user\_info** | **Boolean** | IgnoreUserInfo &amp; UserInfoFromAccessToken are mutually exclusive. Only 1 can be set to true. Ideally this would be an enum with 3 states, [&#39;none&#39;, &#39;userinfo&#39;, &#39;access_token&#39;]. However, for backward compatibility, &#x60;ignore_user_info&#x60; must remain. And &#x60;access_token&#x60; is a niche, non-spec compliant edge case. So it&#39;s use is rare, and should not be advised. | [optional] [default to null] |
| **issuer\_url** | **String** |  | [optional] [default to null] |
| **name\_field** | **String** |  | [optional] [default to null] |
| **organization\_assign\_default** | **Boolean** |  | [optional] [default to null] |
| **organization\_field** | **String** |  | [optional] [default to null] |
| **organization\_mapping** | [**Object**](.md) |  | [optional] [default to null] |
| **scopes** | **List** |  | [optional] [default to null] |
| **sign\_in\_text** | **String** |  | [optional] [default to null] |
| **signups\_disabled\_text** | **String** |  | [optional] [default to null] |
| **skip\_issuer\_checks** | **Boolean** |  | [optional] [default to null] |
| **source\_user\_info\_from\_access\_token** | **Boolean** | UserInfoFromAccessToken as mentioned above is an edge case. This allows sourcing the user_info from the access token itself instead of a user_info endpoint. This assumes the access token is a valid JWT with a set of claims to be merged with the id_token. | [optional] [default to null] |
| **user\_role\_field** | **String** |  | [optional] [default to null] |
| **user\_role\_mapping** | [**Object**](.md) |  | [optional] [default to null] |
| **user\_roles\_default** | **List** |  | [optional] [default to null] |
| **username\_field** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.OptionType

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.Organization

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **created\_at** | **Date** |  | [default to null] |
| **description** | **String** |  | [optional] [default to null] |
| **display\_name** | **String** |  | [optional] [default to null] |
| **icon** | **String** |  | [optional] [default to null] |
| **id** | **UUID** |  | [default to null] |
| **is\_default** | **Boolean** |  | [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **updated\_at** | **Date** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.OrganizationMember

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **created\_at** | **Date** |  | [optional] [default to null] |
| **organization\_id** | **UUID** |  | [optional] [default to null] |
| **roles** | [**List**](codersdk.SlimRole.md) |  | [optional] [default to null] |
| **updated\_at** | **Date** |  | [optional] [default to null] |
| **user\_id** | **UUID** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.OrganizationMemberWithUserData

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **avatar\_url** | **String** |  | [optional] [default to null] |
| **created\_at** | **Date** |  | [optional] [default to null] |
| **email** | **String** |  | [optional] [default to null] |
| **global\_roles** | [**List**](codersdk.SlimRole.md) |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **organization\_id** | **UUID** |  | [optional] [default to null] |
| **roles** | [**List**](codersdk.SlimRole.md) |  | [optional] [default to null] |
| **updated\_at** | **Date** |  | [optional] [default to null] |
| **user\_id** | **UUID** |  | [optional] [default to null] |
| **username** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.OrganizationSyncSettings

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **field** | **String** | Field selects the claim field to be used as the created user&#39;s organizations. If the field is the empty string, then no organization updates will ever come from the OIDC provider. | [optional] [default to null] |
| **mapping** | [**Map**](array.md) | Mapping maps from an OIDC claim --&gt; Coder organization uuid | [optional] [default to null] |
| **organization\_assign\_default** | **Boolean** | AssignDefault will ensure the default org is always included for every user, regardless of their claims. This preserves legacy behavior. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PaginatedMembersResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **count** | **Integer** |  | [optional] [default to null] |
| **members** | [**List**](codersdk.OrganizationMemberWithUserData.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ParameterFormType

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PatchGroupIDPSyncConfigRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **auto\_create\_missing\_groups** | **Boolean** |  | [optional] [default to null] |
| **field** | **String** |  | [optional] [default to null] |
| **regex\_filter** | [**Object**](.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PatchGroupIDPSyncMappingRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **add** | [**List**](codersdk_PatchGroupIDPSyncMappingRequest_add_inner.md) |  | [optional] [default to null] |
| **remove** | [**List**](codersdk_PatchGroupIDPSyncMappingRequest_add_inner.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PatchGroupRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **add\_users** | **List** |  | [optional] [default to null] |
| **avatar\_url** | **String** |  | [optional] [default to null] |
| **display\_name** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **quota\_allowance** | **Integer** |  | [optional] [default to null] |
| **remove\_users** | **List** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PatchOrganizationIDPSyncConfigRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **assign\_default** | **Boolean** |  | [optional] [default to null] |
| **field** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PatchOrganizationIDPSyncMappingRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **add** | [**List**](codersdk_PatchGroupIDPSyncMappingRequest_add_inner.md) |  | [optional] [default to null] |
| **remove** | [**List**](codersdk_PatchGroupIDPSyncMappingRequest_add_inner.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PatchRoleIDPSyncConfigRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **field** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PatchRoleIDPSyncMappingRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **add** | [**List**](codersdk_PatchGroupIDPSyncMappingRequest_add_inner.md) |  | [optional] [default to null] |
| **remove** | [**List**](codersdk_PatchGroupIDPSyncMappingRequest_add_inner.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PatchTemplateVersionRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **message** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PatchWorkspaceProxy

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **display\_name** | **String** |  | [default to null] |
| **icon** | **String** |  | [default to null] |
| **id** | **UUID** |  | [default to null] |
| **name** | **String** |  | [default to null] |
| **regenerate\_token** | **Boolean** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.Permission

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **action** | [**codersdk.RBACAction**](codersdk.RBACAction.md) |  | [optional] [default to null] |
| **negate** | **Boolean** | Negate makes this a negative permission | [optional] [default to null] |
| **resource\_type** | [**codersdk.RBACResource**](codersdk.RBACResource.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PostOAuth2ProviderAppRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **callback\_url** | **String** |  | [default to null] |
| **icon** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PostWorkspaceUsageRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **agent\_id** | **UUID** |  | [optional] [default to null] |
| **app\_name** | [**codersdk.UsageAppName**](codersdk.UsageAppName.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PprofConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **address** | [**serpent.HostPort**](serpent.HostPort.md) |  | [optional] [default to null] |
| **enable** | **Boolean** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PrebuildsConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **failure\_hard\_limit** | **Integer** | FailureHardLimit defines the maximum number of consecutive failed prebuild attempts allowed before a preset is considered to be in a hard limit state. When a preset hits this limit, no new prebuilds will be created until the limit is reset. FailureHardLimit is disabled when set to zero. | [optional] [default to null] |
| **reconciliation\_backoff\_interval** | **Integer** | ReconciliationBackoffInterval specifies the amount of time to increase the backoff interval when errors occur during reconciliation. | [optional] [default to null] |
| **reconciliation\_backoff\_lookback** | **Integer** | ReconciliationBackoffLookback determines the time window to look back when calculating the number of failed prebuilds, which influences the backoff strategy. | [optional] [default to null] |
| **reconciliation\_interval** | **Integer** | ReconciliationInterval defines how often the workspace prebuilds state should be reconciled. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PrebuildsSettings

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **reconciliation\_paused** | **Boolean** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.Preset

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **default** | **Boolean** |  | [optional] [default to null] |
| **description** | **String** |  | [optional] [default to null] |
| **desiredPrebuildInstances** | **Integer** |  | [optional] [default to null] |
| **icon** | **String** |  | [optional] [default to null] |
| **id** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **parameters** | [**List**](codersdk.PresetParameter.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PresetParameter

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **name** | **String** |  | [optional] [default to null] |
| **value** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PreviewParameter

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **default\_value** | [**codersdk.NullHCLString**](codersdk.NullHCLString.md) |  | [optional] [default to null] |
| **description** | **String** |  | [optional] [default to null] |
| **diagnostics** | [**List**](codersdk.FriendlyDiagnostic.md) |  | [optional] [default to null] |
| **display\_name** | **String** |  | [optional] [default to null] |
| **ephemeral** | **Boolean** |  | [optional] [default to null] |
| **form\_type** | [**codersdk.ParameterFormType**](codersdk.ParameterFormType.md) |  | [optional] [default to null] |
| **icon** | **String** |  | [optional] [default to null] |
| **mutable** | **Boolean** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **options** | [**List**](codersdk.PreviewParameterOption.md) |  | [optional] [default to null] |
| **order** | **Integer** | legacy_variable_name was removed (&#x3D; 14) | [optional] [default to null] |
| **required** | **Boolean** |  | [optional] [default to null] |
| **styling** | [**codersdk.PreviewParameterStyling**](codersdk.PreviewParameterStyling.md) |  | [optional] [default to null] |
| **type** | [**codersdk.OptionType**](codersdk.OptionType.md) |  | [optional] [default to null] |
| **validations** | [**List**](codersdk.PreviewParameterValidation.md) |  | [optional] [default to null] |
| **value** | [**codersdk.NullHCLString**](codersdk.NullHCLString.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PreviewParameterOption

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **description** | **String** |  | [optional] [default to null] |
| **icon** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **value** | [**codersdk.NullHCLString**](codersdk.NullHCLString.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PreviewParameterStyling

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **disabled** | **Boolean** |  | [optional] [default to null] |
| **label** | **String** |  | [optional] [default to null] |
| **mask\_input** | **Boolean** |  | [optional] [default to null] |
| **placeholder** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PreviewParameterValidation

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **validation\_error** | **String** |  | [optional] [default to null] |
| **validation\_max** | **Integer** |  | [optional] [default to null] |
| **validation\_min** | **Integer** |  | [optional] [default to null] |
| **validation\_monotonic** | **String** |  | [optional] [default to null] |
| **validation\_regex** | **String** | All validation attributes are optional. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PrometheusConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **address** | [**serpent.HostPort**](serpent.HostPort.md) |  | [optional] [default to null] |
| **aggregate\_agent\_stats\_by** | **List** |  | [optional] [default to null] |
| **collect\_agent\_stats** | **Boolean** |  | [optional] [default to null] |
| **collect\_db\_metrics** | **Boolean** |  | [optional] [default to null] |
| **enable** | **Boolean** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ProvisionerConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **daemon\_poll\_interval** | **Integer** |  | [optional] [default to null] |
| **daemon\_poll\_jitter** | **Integer** |  | [optional] [default to null] |
| **daemon\_psk** | **String** |  | [optional] [default to null] |
| **daemon\_types** | **List** |  | [optional] [default to null] |
| **daemons** | **Integer** | Daemons is the number of built-in terraform provisioners. | [optional] [default to null] |
| **force\_cancel\_interval** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ProvisionerDaemon

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **api\_version** | **String** |  | [optional] [default to null] |
| **created\_at** | **Date** |  | [optional] [default to null] |
| **current\_job** | [**codersdk.ProvisionerDaemonJob**](codersdk.ProvisionerDaemonJob.md) |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **key\_id** | **UUID** |  | [optional] [default to null] |
| **key\_name** | **String** | Optional fields. | [optional] [default to null] |
| **last\_seen\_at** | **Date** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **organization\_id** | **UUID** |  | [optional] [default to null] |
| **previous\_job** | [**codersdk.ProvisionerDaemonJob**](codersdk.ProvisionerDaemonJob.md) |  | [optional] [default to null] |
| **provisioners** | **List** |  | [optional] [default to null] |
| **status** | [**codersdk.ProvisionerDaemonStatus**](codersdk.ProvisionerDaemonStatus.md) |  | [optional] [default to null] |
| **tags** | **Map** |  | [optional] [default to null] |
| **version** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ProvisionerDaemonJob

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **UUID** |  | [optional] [default to null] |
| **status** | [**codersdk.ProvisionerJobStatus**](codersdk.ProvisionerJobStatus.md) |  | [optional] [default to null] |
| **template\_display\_name** | **String** |  | [optional] [default to null] |
| **template\_icon** | **String** |  | [optional] [default to null] |
| **template\_name** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ProvisionerDaemonStatus

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ProvisionerJob

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **available\_workers** | **List** |  | [optional] [default to null] |
| **canceled\_at** | **Date** |  | [optional] [default to null] |
| **completed\_at** | **Date** |  | [optional] [default to null] |
| **created\_at** | **Date** |  | [optional] [default to null] |
| **error** | **String** |  | [optional] [default to null] |
| **error\_code** | [**codersdk.JobErrorCode**](codersdk.JobErrorCode.md) |  | [optional] [default to null] |
| **file\_id** | **UUID** |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **input** | [**codersdk.ProvisionerJobInput**](codersdk.ProvisionerJobInput.md) |  | [optional] [default to null] |
| **logs\_overflowed** | **Boolean** |  | [optional] [default to null] |
| **metadata** | [**codersdk.ProvisionerJobMetadata**](codersdk.ProvisionerJobMetadata.md) |  | [optional] [default to null] |
| **organization\_id** | **UUID** |  | [optional] [default to null] |
| **queue\_position** | **Integer** |  | [optional] [default to null] |
| **queue\_size** | **Integer** |  | [optional] [default to null] |
| **started\_at** | **Date** |  | [optional] [default to null] |
| **status** | [**codersdk.ProvisionerJobStatus**](codersdk.ProvisionerJobStatus.md) |  | [optional] [default to null] |
| **tags** | **Map** |  | [optional] [default to null] |
| **type** | [**codersdk.ProvisionerJobType**](codersdk.ProvisionerJobType.md) |  | [optional] [default to null] |
| **worker\_id** | **UUID** |  | [optional] [default to null] |
| **worker\_name** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ProvisionerJobInput

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **error** | **String** |  | [optional] [default to null] |
| **template\_version\_id** | **UUID** |  | [optional] [default to null] |
| **workspace\_build\_id** | **UUID** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ProvisionerJobLog

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **created\_at** | **Date** |  | [optional] [default to null] |
| **id** | **Integer** |  | [optional] [default to null] |
| **log\_level** | [**codersdk.LogLevel**](codersdk.LogLevel.md) |  | [optional] [default to null] |
| **log\_source** | [**codersdk.LogSource**](codersdk.LogSource.md) |  | [optional] [default to null] |
| **output** | **String** |  | [optional] [default to null] |
| **stage** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ProvisionerJobMetadata

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **template\_display\_name** | **String** |  | [optional] [default to null] |
| **template\_icon** | **String** |  | [optional] [default to null] |
| **template\_id** | **UUID** |  | [optional] [default to null] |
| **template\_name** | **String** |  | [optional] [default to null] |
| **template\_version\_name** | **String** |  | [optional] [default to null] |
| **workspace\_id** | **UUID** |  | [optional] [default to null] |
| **workspace\_name** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ProvisionerJobStatus

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ProvisionerJobType

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ProvisionerKey

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **created\_at** | **Date** |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **organization** | **UUID** |  | [optional] [default to null] |
| **tags** | **Map** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ProvisionerKeyDaemons

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **daemons** | [**List**](codersdk.ProvisionerDaemon.md) |  | [optional] [default to null] |
| **key** | [**codersdk.ProvisionerKey**](codersdk.ProvisionerKey.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ProvisionerLogLevel

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ProvisionerStorageMethod

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ProvisionerTiming

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **action** | **String** |  | [optional] [default to null] |
| **ended\_at** | **Date** |  | [optional] [default to null] |
| **job\_id** | **UUID** |  | [optional] [default to null] |
| **resource** | **String** |  | [optional] [default to null] |
| **source** | **String** |  | [optional] [default to null] |
| **stage** | [**codersdk.TimingStage**](codersdk.TimingStage.md) |  | [optional] [default to null] |
| **started\_at** | **Date** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ProxyHealthReport

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **errors** | **List** | Errors are problems that prevent the workspace proxy from being healthy | [optional] [default to null] |
| **warnings** | **List** | Warnings do not prevent the workspace proxy from being healthy, but should be addressed. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ProxyHealthStatus

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PutExtendWorkspaceRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **deadline** | **Date** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.PutOAuth2ProviderAppRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **callback\_url** | **String** |  | [default to null] |
| **icon** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.RBACAction

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.RBACResource

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.RateLimitConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **api** | **Integer** |  | [optional] [default to null] |
| **disable\_all** | **Boolean** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ReducedUser

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **avatar\_url** | **URI** |  | [optional] [default to null] |
| **created\_at** | **Date** |  | [default to null] |
| **email** | **String** |  | [default to null] |
| **id** | **UUID** |  | [default to null] |
| **last\_seen\_at** | **Date** |  | [optional] [default to null] |
| **login\_type** | [**codersdk.LoginType**](codersdk.LoginType.md) |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **status** | [**codersdk.UserStatus**](codersdk.UserStatus.md) |  | [optional] [default to null] |
| **theme\_preference** | **String** | Deprecated: this value should be retrieved from &#x60;codersdk.UserPreferenceSettings&#x60; instead. | [optional] [default to null] |
| **updated\_at** | **Date** |  | [optional] [default to null] |
| **username** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.Region

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **display\_name** | **String** |  | [optional] [default to null] |
| **healthy** | **Boolean** |  | [optional] [default to null] |
| **icon\_url** | **String** |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **path\_app\_url** | **String** | PathAppURL is the URL to the base path for path apps. Optional unless wildcard_hostname is set. E.g. https://us.example.com | [optional] [default to null] |
| **wildcard\_hostname** | **String** | WildcardHostname is the wildcard hostname for subdomain apps. E.g. *.us.example.com E.g. *--suffix.au.example.com Optional. Does not need to be on the same domain as PathAppURL. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.RegionsResponse-codersdk_Region

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **regions** | [**List**](codersdk.Region.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.RegionsResponse-codersdk_WorkspaceProxy

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **regions** | [**List**](codersdk.WorkspaceProxy.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.Replica

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **created\_at** | **Date** | CreatedAt is the timestamp when the replica was first seen. | [optional] [default to null] |
| **database\_latency** | **Integer** | DatabaseLatency is the latency in microseconds to the database. | [optional] [default to null] |
| **error** | **String** | Error is the replica error. | [optional] [default to null] |
| **hostname** | **String** | Hostname is the hostname of the replica. | [optional] [default to null] |
| **id** | **UUID** | ID is the unique identifier for the replica. | [optional] [default to null] |
| **region\_id** | **Integer** | RegionID is the region of the replica. | [optional] [default to null] |
| **relay\_address** | **String** | RelayAddress is the accessible address to relay DERP connections. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.RequestOneTimePasscodeRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **email** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ResolveAutostartResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **parameter\_mismatch** | **Boolean** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ResourceType

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.Response

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **detail** | **String** | Detail is a debug message that provides further insight into why the action failed. This information can be technical and a regular golang err.Error() text. - \&quot;database: too many open connections\&quot; - \&quot;stat: too many open files\&quot; | [optional] [default to null] |
| **message** | **String** | Message is an actionable message that depicts actions the request took. These messages should be fully formed sentences with proper punctuation. Examples: - \&quot;A user has been created.\&quot; - \&quot;Failed to create a user.\&quot; | [optional] [default to null] |
| **validations** | [**List**](codersdk.ValidationError.md) | Validations are form field-specific friendly error messages. They will be shown on a form field in the UI. These can also be used to add additional context if there is a set of errors in the primary &#39;Message&#39;. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.Role

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **display\_name** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **organization\_id** | **UUID** |  | [optional] [default to null] |
| **organization\_permissions** | [**List**](codersdk.Permission.md) | OrganizationPermissions are specific for the organization in the field &#39;OrganizationID&#39; above. | [optional] [default to null] |
| **site\_permissions** | [**List**](codersdk.Permission.md) |  | [optional] [default to null] |
| **user\_permissions** | [**List**](codersdk.Permission.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.RoleSyncSettings

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **field** | **String** | Field is the name of the claim field that specifies what organization roles a user should be given. If empty, no roles will be synced. | [optional] [default to null] |
| **mapping** | [**Map**](array.md) | Mapping is a map from OIDC groups to Coder organization roles. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.SSHConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **deploymentName** | **String** | DeploymentName is the config-ssh Hostname prefix | [optional] [default to null] |
| **sshconfigOptions** | **List** | SSHConfigOptions are additional options to add to the ssh config file. This will override defaults. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.SSHConfigResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **hostname\_prefix** | **String** | HostnamePrefix is the prefix we append to workspace names for SSH hostnames. Deprecated: use HostnameSuffix instead. | [optional] [default to null] |
| **hostname\_suffix** | **String** | HostnameSuffix is the suffix to append to workspace names for SSH hostnames. | [optional] [default to null] |
| **ssh\_config\_options** | **Map** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ServerSentEvent

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **data** | [**Object**](.md) |  | [optional] [default to null] |
| **type** | [**codersdk.ServerSentEventType**](codersdk.ServerSentEventType.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ServerSentEventType

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.SessionCountDeploymentStats

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **jetbrains** | **Integer** |  | [optional] [default to null] |
| **reconnecting\_pty** | **Integer** |  | [optional] [default to null] |
| **ssh** | **Integer** |  | [optional] [default to null] |
| **vscode** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.SessionLifetime

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **default\_duration** | **Integer** | DefaultDuration is only for browser, workspace app and oauth sessions. | [optional] [default to null] |
| **default\_token\_lifetime** | **Integer** |  | [optional] [default to null] |
| **disable\_expiry\_refresh** | **Boolean** | DisableExpiryRefresh will disable automatically refreshing api keys when they are used from the api. This means the api key lifetime at creation is the lifetime of the api key. | [optional] [default to null] |
| **max\_admin\_token\_lifetime** | **Integer** |  | [optional] [default to null] |
| **max\_token\_lifetime** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.SlimRole

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **display\_name** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **organization\_id** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.SupportConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **links** | [**serpent.Struct-array_codersdk_LinkConfig**](serpent.Struct-array_codersdk_LinkConfig.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.SwaggerConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **enable** | **Boolean** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TLSConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **address** | [**serpent.HostPort**](serpent.HostPort.md) |  | [optional] [default to null] |
| **allow\_insecure\_ciphers** | **Boolean** |  | [optional] [default to null] |
| **cert\_file** | **List** |  | [optional] [default to null] |
| **client\_auth** | **String** |  | [optional] [default to null] |
| **client\_ca\_file** | **String** |  | [optional] [default to null] |
| **client\_cert\_file** | **String** |  | [optional] [default to null] |
| **client\_key\_file** | **String** |  | [optional] [default to null] |
| **enable** | **Boolean** |  | [optional] [default to null] |
| **key\_file** | **List** |  | [optional] [default to null] |
| **min\_version** | **String** |  | [optional] [default to null] |
| **redirect\_http** | **Boolean** |  | [optional] [default to null] |
| **supported\_ciphers** | **List** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TelemetryConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **enable** | **Boolean** |  | [optional] [default to null] |
| **trace** | **Boolean** |  | [optional] [default to null] |
| **url** | [**serpent.URL**](serpent.URL.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.Template

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **active\_user\_count** | **Integer** | ActiveUserCount is set to -1 when loading. | [optional] [default to null] |
| **active\_version\_id** | **UUID** |  | [optional] [default to null] |
| **activity\_bump\_ms** | **Integer** |  | [optional] [default to null] |
| **allow\_user\_autostart** | **Boolean** | AllowUserAutostart and AllowUserAutostop are enterprise-only. Their values are only used if your license is entitled to use the advanced template scheduling feature. | [optional] [default to null] |
| **allow\_user\_autostop** | **Boolean** |  | [optional] [default to null] |
| **allow\_user\_cancel\_workspace\_jobs** | **Boolean** |  | [optional] [default to null] |
| **autostart\_requirement** | [**codersdk.TemplateAutostartRequirement**](codersdk.TemplateAutostartRequirement.md) |  | [optional] [default to null] |
| **autostop\_requirement** | [**codersdk.TemplateAutostopRequirement**](codersdk.TemplateAutostopRequirement.md) | AutostopRequirement and AutostartRequirement are enterprise features. Its value is only used if your license is entitled to use the advanced template scheduling feature. | [optional] [default to null] |
| **build\_time\_stats** | [**Map**](codersdk.TransitionStats.md) |  | [optional] [default to null] |
| **cors\_behavior** | [**codersdk.CORSBehavior**](codersdk.CORSBehavior.md) |  | [optional] [default to null] |
| **created\_at** | **Date** |  | [optional] [default to null] |
| **created\_by\_id** | **UUID** |  | [optional] [default to null] |
| **created\_by\_name** | **String** |  | [optional] [default to null] |
| **default\_ttl\_ms** | **Integer** |  | [optional] [default to null] |
| **deprecated** | **Boolean** |  | [optional] [default to null] |
| **deprecation\_message** | **String** |  | [optional] [default to null] |
| **description** | **String** |  | [optional] [default to null] |
| **display\_name** | **String** |  | [optional] [default to null] |
| **failure\_ttl\_ms** | **Integer** | FailureTTLMillis, TimeTilDormantMillis, and TimeTilDormantAutoDeleteMillis are enterprise-only. Their values are used if your license is entitled to use the advanced template scheduling feature. | [optional] [default to null] |
| **icon** | **String** |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **max\_port\_share\_level** | [**codersdk.WorkspaceAgentPortShareLevel**](codersdk.WorkspaceAgentPortShareLevel.md) |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **organization\_display\_name** | **String** |  | [optional] [default to null] |
| **organization\_icon** | **String** |  | [optional] [default to null] |
| **organization\_id** | **UUID** |  | [optional] [default to null] |
| **organization\_name** | **String** |  | [optional] [default to null] |
| **provisioner** | **String** |  | [optional] [default to null] |
| **require\_active\_version** | **Boolean** | RequireActiveVersion mandates that workspaces are built with the active template version. | [optional] [default to null] |
| **time\_til\_dormant\_autodelete\_ms** | **Integer** |  | [optional] [default to null] |
| **time\_til\_dormant\_ms** | **Integer** |  | [optional] [default to null] |
| **updated\_at** | **Date** |  | [optional] [default to null] |
| **use\_classic\_parameter\_flow** | **Boolean** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TemplateACL

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **group** | [**List**](codersdk.TemplateGroup.md) |  | [optional] [default to null] |
| **users** | [**List**](codersdk.TemplateUser.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TemplateAppUsage

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **display\_name** | **String** |  | [optional] [default to null] |
| **icon** | **String** |  | [optional] [default to null] |
| **seconds** | **Integer** |  | [optional] [default to null] |
| **slug** | **String** |  | [optional] [default to null] |
| **template\_ids** | **List** |  | [optional] [default to null] |
| **times\_used** | **Integer** |  | [optional] [default to null] |
| **type** | [**codersdk.TemplateAppsType**](codersdk.TemplateAppsType.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TemplateAppsType

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TemplateAutostartRequirement

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **days\_of\_week** | **List** | DaysOfWeek is a list of days of the week in which autostart is allowed to happen. If no days are specified, autostart is not allowed. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TemplateAutostopRequirement

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **days\_of\_week** | **List** | DaysOfWeek is a list of days of the week on which restarts are required. Restarts happen within the user&#39;s quiet hours (in their configured timezone). If no days are specified, restarts are not required. Weekdays cannot be specified twice.  Restarts will only happen on weekdays in this list on weeks which line up with Weeks. | [optional] [default to null] |
| **weeks** | **Integer** | Weeks is the number of weeks between required restarts. Weeks are synced across all workspaces (and Coder deployments) using modulo math on a hardcoded epoch week of January 2nd, 2023 (the first Monday of 2023). Values of 0 or 1 indicate weekly restarts. Values of 2 indicate fortnightly restarts, etc. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TemplateExample

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **description** | **String** |  | [optional] [default to null] |
| **icon** | **String** |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **markdown** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **tags** | **List** |  | [optional] [default to null] |
| **url** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TemplateGroup

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **avatar\_url** | **URI** |  | [optional] [default to null] |
| **display\_name** | **String** |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **members** | [**List**](codersdk.ReducedUser.md) |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **organization\_display\_name** | **String** |  | [optional] [default to null] |
| **organization\_id** | **UUID** |  | [optional] [default to null] |
| **organization\_name** | **String** |  | [optional] [default to null] |
| **quota\_allowance** | **Integer** |  | [optional] [default to null] |
| **role** | [**codersdk.TemplateRole**](codersdk.TemplateRole.md) |  | [optional] [default to null] |
| **source** | [**codersdk.GroupSource**](codersdk.GroupSource.md) |  | [optional] [default to null] |
| **total\_member\_count** | **Integer** | How many members are in this group. Shows the total count, even if the user is not authorized to read group member details. May be greater than &#x60;len(Group.Members)&#x60;. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TemplateInsightsIntervalReport

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **active\_users** | **Integer** |  | [optional] [default to null] |
| **end\_time** | **Date** |  | [optional] [default to null] |
| **interval** | [**codersdk.InsightsReportInterval**](codersdk.InsightsReportInterval.md) |  | [optional] [default to null] |
| **start\_time** | **Date** |  | [optional] [default to null] |
| **template\_ids** | **List** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TemplateInsightsReport

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **active\_users** | **Integer** |  | [optional] [default to null] |
| **apps\_usage** | [**List**](codersdk.TemplateAppUsage.md) |  | [optional] [default to null] |
| **end\_time** | **Date** |  | [optional] [default to null] |
| **parameters\_usage** | [**List**](codersdk.TemplateParameterUsage.md) |  | [optional] [default to null] |
| **start\_time** | **Date** |  | [optional] [default to null] |
| **template\_ids** | **List** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TemplateInsightsResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **interval\_reports** | [**List**](codersdk.TemplateInsightsIntervalReport.md) |  | [optional] [default to null] |
| **report** | [**codersdk.TemplateInsightsReport**](codersdk.TemplateInsightsReport.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TemplateParameterUsage

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **description** | **String** |  | [optional] [default to null] |
| **display\_name** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **options** | [**List**](codersdk.TemplateVersionParameterOption.md) |  | [optional] [default to null] |
| **template\_ids** | **List** |  | [optional] [default to null] |
| **type** | **String** |  | [optional] [default to null] |
| **values** | [**List**](codersdk.TemplateParameterValue.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TemplateParameterValue

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **count** | **Integer** |  | [optional] [default to null] |
| **value** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TemplateRole

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TemplateUser

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **avatar\_url** | **URI** |  | [optional] [default to null] |
| **created\_at** | **Date** |  | [default to null] |
| **email** | **String** |  | [default to null] |
| **id** | **UUID** |  | [default to null] |
| **last\_seen\_at** | **Date** |  | [optional] [default to null] |
| **login\_type** | [**codersdk.LoginType**](codersdk.LoginType.md) |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **organization\_ids** | **List** |  | [optional] [default to null] |
| **role** | [**codersdk.TemplateRole**](codersdk.TemplateRole.md) |  | [optional] [default to null] |
| **roles** | [**List**](codersdk.SlimRole.md) |  | [optional] [default to null] |
| **status** | [**codersdk.UserStatus**](codersdk.UserStatus.md) |  | [optional] [default to null] |
| **theme\_preference** | **String** | Deprecated: this value should be retrieved from &#x60;codersdk.UserPreferenceSettings&#x60; instead. | [optional] [default to null] |
| **updated\_at** | **Date** |  | [optional] [default to null] |
| **username** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TemplateVersion

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **archived** | **Boolean** |  | [optional] [default to null] |
| **created\_at** | **Date** |  | [optional] [default to null] |
| **created\_by** | [**codersdk.MinimalUser**](codersdk.MinimalUser.md) |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **job** | [**codersdk.ProvisionerJob**](codersdk.ProvisionerJob.md) |  | [optional] [default to null] |
| **matched\_provisioners** | [**codersdk.MatchedProvisioners**](codersdk.MatchedProvisioners.md) |  | [optional] [default to null] |
| **message** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **organization\_id** | **UUID** |  | [optional] [default to null] |
| **readme** | **String** |  | [optional] [default to null] |
| **template\_id** | **UUID** |  | [optional] [default to null] |
| **updated\_at** | **Date** |  | [optional] [default to null] |
| **warnings** | [**List**](codersdk.TemplateVersionWarning.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TemplateVersionExternalAuth

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **authenticate\_url** | **String** |  | [optional] [default to null] |
| **authenticated** | **Boolean** |  | [optional] [default to null] |
| **display\_icon** | **String** |  | [optional] [default to null] |
| **display\_name** | **String** |  | [optional] [default to null] |
| **id** | **String** |  | [optional] [default to null] |
| **optional** | **Boolean** |  | [optional] [default to null] |
| **type** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TemplateVersionParameter

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **default\_value** | **String** |  | [optional] [default to null] |
| **description** | **String** |  | [optional] [default to null] |
| **description\_plaintext** | **String** |  | [optional] [default to null] |
| **display\_name** | **String** |  | [optional] [default to null] |
| **ephemeral** | **Boolean** |  | [optional] [default to null] |
| **form\_type** | **String** | FormType has an enum value of empty string, &#x60;\&quot;\&quot;&#x60;. Keep the leading comma in the enums struct tag. | [optional] [default to null] |
| **icon** | **String** |  | [optional] [default to null] |
| **mutable** | **Boolean** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **options** | [**List**](codersdk.TemplateVersionParameterOption.md) |  | [optional] [default to null] |
| **required** | **Boolean** |  | [optional] [default to null] |
| **type** | **String** |  | [optional] [default to null] |
| **validation\_error** | **String** |  | [optional] [default to null] |
| **validation\_max** | **Integer** |  | [optional] [default to null] |
| **validation\_min** | **Integer** |  | [optional] [default to null] |
| **validation\_monotonic** | [**codersdk.ValidationMonotonicOrder**](codersdk.ValidationMonotonicOrder.md) |  | [optional] [default to null] |
| **validation\_regex** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TemplateVersionParameterOption

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **description** | **String** |  | [optional] [default to null] |
| **icon** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **value** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TemplateVersionVariable

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **default\_value** | **String** |  | [optional] [default to null] |
| **description** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **required** | **Boolean** |  | [optional] [default to null] |
| **sensitive** | **Boolean** |  | [optional] [default to null] |
| **type** | **String** |  | [optional] [default to null] |
| **value** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TemplateVersionWarning

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TerminalFontName

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TimingStage

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TokenConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **max\_token\_lifetime** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TraceConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **capture\_logs** | **Boolean** |  | [optional] [default to null] |
| **data\_dog** | **Boolean** |  | [optional] [default to null] |
| **enable** | **Boolean** |  | [optional] [default to null] |
| **honeycomb\_api\_key** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.TransitionStats

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **p50** | **Integer** |  | [optional] [default to null] |
| **p95** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UpdateActiveTemplateVersion

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **UUID** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UpdateAppearanceConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **announcement\_banners** | [**List**](codersdk.BannerConfig.md) |  | [optional] [default to null] |
| **application\_name** | **String** |  | [optional] [default to null] |
| **logo\_url** | **String** |  | [optional] [default to null] |
| **service\_banner** | [**codersdk.BannerConfig**](codersdk.BannerConfig.md) | Deprecated: ServiceBanner has been replaced by AnnouncementBanners. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UpdateCheckResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **current** | **Boolean** | Current indicates whether the server version is the same as the latest. | [optional] [default to null] |
| **url** | **String** | URL to download the latest release of Coder. | [optional] [default to null] |
| **version** | **String** | Version is the semantic version for the latest release of Coder. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UpdateOrganizationRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **description** | **String** |  | [optional] [default to null] |
| **display\_name** | **String** |  | [optional] [default to null] |
| **icon** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UpdateRoles

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **roles** | **List** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UpdateTemplateACL

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **group\_perms** | [**Map**](codersdk.TemplateRole.md) | GroupPerms should be a mapping of group id to role. | [optional] [default to null] |
| **user\_perms** | [**Map**](codersdk.TemplateRole.md) | UserPerms should be a mapping of user id to role. The user id must be the uuid of the user, not a username or email address. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UpdateUserAppearanceSettingsRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **terminal\_font** | [**codersdk.TerminalFontName**](codersdk.TerminalFontName.md) |  | [default to null] |
| **theme\_preference** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UpdateUserNotificationPreferences

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **template\_disabled\_map** | **Map** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UpdateUserPasswordRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **old\_password** | **String** |  | [optional] [default to null] |
| **password** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UpdateUserProfileRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **name** | **String** |  | [optional] [default to null] |
| **username** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UpdateUserQuietHoursScheduleRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **schedule** | **String** | Schedule is a cron expression that defines when the user&#39;s quiet hours window is. Schedule must not be empty. For new users, the schedule is set to 2am in their browser or computer&#39;s timezone. The schedule denotes the beginning of a 4 hour window where the workspace is allowed to automatically stop or restart due to maintenance or template schedule.  The schedule must be daily with a single time, and should have a timezone specified via a CRON_TZ prefix (otherwise UTC will be used).  If the schedule is empty, the user will be updated to use the default schedule. | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UpdateWorkspaceACL

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **group\_roles** | [**Map**](codersdk.WorkspaceRole.md) |  | [optional] [default to null] |
| **user\_roles** | [**Map**](codersdk.WorkspaceRole.md) | Keys must be valid UUIDs. To remove a user/group from the ACL use \&quot;\&quot; as the role name (available as a constant named &#x60;codersdk.WorkspaceRoleDeleted&#x60;) | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UpdateWorkspaceAutomaticUpdatesRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **automatic\_updates** | [**codersdk.AutomaticUpdates**](codersdk.AutomaticUpdates.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UpdateWorkspaceAutostartRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **schedule** | **String** | Schedule is expected to be of the form &#x60;CRON_TZ&#x3D;&lt;IANA Timezone&gt; &lt;min&gt; &lt;hour&gt; * * &lt;dow&gt;&#x60; Example: &#x60;CRON_TZ&#x3D;US/Central 30 9 * * 1-5&#x60; represents 0930 in the timezone US/Central on weekdays (Mon-Fri). &#x60;CRON_TZ&#x60; defaults to UTC if not present. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UpdateWorkspaceDormancy

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **dormant** | **Boolean** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UpdateWorkspaceRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **name** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UpdateWorkspaceTTLRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **ttl\_ms** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UploadResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **hash** | **UUID** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UpsertWorkspaceAgentPortShareRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **agent\_name** | **String** |  | [optional] [default to null] |
| **port** | **Integer** |  | [optional] [default to null] |
| **protocol** | [**codersdk.WorkspaceAgentPortShareProtocol**](codersdk.WorkspaceAgentPortShareProtocol.md) |  | [optional] [default to null] |
| **share\_level** | [**codersdk.WorkspaceAgentPortShareLevel**](codersdk.WorkspaceAgentPortShareLevel.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UsageAppName

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UsagePeriod

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **end** | **Date** |  | [optional] [default to null] |
| **issued\_at** | **Date** |  | [optional] [default to null] |
| **start** | **Date** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.User

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **avatar\_url** | **URI** |  | [optional] [default to null] |
| **created\_at** | **Date** |  | [default to null] |
| **email** | **String** |  | [default to null] |
| **id** | **UUID** |  | [default to null] |
| **last\_seen\_at** | **Date** |  | [optional] [default to null] |
| **login\_type** | [**codersdk.LoginType**](codersdk.LoginType.md) |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **organization\_ids** | **List** |  | [optional] [default to null] |
| **roles** | [**List**](codersdk.SlimRole.md) |  | [optional] [default to null] |
| **status** | [**codersdk.UserStatus**](codersdk.UserStatus.md) |  | [optional] [default to null] |
| **theme\_preference** | **String** | Deprecated: this value should be retrieved from &#x60;codersdk.UserPreferenceSettings&#x60; instead. | [optional] [default to null] |
| **updated\_at** | **Date** |  | [optional] [default to null] |
| **username** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UserActivity

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **avatar\_url** | **URI** |  | [optional] [default to null] |
| **seconds** | **Integer** |  | [optional] [default to null] |
| **template\_ids** | **List** |  | [optional] [default to null] |
| **user\_id** | **UUID** |  | [optional] [default to null] |
| **username** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UserActivityInsightsReport

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **end\_time** | **Date** |  | [optional] [default to null] |
| **start\_time** | **Date** |  | [optional] [default to null] |
| **template\_ids** | **List** |  | [optional] [default to null] |
| **users** | [**List**](codersdk.UserActivity.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UserActivityInsightsResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **report** | [**codersdk.UserActivityInsightsReport**](codersdk.UserActivityInsightsReport.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UserAppearanceSettings

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **terminal\_font** | [**codersdk.TerminalFontName**](codersdk.TerminalFontName.md) |  | [optional] [default to null] |
| **theme\_preference** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UserLatency

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **avatar\_url** | **URI** |  | [optional] [default to null] |
| **latency\_ms** | [**codersdk.ConnectionLatency**](codersdk.ConnectionLatency.md) |  | [optional] [default to null] |
| **template\_ids** | **List** |  | [optional] [default to null] |
| **user\_id** | **UUID** |  | [optional] [default to null] |
| **username** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UserLatencyInsightsReport

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **end\_time** | **Date** |  | [optional] [default to null] |
| **start\_time** | **Date** |  | [optional] [default to null] |
| **template\_ids** | **List** |  | [optional] [default to null] |
| **users** | [**List**](codersdk.UserLatency.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UserLatencyInsightsResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **report** | [**codersdk.UserLatencyInsightsReport**](codersdk.UserLatencyInsightsReport.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UserLoginType

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **login\_type** | [**codersdk.LoginType**](codersdk.LoginType.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UserParameter

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **name** | **String** |  | [optional] [default to null] |
| **value** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UserQuietHoursScheduleConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **allow\_user\_custom** | **Boolean** |  | [optional] [default to null] |
| **default\_schedule** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UserQuietHoursScheduleResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **next** | **Date** | Next is the next time that the quiet hours window will start. | [optional] [default to null] |
| **raw\_schedule** | **String** |  | [optional] [default to null] |
| **time** | **String** | Time is the time of day that the quiet hours window starts in the given Timezone each day. | [optional] [default to null] |
| **timezone** | **String** | raw format from the cron expression, UTC if unspecified | [optional] [default to null] |
| **user\_can\_set** | **Boolean** | UserCanSet is true if the user is allowed to set their own quiet hours schedule. If false, the user cannot set a custom schedule and the default schedule will always be used. | [optional] [default to null] |
| **user\_set** | **Boolean** | UserSet is true if the user has set their own quiet hours schedule. If false, the user is using the default schedule. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UserStatus

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.UserStatusChangeCount

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **count** | **Integer** |  | [optional] [default to null] |
| **date** | **Date** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ValidateUserPasswordRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **password** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ValidateUserPasswordResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **details** | **String** |  | [optional] [default to null] |
| **valid** | **Boolean** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ValidationError

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **detail** | **String** |  | [default to null] |
| **field** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.ValidationMonotonicOrder

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.VariableValue

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **name** | **String** |  | [optional] [default to null] |
| **value** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WebpushSubscription

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **auth\_key** | **String** |  | [optional] [default to null] |
| **endpoint** | **String** |  | [optional] [default to null] |
| **p256dh\_key** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.Workspace

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **allow\_renames** | **Boolean** |  | [optional] [default to null] |
| **automatic\_updates** | [**codersdk.AutomaticUpdates**](codersdk.AutomaticUpdates.md) |  | [optional] [default to null] |
| **autostart\_schedule** | **String** |  | [optional] [default to null] |
| **created\_at** | **Date** |  | [optional] [default to null] |
| **deleting\_at** | **Date** | DeletingAt indicates the time at which the workspace will be permanently deleted. A workspace is eligible for deletion if it is dormant (a non-nil dormant_at value) and a value has been specified for time_til_dormant_autodelete on its template. | [optional] [default to null] |
| **dormant\_at** | **Date** | DormantAt being non-nil indicates a workspace that is dormant. A dormant workspace is no longer accessible must be activated. It is subject to deletion if it breaches the duration of the time_til_ field on its template. | [optional] [default to null] |
| **favorite** | **Boolean** |  | [optional] [default to null] |
| **health** | [**codersdk.WorkspaceHealth**](codersdk.WorkspaceHealth.md) | Health shows the health of the workspace and information about what is causing an unhealthy status. | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **is\_prebuild** | **Boolean** | IsPrebuild indicates whether the workspace is a prebuilt workspace. Prebuilt workspaces are owned by the prebuilds system user and have specific behavior, such as being managed differently from regular workspaces. Once a prebuilt workspace is claimed by a user, it transitions to a regular workspace, and IsPrebuild returns false. | [optional] [default to null] |
| **last\_used\_at** | **Date** |  | [optional] [default to null] |
| **latest\_app\_status** | [**codersdk.WorkspaceAppStatus**](codersdk.WorkspaceAppStatus.md) |  | [optional] [default to null] |
| **latest\_build** | [**codersdk.WorkspaceBuild**](codersdk.WorkspaceBuild.md) |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **next\_start\_at** | **Date** |  | [optional] [default to null] |
| **organization\_id** | **UUID** |  | [optional] [default to null] |
| **organization\_name** | **String** |  | [optional] [default to null] |
| **outdated** | **Boolean** |  | [optional] [default to null] |
| **owner\_avatar\_url** | **String** |  | [optional] [default to null] |
| **owner\_id** | **UUID** |  | [optional] [default to null] |
| **owner\_name** | **String** | OwnerName is the username of the owner of the workspace. | [optional] [default to null] |
| **template\_active\_version\_id** | **UUID** |  | [optional] [default to null] |
| **template\_allow\_user\_cancel\_workspace\_jobs** | **Boolean** |  | [optional] [default to null] |
| **template\_display\_name** | **String** |  | [optional] [default to null] |
| **template\_icon** | **String** |  | [optional] [default to null] |
| **template\_id** | **UUID** |  | [optional] [default to null] |
| **template\_name** | **String** |  | [optional] [default to null] |
| **template\_require\_active\_version** | **Boolean** |  | [optional] [default to null] |
| **template\_use\_classic\_parameter\_flow** | **Boolean** |  | [optional] [default to null] |
| **ttl\_ms** | **Integer** |  | [optional] [default to null] |
| **updated\_at** | **Date** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAgent

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **api\_version** | **String** |  | [optional] [default to null] |
| **apps** | [**List**](codersdk.WorkspaceApp.md) |  | [optional] [default to null] |
| **architecture** | **String** |  | [optional] [default to null] |
| **connection\_timeout\_seconds** | **Integer** |  | [optional] [default to null] |
| **created\_at** | **Date** |  | [optional] [default to null] |
| **directory** | **String** |  | [optional] [default to null] |
| **disconnected\_at** | **Date** |  | [optional] [default to null] |
| **display\_apps** | [**List**](codersdk.DisplayApp.md) |  | [optional] [default to null] |
| **environment\_variables** | **Map** |  | [optional] [default to null] |
| **expanded\_directory** | **String** |  | [optional] [default to null] |
| **first\_connected\_at** | **Date** |  | [optional] [default to null] |
| **health** | [**codersdk.WorkspaceAgentHealth**](codersdk.WorkspaceAgentHealth.md) | Health reports the health of the agent. | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **instance\_id** | **String** |  | [optional] [default to null] |
| **last\_connected\_at** | **Date** |  | [optional] [default to null] |
| **latency** | [**Map**](codersdk.DERPRegion.md) | DERPLatency is mapped by region name (e.g. \&quot;New York City\&quot;, \&quot;Seattle\&quot;). | [optional] [default to null] |
| **lifecycle\_state** | [**codersdk.WorkspaceAgentLifecycle**](codersdk.WorkspaceAgentLifecycle.md) |  | [optional] [default to null] |
| **log\_sources** | [**List**](codersdk.WorkspaceAgentLogSource.md) |  | [optional] [default to null] |
| **logs\_length** | **Integer** |  | [optional] [default to null] |
| **logs\_overflowed** | **Boolean** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **operating\_system** | **String** |  | [optional] [default to null] |
| **parent\_id** | [**uuid.NullUUID**](uuid.NullUUID.md) |  | [optional] [default to null] |
| **ready\_at** | **Date** |  | [optional] [default to null] |
| **resource\_id** | **UUID** |  | [optional] [default to null] |
| **scripts** | [**List**](codersdk.WorkspaceAgentScript.md) |  | [optional] [default to null] |
| **started\_at** | **Date** |  | [optional] [default to null] |
| **startup\_script\_behavior** | [**codersdk.WorkspaceAgentStartupScriptBehavior**](codersdk.WorkspaceAgentStartupScriptBehavior.md) | StartupScriptBehavior is a legacy field that is deprecated in favor of the &#x60;coder_script&#x60; resource. It&#39;s only referenced by old clients. Deprecated: Remove in the future! | [optional] [default to null] |
| **status** | [**codersdk.WorkspaceAgentStatus**](codersdk.WorkspaceAgentStatus.md) |  | [optional] [default to null] |
| **subsystems** | [**List**](codersdk.AgentSubsystem.md) |  | [optional] [default to null] |
| **troubleshooting\_url** | **String** |  | [optional] [default to null] |
| **updated\_at** | **Date** |  | [optional] [default to null] |
| **version** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAgentContainer

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **created\_at** | **Date** | CreatedAt is the time the container was created. | [optional] [default to null] |
| **id** | **String** | ID is the unique identifier of the container. | [optional] [default to null] |
| **image** | **String** | Image is the name of the container image. | [optional] [default to null] |
| **labels** | **Map** | Labels is a map of key-value pairs of container labels. | [optional] [default to null] |
| **name** | **String** | FriendlyName is the human-readable name of the container. | [optional] [default to null] |
| **ports** | [**List**](codersdk.WorkspaceAgentContainerPort.md) | Ports includes ports exposed by the container. | [optional] [default to null] |
| **running** | **Boolean** | Running is true if the container is currently running. | [optional] [default to null] |
| **status** | **String** | Status is the current status of the container. This is somewhat implementation-dependent, but should generally be a human-readable string. | [optional] [default to null] |
| **volumes** | **Map** | Volumes is a map of \&quot;things\&quot; mounted into the container. Again, this is somewhat implementation-dependent. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAgentContainerPort

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **host\_ip** | **String** | HostIP is the IP address of the host interface to which the port is bound. Note that this can be an IPv4 or IPv6 address. | [optional] [default to null] |
| **host\_port** | **Integer** | HostPort is the port number *outside* the container. | [optional] [default to null] |
| **network** | **String** | Network is the network protocol used by the port (tcp, udp, etc). | [optional] [default to null] |
| **port** | **Integer** | Port is the port number *inside* the container. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAgentDevcontainer

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **agent** | [**codersdk.WorkspaceAgentDevcontainerAgent**](codersdk.WorkspaceAgentDevcontainerAgent.md) |  | [optional] [default to null] |
| **config\_path** | **String** |  | [optional] [default to null] |
| **container** | [**codersdk.WorkspaceAgentContainer**](codersdk.WorkspaceAgentContainer.md) |  | [optional] [default to null] |
| **dirty** | **Boolean** |  | [optional] [default to null] |
| **error** | **String** |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **status** | [**codersdk.WorkspaceAgentDevcontainerStatus**](codersdk.WorkspaceAgentDevcontainerStatus.md) | Additional runtime fields. | [optional] [default to null] |
| **workspace\_folder** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAgentDevcontainerAgent

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **directory** | **String** |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAgentDevcontainerStatus

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAgentHealth

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **healthy** | **Boolean** | Healthy is true if the agent is healthy. | [optional] [default to null] |
| **reason** | **String** | Reason is a human-readable explanation of the agent&#39;s health. It is empty if Healthy is true. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAgentLifecycle

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAgentListContainersResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **containers** | [**List**](codersdk.WorkspaceAgentContainer.md) | Containers is a list of containers visible to the workspace agent. | [optional] [default to null] |
| **devcontainers** | [**List**](codersdk.WorkspaceAgentDevcontainer.md) | Devcontainers is a list of devcontainers visible to the workspace agent. | [optional] [default to null] |
| **warnings** | **List** | Warnings is a list of warnings that may have occurred during the process of listing containers. This should not include fatal errors. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAgentListeningPort

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **network** | **String** | only \&quot;tcp\&quot; at the moment | [optional] [default to null] |
| **port** | **Integer** |  | [optional] [default to null] |
| **process\_name** | **String** | may be empty | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAgentListeningPortsResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **ports** | [**List**](codersdk.WorkspaceAgentListeningPort.md) | If there are no ports in the list, nothing should be displayed in the UI. There must not be a \&quot;no ports available\&quot; message or anything similar, as there will always be no ports displayed on platforms where our port detection logic is unsupported. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAgentLog

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **created\_at** | **Date** |  | [optional] [default to null] |
| **id** | **Integer** |  | [optional] [default to null] |
| **level** | [**codersdk.LogLevel**](codersdk.LogLevel.md) |  | [optional] [default to null] |
| **output** | **String** |  | [optional] [default to null] |
| **source\_id** | **UUID** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAgentLogSource

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **created\_at** | **Date** |  | [optional] [default to null] |
| **display\_name** | **String** |  | [optional] [default to null] |
| **icon** | **String** |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **workspace\_agent\_id** | **UUID** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAgentPortShare

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **agent\_name** | **String** |  | [optional] [default to null] |
| **port** | **Integer** |  | [optional] [default to null] |
| **protocol** | [**codersdk.WorkspaceAgentPortShareProtocol**](codersdk.WorkspaceAgentPortShareProtocol.md) |  | [optional] [default to null] |
| **share\_level** | [**codersdk.WorkspaceAgentPortShareLevel**](codersdk.WorkspaceAgentPortShareLevel.md) |  | [optional] [default to null] |
| **workspace\_id** | **UUID** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAgentPortShareLevel

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAgentPortShareProtocol

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAgentPortShares

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **shares** | [**List**](codersdk.WorkspaceAgentPortShare.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAgentScript

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **cron** | **String** |  | [optional] [default to null] |
| **display\_name** | **String** |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **log\_path** | **String** |  | [optional] [default to null] |
| **log\_source\_id** | **UUID** |  | [optional] [default to null] |
| **run\_on\_start** | **Boolean** |  | [optional] [default to null] |
| **run\_on\_stop** | **Boolean** |  | [optional] [default to null] |
| **script** | **String** |  | [optional] [default to null] |
| **start\_blocks\_login** | **Boolean** |  | [optional] [default to null] |
| **timeout** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAgentStartupScriptBehavior

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAgentStatus

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceApp

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **command** | **String** |  | [optional] [default to null] |
| **display\_name** | **String** | DisplayName is a friendly name for the app. | [optional] [default to null] |
| **external** | **Boolean** | External specifies whether the URL should be opened externally on the client or not. | [optional] [default to null] |
| **group** | **String** |  | [optional] [default to null] |
| **health** | [**codersdk.WorkspaceAppHealth**](codersdk.WorkspaceAppHealth.md) |  | [optional] [default to null] |
| **healthcheck** | [**codersdk.Healthcheck**](codersdk.Healthcheck.md) | Healthcheck specifies the configuration for checking app health. | [optional] [default to null] |
| **hidden** | **Boolean** |  | [optional] [default to null] |
| **icon** | **String** | Icon is a relative path or external URL that specifies an icon to be displayed in the dashboard. | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **open\_in** | [**codersdk.WorkspaceAppOpenIn**](codersdk.WorkspaceAppOpenIn.md) |  | [optional] [default to null] |
| **sharing\_level** | [**codersdk.WorkspaceAppSharingLevel**](codersdk.WorkspaceAppSharingLevel.md) |  | [optional] [default to null] |
| **slug** | **String** | Slug is a unique identifier within the agent. | [optional] [default to null] |
| **statuses** | [**List**](codersdk.WorkspaceAppStatus.md) | Statuses is a list of statuses for the app. | [optional] [default to null] |
| **subdomain** | **Boolean** | Subdomain denotes whether the app should be accessed via a path on the &#x60;coder server&#x60; or via a hostname-based dev URL. If this is set to true and there is no app wildcard configured on the server, the app will not be accessible in the UI. | [optional] [default to null] |
| **subdomain\_name** | **String** | SubdomainName is the application domain exposed on the &#x60;coder server&#x60;. | [optional] [default to null] |
| **url** | **String** | URL is the address being proxied to inside the workspace. If external is specified, this will be opened on the client. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAppHealth

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAppOpenIn

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAppSharingLevel

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAppStatus

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **agent\_id** | **UUID** |  | [optional] [default to null] |
| **app\_id** | **UUID** |  | [optional] [default to null] |
| **created\_at** | **Date** |  | [optional] [default to null] |
| **icon** | **String** | Deprecated: This field is unused and will be removed in a future version. Icon is an external URL to an icon that will be rendered in the UI. | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **message** | **String** |  | [optional] [default to null] |
| **needs\_user\_attention** | **Boolean** | Deprecated: This field is unused and will be removed in a future version. NeedsUserAttention specifies whether the status needs user attention. | [optional] [default to null] |
| **state** | [**codersdk.WorkspaceAppStatusState**](codersdk.WorkspaceAppStatusState.md) |  | [optional] [default to null] |
| **uri** | **String** | URI is the URI of the resource that the status is for. e.g. https://github.com/org/repo/pull/123 e.g. file:///path/to/file | [optional] [default to null] |
| **workspace\_id** | **UUID** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceAppStatusState

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceBuild

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **ai\_task\_sidebar\_app\_id** | **UUID** |  | [optional] [default to null] |
| **build\_number** | **Integer** |  | [optional] [default to null] |
| **created\_at** | **Date** |  | [optional] [default to null] |
| **daily\_cost** | **Integer** |  | [optional] [default to null] |
| **deadline** | **Date** |  | [optional] [default to null] |
| **has\_ai\_task** | **Boolean** |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **initiator\_id** | **UUID** |  | [optional] [default to null] |
| **initiator\_name** | **String** |  | [optional] [default to null] |
| **job** | [**codersdk.ProvisionerJob**](codersdk.ProvisionerJob.md) |  | [optional] [default to null] |
| **matched\_provisioners** | [**codersdk.MatchedProvisioners**](codersdk.MatchedProvisioners.md) |  | [optional] [default to null] |
| **max\_deadline** | **Date** |  | [optional] [default to null] |
| **reason** | [**codersdk.BuildReason**](codersdk.BuildReason.md) |  | [optional] [default to null] |
| **resources** | [**List**](codersdk.WorkspaceResource.md) |  | [optional] [default to null] |
| **status** | [**codersdk.WorkspaceStatus**](codersdk.WorkspaceStatus.md) |  | [optional] [default to null] |
| **template\_version\_id** | **UUID** |  | [optional] [default to null] |
| **template\_version\_name** | **String** |  | [optional] [default to null] |
| **template\_version\_preset\_id** | **UUID** |  | [optional] [default to null] |
| **transition** | [**codersdk.WorkspaceTransition**](codersdk.WorkspaceTransition.md) |  | [optional] [default to null] |
| **updated\_at** | **Date** |  | [optional] [default to null] |
| **workspace\_id** | **UUID** |  | [optional] [default to null] |
| **workspace\_name** | **String** |  | [optional] [default to null] |
| **workspace\_owner\_avatar\_url** | **String** |  | [optional] [default to null] |
| **workspace\_owner\_id** | **UUID** |  | [optional] [default to null] |
| **workspace\_owner\_name** | **String** | WorkspaceOwnerName is the username of the owner of the workspace. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceBuildParameter

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **name** | **String** |  | [optional] [default to null] |
| **value** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceBuildTimings

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **agent\_connection\_timings** | [**List**](codersdk.AgentConnectionTiming.md) |  | [optional] [default to null] |
| **agent\_script\_timings** | [**List**](codersdk.AgentScriptTiming.md) | TODO: Consolidate agent-related timing metrics into a single struct when updating the API version | [optional] [default to null] |
| **provisioner\_timings** | [**List**](codersdk.ProvisionerTiming.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceConnectionLatencyMS

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **p50** | **BigDecimal** |  | [optional] [default to null] |
| **p95** | **BigDecimal** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceDeploymentStats

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **building** | **Integer** |  | [optional] [default to null] |
| **connection\_latency\_ms** | [**codersdk.WorkspaceConnectionLatencyMS**](codersdk.WorkspaceConnectionLatencyMS.md) |  | [optional] [default to null] |
| **failed** | **Integer** |  | [optional] [default to null] |
| **pending** | **Integer** |  | [optional] [default to null] |
| **running** | **Integer** |  | [optional] [default to null] |
| **rx\_bytes** | **Integer** |  | [optional] [default to null] |
| **stopped** | **Integer** |  | [optional] [default to null] |
| **tx\_bytes** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceHealth

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **failing\_agents** | **List** | FailingAgents lists the IDs of the agents that are failing, if any. | [optional] [default to null] |
| **healthy** | **Boolean** | Healthy is true if the workspace is healthy. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceProxy

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **created\_at** | **Date** |  | [optional] [default to null] |
| **deleted** | **Boolean** |  | [optional] [default to null] |
| **derp\_enabled** | **Boolean** |  | [optional] [default to null] |
| **derp\_only** | **Boolean** |  | [optional] [default to null] |
| **display\_name** | **String** |  | [optional] [default to null] |
| **healthy** | **Boolean** |  | [optional] [default to null] |
| **icon\_url** | **String** |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **path\_app\_url** | **String** | PathAppURL is the URL to the base path for path apps. Optional unless wildcard_hostname is set. E.g. https://us.example.com | [optional] [default to null] |
| **status** | [**codersdk.WorkspaceProxyStatus**](codersdk.WorkspaceProxyStatus.md) | Status is the latest status check of the proxy. This will be empty for deleted proxies. This value can be used to determine if a workspace proxy is healthy and ready to use. | [optional] [default to null] |
| **updated\_at** | **Date** |  | [optional] [default to null] |
| **version** | **String** |  | [optional] [default to null] |
| **wildcard\_hostname** | **String** | WildcardHostname is the wildcard hostname for subdomain apps. E.g. *.us.example.com E.g. *--suffix.au.example.com Optional. Does not need to be on the same domain as PathAppURL. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceProxyStatus

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **checked\_at** | **Date** |  | [optional] [default to null] |
| **report** | [**codersdk.ProxyHealthReport**](codersdk.ProxyHealthReport.md) | Report provides more information about the health of the workspace proxy. | [optional] [default to null] |
| **status** | [**codersdk.ProxyHealthStatus**](codersdk.ProxyHealthStatus.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceQuota

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **budget** | **Integer** |  | [optional] [default to null] |
| **credits\_consumed** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceResource

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **agents** | [**List**](codersdk.WorkspaceAgent.md) |  | [optional] [default to null] |
| **created\_at** | **Date** |  | [optional] [default to null] |
| **daily\_cost** | **Integer** |  | [optional] [default to null] |
| **hide** | **Boolean** |  | [optional] [default to null] |
| **icon** | **String** |  | [optional] [default to null] |
| **id** | **UUID** |  | [optional] [default to null] |
| **job\_id** | **UUID** |  | [optional] [default to null] |
| **metadata** | [**List**](codersdk.WorkspaceResourceMetadata.md) |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **type** | **String** |  | [optional] [default to null] |
| **workspace\_transition** | [**codersdk.WorkspaceTransition**](codersdk.WorkspaceTransition.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceResourceMetadata

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **key** | **String** |  | [optional] [default to null] |
| **sensitive** | **Boolean** |  | [optional] [default to null] |
| **value** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceRole

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceStatus

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspaceTransition

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk.WorkspacesResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **count** | **Integer** |  | [optional] [default to null] |
| **workspaces** | [**List**](codersdk.Workspace.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### codersdk_PatchGroupIDPSyncMappingRequest_add_inner

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **gets** | **String** | The ID of the Coder resource the user should be added to | [optional] [default to null] |
| **given** | **String** | The IdP claim the user has | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### derp.BytesSentRecv

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **key** | [**Object**](.md) | Key is the public key of the client which sent/received these bytes. | [optional] [default to null] |
| **recv** | **Integer** |  | [optional] [default to null] |
| **sent** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### derp.ServerInfoMessage

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **tokenBucketBytesBurst** | **Integer** | TokenBucketBytesBurst is how many bytes the server will allow to burst, temporarily violating TokenBucketBytesPerSecond.  Zero means unspecified. There might be a limit, but the client need not try to respect it. | [optional] [default to null] |
| **tokenBucketBytesPerSecond** | **Integer** | TokenBucketBytesPerSecond is how many bytes per second the server says it will accept, including all framing bytes.  Zero means unspecified. There might be a limit, but the client need not try to respect it. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### health.Code

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### health.Message

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **code** | [**health.Code**](health.Code.md) |  | [optional] [default to null] |
| **message** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### health.Severity

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### healthsdk.AccessURLReport

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **access\_url** | **String** |  | [optional] [default to null] |
| **dismissed** | **Boolean** |  | [optional] [default to null] |
| **error** | **String** |  | [optional] [default to null] |
| **healthy** | **Boolean** | Healthy is deprecated and left for backward compatibility purposes, use &#x60;Severity&#x60; instead. | [optional] [default to null] |
| **healthz\_response** | **String** |  | [optional] [default to null] |
| **reachable** | **Boolean** |  | [optional] [default to null] |
| **severity** | [**health.Severity**](health.Severity.md) |  | [optional] [default to null] |
| **status\_code** | **Integer** |  | [optional] [default to null] |
| **warnings** | [**List**](health.Message.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### healthsdk.DERPHealthReport

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **dismissed** | **Boolean** |  | [optional] [default to null] |
| **error** | **String** |  | [optional] [default to null] |
| **healthy** | **Boolean** | Healthy is deprecated and left for backward compatibility purposes, use &#x60;Severity&#x60; instead. | [optional] [default to null] |
| **netcheck** | [**netcheck.Report**](netcheck.Report.md) |  | [optional] [default to null] |
| **netcheck\_err** | **String** |  | [optional] [default to null] |
| **netcheck\_logs** | **List** |  | [optional] [default to null] |
| **regions** | [**Map**](healthsdk.DERPRegionReport.md) |  | [optional] [default to null] |
| **severity** | [**health.Severity**](health.Severity.md) |  | [optional] [default to null] |
| **warnings** | [**List**](health.Message.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### healthsdk.DERPNodeReport

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **can\_exchange\_messages** | **Boolean** |  | [optional] [default to null] |
| **client\_errs** | [**List**](array.md) |  | [optional] [default to null] |
| **client\_logs** | [**List**](array.md) |  | [optional] [default to null] |
| **error** | **String** |  | [optional] [default to null] |
| **healthy** | **Boolean** | Healthy is deprecated and left for backward compatibility purposes, use &#x60;Severity&#x60; instead. | [optional] [default to null] |
| **node** | [**tailcfg.DERPNode**](tailcfg.DERPNode.md) |  | [optional] [default to null] |
| **node\_info** | [**derp.ServerInfoMessage**](derp.ServerInfoMessage.md) |  | [optional] [default to null] |
| **round\_trip\_ping** | **String** |  | [optional] [default to null] |
| **round\_trip\_ping\_ms** | **Integer** |  | [optional] [default to null] |
| **severity** | [**health.Severity**](health.Severity.md) |  | [optional] [default to null] |
| **stun** | [**healthsdk.STUNReport**](healthsdk.STUNReport.md) |  | [optional] [default to null] |
| **uses\_websocket** | **Boolean** |  | [optional] [default to null] |
| **warnings** | [**List**](health.Message.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### healthsdk.DERPRegionReport

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **error** | **String** |  | [optional] [default to null] |
| **healthy** | **Boolean** | Healthy is deprecated and left for backward compatibility purposes, use &#x60;Severity&#x60; instead. | [optional] [default to null] |
| **node\_reports** | [**List**](healthsdk.DERPNodeReport.md) |  | [optional] [default to null] |
| **region** | [**tailcfg.DERPRegion**](tailcfg.DERPRegion.md) |  | [optional] [default to null] |
| **severity** | [**health.Severity**](health.Severity.md) |  | [optional] [default to null] |
| **warnings** | [**List**](health.Message.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### healthsdk.DatabaseReport

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **dismissed** | **Boolean** |  | [optional] [default to null] |
| **error** | **String** |  | [optional] [default to null] |
| **healthy** | **Boolean** | Healthy is deprecated and left for backward compatibility purposes, use &#x60;Severity&#x60; instead. | [optional] [default to null] |
| **latency** | **String** |  | [optional] [default to null] |
| **latency\_ms** | **Integer** |  | [optional] [default to null] |
| **reachable** | **Boolean** |  | [optional] [default to null] |
| **severity** | [**health.Severity**](health.Severity.md) |  | [optional] [default to null] |
| **threshold\_ms** | **Integer** |  | [optional] [default to null] |
| **warnings** | [**List**](health.Message.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### healthsdk.HealthSection

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### healthsdk.HealthSettings

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **dismissed\_healthchecks** | [**List**](healthsdk.HealthSection.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### healthsdk.HealthcheckReport

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **access\_url** | [**healthsdk.AccessURLReport**](healthsdk.AccessURLReport.md) |  | [optional] [default to null] |
| **coder\_version** | **String** | The Coder version of the server that the report was generated on. | [optional] [default to null] |
| **database** | [**healthsdk.DatabaseReport**](healthsdk.DatabaseReport.md) |  | [optional] [default to null] |
| **derp** | [**healthsdk.DERPHealthReport**](healthsdk.DERPHealthReport.md) |  | [optional] [default to null] |
| **healthy** | **Boolean** | Healthy is true if the report returns no errors. Deprecated: use &#x60;Severity&#x60; instead | [optional] [default to null] |
| **provisioner\_daemons** | [**healthsdk.ProvisionerDaemonsReport**](healthsdk.ProvisionerDaemonsReport.md) |  | [optional] [default to null] |
| **severity** | [**health.Severity**](health.Severity.md) | Severity indicates the status of Coder health. | [optional] [default to null] |
| **time** | **Date** | Time is the time the report was generated at. | [optional] [default to null] |
| **websocket** | [**healthsdk.WebsocketReport**](healthsdk.WebsocketReport.md) |  | [optional] [default to null] |
| **workspace\_proxy** | [**healthsdk.WorkspaceProxyReport**](healthsdk.WorkspaceProxyReport.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### healthsdk.ProvisionerDaemonsReport

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **dismissed** | **Boolean** |  | [optional] [default to null] |
| **error** | **String** |  | [optional] [default to null] |
| **items** | [**List**](healthsdk.ProvisionerDaemonsReportItem.md) |  | [optional] [default to null] |
| **severity** | [**health.Severity**](health.Severity.md) |  | [optional] [default to null] |
| **warnings** | [**List**](health.Message.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### healthsdk.ProvisionerDaemonsReportItem

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **provisioner\_daemon** | [**codersdk.ProvisionerDaemon**](codersdk.ProvisionerDaemon.md) |  | [optional] [default to null] |
| **warnings** | [**List**](health.Message.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### healthsdk.STUNReport

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **canSTUN** | **Boolean** |  | [optional] [default to null] |
| **enabled** | **Boolean** |  | [optional] [default to null] |
| **error** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### healthsdk.UpdateHealthSettings

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **dismissed\_healthchecks** | [**List**](healthsdk.HealthSection.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### healthsdk.WebsocketReport

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **body** | **String** |  | [optional] [default to null] |
| **code** | **Integer** |  | [optional] [default to null] |
| **dismissed** | **Boolean** |  | [optional] [default to null] |
| **error** | **String** |  | [optional] [default to null] |
| **healthy** | **Boolean** | Healthy is deprecated and left for backward compatibility purposes, use &#x60;Severity&#x60; instead. | [optional] [default to null] |
| **severity** | [**health.Severity**](health.Severity.md) |  | [optional] [default to null] |
| **warnings** | [**List**](health.Message.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### healthsdk.WorkspaceProxyReport

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **dismissed** | **Boolean** |  | [optional] [default to null] |
| **error** | **String** |  | [optional] [default to null] |
| **healthy** | **Boolean** | Healthy is deprecated and left for backward compatibility purposes, use &#x60;Severity&#x60; instead. | [optional] [default to null] |
| **severity** | [**health.Severity**](health.Severity.md) |  | [optional] [default to null] |
| **warnings** | [**List**](health.Message.md) |  | [optional] [default to null] |
| **workspace\_proxies** | [**codersdk.RegionsResponse-codersdk_WorkspaceProxy**](codersdk.RegionsResponse-codersdk_WorkspaceProxy.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### netcheck.Report

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **captivePortal** | **String** | CaptivePortal is set when we think there&#39;s a captive portal that is intercepting HTTP traffic. | [optional] [default to null] |
| **globalV4** | **String** | ip:port of global IPv4 | [optional] [default to null] |
| **globalV6** | **String** | [ip]:port of global IPv6 | [optional] [default to null] |
| **hairPinning** | **String** | HairPinning is whether the router supports communicating between two local devices through the NATted public IP address (on IPv4). | [optional] [default to null] |
| **icmpv4** | **Boolean** | an ICMPv4 round trip completed | [optional] [default to null] |
| **ipv4** | **Boolean** | an IPv4 STUN round trip completed | [optional] [default to null] |
| **ipv4CanSend** | **Boolean** | an IPv4 packet was able to be sent | [optional] [default to null] |
| **ipv6** | **Boolean** | an IPv6 STUN round trip completed | [optional] [default to null] |
| **ipv6CanSend** | **Boolean** | an IPv6 packet was able to be sent | [optional] [default to null] |
| **mappingVariesByDestIP** | **String** | MappingVariesByDestIP is whether STUN results depend which STUN server you&#39;re talking to (on IPv4). | [optional] [default to null] |
| **oshasIPv6** | **Boolean** | could bind a socket to ::1 | [optional] [default to null] |
| **pcp** | **String** | PCP is whether PCP appears present on the LAN. Empty means not checked. | [optional] [default to null] |
| **pmp** | **String** | PMP is whether NAT-PMP appears present on the LAN. Empty means not checked. | [optional] [default to null] |
| **preferredDERP** | **Integer** | or 0 for unknown | [optional] [default to null] |
| **regionLatency** | **Map** | keyed by DERP Region ID | [optional] [default to null] |
| **regionV4Latency** | **Map** | keyed by DERP Region ID | [optional] [default to null] |
| **regionV6Latency** | **Map** | keyed by DERP Region ID | [optional] [default to null] |
| **udp** | **Boolean** | a UDP STUN round trip completed | [optional] [default to null] |
| **upnP** | **String** | UPnP is whether UPnP appears present on the LAN. Empty means not checked. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### oauth2.Token

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **access\_token** | **String** | AccessToken is the token that authorizes and authenticates the requests. | [optional] [default to null] |
| **expires\_in** | **Integer** | ExpiresIn is the OAuth2 wire format \&quot;expires_in\&quot; field, which specifies how many seconds later the token expires, relative to an unknown time base approximately around \&quot;now\&quot;. It is the application&#39;s responsibility to populate &#x60;Expiry&#x60; from &#x60;ExpiresIn&#x60; when required. | [optional] [default to null] |
| **expiry** | **String** | Expiry is the optional expiration time of the access token.  If zero, [TokenSource] implementations will reuse the same token forever and RefreshToken or equivalent mechanisms for that TokenSource will not be used. | [optional] [default to null] |
| **refresh\_token** | **String** | RefreshToken is a token that&#39;s used by the application (as opposed to the user) to refresh the access token if it expires. | [optional] [default to null] |
| **token\_type** | **String** | TokenType is the type of token. The Type method returns either this or \&quot;Bearer\&quot;, the default. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### serpent.Group

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **description** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **parent** | [**serpent.Group**](serpent.Group.md) |  | [optional] [default to null] |
| **yaml** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### serpent.HostPort

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **host** | **String** |  | [optional] [default to null] |
| **port** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### serpent.Option

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **annotations** | **Map** | Annotations enable extensions to serpent higher up in the stack. It&#39;s useful for help formatting and documentation generation. | [optional] [default to null] |
| **default** | **String** | Default is parsed into Value if set. | [optional] [default to null] |
| **description** | **String** |  | [optional] [default to null] |
| **env** | **String** | Env is the environment variable used to configure this option. If unset, environment configuring is disabled. | [optional] [default to null] |
| **flag** | **String** | Flag is the long name of the flag used to configure this option. If unset, flag configuring is disabled. | [optional] [default to null] |
| **flag\_shorthand** | **String** | FlagShorthand is the one-character shorthand for the flag. If unset, no shorthand is used. | [optional] [default to null] |
| **group** | [**serpent.Group**](serpent.Group.md) | Group is a group hierarchy that helps organize this option in help, configs and other documentation. | [optional] [default to null] |
| **hidden** | **Boolean** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **required** | **Boolean** | Required means this value must be set by some means. It requires &#x60;ValueSource !&#x3D; ValueSourceNone&#x60; If &#x60;Default&#x60; is set, then &#x60;Required&#x60; is ignored. | [optional] [default to null] |
| **use\_instead** | [**List**](serpent.Option.md) | UseInstead is a list of options that should be used instead of this one. The field is used to generate a deprecation warning. | [optional] [default to null] |
| **value** | [**Object**](.md) | Value includes the types listed in values.go. | [optional] [default to null] |
| **value\_source** | [**serpent.ValueSource**](serpent.ValueSource.md) |  | [optional] [default to null] |
| **yaml** | **String** | YAML is the YAML key used to configure this option. If unset, YAML configuring is disabled. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### serpent.Struct-array_codersdk_ExternalAuthConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **value** | [**List**](codersdk.ExternalAuthConfig.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### serpent.Struct-array_codersdk_LinkConfig

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **value** | [**List**](codersdk.LinkConfig.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### serpent.URL

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **forceQuery** | **Boolean** | append a query (&#39;?&#39;) even if RawQuery is empty | [optional] [default to null] |
| **fragment** | **String** | fragment for references, without &#39;#&#39; | [optional] [default to null] |
| **host** | **String** | host or host:port (see Hostname and Port methods) | [optional] [default to null] |
| **omitHost** | **Boolean** | do not emit empty host (authority) | [optional] [default to null] |
| **opaque** | **String** | encoded opaque data | [optional] [default to null] |
| **path** | **String** | path (relative paths may omit leading slash) | [optional] [default to null] |
| **rawFragment** | **String** | encoded fragment hint (see EscapedFragment method) | [optional] [default to null] |
| **rawPath** | **String** | encoded path hint (see EscapedPath method) | [optional] [default to null] |
| **rawQuery** | **String** | encoded query values, without &#39;?&#39; | [optional] [default to null] |
| **scheme** | **String** |  | [optional] [default to null] |
| **user** | [**Object**](.md) | username and password information | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### serpent.ValueSource

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### tailcfg.DERPHomeParams

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **regionScore** | **Map** | RegionScore scales latencies of DERP regions by a given scaling factor when determining which region to use as the home (\&quot;preferred\&quot;) DERP. Scores in the range (0, 1) will cause this region to be proportionally more preferred, and scores in the range (1, ∞) will penalize a region.  If a region is not present in this map, it is treated as having a score of 1.0.  Scores should not be 0 or negative; such scores will be ignored.  A nil map means no change from the previous value (if any); an empty non-nil map can be sent to reset all scores back to 1.0. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### tailcfg.DERPMap

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **homeParams** | [**tailcfg.DERPHomeParams**](tailcfg.DERPHomeParams.md) | HomeParams, if non-nil, is a change in home parameters.  The rest of the DEPRMap fields, if zero, means unchanged. | [optional] [default to null] |
| **omitDefaultRegions** | **Boolean** | OmitDefaultRegions specifies to not use Tailscale&#39;s DERP servers, and only use those specified in this DERPMap. If there are none set outside of the defaults, this is a noop.  This field is only meaningful if the Regions map is non-nil (indicating a change). | [optional] [default to null] |
| **regions** | [**Map**](tailcfg.DERPRegion.md) | Regions is the set of geographic regions running DERP node(s).  It&#39;s keyed by the DERPRegion.RegionID.  The numbers are not necessarily contiguous. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### tailcfg.DERPNode

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **canPort80** | **Boolean** | CanPort80 specifies whether this DERP node is accessible over HTTP on port 80 specifically. This is used for captive portal checks. | [optional] [default to null] |
| **certName** | **String** | CertName optionally specifies the expected TLS cert common name. If empty, HostName is used. If CertName is non-empty, HostName is only used for the TCP dial (if IPv4/IPv6 are not present) + TLS ClientHello. | [optional] [default to null] |
| **derpport** | **Integer** | DERPPort optionally provides an alternate TLS port number for the DERP HTTPS server.  If zero, 443 is used. | [optional] [default to null] |
| **forceHTTP** | **Boolean** | ForceHTTP is used by unit tests to force HTTP. It should not be set by users. | [optional] [default to null] |
| **hostName** | **String** | HostName is the DERP node&#39;s hostname.  It is required but need not be unique; multiple nodes may have the same HostName but vary in configuration otherwise. | [optional] [default to null] |
| **insecureForTests** | **Boolean** | InsecureForTests is used by unit tests to disable TLS verification. It should not be set by users. | [optional] [default to null] |
| **ipv4** | **String** | IPv4 optionally forces an IPv4 address to use, instead of using DNS. If empty, A record(s) from DNS lookups of HostName are used. If the string is not an IPv4 address, IPv4 is not used; the conventional string to disable IPv4 (and not use DNS) is \&quot;none\&quot;. | [optional] [default to null] |
| **ipv6** | **String** | IPv6 optionally forces an IPv6 address to use, instead of using DNS. If empty, AAAA record(s) from DNS lookups of HostName are used. If the string is not an IPv6 address, IPv6 is not used; the conventional string to disable IPv6 (and not use DNS) is \&quot;none\&quot;. | [optional] [default to null] |
| **name** | **String** | Name is a unique node name (across all regions). It is not a host name. It&#39;s typically of the form \&quot;1b\&quot;, \&quot;2a\&quot;, \&quot;3b\&quot;, etc. (region ID + suffix within that region) | [optional] [default to null] |
| **regionID** | **Integer** | RegionID is the RegionID of the DERPRegion that this node is running in. | [optional] [default to null] |
| **stunonly** | **Boolean** | STUNOnly marks a node as only a STUN server and not a DERP server. | [optional] [default to null] |
| **stunport** | **Integer** | Port optionally specifies a STUN port to use. Zero means 3478. To disable STUN on this node, use -1. | [optional] [default to null] |
| **stuntestIP** | **String** | STUNTestIP is used in tests to override the STUN server&#39;s IP. If empty, it&#39;s assumed to be the same as the DERP server. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### tailcfg.DERPRegion

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **avoid** | **Boolean** | Avoid is whether the client should avoid picking this as its home region. The region should only be used if a peer is there. Clients already using this region as their home should migrate away to a new region without Avoid set. | [optional] [default to null] |
| **embeddedRelay** | **Boolean** | EmbeddedRelay is true when the region is bundled with the Coder control plane. | [optional] [default to null] |
| **nodes** | [**List**](tailcfg.DERPNode.md) | Nodes are the DERP nodes running in this region, in priority order for the current client. Client TLS connections should ideally only go to the first entry (falling back to the second if necessary). STUN packets should go to the first 1 or 2.  If nodes within a region route packets amongst themselves, but not to other regions. That said, each user/domain should get a the same preferred node order, so if all nodes for a user/network pick the first one (as they should, when things are healthy), the inter-cluster routing is minimal to zero. | [optional] [default to null] |
| **regionCode** | **String** | RegionCode is a short name for the region. It&#39;s usually a popular city or airport code in the region: \&quot;nyc\&quot;, \&quot;sf\&quot;, \&quot;sin\&quot;, \&quot;fra\&quot;, etc. | [optional] [default to null] |
| **regionID** | **Integer** | RegionID is a unique integer for a geographic region.  It corresponds to the legacy derpN.tailscale.com hostnames used by older clients. (Older clients will continue to resolve derpN.tailscale.com when contacting peers, rather than use the server-provided DERPMap)  RegionIDs must be non-zero, positive, and guaranteed to fit in a JavaScript number.  RegionIDs in range 900-999 are reserved for end users to run their own DERP nodes. | [optional] [default to null] |
| **regionName** | **String** | RegionName is a long English name for the region: \&quot;New York City\&quot;, \&quot;San Francisco\&quot;, \&quot;Singapore\&quot;, \&quot;Frankfurt\&quot;, etc. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### upload_file_request

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **file** | **File** | File to be uploaded. If using tar format, file must conform to ustar (pax may cause problems). | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### uuid.NullUUID

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **uuid** | **String** |  | [optional] [default to null] |
| **valid** | **Boolean** | Valid is true if UUID is not NULL | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### workspaceapps.AccessMethod

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### workspaceapps.IssueTokenRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **app\_hostname** | **String** | AppHostname is the optional hostname for subdomain apps on the external proxy. It must start with an asterisk. | [optional] [default to null] |
| **app\_path** | **String** | AppPath is the path of the user underneath the app base path. | [optional] [default to null] |
| **app\_query** | **String** | AppQuery is the query parameters the user provided in the app request. | [optional] [default to null] |
| **app\_request** | [**workspaceapps.Request**](workspaceapps.Request.md) |  | [optional] [default to null] |
| **path\_app\_base\_url** | **String** | PathAppBaseURL is required. | [optional] [default to null] |
| **session\_token** | **String** | SessionToken is the session token provided by the user. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### workspaceapps.Request

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **access\_method** | [**workspaceapps.AccessMethod**](workspaceapps.AccessMethod.md) |  | [optional] [default to null] |
| **agent\_name\_or\_id** | **String** | AgentNameOrID is not required if the workspace has only one agent. | [optional] [default to null] |
| **app\_prefix** | **String** | Prefix is the prefix of the subdomain app URL. Prefix should have a trailing \&quot;---\&quot; if set. | [optional] [default to null] |
| **app\_slug\_or\_port** | **String** |  | [optional] [default to null] |
| **base\_path** | **String** | BasePath of the app. For path apps, this is the path prefix in the router for this particular app. For subdomain apps, this should be \&quot;/\&quot;. This is used for setting the cookie path. | [optional] [default to null] |
| **username\_or\_id** | **String** | For the following fields, if the AccessMethod is AccessMethodTerminal, then only AgentNameOrID may be set and it must be a UUID. The other fields must be left blank. | [optional] [default to null] |
| **workspace\_name\_or\_id** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### workspaceapps.StatsReport

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **access\_method** | [**workspaceapps.AccessMethod**](workspaceapps.AccessMethod.md) |  | [optional] [default to null] |
| **agent\_id** | **String** |  | [optional] [default to null] |
| **requests** | **Integer** |  | [optional] [default to null] |
| **session\_ended\_at** | **String** | Updated periodically while app is in use active and when the last connection is closed. | [optional] [default to null] |
| **session\_id** | **String** |  | [optional] [default to null] |
| **session\_started\_at** | **String** |  | [optional] [default to null] |
| **slug\_or\_port** | **String** |  | [optional] [default to null] |
| **user\_id** | **String** |  | [optional] [default to null] |
| **workspace\_id** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### workspacesdk.AgentConnectionInfo

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **derp\_force\_websockets** | **Boolean** |  | [optional] [default to null] |
| **derp\_map** | [**tailcfg.DERPMap**](tailcfg.DERPMap.md) |  | [optional] [default to null] |
| **disable\_direct\_connections** | **Boolean** |  | [optional] [default to null] |
| **hostname\_suffix** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### wsproxysdk.CryptoKeysResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **crypto\_keys** | [**List**](codersdk.CryptoKey.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### wsproxysdk.DeregisterWorkspaceProxyRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **replica\_id** | **String** | ReplicaID is a unique identifier for the replica of the proxy that is deregistering. It should be generated by the client on startup and should&#39;ve already been passed to the register endpoint. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### wsproxysdk.IssueSignedAppTokenResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **signed\_token\_str** | **String** | SignedTokenStr should be set as a cookie on the response. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### wsproxysdk.RegisterWorkspaceProxyRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **access\_url** | **String** | AccessURL that hits the workspace proxy api. | [optional] [default to null] |
| **derp\_enabled** | **Boolean** | DerpEnabled indicates whether the proxy should be included in the DERP map or not. | [optional] [default to null] |
| **derp\_only** | **Boolean** | DerpOnly indicates whether the proxy should only be included in the DERP map and should not be used for serving apps. | [optional] [default to null] |
| **hostname** | **String** | ReplicaHostname is the OS hostname of the machine that the proxy is running on.  This is only used for tracking purposes in the replicas table. | [optional] [default to null] |
| **replica\_error** | **String** | ReplicaError is the error that the replica encountered when trying to dial it&#39;s peers. This is stored in the replicas table for debugging purposes but does not affect the proxy&#39;s ability to register.  This value is only stored on subsequent requests to the register endpoint, not the first request. | [optional] [default to null] |
| **replica\_id** | **String** | ReplicaID is a unique identifier for the replica of the proxy that is registering. It should be generated by the client on startup and persisted (in memory only) until the process is restarted. | [optional] [default to null] |
| **replica\_relay\_address** | **String** | ReplicaRelayAddress is the DERP address of the replica that other replicas may use to connect internally for DERP meshing. | [optional] [default to null] |
| **version** | **String** | Version is the Coder version of the proxy. | [optional] [default to null] |
| **wildcard\_hostname** | **String** | WildcardHostname that the workspace proxy api is serving for subdomain apps. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### wsproxysdk.RegisterWorkspaceProxyResponse

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **derp\_force\_websockets** | **Boolean** |  | [optional] [default to null] |
| **derp\_map** | [**tailcfg.DERPMap**](tailcfg.DERPMap.md) |  | [optional] [default to null] |
| **derp\_mesh\_key** | **String** |  | [optional] [default to null] |
| **derp\_region\_id** | **Integer** |  | [optional] [default to null] |
| **sibling\_replicas** | [**List**](codersdk.Replica.md) | SiblingReplicas is a list of all other replicas of the proxy that have not timed out. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


### wsproxysdk.ReportAppStatsRequest

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **stats** | [**List**](workspaceapps.StatsReport.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


