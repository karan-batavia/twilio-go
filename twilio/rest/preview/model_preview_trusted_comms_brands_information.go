/*
 * Twilio - Preview
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

// PreviewTrustedCommsBrandsInformation struct for PreviewTrustedCommsBrandsInformation
type PreviewTrustedCommsBrandsInformation struct {
	FileLink             *string    `json:"FileLink,omitempty"`
	FileLinkTtlInSeconds *string    `json:"FileLinkTtlInSeconds,omitempty"`
	UpdateTime           *time.Time `json:"UpdateTime,omitempty"`
	Url                  *string    `json:"Url,omitempty"`
}
