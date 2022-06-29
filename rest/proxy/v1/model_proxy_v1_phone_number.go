/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Proxy
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// ProxyV1PhoneNumber struct for ProxyV1PhoneNumber
type ProxyV1PhoneNumber struct {
	// The SID of the Account that created the resource
	AccountSid   *string                                `json:"account_sid,omitempty"`
	Capabilities *ProxyV1ServicePhoneNumberCapabilities `json:"capabilities,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The number of open session assigned to the number.
	InUse *int `json:"in_use,omitempty"`
	// Reserve the phone number for manual assignment to participants only
	IsReserved *bool `json:"is_reserved,omitempty"`
	// The ISO Country Code
	IsoCountry *string `json:"iso_country,omitempty"`
	// The phone number in E.164 format
	PhoneNumber *string `json:"phone_number,omitempty"`
	// The SID of the PhoneNumber resource's parent Service resource
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the PhoneNumber resource
	Url *string `json:"url,omitempty"`
}
