# Agents


All URIs are relative to */api/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**authenticateAgentOnAwsInstance**](AgentsApi.md#authenticateAgentOnAwsInstance) | **POST** /workspaceagents/aws-instance-identity | Authenticate agent on AWS instance |
| [**authenticateAgentOnAzureInstance**](AgentsApi.md#authenticateAgentOnAzureInstance) | **POST** /workspaceagents/azure-instance-identity | Authenticate agent on Azure instance |
| [**authenticateAgentOnGoogleCloudInstance**](AgentsApi.md#authenticateAgentOnGoogleCloudInstance) | **POST** /workspaceagents/google-instance-identity | Authenticate agent on Google Cloud instance |
| [**coordinateWorkspaceAgent**](AgentsApi.md#coordinateWorkspaceAgent) | **GET** /workspaceagents/{workspaceagent}/coordinate | Coordinate workspace agent |
| [**debugOidcContextForAUser**](AgentsApi.md#debugOidcContextForAUser) | **GET** /debug/{user}/debug-link | Debug OIDC context for a user |
| [**getConnectionInfoForWorkspaceAgent**](AgentsApi.md#getConnectionInfoForWorkspaceAgent) | **GET** /workspaceagents/{workspaceagent}/connection | Get connection info for workspace agent |
| [**getConnectionInfoForWorkspaceAgentGeneric**](AgentsApi.md#getConnectionInfoForWorkspaceAgentGeneric) | **GET** /workspaceagents/connection | Get connection info for workspace agent generic |
| [**getDerpMapUpdates**](AgentsApi.md#getDerpMapUpdates) | **GET** /derp-map | Get DERP map updates |
| [**getListeningPortsForWorkspaceAgent**](AgentsApi.md#getListeningPortsForWorkspaceAgent) | **GET** /workspaceagents/{workspaceagent}/listening-ports | Get listening ports for workspace agent |
| [**getLogsByWorkspaceAgent**](AgentsApi.md#getLogsByWorkspaceAgent) | **GET** /workspaceagents/{workspaceagent}/logs | Get logs by workspace agent |
| [**getRunningContainersForWorkspaceAgent**](AgentsApi.md#getRunningContainersForWorkspaceAgent) | **GET** /workspaceagents/{workspaceagent}/containers | Get running containers for workspace agent |
| [**getWorkspaceAgentById**](AgentsApi.md#getWorkspaceAgentById) | **GET** /workspaceagents/{workspaceagent} | Get workspace agent by ID |
| [**getWorkspaceAgentExternalAuth**](AgentsApi.md#getWorkspaceAgentExternalAuth) | **GET** /workspaceagents/me/external-auth | Get workspace agent external auth |
| [**getWorkspaceAgentGitSshKey**](AgentsApi.md#getWorkspaceAgentGitSshKey) | **GET** /workspaceagents/me/gitsshkey | Get workspace agent Git SSH key |
| [**getWorkspaceAgentReinitialization**](AgentsApi.md#getWorkspaceAgentReinitialization) | **GET** /workspaceagents/me/reinit | Get workspace agent reinitialization |
| [**openPtyToWorkspaceAgent**](AgentsApi.md#openPtyToWorkspaceAgent) | **GET** /workspaceagents/{workspaceagent}/pty | Open PTY to workspace agent |
| [**patchWorkspaceAgentAppStatus**](AgentsApi.md#patchWorkspaceAgentAppStatus) | **PATCH** /workspaceagents/me/app-status | Patch workspace agent app status |
| [**patchWorkspaceAgentLogs**](AgentsApi.md#patchWorkspaceAgentLogs) | **PATCH** /workspaceagents/me/logs | Patch workspace agent logs |
| [**postWorkspaceAgentLogSource**](AgentsApi.md#postWorkspaceAgentLogSource) | **POST** /workspaceagents/me/log-source | Post workspace agent log source |
| [**recreateDevcontainerForWorkspaceAgent**](AgentsApi.md#recreateDevcontainerForWorkspaceAgent) | **POST** /workspaceagents/{workspaceagent}/containers/devcontainers/{devcontainer}/recreate | Recreate devcontainer for workspace agent |
| [**removedGetLogsByWorkspaceAgent**](AgentsApi.md#removedGetLogsByWorkspaceAgent) | **GET** /workspaceagents/{workspaceagent}/startup-logs | Removed: Get logs by workspace agent |
| [**removedGetWorkspaceAgentGitAuth**](AgentsApi.md#removedGetWorkspaceAgentGitAuth) | **GET** /workspaceagents/me/gitauth | Removed: Get workspace agent git auth |
| [**userScopedTailnetRpcConnection**](AgentsApi.md#userScopedTailnetRpcConnection) | **GET** /tailnet | User-scoped tailnet RPC connection |
| [**watchForWorkspaceAgentMetadataUpdates**](AgentsApi.md#watchForWorkspaceAgentMetadataUpdates) | **GET** /workspaceagents/{workspaceagent}/watch-metadata | Watch for workspace agent metadata updates |
| [**watchForWorkspaceAgentMetadataUpdatesViaWebsockets**](AgentsApi.md#watchForWorkspaceAgentMetadataUpdatesViaWebsockets) | **GET** /workspaceagents/{workspaceagent}/watch-metadata-ws | Watch for workspace agent metadata updates via WebSockets |
| [**watchWorkspaceAgentForContainerUpdates**](AgentsApi.md#watchWorkspaceAgentForContainerUpdates) | **GET** /workspaceagents/{workspaceagent}/containers/watch | Watch workspace agent for container updates. |
| [**workspaceAgentRpcApi**](AgentsApi.md#workspaceAgentRpcApi) | **GET** /workspaceagents/me/rpc | Workspace agent RPC API |


<a name="authenticateAgentOnAwsInstance"></a>
# **authenticateAgentOnAwsInstance**
> agentsdk.AuthenticateResponse authenticateAgentOnAwsInstance(request)

Authenticate agent on AWS instance

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**agentsdk.AWSInstanceIdentityToken**](../Models/agentsdk.AWSInstanceIdentityToken.md)| Instance identity token | |

### Return type

[**agentsdk.AuthenticateResponse**](../Models/agentsdk.AuthenticateResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="authenticateAgentOnAzureInstance"></a>
# **authenticateAgentOnAzureInstance**
> agentsdk.AuthenticateResponse authenticateAgentOnAzureInstance(request)

Authenticate agent on Azure instance

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**agentsdk.AzureInstanceIdentityToken**](../Models/agentsdk.AzureInstanceIdentityToken.md)| Instance identity token | |

### Return type

[**agentsdk.AuthenticateResponse**](../Models/agentsdk.AuthenticateResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="authenticateAgentOnGoogleCloudInstance"></a>
# **authenticateAgentOnGoogleCloudInstance**
> agentsdk.AuthenticateResponse authenticateAgentOnGoogleCloudInstance(request)

Authenticate agent on Google Cloud instance

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**agentsdk.GoogleInstanceIdentityToken**](../Models/agentsdk.GoogleInstanceIdentityToken.md)| Instance identity token | |

### Return type

[**agentsdk.AuthenticateResponse**](../Models/agentsdk.AuthenticateResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="coordinateWorkspaceAgent"></a>
# **coordinateWorkspaceAgent**
> coordinateWorkspaceAgent(workspaceagent)

Coordinate workspace agent

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspaceagent** | **UUID**| Workspace agent ID | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="debugOidcContextForAUser"></a>
# **debugOidcContextForAUser**
> debugOidcContextForAUser(user)

Debug OIDC context for a user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="getConnectionInfoForWorkspaceAgent"></a>
# **getConnectionInfoForWorkspaceAgent**
> workspacesdk.AgentConnectionInfo getConnectionInfoForWorkspaceAgent(workspaceagent)

Get connection info for workspace agent

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspaceagent** | **UUID**| Workspace agent ID | [default to null] |

### Return type

[**workspacesdk.AgentConnectionInfo**](../Models/workspacesdk.AgentConnectionInfo.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getConnectionInfoForWorkspaceAgentGeneric"></a>
# **getConnectionInfoForWorkspaceAgentGeneric**
> workspacesdk.AgentConnectionInfo getConnectionInfoForWorkspaceAgentGeneric()

Get connection info for workspace agent generic

### Parameters
This endpoint does not need any parameter.

### Return type

[**workspacesdk.AgentConnectionInfo**](../Models/workspacesdk.AgentConnectionInfo.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getDerpMapUpdates"></a>
# **getDerpMapUpdates**
> getDerpMapUpdates()

Get DERP map updates

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="getListeningPortsForWorkspaceAgent"></a>
# **getListeningPortsForWorkspaceAgent**
> codersdk.WorkspaceAgentListeningPortsResponse getListeningPortsForWorkspaceAgent(workspaceagent)

Get listening ports for workspace agent

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspaceagent** | **UUID**| Workspace agent ID | [default to null] |

### Return type

[**codersdk.WorkspaceAgentListeningPortsResponse**](../Models/codersdk.WorkspaceAgentListeningPortsResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getLogsByWorkspaceAgent"></a>
# **getLogsByWorkspaceAgent**
> List getLogsByWorkspaceAgent(workspaceagent, before, after, follow, no\_compression)

Get logs by workspace agent

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspaceagent** | **UUID**| Workspace agent ID | [default to null] |
| **before** | **Integer**| Before log id | [optional] [default to null] |
| **after** | **Integer**| After log id | [optional] [default to null] |
| **follow** | **Boolean**| Follow log stream | [optional] [default to null] |
| **no\_compression** | **Boolean**| Disable compression for WebSocket connection | [optional] [default to null] |

### Return type

[**List**](../Models/codersdk.WorkspaceAgentLog.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getRunningContainersForWorkspaceAgent"></a>
# **getRunningContainersForWorkspaceAgent**
> codersdk.WorkspaceAgentListContainersResponse getRunningContainersForWorkspaceAgent(workspaceagent, label)

Get running containers for workspace agent

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspaceagent** | **UUID**| Workspace agent ID | [default to null] |
| **label** | **String**| Labels | [default to null] |

### Return type

[**codersdk.WorkspaceAgentListContainersResponse**](../Models/codersdk.WorkspaceAgentListContainersResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getWorkspaceAgentById"></a>
# **getWorkspaceAgentById**
> codersdk.WorkspaceAgent getWorkspaceAgentById(workspaceagent)

Get workspace agent by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspaceagent** | **UUID**| Workspace agent ID | [default to null] |

### Return type

[**codersdk.WorkspaceAgent**](../Models/codersdk.WorkspaceAgent.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getWorkspaceAgentExternalAuth"></a>
# **getWorkspaceAgentExternalAuth**
> agentsdk.ExternalAuthResponse getWorkspaceAgentExternalAuth(match, id, listen)

Get workspace agent external auth

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **match** | **String**| Match | [default to null] |
| **id** | **String**| Provider ID | [default to null] |
| **listen** | **Boolean**| Wait for a new token to be issued | [optional] [default to null] |

### Return type

[**agentsdk.ExternalAuthResponse**](../Models/agentsdk.ExternalAuthResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getWorkspaceAgentGitSshKey"></a>
# **getWorkspaceAgentGitSshKey**
> agentsdk.GitSSHKey getWorkspaceAgentGitSshKey()

Get workspace agent Git SSH key

### Parameters
This endpoint does not need any parameter.

### Return type

[**agentsdk.GitSSHKey**](../Models/agentsdk.GitSSHKey.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getWorkspaceAgentReinitialization"></a>
# **getWorkspaceAgentReinitialization**
> agentsdk.ReinitializationEvent getWorkspaceAgentReinitialization()

Get workspace agent reinitialization

### Parameters
This endpoint does not need any parameter.

### Return type

[**agentsdk.ReinitializationEvent**](../Models/agentsdk.ReinitializationEvent.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="openPtyToWorkspaceAgent"></a>
# **openPtyToWorkspaceAgent**
> openPtyToWorkspaceAgent(workspaceagent)

Open PTY to workspace agent

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspaceagent** | **UUID**| Workspace agent ID | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="patchWorkspaceAgentAppStatus"></a>
# **patchWorkspaceAgentAppStatus**
> codersdk.Response patchWorkspaceAgentAppStatus(request)

Patch workspace agent app status

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**agentsdk.PatchAppStatus**](../Models/agentsdk.PatchAppStatus.md)| app status | |

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="patchWorkspaceAgentLogs"></a>
# **patchWorkspaceAgentLogs**
> codersdk.Response patchWorkspaceAgentLogs(request)

Patch workspace agent logs

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**agentsdk.PatchLogs**](../Models/agentsdk.PatchLogs.md)| logs | |

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="postWorkspaceAgentLogSource"></a>
# **postWorkspaceAgentLogSource**
> codersdk.WorkspaceAgentLogSource postWorkspaceAgentLogSource(request)

Post workspace agent log source

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**agentsdk.PostLogSourceRequest**](../Models/agentsdk.PostLogSourceRequest.md)| Log source request | |

### Return type

[**codersdk.WorkspaceAgentLogSource**](../Models/codersdk.WorkspaceAgentLogSource.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="recreateDevcontainerForWorkspaceAgent"></a>
# **recreateDevcontainerForWorkspaceAgent**
> codersdk.Response recreateDevcontainerForWorkspaceAgent(workspaceagent, devcontainer)

Recreate devcontainer for workspace agent

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspaceagent** | **UUID**| Workspace agent ID | [default to null] |
| **devcontainer** | **String**| Devcontainer ID | [default to null] |

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="removedGetLogsByWorkspaceAgent"></a>
# **removedGetLogsByWorkspaceAgent**
> List removedGetLogsByWorkspaceAgent(workspaceagent, before, after, follow, no\_compression)

Removed: Get logs by workspace agent

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspaceagent** | **UUID**| Workspace agent ID | [default to null] |
| **before** | **Integer**| Before log id | [optional] [default to null] |
| **after** | **Integer**| After log id | [optional] [default to null] |
| **follow** | **Boolean**| Follow log stream | [optional] [default to null] |
| **no\_compression** | **Boolean**| Disable compression for WebSocket connection | [optional] [default to null] |

### Return type

[**List**](../Models/codersdk.WorkspaceAgentLog.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="removedGetWorkspaceAgentGitAuth"></a>
# **removedGetWorkspaceAgentGitAuth**
> agentsdk.ExternalAuthResponse removedGetWorkspaceAgentGitAuth(match, id, listen)

Removed: Get workspace agent git auth

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **match** | **String**| Match | [default to null] |
| **id** | **String**| Provider ID | [default to null] |
| **listen** | **Boolean**| Wait for a new token to be issued | [optional] [default to null] |

### Return type

[**agentsdk.ExternalAuthResponse**](../Models/agentsdk.ExternalAuthResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="userScopedTailnetRpcConnection"></a>
# **userScopedTailnetRpcConnection**
> userScopedTailnetRpcConnection()

User-scoped tailnet RPC connection

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="watchForWorkspaceAgentMetadataUpdates"></a>
# **watchForWorkspaceAgentMetadataUpdates**
> watchForWorkspaceAgentMetadataUpdates(workspaceagent)

Watch for workspace agent metadata updates

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspaceagent** | **UUID**| Workspace agent ID | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="watchForWorkspaceAgentMetadataUpdatesViaWebsockets"></a>
# **watchForWorkspaceAgentMetadataUpdatesViaWebsockets**
> codersdk.ServerSentEvent watchForWorkspaceAgentMetadataUpdatesViaWebsockets(workspaceagent)

Watch for workspace agent metadata updates via WebSockets

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspaceagent** | **UUID**| Workspace agent ID | [default to null] |

### Return type

[**codersdk.ServerSentEvent**](../Models/codersdk.ServerSentEvent.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="watchWorkspaceAgentForContainerUpdates"></a>
# **watchWorkspaceAgentForContainerUpdates**
> codersdk.WorkspaceAgentListContainersResponse watchWorkspaceAgentForContainerUpdates(workspaceagent)

Watch workspace agent for container updates.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspaceagent** | **UUID**| Workspace agent ID | [default to null] |

### Return type

[**codersdk.WorkspaceAgentListContainersResponse**](../Models/codersdk.WorkspaceAgentListContainersResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="workspaceAgentRpcApi"></a>
# **workspaceAgentRpcApi**
> workspaceAgentRpcApi()

Workspace agent RPC API

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


