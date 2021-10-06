/*
 * Twilio - Trunking
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.21.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListPhoneNumberResponse struct for ListPhoneNumberResponse
type ListPhoneNumberResponse struct {
	Meta         ListTrunkResponseMeta   `json:"meta,omitempty"`
	PhoneNumbers []TrunkingV1PhoneNumber `json:"phone_numbers,omitempty"`
}
