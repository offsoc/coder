# TemplatesApi


All URIs are relative to */api/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**archiveTemplateUnusedVersionsByTemplateId**](TemplatesApi.md#archiveTemplateUnusedVersionsByTemplateId) | **POST** /templates/{template}/versions/archive | Archive template unused versions by template id |
| [**archiveTemplateVersion**](TemplatesApi.md#archiveTemplateVersion) | **POST** /templateversions/{templateversion}/archive | Archive template version |
| [**cancelTemplateVersionById**](TemplatesApi.md#cancelTemplateVersionById) | **PATCH** /templateversions/{templateversion}/cancel | Cancel template version by ID |
| [**cancelTemplateVersionDryRunByJobId**](TemplatesApi.md#cancelTemplateVersionDryRunByJobId) | **PATCH** /templateversions/{templateversion}/dry-run/{jobID}/cancel | Cancel template version dry-run by job ID |
| [**createTemplateByOrganization**](TemplatesApi.md#createTemplateByOrganization) | **POST** /organizations/{organization}/templates | Create template by organization |
| [**createTemplateVersionByOrganization**](TemplatesApi.md#createTemplateVersionByOrganization) | **POST** /organizations/{organization}/templateversions | Create template version by organization |
| [**createTemplateVersionDryRun**](TemplatesApi.md#createTemplateVersionDryRun) | **POST** /templateversions/{templateversion}/dry-run | Create template version dry-run |
| [**deleteTemplateById**](TemplatesApi.md#deleteTemplateById) | **DELETE** /templates/{template} | Delete template by ID |
| [**evaluateDynamicParametersForTemplateVersion**](TemplatesApi.md#evaluateDynamicParametersForTemplateVersion) | **POST** /templateversions/{templateversion}/dynamic-parameters/evaluate | Evaluate dynamic parameters for template version |
| [**getAllTemplates**](TemplatesApi.md#getAllTemplates) | **GET** /templates | Get all templates |
| [**getExternalAuthByTemplateVersion**](TemplatesApi.md#getExternalAuthByTemplateVersion) | **GET** /templateversions/{templateversion}/external-auth | Get external auth by template version |
| [**getLogsByTemplateVersion**](TemplatesApi.md#getLogsByTemplateVersion) | **GET** /templateversions/{templateversion}/logs | Get logs by template version |
| [**getPreviousTemplateVersionByOrganizationTemplateAndName**](TemplatesApi.md#getPreviousTemplateVersionByOrganizationTemplateAndName) | **GET** /organizations/{organization}/templates/{templatename}/versions/{templateversionname}/previous | Get previous template version by organization, template, and name |
| [**getResourcesByTemplateVersion**](TemplatesApi.md#getResourcesByTemplateVersion) | **GET** /templateversions/{templateversion}/resources | Get resources by template version |
| [**getRichParametersByTemplateVersion**](TemplatesApi.md#getRichParametersByTemplateVersion) | **GET** /templateversions/{templateversion}/rich-parameters | Get rich parameters by template version |
| [**getTemplateDausById**](TemplatesApi.md#getTemplateDausById) | **GET** /templates/{template}/daus | Get template DAUs by ID |
| [**getTemplateExamples**](TemplatesApi.md#getTemplateExamples) | **GET** /templates/examples | Get template examples |
| [**getTemplateExamplesByOrganization**](TemplatesApi.md#getTemplateExamplesByOrganization) | **GET** /organizations/{organization}/templates/examples | Get template examples by organization |
| [**getTemplateMetadataById**](TemplatesApi.md#getTemplateMetadataById) | **GET** /templates/{template} | Get template metadata by ID |
| [**getTemplateVariablesByTemplateVersion**](TemplatesApi.md#getTemplateVariablesByTemplateVersion) | **GET** /templateversions/{templateversion}/variables | Get template variables by template version |
| [**getTemplateVersionById**](TemplatesApi.md#getTemplateVersionById) | **GET** /templateversions/{templateversion} | Get template version by ID |
| [**getTemplateVersionByOrganizationTemplateAndName**](TemplatesApi.md#getTemplateVersionByOrganizationTemplateAndName) | **GET** /organizations/{organization}/templates/{templatename}/versions/{templateversionname} | Get template version by organization, template, and name |
| [**getTemplateVersionByTemplateIdAndName**](TemplatesApi.md#getTemplateVersionByTemplateIdAndName) | **GET** /templates/{template}/versions/{templateversionname} | Get template version by template ID and name |
| [**getTemplateVersionDryRunByJobId**](TemplatesApi.md#getTemplateVersionDryRunByJobId) | **GET** /templateversions/{templateversion}/dry-run/{jobID} | Get template version dry-run by job ID |
| [**getTemplateVersionDryRunLogsByJobId**](TemplatesApi.md#getTemplateVersionDryRunLogsByJobId) | **GET** /templateversions/{templateversion}/dry-run/{jobID}/logs | Get template version dry-run logs by job ID |
| [**getTemplateVersionDryRunMatchedProvisioners**](TemplatesApi.md#getTemplateVersionDryRunMatchedProvisioners) | **GET** /templateversions/{templateversion}/dry-run/{jobID}/matched-provisioners | Get template version dry-run matched provisioners |
| [**getTemplateVersionDryRunResourcesByJobId**](TemplatesApi.md#getTemplateVersionDryRunResourcesByJobId) | **GET** /templateversions/{templateversion}/dry-run/{jobID}/resources | Get template version dry-run resources by job ID |
| [**getTemplateVersionPresets**](TemplatesApi.md#getTemplateVersionPresets) | **GET** /templateversions/{templateversion}/presets | Get template version presets |
| [**getTemplatesByOrganization**](TemplatesApi.md#getTemplatesByOrganization) | **GET** /organizations/{organization}/templates | Get templates by organization |
| [**getTemplatesByOrganizationAndTemplateName**](TemplatesApi.md#getTemplatesByOrganizationAndTemplateName) | **GET** /organizations/{organization}/templates/{templatename} | Get templates by organization and template name |
| [**listTemplateVersionsByTemplateId**](TemplatesApi.md#listTemplateVersionsByTemplateId) | **GET** /templates/{template}/versions | List template versions by template ID |
| [**openDynamicParametersWebsocketByTemplateVersion**](TemplatesApi.md#openDynamicParametersWebsocketByTemplateVersion) | **GET** /templateversions/{templateversion}/dynamic-parameters | Open dynamic parameters WebSocket by template version |
| [**patchTemplateVersionById**](TemplatesApi.md#patchTemplateVersionById) | **PATCH** /templateversions/{templateversion} | Patch template version by ID |
| [**removedGetParametersByTemplateVersion**](TemplatesApi.md#removedGetParametersByTemplateVersion) | **GET** /templateversions/{templateversion}/parameters | Removed: Get parameters by template version |
| [**removedGetSchemaByTemplateVersion**](TemplatesApi.md#removedGetSchemaByTemplateVersion) | **GET** /templateversions/{templateversion}/schema | Removed: Get schema by template version |
| [**unarchiveTemplateVersion**](TemplatesApi.md#unarchiveTemplateVersion) | **POST** /templateversions/{templateversion}/unarchive | Unarchive template version |
| [**updateActiveTemplateVersionByTemplateId**](TemplatesApi.md#updateActiveTemplateVersionByTemplateId) | **PATCH** /templates/{template}/versions | Update active template version by template ID |
| [**updateTemplateMetadataById**](TemplatesApi.md#updateTemplateMetadataById) | **PATCH** /templates/{template} | Update template metadata by ID |


<a name="archiveTemplateUnusedVersionsByTemplateId"></a>
# **archiveTemplateUnusedVersionsByTemplateId**
> codersdk.Response archiveTemplateUnusedVersionsByTemplateId(template, request)

Archive template unused versions by template id

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **template** | **UUID**| Template ID | [default to null] |
| **request** | [**codersdk.ArchiveTemplateVersionsRequest**](../Models/codersdk.ArchiveTemplateVersionsRequest.md)| Archive request | |

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="archiveTemplateVersion"></a>
# **archiveTemplateVersion**
> codersdk.Response archiveTemplateVersion(templateversion)

Archive template version

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateversion** | **UUID**| Template version ID | [default to null] |

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="cancelTemplateVersionById"></a>
# **cancelTemplateVersionById**
> codersdk.Response cancelTemplateVersionById(templateversion)

Cancel template version by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateversion** | **UUID**| Template version ID | [default to null] |

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="cancelTemplateVersionDryRunByJobId"></a>
# **cancelTemplateVersionDryRunByJobId**
> codersdk.Response cancelTemplateVersionDryRunByJobId(jobID, templateversion)

Cancel template version dry-run by job ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **jobID** | **UUID**| Job ID | [default to null] |
| **templateversion** | **UUID**| Template version ID | [default to null] |

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="createTemplateByOrganization"></a>
# **createTemplateByOrganization**
> codersdk.Template createTemplateByOrganization(organization, request)

Create template by organization

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **String**| Organization ID | [default to null] |
| **request** | [**codersdk.CreateTemplateRequest**](../Models/codersdk.CreateTemplateRequest.md)| Request body | |

### Return type

[**codersdk.Template**](../Models/codersdk.Template.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createTemplateVersionByOrganization"></a>
# **createTemplateVersionByOrganization**
> codersdk.TemplateVersion createTemplateVersionByOrganization(organization, request)

Create template version by organization

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |
| **request** | [**codersdk.CreateTemplateVersionRequest**](../Models/codersdk.CreateTemplateVersionRequest.md)| Create template version request | |

### Return type

[**codersdk.TemplateVersion**](../Models/codersdk.TemplateVersion.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createTemplateVersionDryRun"></a>
# **createTemplateVersionDryRun**
> codersdk.ProvisionerJob createTemplateVersionDryRun(templateversion, request)

Create template version dry-run

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateversion** | **UUID**| Template version ID | [default to null] |
| **request** | [**codersdk.CreateTemplateVersionDryRunRequest**](../Models/codersdk.CreateTemplateVersionDryRunRequest.md)| Dry-run request | |

### Return type

[**codersdk.ProvisionerJob**](../Models/codersdk.ProvisionerJob.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteTemplateById"></a>
# **deleteTemplateById**
> codersdk.Response deleteTemplateById(template)

Delete template by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **template** | **UUID**| Template ID | [default to null] |

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="evaluateDynamicParametersForTemplateVersion"></a>
# **evaluateDynamicParametersForTemplateVersion**
> codersdk.DynamicParametersResponse evaluateDynamicParametersForTemplateVersion(templateversion, request)

Evaluate dynamic parameters for template version

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateversion** | **UUID**| Template version ID | [default to null] |
| **request** | [**codersdk.DynamicParametersRequest**](../Models/codersdk.DynamicParametersRequest.md)| Initial parameter values | |

### Return type

[**codersdk.DynamicParametersResponse**](../Models/codersdk.DynamicParametersResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="getAllTemplates"></a>
# **getAllTemplates**
> List getAllTemplates()

Get all templates

    Returns a list of templates. By default, only non-deprecated templates are returned. To include deprecated templates, specify &#x60;deprecated:true&#x60; in the search query.

### Parameters
This endpoint does not need any parameter.

### Return type

[**List**](../Models/codersdk.Template.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getExternalAuthByTemplateVersion"></a>
# **getExternalAuthByTemplateVersion**
> List getExternalAuthByTemplateVersion(templateversion)

Get external auth by template version

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateversion** | **UUID**| Template version ID | [default to null] |

### Return type

[**List**](../Models/codersdk.TemplateVersionExternalAuth.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getLogsByTemplateVersion"></a>
# **getLogsByTemplateVersion**
> List getLogsByTemplateVersion(templateversion, before, after, follow)

Get logs by template version

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateversion** | **UUID**| Template version ID | [default to null] |
| **before** | **Integer**| Before log id | [optional] [default to null] |
| **after** | **Integer**| After log id | [optional] [default to null] |
| **follow** | **Boolean**| Follow log stream | [optional] [default to null] |

### Return type

[**List**](../Models/codersdk.ProvisionerJobLog.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getPreviousTemplateVersionByOrganizationTemplateAndName"></a>
# **getPreviousTemplateVersionByOrganizationTemplateAndName**
> codersdk.TemplateVersion getPreviousTemplateVersionByOrganizationTemplateAndName(organization, templatename, templateversionname)

Get previous template version by organization, template, and name

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |
| **templatename** | **String**| Template name | [default to null] |
| **templateversionname** | **String**| Template version name | [default to null] |

### Return type

[**codersdk.TemplateVersion**](../Models/codersdk.TemplateVersion.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getResourcesByTemplateVersion"></a>
# **getResourcesByTemplateVersion**
> List getResourcesByTemplateVersion(templateversion)

Get resources by template version

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateversion** | **UUID**| Template version ID | [default to null] |

### Return type

[**List**](../Models/codersdk.WorkspaceResource.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getRichParametersByTemplateVersion"></a>
# **getRichParametersByTemplateVersion**
> List getRichParametersByTemplateVersion(templateversion)

Get rich parameters by template version

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateversion** | **UUID**| Template version ID | [default to null] |

### Return type

[**List**](../Models/codersdk.TemplateVersionParameter.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTemplateDausById"></a>
# **getTemplateDausById**
> codersdk.DAUsResponse getTemplateDausById(template)

Get template DAUs by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **template** | **UUID**| Template ID | [default to null] |

### Return type

[**codersdk.DAUsResponse**](../Models/codersdk.DAUsResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTemplateExamples"></a>
# **getTemplateExamples**
> List getTemplateExamples()

Get template examples

### Parameters
This endpoint does not need any parameter.

### Return type

[**List**](../Models/codersdk.TemplateExample.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTemplateExamplesByOrganization"></a>
# **getTemplateExamplesByOrganization**
> List getTemplateExamplesByOrganization(organization)

Get template examples by organization

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |

### Return type

[**List**](../Models/codersdk.TemplateExample.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTemplateMetadataById"></a>
# **getTemplateMetadataById**
> codersdk.Template getTemplateMetadataById(template)

Get template metadata by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **template** | **UUID**| Template ID | [default to null] |

### Return type

[**codersdk.Template**](../Models/codersdk.Template.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTemplateVariablesByTemplateVersion"></a>
# **getTemplateVariablesByTemplateVersion**
> List getTemplateVariablesByTemplateVersion(templateversion)

Get template variables by template version

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateversion** | **UUID**| Template version ID | [default to null] |

### Return type

[**List**](../Models/codersdk.TemplateVersionVariable.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTemplateVersionById"></a>
# **getTemplateVersionById**
> codersdk.TemplateVersion getTemplateVersionById(templateversion)

Get template version by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateversion** | **UUID**| Template version ID | [default to null] |

### Return type

[**codersdk.TemplateVersion**](../Models/codersdk.TemplateVersion.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTemplateVersionByOrganizationTemplateAndName"></a>
# **getTemplateVersionByOrganizationTemplateAndName**
> codersdk.TemplateVersion getTemplateVersionByOrganizationTemplateAndName(organization, templatename, templateversionname)

Get template version by organization, template, and name

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |
| **templatename** | **String**| Template name | [default to null] |
| **templateversionname** | **String**| Template version name | [default to null] |

### Return type

[**codersdk.TemplateVersion**](../Models/codersdk.TemplateVersion.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTemplateVersionByTemplateIdAndName"></a>
# **getTemplateVersionByTemplateIdAndName**
> List getTemplateVersionByTemplateIdAndName(template, templateversionname)

Get template version by template ID and name

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **template** | **UUID**| Template ID | [default to null] |
| **templateversionname** | **String**| Template version name | [default to null] |

### Return type

[**List**](../Models/codersdk.TemplateVersion.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTemplateVersionDryRunByJobId"></a>
# **getTemplateVersionDryRunByJobId**
> codersdk.ProvisionerJob getTemplateVersionDryRunByJobId(templateversion, jobID)

Get template version dry-run by job ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateversion** | **UUID**| Template version ID | [default to null] |
| **jobID** | **UUID**| Job ID | [default to null] |

### Return type

[**codersdk.ProvisionerJob**](../Models/codersdk.ProvisionerJob.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTemplateVersionDryRunLogsByJobId"></a>
# **getTemplateVersionDryRunLogsByJobId**
> List getTemplateVersionDryRunLogsByJobId(templateversion, jobID, before, after, follow)

Get template version dry-run logs by job ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateversion** | **UUID**| Template version ID | [default to null] |
| **jobID** | **UUID**| Job ID | [default to null] |
| **before** | **Integer**| Before Unix timestamp | [optional] [default to null] |
| **after** | **Integer**| After Unix timestamp | [optional] [default to null] |
| **follow** | **Boolean**| Follow log stream | [optional] [default to null] |

### Return type

[**List**](../Models/codersdk.ProvisionerJobLog.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTemplateVersionDryRunMatchedProvisioners"></a>
# **getTemplateVersionDryRunMatchedProvisioners**
> codersdk.MatchedProvisioners getTemplateVersionDryRunMatchedProvisioners(templateversion, jobID)

Get template version dry-run matched provisioners

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateversion** | **UUID**| Template version ID | [default to null] |
| **jobID** | **UUID**| Job ID | [default to null] |

### Return type

[**codersdk.MatchedProvisioners**](../Models/codersdk.MatchedProvisioners.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTemplateVersionDryRunResourcesByJobId"></a>
# **getTemplateVersionDryRunResourcesByJobId**
> List getTemplateVersionDryRunResourcesByJobId(templateversion, jobID)

Get template version dry-run resources by job ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateversion** | **UUID**| Template version ID | [default to null] |
| **jobID** | **UUID**| Job ID | [default to null] |

### Return type

[**List**](../Models/codersdk.WorkspaceResource.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTemplateVersionPresets"></a>
# **getTemplateVersionPresets**
> List getTemplateVersionPresets(templateversion)

Get template version presets

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateversion** | **UUID**| Template version ID | [default to null] |

### Return type

[**List**](../Models/codersdk.Preset.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTemplatesByOrganization"></a>
# **getTemplatesByOrganization**
> List getTemplatesByOrganization(organization)

Get templates by organization

    Returns a list of templates for the specified organization. By default, only non-deprecated templates are returned. To include deprecated templates, specify &#x60;deprecated:true&#x60; in the search query.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |

### Return type

[**List**](../Models/codersdk.Template.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTemplatesByOrganizationAndTemplateName"></a>
# **getTemplatesByOrganizationAndTemplateName**
> codersdk.Template getTemplatesByOrganizationAndTemplateName(organization, templatename)

Get templates by organization and template name

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |
| **templatename** | **String**| Template name | [default to null] |

### Return type

[**codersdk.Template**](../Models/codersdk.Template.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listTemplateVersionsByTemplateId"></a>
# **listTemplateVersionsByTemplateId**
> List listTemplateVersionsByTemplateId(template, after\_id, include\_archived, limit, offset)

List template versions by template ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **template** | **UUID**| Template ID | [default to null] |
| **after\_id** | **UUID**| After ID | [optional] [default to null] |
| **include\_archived** | **Boolean**| Include archived versions in the list | [optional] [default to null] |
| **limit** | **Integer**| Page limit | [optional] [default to null] |
| **offset** | **Integer**| Page offset | [optional] [default to null] |

### Return type

[**List**](../Models/codersdk.TemplateVersion.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="openDynamicParametersWebsocketByTemplateVersion"></a>
# **openDynamicParametersWebsocketByTemplateVersion**
> openDynamicParametersWebsocketByTemplateVersion(templateversion)

Open dynamic parameters WebSocket by template version

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateversion** | **UUID**| Template version ID | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="patchTemplateVersionById"></a>
# **patchTemplateVersionById**
> codersdk.TemplateVersion patchTemplateVersionById(templateversion, request)

Patch template version by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateversion** | **UUID**| Template version ID | [default to null] |
| **request** | [**codersdk.PatchTemplateVersionRequest**](../Models/codersdk.PatchTemplateVersionRequest.md)| Patch template version request | |

### Return type

[**codersdk.TemplateVersion**](../Models/codersdk.TemplateVersion.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="removedGetParametersByTemplateVersion"></a>
# **removedGetParametersByTemplateVersion**
> removedGetParametersByTemplateVersion(templateversion)

Removed: Get parameters by template version

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateversion** | **UUID**| Template version ID | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="removedGetSchemaByTemplateVersion"></a>
# **removedGetSchemaByTemplateVersion**
> removedGetSchemaByTemplateVersion(templateversion)

Removed: Get schema by template version

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateversion** | **UUID**| Template version ID | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="unarchiveTemplateVersion"></a>
# **unarchiveTemplateVersion**
> codersdk.Response unarchiveTemplateVersion(templateversion)

Unarchive template version

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateversion** | **UUID**| Template version ID | [default to null] |

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateActiveTemplateVersionByTemplateId"></a>
# **updateActiveTemplateVersionByTemplateId**
> codersdk.Response updateActiveTemplateVersionByTemplateId(template, request)

Update active template version by template ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **template** | **UUID**| Template ID | [default to null] |
| **request** | [**codersdk.UpdateActiveTemplateVersion**](../Models/codersdk.UpdateActiveTemplateVersion.md)| Modified template version | |

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateTemplateMetadataById"></a>
# **updateTemplateMetadataById**
> codersdk.Template updateTemplateMetadataById(template)

Update template metadata by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **template** | **UUID**| Template ID | [default to null] |

### Return type

[**codersdk.Template**](../Models/codersdk.Template.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


