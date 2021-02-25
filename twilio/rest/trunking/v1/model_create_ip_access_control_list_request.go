/*
 * Twilio - Trunking
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateIpAccessControlListRequest struct for CreateIpAccessControlListRequest
type CreateIpAccessControlListRequest struct {
	// The SID of the [IP Access Control List](https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource) that you want to associate with the trunk.
	IpAccessControlListSid string `json:"IpAccessControlListSid"`
}
