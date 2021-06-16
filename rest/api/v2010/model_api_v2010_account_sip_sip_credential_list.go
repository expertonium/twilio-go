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

// ApiV2010AccountSipSipCredentialList struct for ApiV2010AccountSipSipCredentialList
type ApiV2010AccountSipSipCredentialList struct {
	// The unique sid that identifies this account
	AccountSid *string `json:"account_sid,omitempty"`
	// The date this resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The date this resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// Human readable descriptive text
	FriendlyName *string `json:"friendly_name,omitempty"`
	// A string that uniquely identifies this credential
	Sid *string `json:"sid,omitempty"`
	// The list of credentials associated with this credential list.
	SubresourceUris *map[string]interface{} `json:"subresource_uris,omitempty"`
	// The URI for this resource
	Uri *string `json:"uri,omitempty"`
}