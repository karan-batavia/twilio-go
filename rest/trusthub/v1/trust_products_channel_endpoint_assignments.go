/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Trusthub
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

// Optional parameters for the method 'CreateTrustProductChannelEndpointAssignment'
type CreateTrustProductChannelEndpointAssignmentParams struct {
	// The SID of an channel endpoint
	ChannelEndpointSid *string `json:"ChannelEndpointSid,omitempty"`
	// The type of channel endpoint. eg: phone-number
	ChannelEndpointType *string `json:"ChannelEndpointType,omitempty"`
}

func (params *CreateTrustProductChannelEndpointAssignmentParams) SetChannelEndpointSid(ChannelEndpointSid string) *CreateTrustProductChannelEndpointAssignmentParams {
	params.ChannelEndpointSid = &ChannelEndpointSid
	return params
}
func (params *CreateTrustProductChannelEndpointAssignmentParams) SetChannelEndpointType(ChannelEndpointType string) *CreateTrustProductChannelEndpointAssignmentParams {
	params.ChannelEndpointType = &ChannelEndpointType
	return params
}

// Create a new Assigned Item.
func (c *ApiService) CreateTrustProductChannelEndpointAssignment(TrustProductSid string, params *CreateTrustProductChannelEndpointAssignmentParams) (*TrusthubV1TrustProductChannelEndpointAssignment, error) {
	path := "/v1/TrustProducts/{TrustProductSid}/ChannelEndpointAssignments"
	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ChannelEndpointSid != nil {
		data.Set("ChannelEndpointSid", *params.ChannelEndpointSid)
	}
	if params != nil && params.ChannelEndpointType != nil {
		data.Set("ChannelEndpointType", *params.ChannelEndpointType)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1TrustProductChannelEndpointAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an Assignment Item Instance.
func (c *ApiService) DeleteTrustProductChannelEndpointAssignment(TrustProductSid string, Sid string) error {
	path := "/v1/TrustProducts/{TrustProductSid}/ChannelEndpointAssignments/{Sid}"
	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)
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

// Fetch specific Assigned Item Instance.
func (c *ApiService) FetchTrustProductChannelEndpointAssignment(TrustProductSid string, Sid string) (*TrusthubV1TrustProductChannelEndpointAssignment, error) {
	path := "/v1/TrustProducts/{TrustProductSid}/ChannelEndpointAssignments/{Sid}"
	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1TrustProductChannelEndpointAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListTrustProductChannelEndpointAssignment'
type ListTrustProductChannelEndpointAssignmentParams struct {
	// The SID of an channel endpoint
	ChannelEndpointSid *string `json:"ChannelEndpointSid,omitempty"`
	// comma separated list of channel endpoint sids
	ChannelEndpointSids *string `json:"ChannelEndpointSids,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListTrustProductChannelEndpointAssignmentParams) SetChannelEndpointSid(ChannelEndpointSid string) *ListTrustProductChannelEndpointAssignmentParams {
	params.ChannelEndpointSid = &ChannelEndpointSid
	return params
}
func (params *ListTrustProductChannelEndpointAssignmentParams) SetChannelEndpointSids(ChannelEndpointSids string) *ListTrustProductChannelEndpointAssignmentParams {
	params.ChannelEndpointSids = &ChannelEndpointSids
	return params
}
func (params *ListTrustProductChannelEndpointAssignmentParams) SetPageSize(PageSize int) *ListTrustProductChannelEndpointAssignmentParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListTrustProductChannelEndpointAssignmentParams) SetLimit(Limit int) *ListTrustProductChannelEndpointAssignmentParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of TrustProductChannelEndpointAssignment records from the API. Request is executed immediately.
func (c *ApiService) PageTrustProductChannelEndpointAssignment(TrustProductSid string, params *ListTrustProductChannelEndpointAssignmentParams, pageToken, pageNumber string) (*ListTrustProductChannelEndpointAssignmentResponse, error) {
	path := "/v1/TrustProducts/{TrustProductSid}/ChannelEndpointAssignments"

	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ChannelEndpointSid != nil {
		data.Set("ChannelEndpointSid", *params.ChannelEndpointSid)
	}
	if params != nil && params.ChannelEndpointSids != nil {
		data.Set("ChannelEndpointSids", *params.ChannelEndpointSids)
	}
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

	ps := &ListTrustProductChannelEndpointAssignmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists TrustProductChannelEndpointAssignment records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListTrustProductChannelEndpointAssignment(TrustProductSid string, params *ListTrustProductChannelEndpointAssignmentParams) ([]TrusthubV1TrustProductChannelEndpointAssignment, error) {
	response, errors := c.StreamTrustProductChannelEndpointAssignment(TrustProductSid, params)

	records := make([]TrusthubV1TrustProductChannelEndpointAssignment, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams TrustProductChannelEndpointAssignment records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamTrustProductChannelEndpointAssignment(TrustProductSid string, params *ListTrustProductChannelEndpointAssignmentParams) (chan TrusthubV1TrustProductChannelEndpointAssignment, chan error) {
	if params == nil {
		params = &ListTrustProductChannelEndpointAssignmentParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan TrusthubV1TrustProductChannelEndpointAssignment, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageTrustProductChannelEndpointAssignment(TrustProductSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamTrustProductChannelEndpointAssignment(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamTrustProductChannelEndpointAssignment(response *ListTrustProductChannelEndpointAssignmentResponse, params *ListTrustProductChannelEndpointAssignmentParams, recordChannel chan TrusthubV1TrustProductChannelEndpointAssignment, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Results
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListTrustProductChannelEndpointAssignmentResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListTrustProductChannelEndpointAssignmentResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListTrustProductChannelEndpointAssignmentResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTrustProductChannelEndpointAssignmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
