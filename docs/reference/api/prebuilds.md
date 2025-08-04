# Prebuilds


All URIs are relative to */api/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getPrebuildsSettings**](PrebuildsApi.md#getPrebuildsSettings) | **GET** /prebuilds/settings | Get prebuilds settings |
| [**updatePrebuildsSettings**](PrebuildsApi.md#updatePrebuildsSettings) | **PUT** /prebuilds/settings | Update prebuilds settings |


<a name="getPrebuildsSettings"></a>
# **getPrebuildsSettings**
> codersdk.PrebuildsSettings getPrebuildsSettings()

Get prebuilds settings

### Parameters
This endpoint does not need any parameter.

### Return type

[**codersdk.PrebuildsSettings**](../Models/codersdk.PrebuildsSettings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updatePrebuildsSettings"></a>
# **updatePrebuildsSettings**
> codersdk.PrebuildsSettings updatePrebuildsSettings(request)

Update prebuilds settings

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**codersdk.PrebuildsSettings**](../Models/codersdk.PrebuildsSettings.md)| Prebuilds settings request | |

### Return type

[**codersdk.PrebuildsSettings**](../Models/codersdk.PrebuildsSettings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


