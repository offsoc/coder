# FilesApi


All URIs are relative to */api/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getFileById**](FilesApi.md#getFileById) | **GET** /files/{fileID} | Get file by ID |
| [**uploadFile**](FilesApi.md#uploadFile) | **POST** /files | Upload file |


<a name="getFileById"></a>
# **getFileById**
> getFileById(fileID)

Get file by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **fileID** | **UUID**| File ID | [default to null] |

### Return type

null (empty response body)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="uploadFile"></a>
# **uploadFile**
> codersdk.UploadResponse uploadFile(Content-Type, upload\_file\_request)

Upload file

    Swagger notice: Swagger 2.0 doesn&#39;t support file upload with a &#x60;content-type&#x60; different than &#x60;application/x-www-form-urlencoded&#x60;.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **Content-Type** | **String**| Content-Type must be &#x60;application/x-tar&#x60; or &#x60;application/zip&#x60; | [default to application/x-tar] |
| **upload\_file\_request** | [**upload_file_request**](../Models/upload_file_request.md)|  | |

### Return type

[**codersdk.UploadResponse**](../Models/codersdk.UploadResponse.md)

### Authorization

[CoderSessionToken](../README.md#CoderSessionToken)

### HTTP request headers

- **Content-Type**: application/x-tar
- **Accept**: application/json


