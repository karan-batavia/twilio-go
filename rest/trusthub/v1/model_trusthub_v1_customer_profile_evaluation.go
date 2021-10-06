/*
 * Twilio - Trusthub
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

// TrusthubV1CustomerProfileEvaluation struct for TrusthubV1CustomerProfileEvaluation
type TrusthubV1CustomerProfileEvaluation struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The unique string that identifies the resource
	CustomerProfileSid *string    `json:"customer_profile_sid,omitempty"`
	DateCreated        *time.Time `json:"date_created,omitempty"`
	// The unique string of a policy
	PolicySid *string `json:"policy_sid,omitempty"`
	// The results of the Evaluation resource
	Results *[]map[string]interface{} `json:"results,omitempty"`
	// The unique string that identifies the Evaluation resource
	Sid *string `json:"sid,omitempty"`
	// The compliance status of the Evaluation resource
	Status *string `json:"status,omitempty"`
	Url    *string `json:"url,omitempty"`
}
