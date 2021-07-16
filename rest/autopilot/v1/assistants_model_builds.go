/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
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

// Optional parameters for the method 'CreateModelBuild'
type CreateModelBuildParams struct {
	// The URL we should call using a POST method to send status information to your application.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// An application-defined string that uniquely identifies the new resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the `sid` in the URL path to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateModelBuildParams) SetStatusCallback(StatusCallback string) *CreateModelBuildParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *CreateModelBuildParams) SetUniqueName(UniqueName string) *CreateModelBuildParams {
	params.UniqueName = &UniqueName
	return params
}

func (c *ApiService) CreateModelBuild(AssistantSid string, params *CreateModelBuildParams) (*AutopilotV1AssistantModelBuild, error) {
	path := "/v1/Assistants/{AssistantSid}/ModelBuilds"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1AssistantModelBuild{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteModelBuild(AssistantSid string, Sid string) error {
	path := "/v1/Assistants/{AssistantSid}/ModelBuilds/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
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

func (c *ApiService) FetchModelBuild(AssistantSid string, Sid string) (*AutopilotV1AssistantModelBuild, error) {
	path := "/v1/Assistants/{AssistantSid}/ModelBuilds/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1AssistantModelBuild{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListModelBuild'
type ListModelBuildParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListModelBuildParams) SetPageSize(PageSize int) *ListModelBuildParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a single page of ModelBuild records from the API. Request is executed immediately.
func (c *ApiService) PageModelBuild(AssistantSid string, params *ListModelBuildParams, pageToken string, pageNumber string) (*ListModelBuildResponse, error) {
	path := "/v1/Assistants/{AssistantSid}/ModelBuilds"

	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageToken != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListModelBuildResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ModelBuild records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListModelBuild(AssistantSid string, params *ListModelBuildParams, limit int) ([]AutopilotV1AssistantModelBuild, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageModelBuild(AssistantSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []AutopilotV1AssistantModelBuild

	for response != nil {
		records = append(records, response.ModelBuilds...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListModelBuildResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListModelBuildResponse)
	}

	return records, err
}

// Streams ModelBuild records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamModelBuild(AssistantSid string, params *ListModelBuildParams, limit int) (chan AutopilotV1AssistantModelBuild, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageModelBuild(AssistantSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan AutopilotV1AssistantModelBuild, 1)

	go func() {
		for response != nil {
			for item := range response.ModelBuilds {
				channel <- response.ModelBuilds[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListModelBuildResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListModelBuildResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListModelBuildResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListModelBuildResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateModelBuild'
type UpdateModelBuildParams struct {
	// An application-defined string that uniquely identifies the resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the `sid` in the URL path to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *UpdateModelBuildParams) SetUniqueName(UniqueName string) *UpdateModelBuildParams {
	params.UniqueName = &UniqueName
	return params
}

func (c *ApiService) UpdateModelBuild(AssistantSid string, Sid string, params *UpdateModelBuildParams) (*AutopilotV1AssistantModelBuild, error) {
	path := "/v1/Assistants/{AssistantSid}/ModelBuilds/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1AssistantModelBuild{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
