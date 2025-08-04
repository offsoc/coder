# Organizations


All URIs are relative to */api/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**addNewLicense**](OrganizationsApi.md#addNewLicense) | **POST** /licenses | Add new license |
| [**createOrganization**](OrganizationsApi.md#createOrganization) | **POST** /organizations | Create organization |
| [**deleteOrganization**](OrganizationsApi.md#deleteOrganization) | **DELETE** /organizations/{organization} | Delete organization |
| [**getOrganizationById**](OrganizationsApi.md#getOrganizationById) | **GET** /organizations/{organization} | Get organization by ID |
| [**getOrganizations**](OrganizationsApi.md#getOrganizations) | **GET** /organizations | Get organizations |
| [**getProvisionerJob**](OrganizationsApi.md#getProvisionerJob) | **GET** /organizations/{organization}/provisionerjobs/{job} | Get provisioner job |
| [**getProvisionerJobs**](OrganizationsApi.md#getProvisionerJobs) | **GET** /organizations/{organization}/provisionerjobs | Get provisioner jobs |
| [**updateLicenseEntitlements**](OrganizationsApi.md#updateLicenseEntitlements) | **POST** /licenses/refresh-entitlements | Update license entitlements |
| [**updateOrganization**](OrganizationsApi.md#updateOrganization) | **PATCH** /organizations/{organization} | Update organization |


<a name="addNewLicense"></a>
# **addNewLicense**
> codersdk.License addNewLicense(request)

Add new license

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**codersdk.AddLicenseRequest**](../Models/codersdk.AddLicenseRequest.md)| Add license request | |

### Return type

[**codersdk.License**](../Models/codersdk.License.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createOrganization"></a>
# **createOrganization**
> codersdk.Organization createOrganization(request)

Create organization

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**codersdk.CreateOrganizationRequest**](../Models/codersdk.CreateOrganizationRequest.md)| Create organization request | |

### Return type

[**codersdk.Organization**](../Models/codersdk.Organization.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteOrganization"></a>
# **deleteOrganization**
> codersdk.Response deleteOrganization(organization)

Delete organization

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **String**| Organization ID or name | [default to null] |

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getOrganizationById"></a>
# **getOrganizationById**
> codersdk.Organization getOrganizationById(organization)

Get organization by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |

### Return type

[**codersdk.Organization**](../Models/codersdk.Organization.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getOrganizations"></a>
# **getOrganizations**
> List getOrganizations()

Get organizations

### Parameters
This endpoint does not need any parameter.

### Return type

[**List**](../Models/codersdk.Organization.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getProvisionerJob"></a>
# **getProvisionerJob**
> codersdk.ProvisionerJob getProvisionerJob(organization, job)

Get provisioner job

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |
| **job** | **UUID**| Job ID | [default to null] |

### Return type

[**codersdk.ProvisionerJob**](../Models/codersdk.ProvisionerJob.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getProvisionerJobs"></a>
# **getProvisionerJobs**
> List getProvisionerJobs(organization, limit, ids, status, tags)

Get provisioner jobs

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |
| **limit** | **Integer**| Page limit | [optional] [default to null] |
| **ids** | [**List**](../Models/String.md)| Filter results by job IDs | [optional] [default to null] |
| **status** | **String**| Filter results by status | [optional] [default to null] [enum: pending, running, succeeded, canceling, canceled, failed, unknown, pending, running, succeeded, canceling, canceled, failed] |
| **tags** | [**Object**](../Models/.md)| Provisioner tags to filter by (JSON of the form {&#39;tag1&#39;:&#39;value1&#39;,&#39;tag2&#39;:&#39;value2&#39;}) | [optional] [default to null] |

### Return type

[**List**](../Models/codersdk.ProvisionerJob.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateLicenseEntitlements"></a>
# **updateLicenseEntitlements**
> codersdk.Response updateLicenseEntitlements()

Update license entitlements

### Parameters
This endpoint does not need any parameter.

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateOrganization"></a>
# **updateOrganization**
> codersdk.Organization updateOrganization(organization, request)

Update organization

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **String**| Organization ID or name | [default to null] |
| **request** | [**codersdk.UpdateOrganizationRequest**](../Models/codersdk.UpdateOrganizationRequest.md)| Patch organization request | |

### Return type

[**codersdk.Organization**](../Models/codersdk.Organization.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


