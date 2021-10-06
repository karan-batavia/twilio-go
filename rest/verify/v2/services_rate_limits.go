/*
 * Twilio - Verify
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

// Optional parameters for the method 'CreateRateLimit'
type CreateRateLimitParams struct {
	// Description of this Rate Limit
	Description *string `json:"Description,omitempty"`
	// Provides a unique and addressable name to be assigned to this Rate Limit, assigned by the developer, to be optionally used in addition to SID. **This value should not contain PII.**
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateRateLimitParams) SetDescription(Description string) *CreateRateLimitParams {
	params.Description = &Description
	return params
}
func (params *CreateRateLimitParams) SetUniqueName(UniqueName string) *CreateRateLimitParams {
	params.UniqueName = &UniqueName
	return params
}

// Create a new Rate Limit for a Service
func (c *ApiService) CreateRateLimit(ServiceSid string, params *CreateRateLimitParams) (*VerifyV2RateLimit, error) {
	path := "/v2/Services/{ServiceSid}/RateLimits"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2RateLimit{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Rate Limit.
func (c *ApiService) DeleteRateLimit(ServiceSid string, Sid string) error {
	path := "/v2/Services/{ServiceSid}/RateLimits/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
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

// Fetch a specific Rate Limit.
func (c *ApiService) FetchRateLimit(ServiceSid string, Sid string) (*VerifyV2RateLimit, error) {
	path := "/v2/Services/{ServiceSid}/RateLimits/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2RateLimit{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRateLimit'
type ListRateLimitParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListRateLimitParams) SetPageSize(PageSize int) *ListRateLimitParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListRateLimitParams) SetLimit(Limit int) *ListRateLimitParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of RateLimit records from the API. Request is executed immediately.
func (c *ApiService) PageRateLimit(ServiceSid string, params *ListRateLimitParams, pageToken, pageNumber string) (*ListRateLimitResponse, error) {
	path := "/v2/Services/{ServiceSid}/RateLimits"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

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

	ps := &ListRateLimitResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists RateLimit records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRateLimit(ServiceSid string, params *ListRateLimitParams) ([]VerifyV2RateLimit, error) {
	if params == nil {
		params = &ListRateLimitParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageRateLimit(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []VerifyV2RateLimit

	for response != nil {
		records = append(records, response.RateLimits...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListRateLimitResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListRateLimitResponse)
	}

	return records, err
}

// Streams RateLimit records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRateLimit(ServiceSid string, params *ListRateLimitParams) (chan VerifyV2RateLimit, error) {
	if params == nil {
		params = &ListRateLimitParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageRateLimit(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan VerifyV2RateLimit, 1)

	go func() {
		for response != nil {
			for item := range response.RateLimits {
				channel <- response.RateLimits[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListRateLimitResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListRateLimitResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListRateLimitResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRateLimitResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateRateLimit'
type UpdateRateLimitParams struct {
	// Description of this Rate Limit
	Description *string `json:"Description,omitempty"`
}

func (params *UpdateRateLimitParams) SetDescription(Description string) *UpdateRateLimitParams {
	params.Description = &Description
	return params
}

// Update a specific Rate Limit.
func (c *ApiService) UpdateRateLimit(ServiceSid string, Sid string, params *UpdateRateLimitParams) (*VerifyV2RateLimit, error) {
	path := "/v2/Services/{ServiceSid}/RateLimits/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2RateLimit{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
