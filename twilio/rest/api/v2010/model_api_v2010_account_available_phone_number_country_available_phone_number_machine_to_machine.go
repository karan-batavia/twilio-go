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

// ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberMachineToMachine struct for ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberMachineToMachine
type ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberMachineToMachine struct {
	AddressRequirements *string                  `json:"AddressRequirements,omitempty"`
	Beta                *bool                    `json:"Beta,omitempty"`
	Capabilities        *PhoneNumberCapabilities `json:"Capabilities,omitempty"`
	FriendlyName        *string                  `json:"FriendlyName,omitempty"`
	IsoCountry          *string                  `json:"IsoCountry,omitempty"`
	Lata                *string                  `json:"Lata,omitempty"`
	Latitude            *float32                 `json:"Latitude,omitempty"`
	Locality            *string                  `json:"Locality,omitempty"`
	Longitude           *float32                 `json:"Longitude,omitempty"`
	PhoneNumber         *string                  `json:"PhoneNumber,omitempty"`
	PostalCode          *string                  `json:"PostalCode,omitempty"`
	RateCenter          *string                  `json:"RateCenter,omitempty"`
	Region              *string                  `json:"Region,omitempty"`
}
