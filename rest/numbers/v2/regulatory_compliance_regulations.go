/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Numbers
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

// Fetch specific Regulation Instance.
func (c *ApiService) FetchRegulation(Sid string) (*NumbersV2Regulation, error) {
	path := "/v2/RegulatoryCompliance/Regulations/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers, queryParams)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2Regulation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRegulation'
type ListRegulationParams struct {
	// The type of End User the regulation requires - can be `individual` or `business`.
	EndUserType *string `json:"EndUserType,omitempty"`
	// The ISO country code of the phone number's country.
	IsoCountry *string `json:"IsoCountry,omitempty"`
	// The type of phone number that the regulatory requiremnt is restricting.
	NumberType *string `json:"NumberType,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListRegulationParams) SetEndUserType(EndUserType string) *ListRegulationParams {
	params.EndUserType = &EndUserType
	return params
}
func (params *ListRegulationParams) SetIsoCountry(IsoCountry string) *ListRegulationParams {
	params.IsoCountry = &IsoCountry
	return params
}
func (params *ListRegulationParams) SetNumberType(NumberType string) *ListRegulationParams {
	params.NumberType = &NumberType
	return params
}
func (params *ListRegulationParams) SetPageSize(PageSize int) *ListRegulationParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListRegulationParams) SetLimit(Limit int) *ListRegulationParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Regulation records from the API. Request is executed immediately.
func (c *ApiService) PageRegulation(params *ListRegulationParams, pageToken, pageNumber string) (*ListRegulationResponse, error) {
	path := "/v2/RegulatoryCompliance/Regulations"

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.EndUserType != nil {
		queryParams.Set("EndUserType", *params.EndUserType)
	}
	if params != nil && params.IsoCountry != nil {
		queryParams.Set("IsoCountry", *params.IsoCountry)
	}
	if params != nil && params.NumberType != nil {
		queryParams.Set("NumberType", *params.NumberType)
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

	ps := &ListRegulationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Regulation records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRegulation(params *ListRegulationParams) ([]NumbersV2Regulation, error) {
	response, errors := c.StreamRegulation(params)

	records := make([]NumbersV2Regulation, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Regulation records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRegulation(params *ListRegulationParams) (chan NumbersV2Regulation, chan error) {
	if params == nil {
		params = &ListRegulationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan NumbersV2Regulation, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageRegulation(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamRegulation(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamRegulation(response *ListRegulationResponse, params *ListRegulationParams, recordChannel chan NumbersV2Regulation, errorChannel chan error) {
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

		record, err := client.GetNext(c.baseURL, response, c.getNextListRegulationResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListRegulationResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListRegulationResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRegulationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
