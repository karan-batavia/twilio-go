/*
 * Twilio - Serverless
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

// ServerlessV1AssetVersion struct for ServerlessV1AssetVersion
type ServerlessV1AssetVersion struct {
	// The SID of the Account that created the Asset Version resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Asset resource that is the parent of the Asset Version
	AssetSid *string `json:"asset_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the Asset Version resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The URL-friendly string by which the Asset Version can be referenced
	Path *string `json:"path,omitempty"`
	// The SID of the Service that the Asset Version resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the Asset Version resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the Asset Version resource
	Url *string `json:"url,omitempty"`
	// The access control that determines how the Asset Version can be accessed
	Visibility *string `json:"visibility,omitempty"`
}
