# Debug


All URIs are relative to */api/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**debugDerpTraffic**](DebugApi.md#debugDerpTraffic) | **GET** /debug/derp/traffic | Debug DERP traffic |
| [**debugExpvar**](DebugApi.md#debugExpvar) | **GET** /debug/expvar | Debug expvar |
| [**debugInfoDeploymentHealth**](DebugApi.md#debugInfoDeploymentHealth) | **GET** /debug/health | Debug Info Deployment Health |
| [**debugInfoTailnet**](DebugApi.md#debugInfoTailnet) | **GET** /debug/tailnet | Debug Info Tailnet |
| [**debugInfoWebsocketTest**](DebugApi.md#debugInfoWebsocketTest) | **GET** /debug/ws | Debug Info Websocket Test |
| [**debugInfoWireguardCoordinator**](DebugApi.md#debugInfoWireguardCoordinator) | **GET** /debug/coordinator | Debug Info Wireguard Coordinator |
| [**getHealthSettings**](DebugApi.md#getHealthSettings) | **GET** /debug/health/settings | Get health settings |
| [**updateHealthSettings**](DebugApi.md#updateHealthSettings) | **PUT** /debug/health/settings | Update health settings |


<a name="debugDerpTraffic"></a>
# **debugDerpTraffic**
> List debugDerpTraffic()

Debug DERP traffic

### Parameters
This endpoint does not need any parameter.

### Return type

[**List**](../Models/derp.BytesSentRecv.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="debugExpvar"></a>
# **debugExpvar**
> Map debugExpvar()

Debug expvar

### Parameters
This endpoint does not need any parameter.

### Return type

[**Map**](../Models/AnyType.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="debugInfoDeploymentHealth"></a>
# **debugInfoDeploymentHealth**
> healthsdk.HealthcheckReport debugInfoDeploymentHealth(force)

Debug Info Deployment Health

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **force** | **Boolean**| Force a healthcheck to run | [optional] [default to null] |

### Return type

[**healthsdk.HealthcheckReport**](../Models/healthsdk.HealthcheckReport.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="debugInfoTailnet"></a>
# **debugInfoTailnet**
> debugInfoTailnet()

Debug Info Tailnet

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="debugInfoWebsocketTest"></a>
# **debugInfoWebsocketTest**
> codersdk.Response debugInfoWebsocketTest()

Debug Info Websocket Test

### Parameters
This endpoint does not need any parameter.

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="debugInfoWireguardCoordinator"></a>
# **debugInfoWireguardCoordinator**
> debugInfoWireguardCoordinator()

Debug Info Wireguard Coordinator

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="getHealthSettings"></a>
# **getHealthSettings**
> healthsdk.HealthSettings getHealthSettings()

Get health settings

### Parameters
This endpoint does not need any parameter.

### Return type

[**healthsdk.HealthSettings**](../Models/healthsdk.HealthSettings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateHealthSettings"></a>
# **updateHealthSettings**
> healthsdk.UpdateHealthSettings updateHealthSettings(request)

Update health settings

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**healthsdk.UpdateHealthSettings**](../Models/healthsdk.UpdateHealthSettings.md)| Update health settings | |

### Return type

[**healthsdk.UpdateHealthSettings**](../Models/healthsdk.UpdateHealthSettings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


