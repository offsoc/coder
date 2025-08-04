# Applications


All URIs are relative to */api/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getApplicationsHost**](ApplicationsApi.md#getApplicationsHost) | **GET** /applications/host | Get applications host |
| [**redirectToUriWithEncryptedApiKey**](ApplicationsApi.md#redirectToUriWithEncryptedApiKey) | **GET** /applications/auth-redirect | Redirect to URI with encrypted API key |


<a name="getApplicationsHost"></a>
# **getApplicationsHost**
> codersdk.AppHostResponse getApplicationsHost()

Get applications host

### Parameters
This endpoint does not need any parameter.

### Return type

[**codersdk.AppHostResponse**](../Models/codersdk.AppHostResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="redirectToUriWithEncryptedApiKey"></a>
# **redirectToUriWithEncryptedApiKey**
> redirectToUriWithEncryptedApiKey(redirect\_uri)

Redirect to URI with encrypted API key

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **redirect\_uri** | **String**| Redirect destination | [optional] [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


