/*
 * Twilio - Bulkexports
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// BulkexportsV1ExportConfiguration struct for BulkexportsV1ExportConfiguration
type BulkexportsV1ExportConfiguration struct {
	Enabled       *bool   `json:"Enabled,omitempty"`
	ResourceType  *string `json:"ResourceType,omitempty"`
	Url           *string `json:"Url,omitempty"`
	WebhookMethod *string `json:"WebhookMethod,omitempty"`
	WebhookUrl    *string `json:"WebhookUrl,omitempty"`
}
