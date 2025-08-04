# Git


All URIs are relative to */api/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**deleteExternalAuthUserLinkById**](GitApi.md#deleteExternalAuthUserLinkById) | **DELETE** /external-auth/{externalauth} | Delete external auth user link by ID |
| [**getExternalAuthById**](GitApi.md#getExternalAuthById) | **GET** /external-auth/{externalauth} | Get external auth by ID |
| [**getExternalAuthDeviceById**](GitApi.md#getExternalAuthDeviceById) | **GET** /external-auth/{externalauth}/device | Get external auth device by ID. |
| [**getUserExternalAuths**](GitApi.md#getUserExternalAuths) | **GET** /external-auth | Get user external auths |
| [**postExternalAuthDeviceById**](GitApi.md#postExternalAuthDeviceById) | **POST** /external-auth/{externalauth}/device | Post external auth device by ID |


<a name="deleteExternalAuthUserLinkById"></a>
# **deleteExternalAuthUserLinkById**
> deleteExternalAuthUserLinkById(externalauth)

Delete external auth user link by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **externalauth** | **String**| Git Provider ID | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="getExternalAuthById"></a>
# **getExternalAuthById**
> codersdk.ExternalAuth getExternalAuthById(externalauth)

Get external auth by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **externalauth** | **String**| Git Provider ID | [default to null] |

### Return type

[**codersdk.ExternalAuth**](../Models/codersdk.ExternalAuth.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getExternalAuthDeviceById"></a>
# **getExternalAuthDeviceById**
> codersdk.ExternalAuthDevice getExternalAuthDeviceById(externalauth)

Get external auth device by ID.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **externalauth** | **String**| Git Provider ID | [default to null] |

### Return type

[**codersdk.ExternalAuthDevice**](../Models/codersdk.ExternalAuthDevice.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserExternalAuths"></a>
# **getUserExternalAuths**
> codersdk.ExternalAuthLink getUserExternalAuths()

Get user external auths

### Parameters
This endpoint does not need any parameter.

### Return type

[**codersdk.ExternalAuthLink**](../Models/codersdk.ExternalAuthLink.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="postExternalAuthDeviceById"></a>
# **postExternalAuthDeviceById**
> postExternalAuthDeviceById(externalauth)

Post external auth device by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **externalauth** | **String**| External Provider ID | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


