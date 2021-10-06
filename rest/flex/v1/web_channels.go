/*
 * Twilio - Flex
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

// Optional parameters for the method 'CreateWebChannel'
type CreateWebChannelParams struct {
	// The chat channel's friendly name.
	ChatFriendlyName *string `json:"ChatFriendlyName,omitempty"`
	// The chat channel's unique name.
	ChatUniqueName *string `json:"ChatUniqueName,omitempty"`
	// The chat participant's friendly name.
	CustomerFriendlyName *string `json:"CustomerFriendlyName,omitempty"`
	// The SID of the Flex Flow.
	FlexFlowSid *string `json:"FlexFlowSid,omitempty"`
	// The chat identity.
	Identity *string `json:"Identity,omitempty"`
	// The pre-engagement data.
	PreEngagementData *string `json:"PreEngagementData,omitempty"`
}

func (params *CreateWebChannelParams) SetChatFriendlyName(ChatFriendlyName string) *CreateWebChannelParams {
	params.ChatFriendlyName = &ChatFriendlyName
	return params
}
func (params *CreateWebChannelParams) SetChatUniqueName(ChatUniqueName string) *CreateWebChannelParams {
	params.ChatUniqueName = &ChatUniqueName
	return params
}
func (params *CreateWebChannelParams) SetCustomerFriendlyName(CustomerFriendlyName string) *CreateWebChannelParams {
	params.CustomerFriendlyName = &CustomerFriendlyName
	return params
}
func (params *CreateWebChannelParams) SetFlexFlowSid(FlexFlowSid string) *CreateWebChannelParams {
	params.FlexFlowSid = &FlexFlowSid
	return params
}
func (params *CreateWebChannelParams) SetIdentity(Identity string) *CreateWebChannelParams {
	params.Identity = &Identity
	return params
}
func (params *CreateWebChannelParams) SetPreEngagementData(PreEngagementData string) *CreateWebChannelParams {
	params.PreEngagementData = &PreEngagementData
	return params
}

func (c *ApiService) CreateWebChannel(params *CreateWebChannelParams) (*FlexV1WebChannel, error) {
	path := "/v1/WebChannels"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ChatFriendlyName != nil {
		data.Set("ChatFriendlyName", *params.ChatFriendlyName)
	}
	if params != nil && params.ChatUniqueName != nil {
		data.Set("ChatUniqueName", *params.ChatUniqueName)
	}
	if params != nil && params.CustomerFriendlyName != nil {
		data.Set("CustomerFriendlyName", *params.CustomerFriendlyName)
	}
	if params != nil && params.FlexFlowSid != nil {
		data.Set("FlexFlowSid", *params.FlexFlowSid)
	}
	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
	}
	if params != nil && params.PreEngagementData != nil {
		data.Set("PreEngagementData", *params.PreEngagementData)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1WebChannel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteWebChannel(Sid string) error {
	path := "/v1/WebChannels/{Sid}"
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

func (c *ApiService) FetchWebChannel(Sid string) (*FlexV1WebChannel, error) {
	path := "/v1/WebChannels/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1WebChannel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListWebChannel'
type ListWebChannelParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListWebChannelParams) SetPageSize(PageSize int) *ListWebChannelParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListWebChannelParams) SetLimit(Limit int) *ListWebChannelParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of WebChannel records from the API. Request is executed immediately.
func (c *ApiService) PageWebChannel(params *ListWebChannelParams, pageToken, pageNumber string) (*ListWebChannelResponse, error) {
	path := "/v1/WebChannels"

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

	ps := &ListWebChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists WebChannel records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListWebChannel(params *ListWebChannelParams) ([]FlexV1WebChannel, error) {
	if params == nil {
		params = &ListWebChannelParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageWebChannel(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []FlexV1WebChannel

	for response != nil {
		records = append(records, response.FlexChatChannels...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListWebChannelResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListWebChannelResponse)
	}

	return records, err
}

// Streams WebChannel records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamWebChannel(params *ListWebChannelParams) (chan FlexV1WebChannel, error) {
	if params == nil {
		params = &ListWebChannelParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageWebChannel(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan FlexV1WebChannel, 1)

	go func() {
		for response != nil {
			for item := range response.FlexChatChannels {
				channel <- response.FlexChatChannels[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListWebChannelResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListWebChannelResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListWebChannelResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListWebChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateWebChannel'
type UpdateWebChannelParams struct {
	// The chat status. Can only be `inactive`.
	ChatStatus *string `json:"ChatStatus,omitempty"`
	// The post-engagement data.
	PostEngagementData *string `json:"PostEngagementData,omitempty"`
}

func (params *UpdateWebChannelParams) SetChatStatus(ChatStatus string) *UpdateWebChannelParams {
	params.ChatStatus = &ChatStatus
	return params
}
func (params *UpdateWebChannelParams) SetPostEngagementData(PostEngagementData string) *UpdateWebChannelParams {
	params.PostEngagementData = &PostEngagementData
	return params
}

func (c *ApiService) UpdateWebChannel(Sid string, params *UpdateWebChannelParams) (*FlexV1WebChannel, error) {
	path := "/v1/WebChannels/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ChatStatus != nil {
		data.Set("ChatStatus", *params.ChatStatus)
	}
	if params != nil && params.PostEngagementData != nil {
		data.Set("PostEngagementData", *params.PostEngagementData)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1WebChannel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
