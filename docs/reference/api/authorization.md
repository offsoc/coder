# Authorization


All URIs are relative to */api/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**changePasswordWithAOneTimePasscode**](AuthorizationApi.md#changePasswordWithAOneTimePasscode) | **POST** /users/otp/change-password | Change password with a one-time passcode |
| [**checkAuthorization**](AuthorizationApi.md#checkAuthorization) | **POST** /authcheck | Check authorization |
| [**convertUserFromPasswordToOauthAuthentication**](AuthorizationApi.md#convertUserFromPasswordToOauthAuthentication) | **POST** /users/{user}/convert-login | Convert user from password to oauth authentication |
| [**logInUser**](AuthorizationApi.md#logInUser) | **POST** /users/login | Log in user |
| [**requestOneTimePasscode**](AuthorizationApi.md#requestOneTimePasscode) | **POST** /users/otp/request | Request one-time passcode |
| [**validateUserPassword**](AuthorizationApi.md#validateUserPassword) | **POST** /users/validate-password | Validate user password |


<a name="changePasswordWithAOneTimePasscode"></a>
# **changePasswordWithAOneTimePasscode**
> changePasswordWithAOneTimePasscode(request)

Change password with a one-time passcode

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**codersdk.ChangePasswordWithOneTimePasscodeRequest**](../Models/codersdk.ChangePasswordWithOneTimePasscodeRequest.md)| Change password request | |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

<a name="checkAuthorization"></a>
# **checkAuthorization**
> Map checkAuthorization(request)

Check authorization

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**codersdk.AuthorizationRequest**](../Models/codersdk.AuthorizationRequest.md)| Authorization request | |

### Return type

**Map**

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="convertUserFromPasswordToOauthAuthentication"></a>
# **convertUserFromPasswordToOauthAuthentication**
> codersdk.OAuthConversionResponse convertUserFromPasswordToOauthAuthentication(user, request)

Convert user from password to oauth authentication

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |
| **request** | [**codersdk.ConvertLoginRequest**](../Models/codersdk.ConvertLoginRequest.md)| Convert request | |

### Return type

[**codersdk.OAuthConversionResponse**](../Models/codersdk.OAuthConversionResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="logInUser"></a>
# **logInUser**
> codersdk.LoginWithPasswordResponse logInUser(request)

Log in user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**codersdk.LoginWithPasswordRequest**](../Models/codersdk.LoginWithPasswordRequest.md)| Login request | |

### Return type

[**codersdk.LoginWithPasswordResponse**](../Models/codersdk.LoginWithPasswordResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="requestOneTimePasscode"></a>
# **requestOneTimePasscode**
> requestOneTimePasscode(request)

Request one-time passcode

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**codersdk.RequestOneTimePasscodeRequest**](../Models/codersdk.RequestOneTimePasscodeRequest.md)| One-time passcode request | |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

<a name="validateUserPassword"></a>
# **validateUserPassword**
> codersdk.ValidateUserPasswordResponse validateUserPassword(request)

Validate user password

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**codersdk.ValidateUserPasswordRequest**](../Models/codersdk.ValidateUserPasswordRequest.md)| Validate user password request | |

### Return type

[**codersdk.ValidateUserPasswordResponse**](../Models/codersdk.ValidateUserPasswordResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


