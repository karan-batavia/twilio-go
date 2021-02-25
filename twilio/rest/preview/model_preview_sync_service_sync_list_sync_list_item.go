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

// PreviewSyncServiceSyncListSyncListItem struct for PreviewSyncServiceSyncListSyncListItem
type PreviewSyncServiceSyncListSyncListItem struct {
	AccountSid  *string                 `json:"AccountSid,omitempty"`
	CreatedBy   *string                 `json:"CreatedBy,omitempty"`
	Data        *map[string]interface{} `json:"Data,omitempty"`
	DateCreated *time.Time              `json:"DateCreated,omitempty"`
	DateUpdated *time.Time              `json:"DateUpdated,omitempty"`
	Index       *int32                  `json:"Index,omitempty"`
	ListSid     *string                 `json:"ListSid,omitempty"`
	Revision    *string                 `json:"Revision,omitempty"`
	ServiceSid  *string                 `json:"ServiceSid,omitempty"`
	Url         *string                 `json:"Url,omitempty"`
}
