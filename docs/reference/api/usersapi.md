# UsersApi


All URIs are relative to */api/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**activateUserAccount**](UsersApi.md#activateUserAccount) | **PUT** /users/{user}/status/activate | Activate user account |
| [**assignRoleToUser**](UsersApi.md#assignRoleToUser) | **PUT** /users/{user}/roles | Assign role to user |
| [**checkInitialUserCreated**](UsersApi.md#checkInitialUserCreated) | **GET** /users/first | Check initial user created |
| [**createInitialUser**](UsersApi.md#createInitialUser) | **POST** /users/first | Create initial user |
| [**createNewSessionKey**](UsersApi.md#createNewSessionKey) | **POST** /users/{user}/keys | Create new session key |
| [**createNewUser**](UsersApi.md#createNewUser) | **POST** /users | Create new user |
| [**createTokenApiKey**](UsersApi.md#createTokenApiKey) | **POST** /users/{user}/keys/tokens | Create token API key |
| [**deleteApiKey**](UsersApi.md#deleteApiKey) | **DELETE** /users/{user}/keys/{keyid} | Delete API key |
| [**deleteUser**](UsersApi.md#deleteUser) | **DELETE** /users/{user} | Delete user |
| [**getApiKeyById**](UsersApi.md#getApiKeyById) | **GET** /users/{user}/keys/{keyid} | Get API key by ID |
| [**getApiKeyByTokenName**](UsersApi.md#getApiKeyByTokenName) | **GET** /users/{user}/keys/tokens/{keyname} | Get API key by token name |
| [**getAuthenticationMethods**](UsersApi.md#getAuthenticationMethods) | **GET** /users/authmethods | Get authentication methods |
| [**getAutofillBuildParametersForUser**](UsersApi.md#getAutofillBuildParametersForUser) | **GET** /users/{user}/autofill-parameters | Get autofill build parameters for user |
| [**getGithubDeviceAuth**](UsersApi.md#getGithubDeviceAuth) | **GET** /users/oauth2/github/device | Get Github device auth. |
| [**getOrganizationByUserAndOrganizationName**](UsersApi.md#getOrganizationByUserAndOrganizationName) | **GET** /users/{user}/organizations/{organizationname} | Get organization by user and organization name |
| [**getOrganizationsByUser**](UsersApi.md#getOrganizationsByUser) | **GET** /users/{user}/organizations | Get organizations by user |
| [**getUserAppearanceSettings**](UsersApi.md#getUserAppearanceSettings) | **GET** /users/{user}/appearance | Get user appearance settings |
| [**getUserByName**](UsersApi.md#getUserByName) | **GET** /users/{user} | Get user by name |
| [**getUserGitSshKey**](UsersApi.md#getUserGitSshKey) | **GET** /users/{user}/gitsshkey | Get user Git SSH key |
| [**getUserLoginType**](UsersApi.md#getUserLoginType) | **GET** /users/{user}/login-type | Get user login type |
| [**getUserRoles**](UsersApi.md#getUserRoles) | **GET** /users/{user}/roles | Get user roles |
| [**getUserTokens**](UsersApi.md#getUserTokens) | **GET** /users/{user}/keys/tokens | Get user tokens |
| [**getUsers**](UsersApi.md#getUsers) | **GET** /users | Get users |
| [**logOutUser**](UsersApi.md#logOutUser) | **POST** /users/logout | Log out user |
| [**oauth20GithubCallback**](UsersApi.md#oauth20GithubCallback) | **GET** /users/oauth2/github/callback | OAuth 2.0 GitHub Callback |
| [**openidConnectCallback**](UsersApi.md#openidConnectCallback) | **GET** /users/oidc/callback | OpenID Connect Callback |
| [**regenerateUserSshKey**](UsersApi.md#regenerateUserSshKey) | **PUT** /users/{user}/gitsshkey | Regenerate user SSH key |
| [**suspendUserAccount**](UsersApi.md#suspendUserAccount) | **PUT** /users/{user}/status/suspend | Suspend user account |
| [**updateUserAppearanceSettings**](UsersApi.md#updateUserAppearanceSettings) | **PUT** /users/{user}/appearance | Update user appearance settings |
| [**updateUserPassword**](UsersApi.md#updateUserPassword) | **PUT** /users/{user}/password | Update user password |
| [**updateUserProfile**](UsersApi.md#updateUserProfile) | **PUT** /users/{user}/profile | Update user profile |


<a name="activateUserAccount"></a>
# **activateUserAccount**
> codersdk.User activateUserAccount(user)

Activate user account

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |

### Return type

[**codersdk.User**](../Models/codersdk.User.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="assignRoleToUser"></a>
# **assignRoleToUser**
> codersdk.User assignRoleToUser(user, request)

Assign role to user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |
| **request** | [**codersdk.UpdateRoles**](../Models/codersdk.UpdateRoles.md)| Update roles request | |

### Return type

[**codersdk.User**](../Models/codersdk.User.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="checkInitialUserCreated"></a>
# **checkInitialUserCreated**
> codersdk.Response checkInitialUserCreated()

Check initial user created

### Parameters
This endpoint does not need any parameter.

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="createInitialUser"></a>
# **createInitialUser**
> codersdk.CreateFirstUserResponse createInitialUser(request)

Create initial user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**codersdk.CreateFirstUserRequest**](../Models/codersdk.CreateFirstUserRequest.md)| First user request | |

### Return type

[**codersdk.CreateFirstUserResponse**](../Models/codersdk.CreateFirstUserResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createNewSessionKey"></a>
# **createNewSessionKey**
> codersdk.GenerateAPIKeyResponse createNewSessionKey(user)

Create new session key

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |

### Return type

[**codersdk.GenerateAPIKeyResponse**](../Models/codersdk.GenerateAPIKeyResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="createNewUser"></a>
# **createNewUser**
> codersdk.User createNewUser(request)

Create new user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**codersdk.CreateUserRequestWithOrgs**](../Models/codersdk.CreateUserRequestWithOrgs.md)| Create user request | |

### Return type

[**codersdk.User**](../Models/codersdk.User.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createTokenApiKey"></a>
# **createTokenApiKey**
> codersdk.GenerateAPIKeyResponse createTokenApiKey(user, request)

Create token API key

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |
| **request** | [**codersdk.CreateTokenRequest**](../Models/codersdk.CreateTokenRequest.md)| Create token request | |

### Return type

[**codersdk.GenerateAPIKeyResponse**](../Models/codersdk.GenerateAPIKeyResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteApiKey"></a>
# **deleteApiKey**
> deleteApiKey(user, keyid)

Delete API key

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |
| **keyid** | **UUID**| Key ID | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="deleteUser"></a>
# **deleteUser**
> deleteUser(user)

Delete user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="getApiKeyById"></a>
# **getApiKeyById**
> codersdk.APIKey getApiKeyById(user, keyid)

Get API key by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |
| **keyid** | **UUID**| Key ID | [default to null] |

### Return type

[**codersdk.APIKey**](../Models/codersdk.APIKey.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getApiKeyByTokenName"></a>
# **getApiKeyByTokenName**
> codersdk.APIKey getApiKeyByTokenName(user, keyname)

Get API key by token name

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |
| **keyname** | **String**| Key Name | [default to null] |

### Return type

[**codersdk.APIKey**](../Models/codersdk.APIKey.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAuthenticationMethods"></a>
# **getAuthenticationMethods**
> codersdk.AuthMethods getAuthenticationMethods()

Get authentication methods

### Parameters
This endpoint does not need any parameter.

### Return type

[**codersdk.AuthMethods**](../Models/codersdk.AuthMethods.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAutofillBuildParametersForUser"></a>
# **getAutofillBuildParametersForUser**
> List getAutofillBuildParametersForUser(user, template\_id)

Get autofill build parameters for user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, username, or me | [default to null] |
| **template\_id** | **String**| Template ID | [default to null] |

### Return type

[**List**](../Models/codersdk.UserParameter.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getGithubDeviceAuth"></a>
# **getGithubDeviceAuth**
> codersdk.ExternalAuthDevice getGithubDeviceAuth()

Get Github device auth.

### Parameters
This endpoint does not need any parameter.

### Return type

[**codersdk.ExternalAuthDevice**](../Models/codersdk.ExternalAuthDevice.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getOrganizationByUserAndOrganizationName"></a>
# **getOrganizationByUserAndOrganizationName**
> codersdk.Organization getOrganizationByUserAndOrganizationName(user, organizationname)

Get organization by user and organization name

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |
| **organizationname** | **String**| Organization name | [default to null] |

### Return type

[**codersdk.Organization**](../Models/codersdk.Organization.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getOrganizationsByUser"></a>
# **getOrganizationsByUser**
> List getOrganizationsByUser(user)

Get organizations by user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |

### Return type

[**List**](../Models/codersdk.Organization.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserAppearanceSettings"></a>
# **getUserAppearanceSettings**
> codersdk.UserAppearanceSettings getUserAppearanceSettings(user)

Get user appearance settings

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |

### Return type

[**codersdk.UserAppearanceSettings**](../Models/codersdk.UserAppearanceSettings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserByName"></a>
# **getUserByName**
> codersdk.User getUserByName(user)

Get user by name

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, username, or me | [default to null] |

### Return type

[**codersdk.User**](../Models/codersdk.User.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserGitSshKey"></a>
# **getUserGitSshKey**
> codersdk.GitSSHKey getUserGitSshKey(user)

Get user Git SSH key

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |

### Return type

[**codersdk.GitSSHKey**](../Models/codersdk.GitSSHKey.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserLoginType"></a>
# **getUserLoginType**
> codersdk.UserLoginType getUserLoginType(user)

Get user login type

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |

### Return type

[**codersdk.UserLoginType**](../Models/codersdk.UserLoginType.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserRoles"></a>
# **getUserRoles**
> codersdk.User getUserRoles(user)

Get user roles

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |

### Return type

[**codersdk.User**](../Models/codersdk.User.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserTokens"></a>
# **getUserTokens**
> List getUserTokens(user)

Get user tokens

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |

### Return type

[**List**](../Models/codersdk.APIKey.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUsers"></a>
# **getUsers**
> codersdk.GetUsersResponse getUsers(q, after\_id, limit, offset)

Get users

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **q** | **String**| Search query | [optional] [default to null] |
| **after\_id** | **UUID**| After ID | [optional] [default to null] |
| **limit** | **Integer**| Page limit | [optional] [default to null] |
| **offset** | **Integer**| Page offset | [optional] [default to null] |

### Return type

[**codersdk.GetUsersResponse**](../Models/codersdk.GetUsersResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="logOutUser"></a>
# **logOutUser**
> codersdk.Response logOutUser()

Log out user

### Parameters
This endpoint does not need any parameter.

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="oauth20GithubCallback"></a>
# **oauth20GithubCallback**
> oauth20GithubCallback()

OAuth 2.0 GitHub Callback

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="openidConnectCallback"></a>
# **openidConnectCallback**
> openidConnectCallback()

OpenID Connect Callback

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="regenerateUserSshKey"></a>
# **regenerateUserSshKey**
> codersdk.GitSSHKey regenerateUserSshKey(user)

Regenerate user SSH key

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |

### Return type

[**codersdk.GitSSHKey**](../Models/codersdk.GitSSHKey.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="suspendUserAccount"></a>
# **suspendUserAccount**
> codersdk.User suspendUserAccount(user)

Suspend user account

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |

### Return type

[**codersdk.User**](../Models/codersdk.User.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateUserAppearanceSettings"></a>
# **updateUserAppearanceSettings**
> codersdk.UserAppearanceSettings updateUserAppearanceSettings(user, request)

Update user appearance settings

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |
| **request** | [**codersdk.UpdateUserAppearanceSettingsRequest**](../Models/codersdk.UpdateUserAppearanceSettingsRequest.md)| New appearance settings | |

### Return type

[**codersdk.UserAppearanceSettings**](../Models/codersdk.UserAppearanceSettings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateUserPassword"></a>
# **updateUserPassword**
> updateUserPassword(user, request)

Update user password

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |
| **request** | [**codersdk.UpdateUserPasswordRequest**](../Models/codersdk.UpdateUserPasswordRequest.md)| Update password request | |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

<a name="updateUserProfile"></a>
# **updateUserProfile**
> codersdk.User updateUserProfile(user, request)

Update user profile

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |
| **request** | [**codersdk.UpdateUserProfileRequest**](../Models/codersdk.UpdateUserProfileRequest.md)| Updated profile | |

### Return type

[**codersdk.User**](../Models/codersdk.User.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


