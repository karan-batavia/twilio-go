/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Supersim
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
	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'ListUsageRecord'
type ListUsageRecordParams struct {
	// SID or unique name of a Sim resource. Only show UsageRecords representing usage incurred by this Super SIM.
	Sim *string `json:"Sim,omitempty"`
	// SID or unique name of a Fleet resource. Only show UsageRecords representing usage for Super SIMs belonging to this Fleet resource at the time the usage occurred.
	Fleet *string `json:"Fleet,omitempty"`
	// SID of a Network resource. Only show UsageRecords representing usage on this network.
	Network *string `json:"Network,omitempty"`
	// Alpha-2 ISO Country Code. Only show UsageRecords representing usage in this country.
	IsoCountry *string `json:"IsoCountry,omitempty"`
	// Dimension over which to aggregate usage records. Can be: `sim`, `fleet`, `network`, `isoCountry`. Default is to not aggregate across any of these dimensions, UsageRecords will be aggregated into the time buckets described by the `Granularity` parameter.
	Group *string `json:"Group,omitempty"`
	// Time-based grouping that UsageRecords should be aggregated by. Can be: `hour`, `day`, or `all`. Default is `all`. `all` returns one UsageRecord that describes the usage for the entire period.
	Granularity *string `json:"Granularity,omitempty"`
	// Only include usage that occurred at or after this time, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Default is one month before the `end_time`.
	StartTime *time.Time `json:"StartTime,omitempty"`
	// Only include usage that occurred before this time (exclusive), specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Default is the current time.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListUsageRecordParams) SetSim(Sim string) *ListUsageRecordParams {
	params.Sim = &Sim
	return params
}
func (params *ListUsageRecordParams) SetFleet(Fleet string) *ListUsageRecordParams {
	params.Fleet = &Fleet
	return params
}
func (params *ListUsageRecordParams) SetNetwork(Network string) *ListUsageRecordParams {
	params.Network = &Network
	return params
}
func (params *ListUsageRecordParams) SetIsoCountry(IsoCountry string) *ListUsageRecordParams {
	params.IsoCountry = &IsoCountry
	return params
}
func (params *ListUsageRecordParams) SetGroup(Group string) *ListUsageRecordParams {
	params.Group = &Group
	return params
}
func (params *ListUsageRecordParams) SetGranularity(Granularity string) *ListUsageRecordParams {
	params.Granularity = &Granularity
	return params
}
func (params *ListUsageRecordParams) SetStartTime(StartTime time.Time) *ListUsageRecordParams {
	params.StartTime = &StartTime
	return params
}
func (params *ListUsageRecordParams) SetEndTime(EndTime time.Time) *ListUsageRecordParams {
	params.EndTime = &EndTime
	return params
}
func (params *ListUsageRecordParams) SetPageSize(PageSize int) *ListUsageRecordParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListUsageRecordParams) SetLimit(Limit int) *ListUsageRecordParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of UsageRecord records from the API. Request is executed immediately.
func (c *ApiService) PageUsageRecord(params *ListUsageRecordParams, pageToken, pageNumber string) (*ListUsageRecordResponse, error) {
	path := "/v1/UsageRecords"

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Sim != nil {
		queryParams.Set("Sim", *params.Sim)
	}
	if params != nil && params.Fleet != nil {
		queryParams.Set("Fleet", *params.Fleet)
	}
	if params != nil && params.Network != nil {
		queryParams.Set("Network", *params.Network)
	}
	if params != nil && params.IsoCountry != nil {
		queryParams.Set("IsoCountry", *params.IsoCountry)
	}
	if params != nil && params.Group != nil {
		queryParams.Set("Group", *params.Group)
	}
	if params != nil && params.Granularity != nil {
		queryParams.Set("Granularity", *params.Granularity)
	}
	if params != nil && params.StartTime != nil {
		queryParams.Set("StartTime", fmt.Sprint((*params.StartTime).Format(time.RFC3339)))
	}
	if params != nil && params.EndTime != nil {
		queryParams.Set("EndTime", fmt.Sprint((*params.EndTime).Format(time.RFC3339)))
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

	ps := &ListUsageRecordResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists UsageRecord records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListUsageRecord(params *ListUsageRecordParams) ([]SupersimV1UsageRecord, error) {
	response, errors := c.StreamUsageRecord(params)

	records := make([]SupersimV1UsageRecord, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams UsageRecord records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamUsageRecord(params *ListUsageRecordParams) (chan SupersimV1UsageRecord, chan error) {
	if params == nil {
		params = &ListUsageRecordParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan SupersimV1UsageRecord, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageUsageRecord(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamUsageRecord(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamUsageRecord(response *ListUsageRecordResponse, params *ListUsageRecordParams, recordChannel chan SupersimV1UsageRecord, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.UsageRecords
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListUsageRecordResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListUsageRecordResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListUsageRecordResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUsageRecordResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
