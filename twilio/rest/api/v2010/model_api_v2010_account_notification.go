/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ApiV2010AccountNotification struct for ApiV2010AccountNotification
type ApiV2010AccountNotification struct {
	AccountSid string `json:"AccountSid,omitempty"`
	ApiVersion string `json:"ApiVersion,omitempty"`
	CallSid string `json:"CallSid,omitempty"`
	DateCreated string `json:"DateCreated,omitempty"`
	DateUpdated string `json:"DateUpdated,omitempty"`
	ErrorCode string `json:"ErrorCode,omitempty"`
	Log string `json:"Log,omitempty"`
	MessageDate string `json:"MessageDate,omitempty"`
	MessageText string `json:"MessageText,omitempty"`
	MoreInfo string `json:"MoreInfo,omitempty"`
	RequestMethod HttpMethod `json:"RequestMethod,omitempty"`
	RequestUrl string `json:"RequestUrl,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Uri string `json:"Uri,omitempty"`
}