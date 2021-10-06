/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.21.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// ProxyV1Service struct for ProxyV1Service
type ProxyV1Service struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The URL we call when the interaction status changes
	CallbackUrl *string `json:"callback_url,omitempty"`
	// The SID of the Chat Service Instance
	ChatInstanceSid *string `json:"chat_instance_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Default TTL for a Session, in seconds
	DefaultTtl *int `json:"default_ttl,omitempty"`
	// Where a proxy number must be located relative to the participant identifier
	GeoMatchLevel *string `json:"geo_match_level,omitempty"`
	// The URL we call on each interaction
	InterceptCallbackUrl *string `json:"intercept_callback_url,omitempty"`
	// The URLs of resources related to the Service
	Links *map[string]interface{} `json:"links,omitempty"`
	// The preference for Proxy Number selection for the Service instance
	NumberSelectionBehavior *string `json:"number_selection_behavior,omitempty"`
	// The URL we call when an inbound call or SMS action occurs on a closed or non-existent Session
	OutOfSessionCallbackUrl *string `json:"out_of_session_callback_url,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"unique_name,omitempty"`
	// The absolute URL of the Service resource
	Url *string `json:"url,omitempty"`
}
