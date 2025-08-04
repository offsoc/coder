# Notifications


All URIs are relative to */api/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createUserWebpushSubscription**](NotificationsApi.md#createUserWebpushSubscription) | **POST** /users/{user}/webpush/subscription | Create user webpush subscription |
| [**deleteUserWebpushSubscription**](NotificationsApi.md#deleteUserWebpushSubscription) | **DELETE** /users/{user}/webpush/subscription | Delete user webpush subscription |
| [**getNotificationDispatchMethods**](NotificationsApi.md#getNotificationDispatchMethods) | **GET** /notifications/dispatch-methods | Get notification dispatch methods |
| [**getNotificationsSettings**](NotificationsApi.md#getNotificationsSettings) | **GET** /notifications/settings | Get notifications settings |
| [**getSystemNotificationTemplates**](NotificationsApi.md#getSystemNotificationTemplates) | **GET** /notifications/templates/system | Get system notification templates |
| [**getUserNotificationPreferences**](NotificationsApi.md#getUserNotificationPreferences) | **GET** /users/{user}/notifications/preferences | Get user notification preferences |
| [**listInboxNotifications**](NotificationsApi.md#listInboxNotifications) | **GET** /notifications/inbox | List inbox notifications |
| [**markAllUnreadNotificationsAsRead**](NotificationsApi.md#markAllUnreadNotificationsAsRead) | **PUT** /notifications/inbox/mark-all-as-read | Mark all unread notifications as read |
| [**sendATestNotification**](NotificationsApi.md#sendATestNotification) | **POST** /notifications/test | Send a test notification |
| [**sendATestPushNotification**](NotificationsApi.md#sendATestPushNotification) | **POST** /users/{user}/webpush/test | Send a test push notification |
| [**updateNotificationsSettings**](NotificationsApi.md#updateNotificationsSettings) | **PUT** /notifications/settings | Update notifications settings |
| [**updateReadStatusOfANotification**](NotificationsApi.md#updateReadStatusOfANotification) | **PUT** /notifications/inbox/{id}/read-status | Update read status of a notification |
| [**updateUserNotificationPreferences**](NotificationsApi.md#updateUserNotificationPreferences) | **PUT** /users/{user}/notifications/preferences | Update user notification preferences |
| [**watchForNewInboxNotifications**](NotificationsApi.md#watchForNewInboxNotifications) | **GET** /notifications/inbox/watch | Watch for new inbox notifications |


<a name="createUserWebpushSubscription"></a>
# **createUserWebpushSubscription**
> createUserWebpushSubscription(user, request)

Create user webpush subscription

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |
| **request** | [**codersdk.WebpushSubscription**](../Models/codersdk.WebpushSubscription.md)| Webpush subscription | |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

<a name="deleteUserWebpushSubscription"></a>
# **deleteUserWebpushSubscription**
> deleteUserWebpushSubscription(user, request)

Delete user webpush subscription

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |
| **request** | [**codersdk.DeleteWebpushSubscription**](../Models/codersdk.DeleteWebpushSubscription.md)| Webpush subscription | |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

<a name="getNotificationDispatchMethods"></a>
# **getNotificationDispatchMethods**
> List getNotificationDispatchMethods()

Get notification dispatch methods

### Parameters
This endpoint does not need any parameter.

### Return type

[**List**](../Models/codersdk.NotificationMethodsResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getNotificationsSettings"></a>
# **getNotificationsSettings**
> codersdk.NotificationsSettings getNotificationsSettings()

Get notifications settings

### Parameters
This endpoint does not need any parameter.

### Return type

[**codersdk.NotificationsSettings**](../Models/codersdk.NotificationsSettings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getSystemNotificationTemplates"></a>
# **getSystemNotificationTemplates**
> List getSystemNotificationTemplates()

Get system notification templates

### Parameters
This endpoint does not need any parameter.

### Return type

[**List**](../Models/codersdk.NotificationTemplate.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserNotificationPreferences"></a>
# **getUserNotificationPreferences**
> List getUserNotificationPreferences(user)

Get user notification preferences

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |

### Return type

[**List**](../Models/codersdk.NotificationPreference.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listInboxNotifications"></a>
# **listInboxNotifications**
> codersdk.ListInboxNotificationsResponse listInboxNotifications(targets, templates, read\_status, starting\_before)

List inbox notifications

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **targets** | **String**| Comma-separated list of target IDs to filter notifications | [optional] [default to null] |
| **templates** | **String**| Comma-separated list of template IDs to filter notifications | [optional] [default to null] |
| **read\_status** | **String**| Filter notifications by read status. Possible values: read, unread, all | [optional] [default to null] |
| **starting\_before** | **UUID**| ID of the last notification from the current page. Notifications returned will be older than the associated one | [optional] [default to null] |

### Return type

[**codersdk.ListInboxNotificationsResponse**](../Models/codersdk.ListInboxNotificationsResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="markAllUnreadNotificationsAsRead"></a>
# **markAllUnreadNotificationsAsRead**
> markAllUnreadNotificationsAsRead()

Mark all unread notifications as read

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="sendATestNotification"></a>
# **sendATestNotification**
> sendATestNotification()

Send a test notification

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="sendATestPushNotification"></a>
# **sendATestPushNotification**
> sendATestPushNotification(user)

Send a test push notification

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

<a name="updateNotificationsSettings"></a>
# **updateNotificationsSettings**
> codersdk.NotificationsSettings updateNotificationsSettings(request)

Update notifications settings

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **request** | [**codersdk.NotificationsSettings**](../Models/codersdk.NotificationsSettings.md)| Notifications settings request | |

### Return type

[**codersdk.NotificationsSettings**](../Models/codersdk.NotificationsSettings.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateReadStatusOfANotification"></a>
# **updateReadStatusOfANotification**
> codersdk.Response updateReadStatusOfANotification(id)

Update read status of a notification

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | **String**| id of the notification | [default to null] |

### Return type

[**codersdk.Response**](../Models/codersdk.Response.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateUserNotificationPreferences"></a>
# **updateUserNotificationPreferences**
> List updateUserNotificationPreferences(user, request)

Update user notification preferences

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| User ID, name, or me | [default to null] |
| **request** | [**codersdk.UpdateUserNotificationPreferences**](../Models/codersdk.UpdateUserNotificationPreferences.md)| Preferences | |

### Return type

[**List**](../Models/codersdk.NotificationPreference.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="watchForNewInboxNotifications"></a>
# **watchForNewInboxNotifications**
> codersdk.GetInboxNotificationResponse watchForNewInboxNotifications(targets, templates, read\_status, format)

Watch for new inbox notifications

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **targets** | **String**| Comma-separated list of target IDs to filter notifications | [optional] [default to null] |
| **templates** | **String**| Comma-separated list of template IDs to filter notifications | [optional] [default to null] |
| **read\_status** | **String**| Filter notifications by read status. Possible values: read, unread, all | [optional] [default to null] |
| **format** | **String**| Define the output format for notifications title and body. | [optional] [default to null] [enum: plaintext, markdown] |

### Return type

[**codersdk.GetInboxNotificationResponse**](../Models/codersdk.GetInboxNotificationResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


