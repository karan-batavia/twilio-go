/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Lookups
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"
	"strings"
)

// Optional parameters for the method 'FetchPhoneNumber'
type FetchPhoneNumberParams struct {
	// A comma-separated list of fields to return. Possible values are validation, caller_name, sim_swap, call_forwarding, line_status, line_type_intelligence, identity_match, reassigned_number, sms_pumping_risk, phone_number_quality_score, pre_fill.
	Fields *string `json:"Fields,omitempty"`
	// The [country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) used if the phone number provided is in national format.
	CountryCode *string `json:"CountryCode,omitempty"`
	// User’s first name. This query parameter is only used (optionally) for identity_match package requests.
	FirstName *string `json:"FirstName,omitempty"`
	// User’s last name. This query parameter is only used (optionally) for identity_match package requests.
	LastName *string `json:"LastName,omitempty"`
	// User’s first address line. This query parameter is only used (optionally) for identity_match package requests.
	AddressLine1 *string `json:"AddressLine1,omitempty"`
	// User’s second address line. This query parameter is only used (optionally) for identity_match package requests.
	AddressLine2 *string `json:"AddressLine2,omitempty"`
	// User’s city. This query parameter is only used (optionally) for identity_match package requests.
	City *string `json:"City,omitempty"`
	// User’s country subdivision, such as state, province, or locality. This query parameter is only used (optionally) for identity_match package requests.
	State *string `json:"State,omitempty"`
	// User’s postal zip code. This query parameter is only used (optionally) for identity_match package requests.
	PostalCode *string `json:"PostalCode,omitempty"`
	// User’s country, up to two characters. This query parameter is only used (optionally) for identity_match package requests.
	AddressCountryCode *string `json:"AddressCountryCode,omitempty"`
	// User’s national ID, such as SSN or Passport ID. This query parameter is only used (optionally) for identity_match package requests.
	NationalId *string `json:"NationalId,omitempty"`
	// User’s date of birth, in YYYYMMDD format. This query parameter is only used (optionally) for identity_match package requests.
	DateOfBirth *string `json:"DateOfBirth,omitempty"`
	// The date you obtained consent to call or text the end-user of the phone number or a date on which you are reasonably certain that the end-user could still be reached at that number. This query parameter is only used (optionally) for reassigned_number package requests.
	LastVerifiedDate *string `json:"LastVerifiedDate,omitempty"`
	// The unique identifier associated with a verification process through verify API. This query parameter is only used (optionally) for pre_fill package requests.
	VerificationSid *string `json:"VerificationSid,omitempty"`
}

func (params *FetchPhoneNumberParams) SetFields(Fields string) *FetchPhoneNumberParams {
	params.Fields = &Fields
	return params
}
func (params *FetchPhoneNumberParams) SetCountryCode(CountryCode string) *FetchPhoneNumberParams {
	params.CountryCode = &CountryCode
	return params
}
func (params *FetchPhoneNumberParams) SetFirstName(FirstName string) *FetchPhoneNumberParams {
	params.FirstName = &FirstName
	return params
}
func (params *FetchPhoneNumberParams) SetLastName(LastName string) *FetchPhoneNumberParams {
	params.LastName = &LastName
	return params
}
func (params *FetchPhoneNumberParams) SetAddressLine1(AddressLine1 string) *FetchPhoneNumberParams {
	params.AddressLine1 = &AddressLine1
	return params
}
func (params *FetchPhoneNumberParams) SetAddressLine2(AddressLine2 string) *FetchPhoneNumberParams {
	params.AddressLine2 = &AddressLine2
	return params
}
func (params *FetchPhoneNumberParams) SetCity(City string) *FetchPhoneNumberParams {
	params.City = &City
	return params
}
func (params *FetchPhoneNumberParams) SetState(State string) *FetchPhoneNumberParams {
	params.State = &State
	return params
}
func (params *FetchPhoneNumberParams) SetPostalCode(PostalCode string) *FetchPhoneNumberParams {
	params.PostalCode = &PostalCode
	return params
}
func (params *FetchPhoneNumberParams) SetAddressCountryCode(AddressCountryCode string) *FetchPhoneNumberParams {
	params.AddressCountryCode = &AddressCountryCode
	return params
}
func (params *FetchPhoneNumberParams) SetNationalId(NationalId string) *FetchPhoneNumberParams {
	params.NationalId = &NationalId
	return params
}
func (params *FetchPhoneNumberParams) SetDateOfBirth(DateOfBirth string) *FetchPhoneNumberParams {
	params.DateOfBirth = &DateOfBirth
	return params
}
func (params *FetchPhoneNumberParams) SetLastVerifiedDate(LastVerifiedDate string) *FetchPhoneNumberParams {
	params.LastVerifiedDate = &LastVerifiedDate
	return params
}
func (params *FetchPhoneNumberParams) SetVerificationSid(VerificationSid string) *FetchPhoneNumberParams {
	params.VerificationSid = &VerificationSid
	return params
}

//
func (c *ApiService) FetchPhoneNumber(PhoneNumber string, params *FetchPhoneNumberParams) (*LookupsV2PhoneNumber, error) {
	path := "/v2/PhoneNumbers/{PhoneNumber}"
	path = strings.Replace(path, "{"+"PhoneNumber"+"}", PhoneNumber, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Fields != nil {
		data.Set("Fields", *params.Fields)
	}
	if params != nil && params.CountryCode != nil {
		data.Set("CountryCode", *params.CountryCode)
	}
	if params != nil && params.FirstName != nil {
		data.Set("FirstName", *params.FirstName)
	}
	if params != nil && params.LastName != nil {
		data.Set("LastName", *params.LastName)
	}
	if params != nil && params.AddressLine1 != nil {
		data.Set("AddressLine1", *params.AddressLine1)
	}
	if params != nil && params.AddressLine2 != nil {
		data.Set("AddressLine2", *params.AddressLine2)
	}
	if params != nil && params.City != nil {
		data.Set("City", *params.City)
	}
	if params != nil && params.State != nil {
		data.Set("State", *params.State)
	}
	if params != nil && params.PostalCode != nil {
		data.Set("PostalCode", *params.PostalCode)
	}
	if params != nil && params.AddressCountryCode != nil {
		data.Set("AddressCountryCode", *params.AddressCountryCode)
	}
	if params != nil && params.NationalId != nil {
		data.Set("NationalId", *params.NationalId)
	}
	if params != nil && params.DateOfBirth != nil {
		data.Set("DateOfBirth", *params.DateOfBirth)
	}
	if params != nil && params.LastVerifiedDate != nil {
		data.Set("LastVerifiedDate", *params.LastVerifiedDate)
	}
	if params != nil && params.VerificationSid != nil {
		data.Set("VerificationSid", *params.VerificationSid)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &LookupsV2PhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
