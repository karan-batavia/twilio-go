/*
 * Twilio - Fax
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

// Delete a specific fax media instance.
func (c *ApiService) DeleteFaxMedia(FaxSid string, Sid string) error {
	path := "/v1/Faxes/{FaxSid}/Media/{Sid}"
	path = strings.Replace(path, "{"+"FaxSid"+"}", FaxSid, -1)
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

// Fetch a specific fax media instance.
func (c *ApiService) FetchFaxMedia(FaxSid string, Sid string) (*FaxV1FaxFaxMedia, error) {
	path := "/v1/Faxes/{FaxSid}/Media/{Sid}"
	path = strings.Replace(path, "{"+"FaxSid"+"}", FaxSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FaxV1FaxFaxMedia{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFaxMedia'
type ListFaxMediaParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListFaxMediaParams) SetPageSize(PageSize int) *ListFaxMediaParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a single page of FaxMedia records from the API. Request is executed immediately.
func (c *ApiService) PageFaxMedia(FaxSid string, params *ListFaxMediaParams, pageToken string, pageNumber string) (*ListFaxMediaResponse, error) {
	path := "/v1/Faxes/{FaxSid}/Media"

	path = strings.Replace(path, "{"+"FaxSid"+"}", FaxSid, -1)

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

	ps := &ListFaxMediaResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists FaxMedia records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListFaxMedia(FaxSid string, params *ListFaxMediaParams, limit int) ([]FaxV1FaxFaxMedia, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageFaxMedia(FaxSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []FaxV1FaxFaxMedia

	for response != nil {
		records = append(records, response.Media...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListFaxMediaResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListFaxMediaResponse)
	}

	return records, err
}

// Streams FaxMedia records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamFaxMedia(FaxSid string, params *ListFaxMediaParams, limit int) (chan FaxV1FaxFaxMedia, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageFaxMedia(FaxSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan FaxV1FaxFaxMedia, 1)

	go func() {
		for response != nil {
			for item := range response.Media {
				channel <- response.Media[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListFaxMediaResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListFaxMediaResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListFaxMediaResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFaxMediaResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
