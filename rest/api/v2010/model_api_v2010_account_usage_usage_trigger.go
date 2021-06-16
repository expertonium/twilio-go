/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountUsageUsageTrigger struct for ApiV2010AccountUsageUsageTrigger
type ApiV2010AccountUsageUsageTrigger struct {
	// The SID of the Account that this trigger monitors
	AccountSid *string `json:"account_sid,omitempty"`
	// The API version used to create the resource
	ApiVersion *string `json:"api_version,omitempty"`
	// The HTTP method we use to call callback_url
	CallbackMethod *string `json:"callback_method,omitempty"`
	// he URL we call when the trigger fires
	CallbackUrl *string `json:"callback_url,omitempty"`
	// The current value of the field the trigger is watching
	CurrentValue *string `json:"current_value,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the trigger was last fired
	DateFired *string `json:"date_fired,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// The string that you assigned to describe the trigger
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The frequency of a recurring UsageTrigger
	Recurring *string `json:"recurring,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The field in the UsageRecord resource that fires the trigger
	TriggerBy *string `json:"trigger_by,omitempty"`
	// The value at which the trigger will fire
	TriggerValue *string `json:"trigger_value,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
	// The usage category the trigger watches
	UsageCategory *string `json:"usage_category,omitempty"`
	// The URI of the UsageRecord resource this trigger watches
	UsageRecordUri *string `json:"usage_record_uri,omitempty"`
}