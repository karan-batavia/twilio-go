/*
 * Twilio - Api
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

	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'DeleteRecordingAddOnResultPayload'
type DeleteRecordingAddOnResultPayloadParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteRecordingAddOnResultPayloadParams) SetPathAccountSid(PathAccountSid string) *DeleteRecordingAddOnResultPayloadParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete a payload from the result along with all associated Data
func (c *ApiService) DeleteRecordingAddOnResultPayload(ReferenceSid string, AddOnResultSid string, Sid string, params *DeleteRecordingAddOnResultPayloadParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{AddOnResultSid}/Payloads/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ReferenceSid"+"}", ReferenceSid, -1)
	path = strings.Replace(path, "{"+"AddOnResultSid"+"}", AddOnResultSid, -1)
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

// Optional parameters for the method 'FetchRecordingAddOnResultPayload'
type FetchRecordingAddOnResultPayloadParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchRecordingAddOnResultPayloadParams) SetPathAccountSid(PathAccountSid string) *FetchRecordingAddOnResultPayloadParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an instance of a result payload
func (c *ApiService) FetchRecordingAddOnResultPayload(ReferenceSid string, AddOnResultSid string, Sid string, params *FetchRecordingAddOnResultPayloadParams) (*ApiV2010RecordingAddOnResultPayload, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{AddOnResultSid}/Payloads/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ReferenceSid"+"}", ReferenceSid, -1)
	path = strings.Replace(path, "{"+"AddOnResultSid"+"}", AddOnResultSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010RecordingAddOnResultPayload{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRecordingAddOnResultPayload'
type ListRecordingAddOnResultPayloadParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListRecordingAddOnResultPayloadParams) SetPathAccountSid(PathAccountSid string) *ListRecordingAddOnResultPayloadParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListRecordingAddOnResultPayloadParams) SetPageSize(PageSize int) *ListRecordingAddOnResultPayloadParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListRecordingAddOnResultPayloadParams) SetLimit(Limit int) *ListRecordingAddOnResultPayloadParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of RecordingAddOnResultPayload records from the API. Request is executed immediately.
func (c *ApiService) PageRecordingAddOnResultPayload(ReferenceSid string, AddOnResultSid string, params *ListRecordingAddOnResultPayloadParams, pageToken, pageNumber string) (*ListRecordingAddOnResultPayloadResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{AddOnResultSid}/Payloads.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ReferenceSid"+"}", ReferenceSid, -1)
	path = strings.Replace(path, "{"+"AddOnResultSid"+"}", AddOnResultSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

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

	ps := &ListRecordingAddOnResultPayloadResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists RecordingAddOnResultPayload records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRecordingAddOnResultPayload(ReferenceSid string, AddOnResultSid string, params *ListRecordingAddOnResultPayloadParams) ([]ApiV2010RecordingAddOnResultPayload, error) {
	if params == nil {
		params = &ListRecordingAddOnResultPayloadParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageRecordingAddOnResultPayload(ReferenceSid, AddOnResultSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010RecordingAddOnResultPayload

	for response != nil {
		records = append(records, response.Payloads...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListRecordingAddOnResultPayloadResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListRecordingAddOnResultPayloadResponse)
	}

	return records, err
}

// Streams RecordingAddOnResultPayload records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRecordingAddOnResultPayload(ReferenceSid string, AddOnResultSid string, params *ListRecordingAddOnResultPayloadParams) (chan ApiV2010RecordingAddOnResultPayload, error) {
	if params == nil {
		params = &ListRecordingAddOnResultPayloadParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageRecordingAddOnResultPayload(ReferenceSid, AddOnResultSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010RecordingAddOnResultPayload, 1)

	go func() {
		for response != nil {
			for item := range response.Payloads {
				channel <- response.Payloads[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListRecordingAddOnResultPayloadResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListRecordingAddOnResultPayloadResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListRecordingAddOnResultPayloadResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRecordingAddOnResultPayloadResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
