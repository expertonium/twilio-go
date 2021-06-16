/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListWorkerReservationResponse struct for ListWorkerReservationResponse
type ListWorkerReservationResponse struct {
	Meta         ListWorkspaceResponseMeta                      `json:"meta,omitempty"`
	Reservations []TaskrouterV1WorkspaceWorkerWorkerReservation `json:"reservations,omitempty"`
}