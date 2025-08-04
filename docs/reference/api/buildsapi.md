# BuildsApi


All URIs are relative to */api/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**cancelWorkspaceBuild**](BuildsApi.md#cancelWorkspaceBuild) | **PATCH** /workspacebuilds/{workspacebuild}/cancel | Cancel workspace build |
| [**createWorkspaceBuild**](BuildsApi.md#createWorkspaceBuild) | **POST** /workspaces/{workspace}/builds | Create workspace build |
| [**getBuildParametersForWorkspaceBuild**](BuildsApi.md#getBuildParametersForWorkspaceBuild) | **GET** /workspacebuilds/{workspacebuild}/parameters | Get build parameters for workspace build |
| [**getProvisionerStateForWorkspaceBuild**](BuildsApi.md#getProvisionerStateForWorkspaceBuild) | **GET** /workspacebuilds/{workspacebuild}/state | Get provisioner state for workspace build |
| [**getWorkspaceBuild**](BuildsApi.md#getWorkspaceBuild) | **GET** /workspacebuilds/{workspacebuild} | Get workspace build |
| [**getWorkspaceBuildByUserWorkspaceNameAndBuildNumber**](BuildsApi.md#getWorkspaceBuildByUserWorkspaceNameAndBuildNumber) | **GET** /users/{user}/workspace/{workspacename}/builds/{buildnumber} | Get workspace build by user, workspace name, and build number |
| [**getWorkspaceBuildLogs**](BuildsApi.md#getWorkspaceBuildLogs) | **GET** /workspacebuilds/{workspacebuild}/logs | Get workspace build logs |
| [**getWorkspaceBuildTimingsById**](BuildsApi.md#getWorkspaceBuildTimingsById) | **GET** /workspacebuilds/{workspacebuild}/timings | Get workspace build timings by ID |
| [**getWorkspaceBuildsByWorkspaceId**](BuildsApi.md#getWorkspaceBuildsByWorkspaceId) | **GET** /workspaces/{workspace}/builds | Get workspace builds by workspace ID |
| [**removedGetWorkspaceResourcesForWorkspaceBuild**](BuildsApi.md#removedGetWorkspaceResourcesForWorkspaceBuild) | **GET** /workspacebuilds/{workspacebuild}/resources | Removed: Get workspace resources for workspace build |


<a name="cancelWorkspaceBuild"></a>
# **cancelWorkspaceBuild**
> codersdk.Response cancelWorkspaceBuild(workspacebuild, expect\_status)

Cancel workspace build

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspacebuild** | **String**| Workspace build ID | [default to null] |
| **expect\_status** | **String**| Expected status of the job. If expect_status is supplied, the request will be rejected with 412 Precondition Failed if the job doesn&#39;t match the state when performing the cancellation. | [optional] [default to null] [enum: running, pending] |

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="createWorkspaceBuild"></a>
# **createWorkspaceBuild**
> codersdk.WorkspaceBuild createWorkspaceBuild(workspace, request)

Create workspace build

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspace** | **UUID**| Workspace ID | [default to null] |
| **request** | [**codersdk.CreateWorkspaceBuildRequest**](../Models/codersdk.CreateWorkspaceBuildRequest.md)| Create workspace build request | |

### Return type

[**codersdk.WorkspaceBuild**](../Models/codersdk.WorkspaceBuild.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="getBuildParametersForWorkspaceBuild"></a>
# **getBuildParametersForWorkspaceBuild**
> List getBuildParametersForWorkspaceBuild(workspacebuild)

Get build parameters for workspace build

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspacebuild** | **String**| Workspace build ID | [default to null] |

### Return type

[**List**](../Models/codersdk.WorkspaceBuildParameter.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getProvisionerStateForWorkspaceBuild"></a>
# **getProvisionerStateForWorkspaceBuild**
> codersdk.WorkspaceBuild getProvisionerStateForWorkspaceBuild(workspacebuild)

Get provisioner state for workspace build

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspacebuild** | **String**| Workspace build ID | [default to null] |

### Return type

[**codersdk.WorkspaceBuild**](../Models/codersdk.WorkspaceBuild.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getWorkspaceBuild"></a>
# **getWorkspaceBuild**
> codersdk.WorkspaceBuild getWorkspaceBuild(workspacebuild)

Get workspace build

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspacebuild** | **String**| Workspace build ID | [default to null] |

### Return type

[**codersdk.WorkspaceBuild**](../Models/codersdk.WorkspaceBuild.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getWorkspaceBuildByUserWorkspaceNameAndBuildNumber"></a>
# **getWorkspaceBuildByUserWorkspaceNameAndBuildNumber**
> codersdk.WorkspaceBuild getWorkspaceBuildByUserWorkspaceNameAndBuildNumber(user, workspacename, buildnumber)

Get workspace build by user, workspace name, and build number

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |
| **workspacename** | **String**| Workspace name | [default to null] |
| **buildnumber** | **BigDecimal**| Build number | [default to null] |

### Return type

[**codersdk.WorkspaceBuild**](../Models/codersdk.WorkspaceBuild.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getWorkspaceBuildLogs"></a>
# **getWorkspaceBuildLogs**
> List getWorkspaceBuildLogs(workspacebuild, before, after, follow)

Get workspace build logs

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspacebuild** | **String**| Workspace build ID | [default to null] |
| **before** | **Integer**| Before log id | [optional] [default to null] |
| **after** | **Integer**| After log id | [optional] [default to null] |
| **follow** | **Boolean**| Follow log stream | [optional] [default to null] |

### Return type

[**List**](../Models/codersdk.ProvisionerJobLog.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getWorkspaceBuildTimingsById"></a>
# **getWorkspaceBuildTimingsById**
> codersdk.WorkspaceBuildTimings getWorkspaceBuildTimingsById(workspacebuild)

Get workspace build timings by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspacebuild** | **UUID**| Workspace build ID | [default to null] |

### Return type

[**codersdk.WorkspaceBuildTimings**](../Models/codersdk.WorkspaceBuildTimings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getWorkspaceBuildsByWorkspaceId"></a>
# **getWorkspaceBuildsByWorkspaceId**
> List getWorkspaceBuildsByWorkspaceId(workspace, after\_id, limit, offset, since)

Get workspace builds by workspace ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspace** | **UUID**| Workspace ID | [default to null] |
| **after\_id** | **UUID**| After ID | [optional] [default to null] |
| **limit** | **Integer**| Page limit | [optional] [default to null] |
| **offset** | **Integer**| Page offset | [optional] [default to null] |
| **since** | **Date**| Since timestamp | [optional] [default to null] |

### Return type

[**List**](../Models/codersdk.WorkspaceBuild.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="removedGetWorkspaceResourcesForWorkspaceBuild"></a>
# **removedGetWorkspaceResourcesForWorkspaceBuild**
> List removedGetWorkspaceResourcesForWorkspaceBuild(workspacebuild)

Removed: Get workspace resources for workspace build

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspacebuild** | **String**| Workspace build ID | [default to null] |

### Return type

[**List**](../Models/codersdk.WorkspaceResource.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


