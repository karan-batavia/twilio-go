/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Insights
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

// Optional parameters for the method 'ListEvent'
type ListEventParams struct {
	// The Edge of this Event. One of `unknown_edge`, `carrier_edge`, `sip_edge`, `sdk_edge` or `client_edge`.
	Edge *string `json:"Edge,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListEventParams) SetEdge(Edge string) *ListEventParams {
	params.Edge = &Edge
	return params
}
func (params *ListEventParams) SetPageSize(PageSize int) *ListEventParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListEventParams) SetLimit(Limit int) *ListEventParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Event records from the API. Request is executed immediately.
func (c *ApiService) PageEvent(CallSid string, params *ListEventParams, pageToken, pageNumber string) (*ListEventResponse, error) {
	path := "/v1/Voice/{CallSid}/Events"

	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Edge != nil {
		queryParams.Set("Edge", *params.Edge)
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

	ps := &ListEventResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Event records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListEvent(CallSid string, params *ListEventParams) ([]InsightsV1Event, error) {
	response, errors := c.StreamEvent(CallSid, params)

	records := make([]InsightsV1Event, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Event records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamEvent(CallSid string, params *ListEventParams) (chan InsightsV1Event, chan error) {
	if params == nil {
		params = &ListEventParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan InsightsV1Event, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageEvent(CallSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamEvent(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamEvent(response *ListEventResponse, params *ListEventParams, recordChannel chan InsightsV1Event, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Events
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListEventResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListEventResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListEventResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEventResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
