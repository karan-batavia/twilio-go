/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// ProxyV1ServiceSessionParticipant struct for ProxyV1ServiceSessionParticipant
type ProxyV1ServiceSessionParticipant struct {
	AccountSid         *string                 `json:"AccountSid,omitempty"`
	DateCreated        *time.Time              `json:"DateCreated,omitempty"`
	DateDeleted        *time.Time              `json:"DateDeleted,omitempty"`
	DateUpdated        *time.Time              `json:"DateUpdated,omitempty"`
	FriendlyName       *string                 `json:"FriendlyName,omitempty"`
	Identifier         *string                 `json:"Identifier,omitempty"`
	Links              *map[string]interface{} `json:"Links,omitempty"`
	ProxyIdentifier    *string                 `json:"ProxyIdentifier,omitempty"`
	ProxyIdentifierSid *string                 `json:"ProxyIdentifierSid,omitempty"`
	ServiceSid         *string                 `json:"ServiceSid,omitempty"`
	SessionSid         *string                 `json:"SessionSid,omitempty"`
	Sid                *string                 `json:"Sid,omitempty"`
	Url                *string                 `json:"Url,omitempty"`
}
