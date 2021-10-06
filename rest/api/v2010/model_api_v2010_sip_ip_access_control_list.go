/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.21.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010SipIpAccessControlList struct for ApiV2010SipIpAccessControlList
type ApiV2010SipIpAccessControlList struct {
	// The unique sid that identifies this account
	AccountSid *string `json:"account_sid,omitempty"`
	// The date this resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The date this resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// A human readable description of this resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// A string that uniquely identifies this resource
	Sid *string `json:"sid,omitempty"`
	// The IP addresses associated with this resource.
	SubresourceUris *map[string]interface{} `json:"subresource_uris,omitempty"`
	// The URI for this resource
	Uri *string `json:"uri,omitempty"`
}
