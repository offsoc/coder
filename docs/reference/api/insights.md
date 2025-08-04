# Insights


All URIs are relative to */api/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getDeploymentDaus**](InsightsApi.md#getDeploymentDaus) | **GET** /insights/daus | Get deployment DAUs |
| [**getInsightsAboutTemplates**](InsightsApi.md#getInsightsAboutTemplates) | **GET** /insights/templates | Get insights about templates |
| [**getInsightsAboutUserActivity**](InsightsApi.md#getInsightsAboutUserActivity) | **GET** /insights/user-activity | Get insights about user activity |
| [**getInsightsAboutUserLatency**](InsightsApi.md#getInsightsAboutUserLatency) | **GET** /insights/user-latency | Get insights about user latency |
| [**getInsightsAboutUserStatusCounts**](InsightsApi.md#getInsightsAboutUserStatusCounts) | **GET** /insights/user-status-counts | Get insights about user status counts |


<a name="getDeploymentDaus"></a>
# **getDeploymentDaus**
> codersdk.DAUsResponse getDeploymentDaus(tz\_offset)

Get deployment DAUs

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **tz\_offset** | **Integer**| Time-zone offset (e.g. -2) | [default to null] |

### Return type

[**codersdk.DAUsResponse**](../Models/codersdk.DAUsResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getInsightsAboutTemplates"></a>
# **getInsightsAboutTemplates**
> codersdk.TemplateInsightsResponse getInsightsAboutTemplates(start\_time, end\_time, interval, template\_ids)

Get insights about templates

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **start\_time** | **Date**| Start time | [default to null] |
| **end\_time** | **Date**| End time | [default to null] |
| **interval** | **String**| Interval | [default to null] [enum: week, day] |
| **template\_ids** | [**List**](../Models/String.md)| Template IDs | [optional] [default to null] |

### Return type

[**codersdk.TemplateInsightsResponse**](../Models/codersdk.TemplateInsightsResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getInsightsAboutUserActivity"></a>
# **getInsightsAboutUserActivity**
> codersdk.UserActivityInsightsResponse getInsightsAboutUserActivity(start\_time, end\_time, template\_ids)

Get insights about user activity

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **start\_time** | **Date**| Start time | [default to null] |
| **end\_time** | **Date**| End time | [default to null] |
| **template\_ids** | [**List**](../Models/String.md)| Template IDs | [optional] [default to null] |

### Return type

[**codersdk.UserActivityInsightsResponse**](../Models/codersdk.UserActivityInsightsResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getInsightsAboutUserLatency"></a>
# **getInsightsAboutUserLatency**
> codersdk.UserLatencyInsightsResponse getInsightsAboutUserLatency(start\_time, end\_time, template\_ids)

Get insights about user latency

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **start\_time** | **Date**| Start time | [default to null] |
| **end\_time** | **Date**| End time | [default to null] |
| **template\_ids** | [**List**](../Models/String.md)| Template IDs | [optional] [default to null] |

### Return type

[**codersdk.UserLatencyInsightsResponse**](../Models/codersdk.UserLatencyInsightsResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getInsightsAboutUserStatusCounts"></a>
# **getInsightsAboutUserStatusCounts**
> codersdk.GetUserStatusCountsResponse getInsightsAboutUserStatusCounts(tz\_offset)

Get insights about user status counts

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **tz\_offset** | **Integer**| Time-zone offset (e.g. -2) | [default to null] |

### Return type

[**codersdk.GetUserStatusCountsResponse**](../Models/codersdk.GetUserStatusCountsResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


