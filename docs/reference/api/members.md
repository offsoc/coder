# Members


All URIs are relative to */api/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**addOrganizationMember**](MembersApi.md#addOrganizationMember) | **POST** /organizations/{organization}/members/{user} | Add organization member |
| [**assignRoleToOrganizationMember**](MembersApi.md#assignRoleToOrganizationMember) | **PUT** /organizations/{organization}/members/{user}/roles | Assign role to organization member |
| [**deleteACustomOrganizationRole**](MembersApi.md#deleteACustomOrganizationRole) | **DELETE** /organizations/{organization}/members/roles/{roleName} | Delete a custom organization role |
| [**getMemberRolesByOrganization**](MembersApi.md#getMemberRolesByOrganization) | **GET** /organizations/{organization}/members/roles | Get member roles by organization |
| [**getSiteMemberRoles**](MembersApi.md#getSiteMemberRoles) | **GET** /users/roles | Get site member roles |
| [**insertACustomOrganizationRole**](MembersApi.md#insertACustomOrganizationRole) | **POST** /organizations/{organization}/members/roles | Insert a custom organization role |
| [**listOrganizationMembers**](MembersApi.md#listOrganizationMembers) | **GET** /organizations/{organization}/members | List organization members |
| [**paginatedOrganizationMembers**](MembersApi.md#paginatedOrganizationMembers) | **GET** /organizations/{organization}/paginated-members | Paginated organization members |
| [**removeOrganizationMember**](MembersApi.md#removeOrganizationMember) | **DELETE** /organizations/{organization}/members/{user} | Remove organization member |
| [**upsertACustomOrganizationRole**](MembersApi.md#upsertACustomOrganizationRole) | **PUT** /organizations/{organization}/members/roles | Upsert a custom organization role |


<a name="addOrganizationMember"></a>
# **addOrganizationMember**
> codersdk.OrganizationMember addOrganizationMember(organization, user)

Add organization member

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **String**| Organization ID | [default to null] |
| **user** | **String**| User ID, name, or me | [default to null] |

### Return type

[**codersdk.OrganizationMember**](../Models/codersdk.OrganizationMember.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="assignRoleToOrganizationMember"></a>
# **assignRoleToOrganizationMember**
> codersdk.OrganizationMember assignRoleToOrganizationMember(organization, user, request)

Assign role to organization member

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **String**| Organization ID | [default to null] |
| **user** | **String**| User ID, name, or me | [default to null] |
| **request** | [**codersdk.UpdateRoles**](../Models/codersdk.UpdateRoles.md)| Update roles request | |

### Return type

[**codersdk.OrganizationMember**](../Models/codersdk.OrganizationMember.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteACustomOrganizationRole"></a>
# **deleteACustomOrganizationRole**
> List deleteACustomOrganizationRole(organization, roleName)

Delete a custom organization role

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |
| **roleName** | **String**| Role name | [default to null] |

### Return type

[**List**](../Models/codersdk.Role.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getMemberRolesByOrganization"></a>
# **getMemberRolesByOrganization**
> List getMemberRolesByOrganization(organization)

Get member roles by organization

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |

### Return type

[**List**](../Models/codersdk.AssignableRoles.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getSiteMemberRoles"></a>
# **getSiteMemberRoles**
> List getSiteMemberRoles()

Get site member roles

### Parameters
This endpoint does not need any parameter.

### Return type

[**List**](../Models/codersdk.AssignableRoles.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="insertACustomOrganizationRole"></a>
# **insertACustomOrganizationRole**
> List insertACustomOrganizationRole(organization, request)

Insert a custom organization role

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |
| **request** | [**codersdk.CustomRoleRequest**](../Models/codersdk.CustomRoleRequest.md)| Insert role request | |

### Return type

[**List**](../Models/codersdk.Role.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="listOrganizationMembers"></a>
# **listOrganizationMembers**
> List listOrganizationMembers(organization)

List organization members

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **String**| Organization ID | [default to null] |

### Return type

[**List**](../Models/codersdk.OrganizationMemberWithUserData.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="paginatedOrganizationMembers"></a>
# **paginatedOrganizationMembers**
> List paginatedOrganizationMembers(organization, limit, offset)

Paginated organization members

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **String**| Organization ID | [default to null] |
| **limit** | **Integer**| Page limit, if 0 returns all members | [optional] [default to null] |
| **offset** | **Integer**| Page offset | [optional] [default to null] |

### Return type

[**List**](../Models/codersdk.PaginatedMembersResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="removeOrganizationMember"></a>
# **removeOrganizationMember**
> removeOrganizationMember(organization, user)

Remove organization member

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **String**| Organization ID | [default to null] |
| **user** | **String**| User ID, name, or me | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="upsertACustomOrganizationRole"></a>
# **upsertACustomOrganizationRole**
> List upsertACustomOrganizationRole(organization, request)

Upsert a custom organization role

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **organization** | **UUID**| Organization ID | [default to null] |
| **request** | [**codersdk.CustomRoleRequest**](../Models/codersdk.CustomRoleRequest.md)| Upsert role request | |

### Return type

[**List**](../Models/codersdk.Role.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


