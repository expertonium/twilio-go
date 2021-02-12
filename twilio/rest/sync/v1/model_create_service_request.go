/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateServiceRequest struct for CreateServiceRequest
type CreateServiceRequest struct {
	// Whether token identities in the Service must be granted access to Sync objects by using the [Permissions](https://www.twilio.com/docs/sync/api/sync-permissions) resource.
	AclEnabled bool `json:"AclEnabled,omitempty"`
	// A string that you assign to describe the resource.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// Whether every `endpoint_disconnected` event should occur after a configurable delay. The default is `false`, where the `endpoint_disconnected` event occurs immediately after disconnection. When `true`, intervening reconnections can prevent the `endpoint_disconnected` event.
	ReachabilityDebouncingEnabled bool `json:"ReachabilityDebouncingEnabled,omitempty"`
	// The reachability event delay in milliseconds if `reachability_debouncing_enabled` = `true`.  Must be between 1,000 and 30,000 and defaults to 5,000. This is the number of milliseconds after the last running client disconnects, and a Sync identity is declared offline, before the `webhook_url` is called if all endpoints remain offline. A reconnection from the same identity by any endpoint during this interval prevents the call to `webhook_url`.
	ReachabilityDebouncingWindow int32 `json:"ReachabilityDebouncingWindow,omitempty"`
	// Whether the service instance should call `webhook_url` when client endpoints connect to Sync. The default is `false`.
	ReachabilityWebhooksEnabled bool `json:"ReachabilityWebhooksEnabled,omitempty"`
	// The URL we should call when Sync objects are manipulated.
	WebhookUrl string `json:"WebhookUrl,omitempty"`
	// Whether the Service instance should call `webhook_url` when the REST API is used to update Sync objects. The default is `false`.
	WebhooksFromRestEnabled bool `json:"WebhooksFromRestEnabled,omitempty"`
}