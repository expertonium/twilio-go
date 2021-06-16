/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// FlexV1WebChannel struct for FlexV1WebChannel
type FlexV1WebChannel struct {
	// The SID of the Account that created the resource and owns this Workflow
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The SID of the Flex Flow
	FlexFlowSid *string `json:"flex_flow_sid,omitempty"`
	// The unique string that identifies the WebChannel resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the WebChannel resource
	Url *string `json:"url,omitempty"`
}