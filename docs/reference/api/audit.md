# Audit


All URIs are relative to */api/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**generateFakeAuditLog**](AuditApi.md#generateFakeAuditLog) | **POST** /audit/testgenerate | Generate fake audit log |
| [**getAuditLogs**](AuditApi.md#getAuditLogs) | **GET** /audit | Get audit logs |


<a name="generateFakeAuditLog"></a>
# **generateFakeAuditLog**
> generateFakeAuditLog(request)

Generate fake audit log

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**codersdk.CreateTestAuditLogRequest**](../Models/codersdk.CreateTestAuditLogRequest.md)| Audit log request | |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

<a name="getAuditLogs"></a>
# **getAuditLogs**
> codersdk.AuditLogResponse getAuditLogs(limit, q, offset)

Get audit logs

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **limit** | **Integer**| Page limit | [default to null] |
| **q** | **String**| Search query | [optional] [default to null] |
| **offset** | **Integer**| Page offset | [optional] [default to null] |

### Return type

[**codersdk.AuditLogResponse**](../Models/codersdk.AuditLogResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


