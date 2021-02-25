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

// ApiV2010AccountSipSipIpAccessControlListSipIpAddress struct for ApiV2010AccountSipSipIpAccessControlListSipIpAddress
type ApiV2010AccountSipSipIpAccessControlListSipIpAddress struct {
	AccountSid             *string `json:"AccountSid,omitempty"`
	CidrPrefixLength       *int32  `json:"CidrPrefixLength,omitempty"`
	DateCreated            *string `json:"DateCreated,omitempty"`
	DateUpdated            *string `json:"DateUpdated,omitempty"`
	FriendlyName           *string `json:"FriendlyName,omitempty"`
	IpAccessControlListSid *string `json:"IpAccessControlListSid,omitempty"`
	IpAddress              *string `json:"IpAddress,omitempty"`
	Sid                    *string `json:"Sid,omitempty"`
	Uri                    *string `json:"Uri,omitempty"`
}
