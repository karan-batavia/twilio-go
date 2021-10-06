/*
 * Twilio - Verify
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

// VerifyV2Verification struct for VerifyV2Verification
type VerifyV2Verification struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The amount of the associated PSD2 compliant transaction.
	Amount *string `json:"amount,omitempty"`
	// The verification method used.
	Channel *string `json:"channel,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Information about the phone number being verified
	Lookup *map[string]interface{} `json:"lookup,omitempty"`
	// The payee of the associated PSD2 compliant transaction
	Payee *string `json:"payee,omitempty"`
	// An array of verification attempt objects.
	SendCodeAttempts *[]map[string]interface{} `json:"send_code_attempts,omitempty"`
	// The SID of the Service that the resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The status of the verification resource
	Status *string `json:"status,omitempty"`
	// The phone number or email being verified
	To *string `json:"to,omitempty"`
	// The absolute URL of the Verification resource
	Url *string `json:"url,omitempty"`
	// Whether the verification was successful
	Valid *bool `json:"valid,omitempty"`
}
