/*
 * Twilio - Wireless
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

	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'ListAccountUsageRecord'
type ListAccountUsageRecordParams struct {
	// Only include usage that has occurred on or before this date. Format is [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html).
	End *time.Time `json:"End,omitempty"`
	// Only include usage that has occurred on or after this date. Format is [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html).
	Start *time.Time `json:"Start,omitempty"`
	// How to summarize the usage by time. Can be: `daily`, `hourly`, or `all`. A value of `all` returns one Usage Record that describes the usage for the entire period.
	Granularity *string `json:"Granularity,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListAccountUsageRecordParams) SetEnd(End time.Time) *ListAccountUsageRecordParams {
	params.End = &End
	return params
}
func (params *ListAccountUsageRecordParams) SetStart(Start time.Time) *ListAccountUsageRecordParams {
	params.Start = &Start
	return params
}
func (params *ListAccountUsageRecordParams) SetGranularity(Granularity string) *ListAccountUsageRecordParams {
	params.Granularity = &Granularity
	return params
}
func (params *ListAccountUsageRecordParams) SetPageSize(PageSize int) *ListAccountUsageRecordParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListAccountUsageRecordParams) SetLimit(Limit int) *ListAccountUsageRecordParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of AccountUsageRecord records from the API. Request is executed immediately.
func (c *ApiService) PageAccountUsageRecord(params *ListAccountUsageRecordParams, pageToken, pageNumber string) (*ListAccountUsageRecordResponse, error) {
	path := "/v1/UsageRecords"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.End != nil {
		data.Set("End", fmt.Sprint((*params.End).Format(time.RFC3339)))
	}
	if params != nil && params.Start != nil {
		data.Set("Start", fmt.Sprint((*params.Start).Format(time.RFC3339)))
	}
	if params != nil && params.Granularity != nil {
		data.Set("Granularity", *params.Granularity)
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

	ps := &ListAccountUsageRecordResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists AccountUsageRecord records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListAccountUsageRecord(params *ListAccountUsageRecordParams) ([]WirelessV1AccountUsageRecord, error) {
	if params == nil {
		params = &ListAccountUsageRecordParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageAccountUsageRecord(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []WirelessV1AccountUsageRecord

	for response != nil {
		records = append(records, response.UsageRecords...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListAccountUsageRecordResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListAccountUsageRecordResponse)
	}

	return records, err
}

// Streams AccountUsageRecord records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamAccountUsageRecord(params *ListAccountUsageRecordParams) (chan WirelessV1AccountUsageRecord, error) {
	if params == nil {
		params = &ListAccountUsageRecordParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageAccountUsageRecord(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan WirelessV1AccountUsageRecord, 1)

	go func() {
		for response != nil {
			for item := range response.UsageRecords {
				channel <- response.UsageRecords[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListAccountUsageRecordResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListAccountUsageRecordResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListAccountUsageRecordResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAccountUsageRecordResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
