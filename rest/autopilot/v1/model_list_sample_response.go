/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSampleResponse struct for ListSampleResponse
type ListSampleResponse struct {
	Meta    ListAssistantResponseMeta        `json:"meta,omitempty"`
	Samples []AutopilotV1AssistantTaskSample `json:"samples,omitempty"`
}