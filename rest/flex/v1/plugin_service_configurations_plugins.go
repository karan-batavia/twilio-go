/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
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

// Optional parameters for the method 'FetchConfiguredPlugin'
type FetchConfiguredPluginParams struct {
	// The Flex-Metadata HTTP request header
	FlexMetadata *string `json:"Flex-Metadata,omitempty"`
}

func (params *FetchConfiguredPluginParams) SetFlexMetadata(FlexMetadata string) *FetchConfiguredPluginParams {
	params.FlexMetadata = &FlexMetadata
	return params
}

func (c *ApiService) FetchConfiguredPlugin(ConfigurationSid string, PluginSid string, params *FetchConfiguredPluginParams) (*FlexV1ConfiguredPlugin, error) {
	path := "/v1/PluginService/Configurations/{ConfigurationSid}/Plugins/{PluginSid}"
	path = strings.Replace(path, "{"+"ConfigurationSid"+"}", ConfigurationSid, -1)
	path = strings.Replace(path, "{"+"PluginSid"+"}", PluginSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FlexMetadata != nil {
		headers["Flex-Metadata"] = *params.FlexMetadata
	}
	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1ConfiguredPlugin{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListConfiguredPlugin'
type ListConfiguredPluginParams struct {
	// The Flex-Metadata HTTP request header
	FlexMetadata *string `json:"Flex-Metadata,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListConfiguredPluginParams) SetFlexMetadata(FlexMetadata string) *ListConfiguredPluginParams {
	params.FlexMetadata = &FlexMetadata
	return params
}
func (params *ListConfiguredPluginParams) SetPageSize(PageSize int) *ListConfiguredPluginParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListConfiguredPluginParams) SetLimit(Limit int) *ListConfiguredPluginParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ConfiguredPlugin records from the API. Request is executed immediately.
func (c *ApiService) PageConfiguredPlugin(ConfigurationSid string, params *ListConfiguredPluginParams, pageToken, pageNumber string) (*ListConfiguredPluginResponse, error) {
	path := "/v1/PluginService/Configurations/{ConfigurationSid}/Plugins"

	path = strings.Replace(path, "{"+"ConfigurationSid"+"}", ConfigurationSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
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

	ps := &ListConfiguredPluginResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ConfiguredPlugin records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListConfiguredPlugin(ConfigurationSid string, params *ListConfiguredPluginParams) ([]FlexV1ConfiguredPlugin, error) {
	response, errors := c.StreamConfiguredPlugin(ConfigurationSid, params)

	records := make([]FlexV1ConfiguredPlugin, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams ConfiguredPlugin records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamConfiguredPlugin(ConfigurationSid string, params *ListConfiguredPluginParams) (chan FlexV1ConfiguredPlugin, chan error) {
	if params == nil {
		params = &ListConfiguredPluginParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan FlexV1ConfiguredPlugin, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageConfiguredPlugin(ConfigurationSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamConfiguredPlugin(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamConfiguredPlugin(response *ListConfiguredPluginResponse, params *ListConfiguredPluginParams, recordChannel chan FlexV1ConfiguredPlugin, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Plugins
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListConfiguredPluginResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListConfiguredPluginResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListConfiguredPluginResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListConfiguredPluginResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
