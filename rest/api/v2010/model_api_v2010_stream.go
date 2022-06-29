/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ApiV2010Stream struct for ApiV2010Stream
type ApiV2010Stream struct {
	// The SID of the Account that created this resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Call the resource is associated with
	CallSid *string `json:"call_sid,omitempty"`
	// The RFC 2822 date and time in GMT that this resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// The name of this resource
	Name *string `json:"name,omitempty"`
	// The SID of the Stream resource.
	Sid *string `json:"sid,omitempty"`
	// The status - one of `stopped`, `in-progress`
	Status *string `json:"status,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
}
