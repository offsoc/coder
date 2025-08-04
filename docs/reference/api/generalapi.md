# GeneralApi


All URIs are relative to */api/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiRootHandler**](GeneralApi.md#apiRootHandler) | **GET** / | API root handler |
| [**buildInfo**](GeneralApi.md#buildInfo) | **GET** /buildinfo | Build info |
| [**getDeploymentConfig**](GeneralApi.md#getDeploymentConfig) | **GET** /deployment/config | Get deployment config |
| [**getDeploymentStats**](GeneralApi.md#getDeploymentStats) | **GET** /deployment/stats | Get deployment stats |
| [**getEnabledExperiments**](GeneralApi.md#getEnabledExperiments) | **GET** /experiments | Get enabled experiments |
| [**getSafeExperiments**](GeneralApi.md#getSafeExperiments) | **GET** /experiments/available | Get safe experiments |
| [**getTokenConfig**](GeneralApi.md#getTokenConfig) | **GET** /users/{user}/keys/tokens/tokenconfig | Get token config |
| [**reportCspViolations**](GeneralApi.md#reportCspViolations) | **POST** /csp/reports | Report CSP violations |
| [**sshConfig**](GeneralApi.md#sshConfig) | **GET** /deployment/ssh | SSH Config |
| [**updateCheck**](GeneralApi.md#updateCheck) | **GET** /updatecheck | Update check |


<a name="apiRootHandler"></a>
# **apiRootHandler**
> codersdk.Response apiRootHandler()

API root handler

### Parameters
This endpoint does not need any parameter.

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="buildInfo"></a>
# **buildInfo**
> codersdk.BuildInfoResponse buildInfo()

Build info

### Parameters
This endpoint does not need any parameter.

### Return type

[**codersdk.BuildInfoResponse**](../Models/codersdk.BuildInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getDeploymentConfig"></a>
# **getDeploymentConfig**
> codersdk.DeploymentConfig getDeploymentConfig()

Get deployment config

### Parameters
This endpoint does not need any parameter.

### Return type

[**codersdk.DeploymentConfig**](../Models/codersdk.DeploymentConfig.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getDeploymentStats"></a>
# **getDeploymentStats**
> codersdk.DeploymentStats getDeploymentStats()

Get deployment stats

### Parameters
This endpoint does not need any parameter.

### Return type

[**codersdk.DeploymentStats**](../Models/codersdk.DeploymentStats.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getEnabledExperiments"></a>
# **getEnabledExperiments**
> List getEnabledExperiments()

Get enabled experiments

### Parameters
This endpoint does not need any parameter.

### Return type

[**List**](../Models/codersdk.Experiment.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getSafeExperiments"></a>
# **getSafeExperiments**
> List getSafeExperiments()

Get safe experiments

### Parameters
This endpoint does not need any parameter.

### Return type

[**List**](../Models/codersdk.Experiment.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTokenConfig"></a>
# **getTokenConfig**
> codersdk.TokenConfig getTokenConfig(user)

Get token config

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |

### Return type

[**codersdk.TokenConfig**](../Models/codersdk.TokenConfig.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="reportCspViolations"></a>
# **reportCspViolations**
> reportCspViolations(request)

Report CSP violations

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**coderd.cspViolation**](../Models/coderd.cspViolation.md)| Violation report | |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

<a name="sshConfig"></a>
# **sshConfig**
> codersdk.SSHConfigResponse sshConfig()

SSH Config

### Parameters
This endpoint does not need any parameter.

### Return type

[**codersdk.SSHConfigResponse**](../Models/codersdk.SSHConfigResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateCheck"></a>
# **updateCheck**
> codersdk.UpdateCheckResponse updateCheck()

Update check

### Parameters
This endpoint does not need any parameter.

### Return type

[**codersdk.UpdateCheckResponse**](../Models/codersdk.UpdateCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


