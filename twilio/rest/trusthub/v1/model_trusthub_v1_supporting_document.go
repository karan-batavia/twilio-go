/*
 * Twilio - Trusthub
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

// TrusthubV1SupportingDocument struct for TrusthubV1SupportingDocument
type TrusthubV1SupportingDocument struct {
	AccountSid   *string                   `json:"AccountSid,omitempty"`
	Attributes   *map[string]interface{}   `json:"Attributes,omitempty"`
	DateCreated  *time.Time                `json:"DateCreated,omitempty"`
	DateUpdated  *time.Time                `json:"DateUpdated,omitempty"`
	FriendlyName *string                   `json:"FriendlyName,omitempty"`
	MimeType     *string                   `json:"MimeType,omitempty"`
	Sid          *string                   `json:"Sid,omitempty"`
	Status       *SupportingDocumentStatus `json:"Status,omitempty"`
	Type         *string                   `json:"Type,omitempty"`
	Url          *string                   `json:"Url,omitempty"`
}
