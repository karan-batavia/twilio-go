/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.21.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListBindingResponse struct for ListBindingResponse
type ListBindingResponse struct {
	Bindings []IpMessagingV2Binding     `json:"bindings,omitempty"`
	Meta     ListCredentialResponseMeta `json:"meta,omitempty"`
}
