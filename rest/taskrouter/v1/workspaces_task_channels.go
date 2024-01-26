/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Taskrouter
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateTaskChannel'
type CreateTaskChannelParams struct {
	// A descriptive string that you create to describe the Task Channel. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// An application-defined string that uniquely identifies the Task Channel, such as `voice` or `sms`.
	UniqueName *string `json:"UniqueName,omitempty"`
	// Whether the Task Channel should prioritize Workers that have been idle. If `true`, Workers that have been idle the longest are prioritized.
	ChannelOptimizedRouting *bool `json:"ChannelOptimizedRouting,omitempty"`
}

func (params *CreateTaskChannelParams) SetFriendlyName(FriendlyName string) *CreateTaskChannelParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateTaskChannelParams) SetUniqueName(UniqueName string) *CreateTaskChannelParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *CreateTaskChannelParams) SetChannelOptimizedRouting(ChannelOptimizedRouting bool) *CreateTaskChannelParams {
	params.ChannelOptimizedRouting = &ChannelOptimizedRouting
	return params
}

func (c *ApiService) CreateTaskChannel(WorkspaceSid string, params *CreateTaskChannelParams) (*TaskrouterV1TaskChannel, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/TaskChannels"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.ChannelOptimizedRouting != nil {
		data.Set("ChannelOptimizedRouting", fmt.Sprint(*params.ChannelOptimizedRouting))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers, queryParams)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1TaskChannel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteTaskChannel(WorkspaceSid string, Sid string) error {
	path := "/v1/Workspaces/{WorkspaceSid}/TaskChannels/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers, queryParams)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchTaskChannel(WorkspaceSid string, Sid string) (*TaskrouterV1TaskChannel, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/TaskChannels/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers, queryParams)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1TaskChannel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListTaskChannel'
type ListTaskChannelParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListTaskChannelParams) SetPageSize(PageSize int) *ListTaskChannelParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListTaskChannelParams) SetLimit(Limit int) *ListTaskChannelParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of TaskChannel records from the API. Request is executed immediately.
func (c *ApiService) PageTaskChannel(WorkspaceSid string, params *ListTaskChannelParams, pageToken, pageNumber string) (*ListTaskChannelResponse, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/TaskChannels"

	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		queryParams.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers, queryParams)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTaskChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists TaskChannel records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListTaskChannel(WorkspaceSid string, params *ListTaskChannelParams) ([]TaskrouterV1TaskChannel, error) {
	response, errors := c.StreamTaskChannel(WorkspaceSid, params)

	records := make([]TaskrouterV1TaskChannel, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams TaskChannel records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamTaskChannel(WorkspaceSid string, params *ListTaskChannelParams) (chan TaskrouterV1TaskChannel, chan error) {
	if params == nil {
		params = &ListTaskChannelParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan TaskrouterV1TaskChannel, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageTaskChannel(WorkspaceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamTaskChannel(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamTaskChannel(response *ListTaskChannelResponse, params *ListTaskChannelParams, recordChannel chan TaskrouterV1TaskChannel, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Channels
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListTaskChannelResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListTaskChannelResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListTaskChannelResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTaskChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateTaskChannel'
type UpdateTaskChannelParams struct {
	// A descriptive string that you create to describe the Task Channel. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether the TaskChannel should prioritize Workers that have been idle. If `true`, Workers that have been idle the longest are prioritized.
	ChannelOptimizedRouting *bool `json:"ChannelOptimizedRouting,omitempty"`
}

func (params *UpdateTaskChannelParams) SetFriendlyName(FriendlyName string) *UpdateTaskChannelParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateTaskChannelParams) SetChannelOptimizedRouting(ChannelOptimizedRouting bool) *UpdateTaskChannelParams {
	params.ChannelOptimizedRouting = &ChannelOptimizedRouting
	return params
}

func (c *ApiService) UpdateTaskChannel(WorkspaceSid string, Sid string, params *UpdateTaskChannelParams) (*TaskrouterV1TaskChannel, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/TaskChannels/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.ChannelOptimizedRouting != nil {
		data.Set("ChannelOptimizedRouting", fmt.Sprint(*params.ChannelOptimizedRouting))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers, queryParams)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1TaskChannel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
