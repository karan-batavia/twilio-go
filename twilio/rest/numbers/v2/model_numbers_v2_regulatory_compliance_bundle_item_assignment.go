/*
 * Twilio - Numbers
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

// NumbersV2RegulatoryComplianceBundleItemAssignment struct for NumbersV2RegulatoryComplianceBundleItemAssignment
type NumbersV2RegulatoryComplianceBundleItemAssignment struct {
	AccountSid  *string    `json:"AccountSid,omitempty"`
	BundleSid   *string    `json:"BundleSid,omitempty"`
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	ObjectSid   *string    `json:"ObjectSid,omitempty"`
	Sid         *string    `json:"Sid,omitempty"`
	Url         *string    `json:"Url,omitempty"`
}
