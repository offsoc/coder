# Provisioning


All URIs are relative to */api/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getProvisionerDaemons**](ProvisioningApi.md#getProvisionerDaemons) | **GET** /organizations/{organization}/provisionerdaemons | Get provisioner daemons |


<a name="getProvisionerDaemons"></a>
# **getProvisionerDaemons**
> List getProvisionerDaemons(organization, limit, ids, status, tags)

Get provisioner daemons

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |
| **limit** | **Integer**| Page limit | [optional] [default to null] |
| **ids** | [**List**](../Models/String.md)| Filter results by job IDs | [optional] [default to null] |
| **status** | **String**| Filter results by status | [optional] [default to null] [enum: pending, running, succeeded, canceling, canceled, failed, unknown, pending, running, succeeded, canceling, canceled, failed] |
| **tags** | [**Object**](../Models/.md)| Provisioner tags to filter by (JSON of the form {&#39;tag1&#39;:&#39;value1&#39;,&#39;tag2&#39;:&#39;value2&#39;}) | [optional] [default to null] |

### Return type

[**List**](../Models/codersdk.ProvisionerDaemon.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


