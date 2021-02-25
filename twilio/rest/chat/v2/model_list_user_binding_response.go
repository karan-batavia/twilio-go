/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListUserBindingResponse struct for ListUserBindingResponse
type ListUserBindingResponse struct {
	Bindings []ChatV2ServiceUserUserBinding `json:"Bindings,omitempty"`
	Meta     ListCredentialResponseMeta     `json:"Meta,omitempty"`
}
