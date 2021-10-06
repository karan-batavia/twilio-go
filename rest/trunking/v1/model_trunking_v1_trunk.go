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

import (
	"time"
)

// TrunkingV1Trunk struct for TrunkingV1Trunk
type TrunkingV1Trunk struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The types of authentication mapped to the domain
	AuthType *string `json:"auth_type,omitempty"`
	// Reserved
	AuthTypeSet *[]string `json:"auth_type_set,omitempty"`
	// Whether Caller ID Name (CNAM) lookup is enabled for the trunk
	CnamLookupEnabled *bool `json:"cnam_lookup_enabled,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The HTTP method we use to call the disaster_recovery_url
	DisasterRecoveryMethod *string `json:"disaster_recovery_method,omitempty"`
	// The HTTP URL that we call if an error occurs while sending SIP traffic towards your configured Origination URL
	DisasterRecoveryUrl *string `json:"disaster_recovery_url,omitempty"`
	// The unique address you reserve on Twilio to which you route your SIP traffic
	DomainName *string `json:"domain_name,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The recording settings for the trunk
	Recording *map[string]interface{} `json:"recording,omitempty"`
	// Whether Secure Trunking is enabled for the trunk
	Secure *bool `json:"secure,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// Caller Id for transfer target
	TransferCallerId *string `json:"transfer_caller_id,omitempty"`
	// The call transfer settings for the trunk
	TransferMode *string `json:"transfer_mode,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
