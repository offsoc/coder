# Workspaces


All URIs are relative to */api/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createUserWorkspace**](WorkspacesApi.md#createUserWorkspace) | **POST** /users/{user}/workspaces | Create user workspace |
| [**createUserWorkspaceByOrganization**](WorkspacesApi.md#createUserWorkspaceByOrganization) | **POST** /organizations/{organization}/members/{user}/workspaces | Create user workspace by organization |
| [**extendWorkspaceDeadlineById**](WorkspacesApi.md#extendWorkspaceDeadlineById) | **PUT** /workspaces/{workspace}/extend | Extend workspace deadline by ID |
| [**favoriteWorkspaceById**](WorkspacesApi.md#favoriteWorkspaceById) | **PUT** /workspaces/{workspace}/favorite | Favorite workspace by ID. |
| [**getWorkspaceMetadataById**](WorkspacesApi.md#getWorkspaceMetadataById) | **GET** /workspaces/{workspace} | Get workspace metadata by ID |
| [**getWorkspaceMetadataByUserAndWorkspaceName**](WorkspacesApi.md#getWorkspaceMetadataByUserAndWorkspaceName) | **GET** /users/{user}/workspace/{workspacename} | Get workspace metadata by user and workspace name |
| [**getWorkspaceTimingsById**](WorkspacesApi.md#getWorkspaceTimingsById) | **GET** /workspaces/{workspace}/timings | Get workspace timings by ID |
| [**listWorkspaces**](WorkspacesApi.md#listWorkspaces) | **GET** /workspaces | List workspaces |
| [**postWorkspaceUsageById**](WorkspacesApi.md#postWorkspaceUsageById) | **POST** /workspaces/{workspace}/usage | Post Workspace Usage by ID |
| [**resolveWorkspaceAutostartById**](WorkspacesApi.md#resolveWorkspaceAutostartById) | **GET** /workspaces/{workspace}/resolve-autostart | Resolve workspace autostart by id. |
| [**unfavoriteWorkspaceById**](WorkspacesApi.md#unfavoriteWorkspaceById) | **DELETE** /workspaces/{workspace}/favorite | Unfavorite workspace by ID. |
| [**updateWorkspaceAcl**](WorkspacesApi.md#updateWorkspaceAcl) | **PATCH** /workspaces/{workspace}/acl | Update workspace ACL |
| [**updateWorkspaceAutomaticUpdatesById**](WorkspacesApi.md#updateWorkspaceAutomaticUpdatesById) | **PUT** /workspaces/{workspace}/autoupdates | Update workspace automatic updates by ID |
| [**updateWorkspaceAutostartScheduleById**](WorkspacesApi.md#updateWorkspaceAutostartScheduleById) | **PUT** /workspaces/{workspace}/autostart | Update workspace autostart schedule by ID |
| [**updateWorkspaceDormancyStatusById**](WorkspacesApi.md#updateWorkspaceDormancyStatusById) | **PUT** /workspaces/{workspace}/dormant | Update workspace dormancy status by id. |
| [**updateWorkspaceMetadataById**](WorkspacesApi.md#updateWorkspaceMetadataById) | **PATCH** /workspaces/{workspace} | Update workspace metadata by ID |
| [**updateWorkspaceTtlById**](WorkspacesApi.md#updateWorkspaceTtlById) | **PUT** /workspaces/{workspace}/ttl | Update workspace TTL by ID |
| [**watchWorkspaceById**](WorkspacesApi.md#watchWorkspaceById) | **GET** /workspaces/{workspace}/watch | Watch workspace by ID |
| [**watchWorkspaceByIdViaWebsockets**](WorkspacesApi.md#watchWorkspaceByIdViaWebsockets) | **GET** /workspaces/{workspace}/watch-ws | Watch workspace by ID via WebSockets |


<a name="createUserWorkspace"></a>
# **createUserWorkspace**
> codersdk.Workspace createUserWorkspace(user, request)

Create user workspace

    Create a new workspace using a template. The request must specify either the Template ID or the Template Version ID, not both. If the Template ID is specified, the active version of the template will be used.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| Username, UUID, or me | [default to null] |
| **request** | [**codersdk.CreateWorkspaceRequest**](../Models/codersdk.CreateWorkspaceRequest.md)| Create workspace request | |

### Return type

[**codersdk.Workspace**](../Models/codersdk.Workspace.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createUserWorkspaceByOrganization"></a>
# **createUserWorkspaceByOrganization**
> codersdk.Workspace createUserWorkspaceByOrganization(organization, user, request)

Create user workspace by organization

    Create a new workspace using a template. The request must specify either the Template ID or the Template Version ID, not both. If the Template ID is specified, the active version of the template will be used.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |
| **user** | **String**| Username, UUID, or me | [default to null] |
| **request** | [**codersdk.CreateWorkspaceRequest**](../Models/codersdk.CreateWorkspaceRequest.md)| Create workspace request | |

### Return type

[**codersdk.Workspace**](../Models/codersdk.Workspace.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="extendWorkspaceDeadlineById"></a>
# **extendWorkspaceDeadlineById**
> codersdk.Response extendWorkspaceDeadlineById(workspace, request)

Extend workspace deadline by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspace** | **UUID**| Workspace ID | [default to null] |
| **request** | [**codersdk.PutExtendWorkspaceRequest**](../Models/codersdk.PutExtendWorkspaceRequest.md)| Extend deadline update request | |

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="favoriteWorkspaceById"></a>
# **favoriteWorkspaceById**
> favoriteWorkspaceById(workspace)

Favorite workspace by ID.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspace** | **UUID**| Workspace ID | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="getWorkspaceMetadataById"></a>
# **getWorkspaceMetadataById**
> codersdk.Workspace getWorkspaceMetadataById(workspace, include\_deleted)

Get workspace metadata by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspace** | **UUID**| Workspace ID | [default to null] |
| **include\_deleted** | **Boolean**| Return data instead of HTTP 404 if the workspace is deleted | [optional] [default to null] |

### Return type

[**codersdk.Workspace**](../Models/codersdk.Workspace.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getWorkspaceMetadataByUserAndWorkspaceName"></a>
# **getWorkspaceMetadataByUserAndWorkspaceName**
> codersdk.Workspace getWorkspaceMetadataByUserAndWorkspaceName(user, workspacename, include\_deleted)

Get workspace metadata by user and workspace name

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |
| **workspacename** | **String**| Workspace name | [default to null] |
| **include\_deleted** | **Boolean**| Return data instead of HTTP 404 if the workspace is deleted | [optional] [default to null] |

### Return type

[**codersdk.Workspace**](../Models/codersdk.Workspace.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getWorkspaceTimingsById"></a>
# **getWorkspaceTimingsById**
> codersdk.WorkspaceBuildTimings getWorkspaceTimingsById(workspace)

Get workspace timings by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspace** | **UUID**| Workspace ID | [default to null] |

### Return type

[**codersdk.WorkspaceBuildTimings**](../Models/codersdk.WorkspaceBuildTimings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listWorkspaces"></a>
# **listWorkspaces**
> codersdk.WorkspacesResponse listWorkspaces(q, limit, offset)

List workspaces

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **q** | **String**| Search query in the format &#x60;key:value&#x60;. Available keys are: owner, template, name, status, has-agent, dormant, last_used_after, last_used_before, has-ai-task. | [optional] [default to null] |
| **limit** | **Integer**| Page limit | [optional] [default to null] |
| **offset** | **Integer**| Page offset | [optional] [default to null] |

### Return type

[**codersdk.WorkspacesResponse**](../Models/codersdk.WorkspacesResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="postWorkspaceUsageById"></a>
# **postWorkspaceUsageById**
> postWorkspaceUsageById(workspace, request)

Post Workspace Usage by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspace** | **UUID**| Workspace ID | [default to null] |
| **request** | [**codersdk.PostWorkspaceUsageRequest**](../Models/codersdk.PostWorkspaceUsageRequest.md)| Post workspace usage request | [optional] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

<a name="resolveWorkspaceAutostartById"></a>
# **resolveWorkspaceAutostartById**
> codersdk.ResolveAutostartResponse resolveWorkspaceAutostartById(workspace)

Resolve workspace autostart by id.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspace** | **UUID**| Workspace ID | [default to null] |

### Return type

[**codersdk.ResolveAutostartResponse**](../Models/codersdk.ResolveAutostartResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="unfavoriteWorkspaceById"></a>
# **unfavoriteWorkspaceById**
> unfavoriteWorkspaceById(workspace)

Unfavorite workspace by ID.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspace** | **UUID**| Workspace ID | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="updateWorkspaceAcl"></a>
# **updateWorkspaceAcl**
> updateWorkspaceAcl(workspace, request)

Update workspace ACL

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspace** | **UUID**| Workspace ID | [default to null] |
| **request** | [**codersdk.UpdateWorkspaceACL**](../Models/codersdk.UpdateWorkspaceACL.md)| Update workspace ACL request | |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

<a name="updateWorkspaceAutomaticUpdatesById"></a>
# **updateWorkspaceAutomaticUpdatesById**
> updateWorkspaceAutomaticUpdatesById(workspace, request)

Update workspace automatic updates by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspace** | **UUID**| Workspace ID | [default to null] |
| **request** | [**codersdk.UpdateWorkspaceAutomaticUpdatesRequest**](../Models/codersdk.UpdateWorkspaceAutomaticUpdatesRequest.md)| Automatic updates request | |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

<a name="updateWorkspaceAutostartScheduleById"></a>
# **updateWorkspaceAutostartScheduleById**
> updateWorkspaceAutostartScheduleById(workspace, request)

Update workspace autostart schedule by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspace** | **UUID**| Workspace ID | [default to null] |
| **request** | [**codersdk.UpdateWorkspaceAutostartRequest**](../Models/codersdk.UpdateWorkspaceAutostartRequest.md)| Schedule update request | |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

<a name="updateWorkspaceDormancyStatusById"></a>
# **updateWorkspaceDormancyStatusById**
> codersdk.Workspace updateWorkspaceDormancyStatusById(workspace, request)

Update workspace dormancy status by id.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspace** | **UUID**| Workspace ID | [default to null] |
| **request** | [**codersdk.UpdateWorkspaceDormancy**](../Models/codersdk.UpdateWorkspaceDormancy.md)| Make a workspace dormant or active | |

### Return type

[**codersdk.Workspace**](../Models/codersdk.Workspace.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateWorkspaceMetadataById"></a>
# **updateWorkspaceMetadataById**
> updateWorkspaceMetadataById(workspace, request)

Update workspace metadata by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspace** | **UUID**| Workspace ID | [default to null] |
| **request** | [**codersdk.UpdateWorkspaceRequest**](../Models/codersdk.UpdateWorkspaceRequest.md)| Metadata update request | |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

<a name="updateWorkspaceTtlById"></a>
# **updateWorkspaceTtlById**
> updateWorkspaceTtlById(workspace, request)

Update workspace TTL by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspace** | **UUID**| Workspace ID | [default to null] |
| **request** | [**codersdk.UpdateWorkspaceTTLRequest**](../Models/codersdk.UpdateWorkspaceTTLRequest.md)| Workspace TTL update request | |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

<a name="watchWorkspaceById"></a>
# **watchWorkspaceById**
> codersdk.Response watchWorkspaceById(workspace)

Watch workspace by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspace** | **UUID**| Workspace ID | [default to null] |

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream

<a name="watchWorkspaceByIdViaWebsockets"></a>
# **watchWorkspaceByIdViaWebsockets**
> codersdk.ServerSentEvent watchWorkspaceByIdViaWebsockets(workspace)

Watch workspace by ID via WebSockets

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspace** | **UUID**| Workspace ID | [default to null] |

### Return type

[**codersdk.ServerSentEvent**](../Models/codersdk.ServerSentEvent.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


