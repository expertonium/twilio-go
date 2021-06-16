/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ChatV1ServiceUserUserChannel struct for ChatV1ServiceUserUserChannel
type ChatV1ServiceUserUserChannel struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Channel the resource belongs to
	ChannelSid *string `json:"channel_sid,omitempty"`
	// The index of the last Message in the Channel the Member has read
	LastConsumedMessageIndex *int32 `json:"last_consumed_message_index,omitempty"`
	// Absolute URLs to access the Members, Messages , Invites and, if it exists, the last Message for the Channel
	Links *map[string]interface{} `json:"links,omitempty"`
	// The SID of the User as a Member in the Channel
	MemberSid *string `json:"member_sid,omitempty"`
	// The SID of the Service that the resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The status of the User on the Channel
	Status *string `json:"status,omitempty"`
	// The number of unread Messages in the Channel for the User
	UnreadMessagesCount *int32 `json:"unread_messages_count,omitempty"`
}