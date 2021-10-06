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

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'FetchIncomingPhoneNumberAssignedAddOnExtension'
type FetchIncomingPhoneNumberAssignedAddOnExtensionParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchIncomingPhoneNumberAssignedAddOnExtensionParams) SetPathAccountSid(PathAccountSid string) *FetchIncomingPhoneNumberAssignedAddOnExtensionParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an instance of an Extension for the Assigned Add-on.
func (c *ApiService) FetchIncomingPhoneNumberAssignedAddOnExtension(ResourceSid string, AssignedAddOnSid string, Sid string, params *FetchIncomingPhoneNumberAssignedAddOnExtensionParams) (*ApiV2010IncomingPhoneNumberAssignedAddOnExtension, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{AssignedAddOnSid}/Extensions/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ResourceSid"+"}", ResourceSid, -1)
	path = strings.Replace(path, "{"+"AssignedAddOnSid"+"}", AssignedAddOnSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010IncomingPhoneNumberAssignedAddOnExtension{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListIncomingPhoneNumberAssignedAddOnExtension'
type ListIncomingPhoneNumberAssignedAddOnExtensionParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListIncomingPhoneNumberAssignedAddOnExtensionParams) SetPathAccountSid(PathAccountSid string) *ListIncomingPhoneNumberAssignedAddOnExtensionParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListIncomingPhoneNumberAssignedAddOnExtensionParams) SetPageSize(PageSize int) *ListIncomingPhoneNumberAssignedAddOnExtensionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListIncomingPhoneNumberAssignedAddOnExtensionParams) SetLimit(Limit int) *ListIncomingPhoneNumberAssignedAddOnExtensionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of IncomingPhoneNumberAssignedAddOnExtension records from the API. Request is executed immediately.
func (c *ApiService) PageIncomingPhoneNumberAssignedAddOnExtension(ResourceSid string, AssignedAddOnSid string, params *ListIncomingPhoneNumberAssignedAddOnExtensionParams, pageToken, pageNumber string) (*ListIncomingPhoneNumberAssignedAddOnExtensionResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{AssignedAddOnSid}/Extensions.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ResourceSid"+"}", ResourceSid, -1)
	path = strings.Replace(path, "{"+"AssignedAddOnSid"+"}", AssignedAddOnSid, -1)

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

	ps := &ListIncomingPhoneNumberAssignedAddOnExtensionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists IncomingPhoneNumberAssignedAddOnExtension records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListIncomingPhoneNumberAssignedAddOnExtension(ResourceSid string, AssignedAddOnSid string, params *ListIncomingPhoneNumberAssignedAddOnExtensionParams) ([]ApiV2010IncomingPhoneNumberAssignedAddOnExtension, error) {
	if params == nil {
		params = &ListIncomingPhoneNumberAssignedAddOnExtensionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageIncomingPhoneNumberAssignedAddOnExtension(ResourceSid, AssignedAddOnSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010IncomingPhoneNumberAssignedAddOnExtension

	for response != nil {
		records = append(records, response.Extensions...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListIncomingPhoneNumberAssignedAddOnExtensionResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListIncomingPhoneNumberAssignedAddOnExtensionResponse)
	}

	return records, err
}

// Streams IncomingPhoneNumberAssignedAddOnExtension records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamIncomingPhoneNumberAssignedAddOnExtension(ResourceSid string, AssignedAddOnSid string, params *ListIncomingPhoneNumberAssignedAddOnExtensionParams) (chan ApiV2010IncomingPhoneNumberAssignedAddOnExtension, error) {
	if params == nil {
		params = &ListIncomingPhoneNumberAssignedAddOnExtensionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageIncomingPhoneNumberAssignedAddOnExtension(ResourceSid, AssignedAddOnSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010IncomingPhoneNumberAssignedAddOnExtension, 1)

	go func() {
		for response != nil {
			for item := range response.Extensions {
				channel <- response.Extensions[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListIncomingPhoneNumberAssignedAddOnExtensionResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListIncomingPhoneNumberAssignedAddOnExtensionResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListIncomingPhoneNumberAssignedAddOnExtensionResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListIncomingPhoneNumberAssignedAddOnExtensionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
