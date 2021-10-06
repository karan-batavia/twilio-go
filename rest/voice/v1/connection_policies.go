/*
 * Twilio - Voice
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

// Optional parameters for the method 'CreateConnectionPolicy'
type CreateConnectionPolicyParams struct {
	// A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateConnectionPolicyParams) SetFriendlyName(FriendlyName string) *CreateConnectionPolicyParams {
	params.FriendlyName = &FriendlyName
	return params
}

func (c *ApiService) CreateConnectionPolicy(params *CreateConnectionPolicyParams) (*VoiceV1ConnectionPolicy, error) {
	path := "/v1/ConnectionPolicies"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1ConnectionPolicy{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteConnectionPolicy(Sid string) error {
	path := "/v1/ConnectionPolicies/{Sid}"
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

func (c *ApiService) FetchConnectionPolicy(Sid string) (*VoiceV1ConnectionPolicy, error) {
	path := "/v1/ConnectionPolicies/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1ConnectionPolicy{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListConnectionPolicy'
type ListConnectionPolicyParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListConnectionPolicyParams) SetPageSize(PageSize int) *ListConnectionPolicyParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListConnectionPolicyParams) SetLimit(Limit int) *ListConnectionPolicyParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ConnectionPolicy records from the API. Request is executed immediately.
func (c *ApiService) PageConnectionPolicy(params *ListConnectionPolicyParams, pageToken, pageNumber string) (*ListConnectionPolicyResponse, error) {
	path := "/v1/ConnectionPolicies"

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

	ps := &ListConnectionPolicyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ConnectionPolicy records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListConnectionPolicy(params *ListConnectionPolicyParams) ([]VoiceV1ConnectionPolicy, error) {
	if params == nil {
		params = &ListConnectionPolicyParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageConnectionPolicy(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []VoiceV1ConnectionPolicy

	for response != nil {
		records = append(records, response.ConnectionPolicies...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListConnectionPolicyResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListConnectionPolicyResponse)
	}

	return records, err
}

// Streams ConnectionPolicy records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamConnectionPolicy(params *ListConnectionPolicyParams) (chan VoiceV1ConnectionPolicy, error) {
	if params == nil {
		params = &ListConnectionPolicyParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageConnectionPolicy(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan VoiceV1ConnectionPolicy, 1)

	go func() {
		for response != nil {
			for item := range response.ConnectionPolicies {
				channel <- response.ConnectionPolicies[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListConnectionPolicyResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListConnectionPolicyResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListConnectionPolicyResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListConnectionPolicyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateConnectionPolicy'
type UpdateConnectionPolicyParams struct {
	// A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateConnectionPolicyParams) SetFriendlyName(FriendlyName string) *UpdateConnectionPolicyParams {
	params.FriendlyName = &FriendlyName
	return params
}

func (c *ApiService) UpdateConnectionPolicy(Sid string, params *UpdateConnectionPolicyParams) (*VoiceV1ConnectionPolicy, error) {
	path := "/v1/ConnectionPolicies/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1ConnectionPolicy{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
