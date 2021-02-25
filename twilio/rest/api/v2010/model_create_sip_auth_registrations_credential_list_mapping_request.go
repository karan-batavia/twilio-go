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

// CreateSipAuthRegistrationsCredentialListMappingRequest struct for CreateSipAuthRegistrationsCredentialListMappingRequest
type CreateSipAuthRegistrationsCredentialListMappingRequest struct {
	// The SID of the CredentialList resource to map to the SIP domain.
	CredentialListSid string `json:"CredentialListSid"`
}
