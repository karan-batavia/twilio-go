/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.21.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListStepResponse struct for ListStepResponse
type ListStepResponse struct {
	Meta  ListFlowResponseMeta `json:"meta,omitempty"`
	Steps []StudioV1Step       `json:"steps,omitempty"`
}
