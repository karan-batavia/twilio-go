/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.21.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListTrunkingCountryResponse struct for ListTrunkingCountryResponse
type ListTrunkingCountryResponse struct {
	Countries []PricingV2TrunkingCountry      `json:"countries,omitempty"`
	Meta      ListTrunkingCountryResponseMeta `json:"meta,omitempty"`
}
