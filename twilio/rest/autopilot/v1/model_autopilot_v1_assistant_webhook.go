/*
 * Twilio - Autopilot
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

// AutopilotV1AssistantWebhook struct for AutopilotV1AssistantWebhook
type AutopilotV1AssistantWebhook struct {
	AccountSid    *string    `json:"AccountSid,omitempty"`
	AssistantSid  *string    `json:"AssistantSid,omitempty"`
	DateCreated   *time.Time `json:"DateCreated,omitempty"`
	DateUpdated   *time.Time `json:"DateUpdated,omitempty"`
	Events        *string    `json:"Events,omitempty"`
	Sid           *string    `json:"Sid,omitempty"`
	UniqueName    *string    `json:"UniqueName,omitempty"`
	Url           *string    `json:"Url,omitempty"`
	WebhookMethod *string    `json:"WebhookMethod,omitempty"`
	WebhookUrl    *string    `json:"WebhookUrl,omitempty"`
}
