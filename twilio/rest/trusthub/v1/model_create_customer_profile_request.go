/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateCustomerProfileRequest struct for CreateCustomerProfileRequest
type CreateCustomerProfileRequest struct {
	// The email address that will receive updates when the Customer-Profile resource changes status.
	Email string `json:"Email"`
	// The string that you assigned to describe the resource.
	FriendlyName string `json:"FriendlyName"`
	// The unique string of a policy that is associated to the Customer-Profile resource.
	PolicySid string `json:"PolicySid"`
	// The URL we call to inform your application of status changes.
	StatusCallback string `json:"StatusCallback,omitempty"`
}
