/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ApiV2010AccountRecordingRecordingAddOnResult struct for ApiV2010AccountRecordingRecordingAddOnResult
type ApiV2010AccountRecordingRecordingAddOnResult struct {
	AccountSid            *string                     `json:"AccountSid,omitempty"`
	AddOnConfigurationSid *string                     `json:"AddOnConfigurationSid,omitempty"`
	AddOnSid              *string                     `json:"AddOnSid,omitempty"`
	DateCompleted         *string                     `json:"DateCompleted,omitempty"`
	DateCreated           *string                     `json:"DateCreated,omitempty"`
	DateUpdated           *string                     `json:"DateUpdated,omitempty"`
	ReferenceSid          *string                     `json:"ReferenceSid,omitempty"`
	Sid                   *string                     `json:"Sid,omitempty"`
	Status                *RecordingAddOnResultStatus `json:"Status,omitempty"`
	SubresourceUris       *map[string]interface{}     `json:"SubresourceUris,omitempty"`
}
