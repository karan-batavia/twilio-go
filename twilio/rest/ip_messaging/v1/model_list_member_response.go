/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListMemberResponse struct for ListMemberResponse
type ListMemberResponse struct {
	Members []IpMessagingV1ServiceChannelMember `json:"Members,omitempty"`
	Meta    ListCredentialResponseMeta          `json:"Meta,omitempty"`
}
