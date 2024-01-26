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

// Optional parameters for the method 'CreateCustomerProfileChannelEndpointAssignment'
type CreateCustomerProfileChannelEndpointAssignmentParams struct {
	// The type of channel endpoint. eg: phone-number
	ChannelEndpointType *string `json:"ChannelEndpointType,omitempty"`
	// The SID of an channel endpoint
	ChannelEndpointSid *string `json:"ChannelEndpointSid,omitempty"`
}

func (params *CreateCustomerProfileChannelEndpointAssignmentParams) SetChannelEndpointType(ChannelEndpointType string) *CreateCustomerProfileChannelEndpointAssignmentParams {
	params.ChannelEndpointType = &ChannelEndpointType
	return params
}
func (params *CreateCustomerProfileChannelEndpointAssignmentParams) SetChannelEndpointSid(ChannelEndpointSid string) *CreateCustomerProfileChannelEndpointAssignmentParams {
	params.ChannelEndpointSid = &ChannelEndpointSid
	return params
}

// Create a new Assigned Item.
func (c *ApiService) CreateCustomerProfileChannelEndpointAssignment(CustomerProfileSid string, params *CreateCustomerProfileChannelEndpointAssignmentParams) (*TrusthubV1CustomerProfileChannelEndpointAssignment, error) {
	path := "/v1/CustomerProfiles/{CustomerProfileSid}/ChannelEndpointAssignments"
	path = strings.Replace(path, "{"+"CustomerProfileSid"+"}", CustomerProfileSid, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ChannelEndpointType != nil {
		data.Set("ChannelEndpointType", *params.ChannelEndpointType)
	}
	if params != nil && params.ChannelEndpointSid != nil {
		data.Set("ChannelEndpointSid", *params.ChannelEndpointSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers, queryParams)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1CustomerProfileChannelEndpointAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an Assignment Item Instance.
func (c *ApiService) DeleteCustomerProfileChannelEndpointAssignment(CustomerProfileSid string, Sid string) error {
	path := "/v1/CustomerProfiles/{CustomerProfileSid}/ChannelEndpointAssignments/{Sid}"
	path = strings.Replace(path, "{"+"CustomerProfileSid"+"}", CustomerProfileSid, -1)
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

// Fetch specific Assigned Item Instance.
func (c *ApiService) FetchCustomerProfileChannelEndpointAssignment(CustomerProfileSid string, Sid string) (*TrusthubV1CustomerProfileChannelEndpointAssignment, error) {
	path := "/v1/CustomerProfiles/{CustomerProfileSid}/ChannelEndpointAssignments/{Sid}"
	path = strings.Replace(path, "{"+"CustomerProfileSid"+"}", CustomerProfileSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers, queryParams)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1CustomerProfileChannelEndpointAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListCustomerProfileChannelEndpointAssignment'
type ListCustomerProfileChannelEndpointAssignmentParams struct {
	// The SID of an channel endpoint
	ChannelEndpointSid *string `json:"ChannelEndpointSid,omitempty"`
	// comma separated list of channel endpoint sids
	ChannelEndpointSids *string `json:"ChannelEndpointSids,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListCustomerProfileChannelEndpointAssignmentParams) SetChannelEndpointSid(ChannelEndpointSid string) *ListCustomerProfileChannelEndpointAssignmentParams {
	params.ChannelEndpointSid = &ChannelEndpointSid
	return params
}
func (params *ListCustomerProfileChannelEndpointAssignmentParams) SetChannelEndpointSids(ChannelEndpointSids string) *ListCustomerProfileChannelEndpointAssignmentParams {
	params.ChannelEndpointSids = &ChannelEndpointSids
	return params
}
func (params *ListCustomerProfileChannelEndpointAssignmentParams) SetPageSize(PageSize int) *ListCustomerProfileChannelEndpointAssignmentParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListCustomerProfileChannelEndpointAssignmentParams) SetLimit(Limit int) *ListCustomerProfileChannelEndpointAssignmentParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of CustomerProfileChannelEndpointAssignment records from the API. Request is executed immediately.
func (c *ApiService) PageCustomerProfileChannelEndpointAssignment(CustomerProfileSid string, params *ListCustomerProfileChannelEndpointAssignmentParams, pageToken, pageNumber string) (*ListCustomerProfileChannelEndpointAssignmentResponse, error) {
	path := "/v1/CustomerProfiles/{CustomerProfileSid}/ChannelEndpointAssignments"

	path = strings.Replace(path, "{"+"CustomerProfileSid"+"}", CustomerProfileSid, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ChannelEndpointSid != nil {
		queryParams.Set("ChannelEndpointSid", *params.ChannelEndpointSid)
	}
	if params != nil && params.ChannelEndpointSids != nil {
		queryParams.Set("ChannelEndpointSids", *params.ChannelEndpointSids)
	}
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

	ps := &ListCustomerProfileChannelEndpointAssignmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists CustomerProfileChannelEndpointAssignment records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCustomerProfileChannelEndpointAssignment(CustomerProfileSid string, params *ListCustomerProfileChannelEndpointAssignmentParams) ([]TrusthubV1CustomerProfileChannelEndpointAssignment, error) {
	response, errors := c.StreamCustomerProfileChannelEndpointAssignment(CustomerProfileSid, params)

	records := make([]TrusthubV1CustomerProfileChannelEndpointAssignment, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams CustomerProfileChannelEndpointAssignment records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCustomerProfileChannelEndpointAssignment(CustomerProfileSid string, params *ListCustomerProfileChannelEndpointAssignmentParams) (chan TrusthubV1CustomerProfileChannelEndpointAssignment, chan error) {
	if params == nil {
		params = &ListCustomerProfileChannelEndpointAssignmentParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan TrusthubV1CustomerProfileChannelEndpointAssignment, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageCustomerProfileChannelEndpointAssignment(CustomerProfileSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamCustomerProfileChannelEndpointAssignment(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamCustomerProfileChannelEndpointAssignment(response *ListCustomerProfileChannelEndpointAssignmentResponse, params *ListCustomerProfileChannelEndpointAssignmentParams, recordChannel chan TrusthubV1CustomerProfileChannelEndpointAssignment, errorChannel chan error) {
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

		record, err := client.GetNext(c.baseURL, response, c.getNextListCustomerProfileChannelEndpointAssignmentResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListCustomerProfileChannelEndpointAssignmentResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListCustomerProfileChannelEndpointAssignmentResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCustomerProfileChannelEndpointAssignmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
