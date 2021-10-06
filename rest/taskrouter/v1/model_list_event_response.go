/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.21.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListEventResponse struct for ListEventResponse
type ListEventResponse struct {
	Events []TaskrouterV1Event       `json:"events,omitempty"`
	Meta   ListWorkspaceResponseMeta `json:"meta,omitempty"`
}
