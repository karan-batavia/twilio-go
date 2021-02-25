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

// ApiV2010AccountCallCallFeedback struct for ApiV2010AccountCallCallFeedback
type ApiV2010AccountCallCallFeedback struct {
	AccountSid   *string   `json:"AccountSid,omitempty"`
	DateCreated  *string   `json:"DateCreated,omitempty"`
	DateUpdated  *string   `json:"DateUpdated,omitempty"`
	Issues       *[]string `json:"Issues,omitempty"`
	QualityScore *int32    `json:"QualityScore,omitempty"`
	Sid          *string   `json:"Sid,omitempty"`
}
