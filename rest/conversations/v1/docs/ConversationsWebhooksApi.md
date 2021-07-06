# ConversationsWebhooksApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConversationScopedWebhook**](ConversationsWebhooksApi.md#CreateConversationScopedWebhook) | **Post** /v1/Conversations/{ConversationSid}/Webhooks | 
[**DeleteConversationScopedWebhook**](ConversationsWebhooksApi.md#DeleteConversationScopedWebhook) | **Delete** /v1/Conversations/{ConversationSid}/Webhooks/{Sid} | 
[**FetchConversationScopedWebhook**](ConversationsWebhooksApi.md#FetchConversationScopedWebhook) | **Get** /v1/Conversations/{ConversationSid}/Webhooks/{Sid} | 
[**ListConversationScopedWebhook**](ConversationsWebhooksApi.md#ListConversationScopedWebhook) | **Get** /v1/Conversations/{ConversationSid}/Webhooks | 
[**UpdateConversationScopedWebhook**](ConversationsWebhooksApi.md#UpdateConversationScopedWebhook) | **Post** /v1/Conversations/{ConversationSid}/Webhooks/{Sid} | 



## CreateConversationScopedWebhook

> ConversationsV1ConversationConversationScopedWebhook CreateConversationScopedWebhook(ctx, ConversationSidoptional)



Create a new webhook scoped to the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.

### Other Parameters

Other parameters are passed through a pointer to a CreateConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**ConfigurationFilters** | **[]string** | The list of events, firing webhook event for this Conversation.
**ConfigurationFlowSid** | **string** | The studio flow SID, where the webhook should be sent to.
**ConfigurationMethod** | **string** | The HTTP method to be used when sending a webhook request.
**ConfigurationReplayAfter** | **int** | The message index for which and it&#39;s successors the webhook will be replayed. Not set by default
**ConfigurationTriggers** | **[]string** | The list of keywords, firing webhook event for this Conversation.
**ConfigurationUrl** | **string** | The absolute url the webhook request should be sent to.
**Target** | **string** | The target of this webhook: &#x60;webhook&#x60;, &#x60;studio&#x60;, &#x60;trigger&#x60;

### Return type

[**ConversationsV1ConversationConversationScopedWebhook**](ConversationsV1ConversationConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConversationScopedWebhook

> DeleteConversationScopedWebhook(ctx, ConversationSidSid)



Remove an existing webhook scoped to the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConversationScopedWebhook

> ConversationsV1ConversationConversationScopedWebhook FetchConversationScopedWebhook(ctx, ConversationSidSid)



Fetch the configuration of a conversation-scoped webhook

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ConversationConversationScopedWebhook**](ConversationsV1ConversationConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversationScopedWebhook

> ListConversationScopedWebhookResponse ListConversationScopedWebhook(ctx, ConversationSidoptional)



Retrieve a list of all webhooks scoped to the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.

### Other Parameters

Other parameters are passed through a pointer to a ListConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListConversationScopedWebhookResponse**](ListConversationScopedWebhookResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConversationScopedWebhook

> ConversationsV1ConversationConversationScopedWebhook UpdateConversationScopedWebhook(ctx, ConversationSidSidoptional)



Update an existing conversation-scoped webhook

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**ConfigurationFilters** | **[]string** | The list of events, firing webhook event for this Conversation.
**ConfigurationFlowSid** | **string** | The studio flow SID, where the webhook should be sent to.
**ConfigurationMethod** | **string** | The HTTP method to be used when sending a webhook request.
**ConfigurationTriggers** | **[]string** | The list of keywords, firing webhook event for this Conversation.
**ConfigurationUrl** | **string** | The absolute url the webhook request should be sent to.

### Return type

[**ConversationsV1ConversationConversationScopedWebhook**](ConversationsV1ConversationConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
