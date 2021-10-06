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
	"encoding/json"
	"fmt"
	"net/url"

	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateTrunk'
type CreateTrunkParams struct {
	// Whether Caller ID Name (CNAM) lookup should be enabled for the trunk. If enabled, all inbound calls to the SIP Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information.
	CnamLookupEnabled *bool `json:"CnamLookupEnabled,omitempty"`
	// The HTTP method we should use to call the `disaster_recovery_url`. Can be: `GET` or `POST`.
	DisasterRecoveryMethod *string `json:"DisasterRecoveryMethod,omitempty"`
	// The URL we should call using the `disaster_recovery_method` if an error occurs while sending SIP traffic towards the configured Origination URL. We retrieve TwiML from the URL and execute the instructions like any other normal TwiML call. See [Disaster Recovery](https://www.twilio.com/docs/sip-trunking#disaster-recovery) for more information.
	DisasterRecoveryUrl *string `json:"DisasterRecoveryUrl,omitempty"`
	// The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and `-` and must end with `pstn.twilio.com`. See [Termination Settings](https://www.twilio.com/docs/sip-trunking#termination) for more information.
	DomainName *string `json:"DomainName,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether Secure Trunking is enabled for the trunk. If enabled, all calls going through the trunk will be secure using SRTP for media and TLS for signaling. If disabled, then RTP will be used for media. See [Secure Trunking](https://www.twilio.com/docs/sip-trunking#securetrunking) for more information.
	Secure *bool `json:"Secure,omitempty"`
	// Caller Id for transfer target. Can be: `from-transferee` (default) or `from-transferor`.
	TransferCallerId *string `json:"TransferCallerId,omitempty"`
	// The call transfer settings for the trunk. Can be: `enable-all`, `sip-only` and `disable-all`. See [Transfer](https://www.twilio.com/docs/sip-trunking/call-transfer) for more information.
	TransferMode *string `json:"TransferMode,omitempty"`
}

func (params *CreateTrunkParams) SetCnamLookupEnabled(CnamLookupEnabled bool) *CreateTrunkParams {
	params.CnamLookupEnabled = &CnamLookupEnabled
	return params
}
func (params *CreateTrunkParams) SetDisasterRecoveryMethod(DisasterRecoveryMethod string) *CreateTrunkParams {
	params.DisasterRecoveryMethod = &DisasterRecoveryMethod
	return params
}
func (params *CreateTrunkParams) SetDisasterRecoveryUrl(DisasterRecoveryUrl string) *CreateTrunkParams {
	params.DisasterRecoveryUrl = &DisasterRecoveryUrl
	return params
}
func (params *CreateTrunkParams) SetDomainName(DomainName string) *CreateTrunkParams {
	params.DomainName = &DomainName
	return params
}
func (params *CreateTrunkParams) SetFriendlyName(FriendlyName string) *CreateTrunkParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateTrunkParams) SetSecure(Secure bool) *CreateTrunkParams {
	params.Secure = &Secure
	return params
}
func (params *CreateTrunkParams) SetTransferCallerId(TransferCallerId string) *CreateTrunkParams {
	params.TransferCallerId = &TransferCallerId
	return params
}
func (params *CreateTrunkParams) SetTransferMode(TransferMode string) *CreateTrunkParams {
	params.TransferMode = &TransferMode
	return params
}

func (c *ApiService) CreateTrunk(params *CreateTrunkParams) (*TrunkingV1Trunk, error) {
	path := "/v1/Trunks"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CnamLookupEnabled != nil {
		data.Set("CnamLookupEnabled", fmt.Sprint(*params.CnamLookupEnabled))
	}
	if params != nil && params.DisasterRecoveryMethod != nil {
		data.Set("DisasterRecoveryMethod", *params.DisasterRecoveryMethod)
	}
	if params != nil && params.DisasterRecoveryUrl != nil {
		data.Set("DisasterRecoveryUrl", *params.DisasterRecoveryUrl)
	}
	if params != nil && params.DomainName != nil {
		data.Set("DomainName", *params.DomainName)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Secure != nil {
		data.Set("Secure", fmt.Sprint(*params.Secure))
	}
	if params != nil && params.TransferCallerId != nil {
		data.Set("TransferCallerId", *params.TransferCallerId)
	}
	if params != nil && params.TransferMode != nil {
		data.Set("TransferMode", *params.TransferMode)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1Trunk{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteTrunk(Sid string) error {
	path := "/v1/Trunks/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchTrunk(Sid string) (*TrunkingV1Trunk, error) {
	path := "/v1/Trunks/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1Trunk{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListTrunk'
type ListTrunkParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListTrunkParams) SetPageSize(PageSize int) *ListTrunkParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListTrunkParams) SetLimit(Limit int) *ListTrunkParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Trunk records from the API. Request is executed immediately.
func (c *ApiService) PageTrunk(params *ListTrunkParams, pageToken, pageNumber string) (*ListTrunkResponse, error) {
	path := "/v1/Trunks"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTrunkResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Trunk records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListTrunk(params *ListTrunkParams) ([]TrunkingV1Trunk, error) {
	if params == nil {
		params = &ListTrunkParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageTrunk(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []TrunkingV1Trunk

	for response != nil {
		records = append(records, response.Trunks...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListTrunkResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListTrunkResponse)
	}

	return records, err
}

// Streams Trunk records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamTrunk(params *ListTrunkParams) (chan TrunkingV1Trunk, error) {
	if params == nil {
		params = &ListTrunkParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageTrunk(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan TrunkingV1Trunk, 1)

	go func() {
		for response != nil {
			for item := range response.Trunks {
				channel <- response.Trunks[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListTrunkResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListTrunkResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListTrunkResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTrunkResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateTrunk'
type UpdateTrunkParams struct {
	// Whether Caller ID Name (CNAM) lookup should be enabled for the trunk. If enabled, all inbound calls to the SIP Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information.
	CnamLookupEnabled *bool `json:"CnamLookupEnabled,omitempty"`
	// The HTTP method we should use to call the `disaster_recovery_url`. Can be: `GET` or `POST`.
	DisasterRecoveryMethod *string `json:"DisasterRecoveryMethod,omitempty"`
	// The URL we should call using the `disaster_recovery_method` if an error occurs while sending SIP traffic towards the configured Origination URL. We retrieve TwiML from the URL and execute the instructions like any other normal TwiML call. See [Disaster Recovery](https://www.twilio.com/docs/sip-trunking#disaster-recovery) for more information.
	DisasterRecoveryUrl *string `json:"DisasterRecoveryUrl,omitempty"`
	// The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and `-` and must end with `pstn.twilio.com`. See [Termination Settings](https://www.twilio.com/docs/sip-trunking#termination) for more information.
	DomainName *string `json:"DomainName,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether Secure Trunking is enabled for the trunk. If enabled, all calls going through the trunk will be secure using SRTP for media and TLS for signaling. If disabled, then RTP will be used for media. See [Secure Trunking](https://www.twilio.com/docs/sip-trunking#securetrunking) for more information.
	Secure *bool `json:"Secure,omitempty"`
	// Caller Id for transfer target. Can be: `from-transferee` (default) or `from-transferor`.
	TransferCallerId *string `json:"TransferCallerId,omitempty"`
	// The call transfer settings for the trunk. Can be: `enable-all`, `sip-only` and `disable-all`. See [Transfer](https://www.twilio.com/docs/sip-trunking/call-transfer) for more information.
	TransferMode *string `json:"TransferMode,omitempty"`
}

func (params *UpdateTrunkParams) SetCnamLookupEnabled(CnamLookupEnabled bool) *UpdateTrunkParams {
	params.CnamLookupEnabled = &CnamLookupEnabled
	return params
}
func (params *UpdateTrunkParams) SetDisasterRecoveryMethod(DisasterRecoveryMethod string) *UpdateTrunkParams {
	params.DisasterRecoveryMethod = &DisasterRecoveryMethod
	return params
}
func (params *UpdateTrunkParams) SetDisasterRecoveryUrl(DisasterRecoveryUrl string) *UpdateTrunkParams {
	params.DisasterRecoveryUrl = &DisasterRecoveryUrl
	return params
}
func (params *UpdateTrunkParams) SetDomainName(DomainName string) *UpdateTrunkParams {
	params.DomainName = &DomainName
	return params
}
func (params *UpdateTrunkParams) SetFriendlyName(FriendlyName string) *UpdateTrunkParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateTrunkParams) SetSecure(Secure bool) *UpdateTrunkParams {
	params.Secure = &Secure
	return params
}
func (params *UpdateTrunkParams) SetTransferCallerId(TransferCallerId string) *UpdateTrunkParams {
	params.TransferCallerId = &TransferCallerId
	return params
}
func (params *UpdateTrunkParams) SetTransferMode(TransferMode string) *UpdateTrunkParams {
	params.TransferMode = &TransferMode
	return params
}

func (c *ApiService) UpdateTrunk(Sid string, params *UpdateTrunkParams) (*TrunkingV1Trunk, error) {
	path := "/v1/Trunks/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CnamLookupEnabled != nil {
		data.Set("CnamLookupEnabled", fmt.Sprint(*params.CnamLookupEnabled))
	}
	if params != nil && params.DisasterRecoveryMethod != nil {
		data.Set("DisasterRecoveryMethod", *params.DisasterRecoveryMethod)
	}
	if params != nil && params.DisasterRecoveryUrl != nil {
		data.Set("DisasterRecoveryUrl", *params.DisasterRecoveryUrl)
	}
	if params != nil && params.DomainName != nil {
		data.Set("DomainName", *params.DomainName)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Secure != nil {
		data.Set("Secure", fmt.Sprint(*params.Secure))
	}
	if params != nil && params.TransferCallerId != nil {
		data.Set("TransferCallerId", *params.TransferCallerId)
	}
	if params != nil && params.TransferMode != nil {
		data.Set("TransferMode", *params.TransferMode)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1Trunk{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
