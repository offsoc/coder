# PortSharing


All URIs are relative to */api/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**deleteWorkspaceAgentPortShare**](PortSharingApi.md#deleteWorkspaceAgentPortShare) | **DELETE** /workspaces/{workspace}/port-share | Delete workspace agent port share |
| [**getWorkspaceAgentPortShares**](PortSharingApi.md#getWorkspaceAgentPortShares) | **GET** /workspaces/{workspace}/port-share | Get workspace agent port shares |
| [**upsertWorkspaceAgentPortShare**](PortSharingApi.md#upsertWorkspaceAgentPortShare) | **POST** /workspaces/{workspace}/port-share | Upsert workspace agent port share |


<a name="deleteWorkspaceAgentPortShare"></a>
# **deleteWorkspaceAgentPortShare**
> deleteWorkspaceAgentPortShare(workspace, request)

Delete workspace agent port share

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspace** | **UUID**| Workspace ID | [default to null] |
| **request** | [**codersdk.DeleteWorkspaceAgentPortShareRequest**](../Models/codersdk.DeleteWorkspaceAgentPortShareRequest.md)| Delete port sharing level request | |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

<a name="getWorkspaceAgentPortShares"></a>
# **getWorkspaceAgentPortShares**
> codersdk.WorkspaceAgentPortShares getWorkspaceAgentPortShares(workspace)

Get workspace agent port shares

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspace** | **UUID**| Workspace ID | [default to null] |

### Return type

[**codersdk.WorkspaceAgentPortShares**](../Models/codersdk.WorkspaceAgentPortShares.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="upsertWorkspaceAgentPortShare"></a>
# **upsertWorkspaceAgentPortShare**
> codersdk.WorkspaceAgentPortShare upsertWorkspaceAgentPortShare(workspace, request)

Upsert workspace agent port share

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspace** | **UUID**| Workspace ID | [default to null] |
| **request** | [**codersdk.UpsertWorkspaceAgentPortShareRequest**](../Models/codersdk.UpsertWorkspaceAgentPortShareRequest.md)| Upsert port sharing level request | |

### Return type

[**codersdk.WorkspaceAgentPortShare**](../Models/codersdk.WorkspaceAgentPortShare.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


