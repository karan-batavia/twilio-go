/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Intelligence
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

// Optional parameters for the method 'FetchOperatorResult'
type FetchOperatorResultParams struct {
	// Grant access to PII redacted/unredacted Language Understanding operator. If redaction is enabled, the default is True.
	Redacted *bool `json:"Redacted,omitempty"`
}

func (params *FetchOperatorResultParams) SetRedacted(Redacted bool) *FetchOperatorResultParams {
	params.Redacted = &Redacted
	return params
}

// Fetch a specific Operator Result for the given Transcript.
func (c *ApiService) FetchOperatorResult(TranscriptSid string, OperatorSid string, params *FetchOperatorResultParams) (*IntelligenceV2OperatorResult, error) {
	path := "/v2/Transcripts/{TranscriptSid}/OperatorResults/{OperatorSid}"
	path = strings.Replace(path, "{"+"TranscriptSid"+"}", TranscriptSid, -1)
	path = strings.Replace(path, "{"+"OperatorSid"+"}", OperatorSid, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Redacted != nil {
		queryParams.Set("Redacted", fmt.Sprint(*params.Redacted))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers, queryParams)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IntelligenceV2OperatorResult{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListOperatorResult'
type ListOperatorResultParams struct {
	// Grant access to PII redacted/unredacted Language Understanding operator. If redaction is enabled, the default is True.
	Redacted *bool `json:"Redacted,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListOperatorResultParams) SetRedacted(Redacted bool) *ListOperatorResultParams {
	params.Redacted = &Redacted
	return params
}
func (params *ListOperatorResultParams) SetPageSize(PageSize int) *ListOperatorResultParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListOperatorResultParams) SetLimit(Limit int) *ListOperatorResultParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of OperatorResult records from the API. Request is executed immediately.
func (c *ApiService) PageOperatorResult(TranscriptSid string, params *ListOperatorResultParams, pageToken, pageNumber string) (*ListOperatorResultResponse, error) {
	path := "/v2/Transcripts/{TranscriptSid}/OperatorResults"

	path = strings.Replace(path, "{"+"TranscriptSid"+"}", TranscriptSid, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Redacted != nil {
		queryParams.Set("Redacted", fmt.Sprint(*params.Redacted))
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

	ps := &ListOperatorResultResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists OperatorResult records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListOperatorResult(TranscriptSid string, params *ListOperatorResultParams) ([]IntelligenceV2OperatorResult, error) {
	response, errors := c.StreamOperatorResult(TranscriptSid, params)

	records := make([]IntelligenceV2OperatorResult, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams OperatorResult records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamOperatorResult(TranscriptSid string, params *ListOperatorResultParams) (chan IntelligenceV2OperatorResult, chan error) {
	if params == nil {
		params = &ListOperatorResultParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan IntelligenceV2OperatorResult, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageOperatorResult(TranscriptSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamOperatorResult(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamOperatorResult(response *ListOperatorResultResponse, params *ListOperatorResultParams, recordChannel chan IntelligenceV2OperatorResult, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.OperatorResults
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListOperatorResultResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListOperatorResultResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListOperatorResultResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListOperatorResultResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
