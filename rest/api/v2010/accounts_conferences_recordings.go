/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
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

// Optional parameters for the method 'DeleteConferenceRecording'
type DeleteConferenceRecordingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteConferenceRecordingParams) SetPathAccountSid(PathAccountSid string) *DeleteConferenceRecordingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete a recording from your account
func (c *ApiService) DeleteConferenceRecording(ConferenceSid string, Sid string, params *DeleteConferenceRecordingParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Recordings/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ConferenceSid"+"}", ConferenceSid, -1)
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

// Optional parameters for the method 'FetchConferenceRecording'
type FetchConferenceRecordingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchConferenceRecordingParams) SetPathAccountSid(PathAccountSid string) *FetchConferenceRecordingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an instance of a recording for a call
func (c *ApiService) FetchConferenceRecording(ConferenceSid string, Sid string, params *FetchConferenceRecordingParams) (*ApiV2010ConferenceRecording, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Recordings/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ConferenceSid"+"}", ConferenceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers, queryParams)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010ConferenceRecording{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListConferenceRecording'
type ListConferenceRecordingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. You can also specify inequality: `DateCreated<=YYYY-MM-DD` will return recordings generated at or before midnight on a given date, and `DateCreated>=YYYY-MM-DD` returns recordings generated at or after midnight on a date.
	DateCreated *string `json:"DateCreated,omitempty"`
	// The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. You can also specify inequality: `DateCreated<=YYYY-MM-DD` will return recordings generated at or before midnight on a given date, and `DateCreated>=YYYY-MM-DD` returns recordings generated at or after midnight on a date.
	DateCreatedBefore *string `json:"DateCreated&lt;,omitempty"`
	// The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. You can also specify inequality: `DateCreated<=YYYY-MM-DD` will return recordings generated at or before midnight on a given date, and `DateCreated>=YYYY-MM-DD` returns recordings generated at or after midnight on a date.
	DateCreatedAfter *string `json:"DateCreated&gt;,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListConferenceRecordingParams) SetPathAccountSid(PathAccountSid string) *ListConferenceRecordingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListConferenceRecordingParams) SetDateCreated(DateCreated string) *ListConferenceRecordingParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *ListConferenceRecordingParams) SetDateCreatedBefore(DateCreatedBefore string) *ListConferenceRecordingParams {
	params.DateCreatedBefore = &DateCreatedBefore
	return params
}
func (params *ListConferenceRecordingParams) SetDateCreatedAfter(DateCreatedAfter string) *ListConferenceRecordingParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *ListConferenceRecordingParams) SetPageSize(PageSize int) *ListConferenceRecordingParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListConferenceRecordingParams) SetLimit(Limit int) *ListConferenceRecordingParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ConferenceRecording records from the API. Request is executed immediately.
func (c *ApiService) PageConferenceRecording(ConferenceSid string, params *ListConferenceRecordingParams, pageToken, pageNumber string) (*ListConferenceRecordingResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Recordings.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ConferenceSid"+"}", ConferenceSid, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.DateCreated != nil {
		queryParams.Set("DateCreated", fmt.Sprint(*params.DateCreated))
	}
	if params != nil && params.DateCreatedBefore != nil {
		queryParams.Set("DateCreated<", fmt.Sprint(*params.DateCreatedBefore))
	}
	if params != nil && params.DateCreatedAfter != nil {
		queryParams.Set("DateCreated>", fmt.Sprint(*params.DateCreatedAfter))
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

	ps := &ListConferenceRecordingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ConferenceRecording records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListConferenceRecording(ConferenceSid string, params *ListConferenceRecordingParams) ([]ApiV2010ConferenceRecording, error) {
	response, errors := c.StreamConferenceRecording(ConferenceSid, params)

	records := make([]ApiV2010ConferenceRecording, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams ConferenceRecording records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamConferenceRecording(ConferenceSid string, params *ListConferenceRecordingParams) (chan ApiV2010ConferenceRecording, chan error) {
	if params == nil {
		params = &ListConferenceRecordingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ApiV2010ConferenceRecording, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageConferenceRecording(ConferenceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamConferenceRecording(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamConferenceRecording(response *ListConferenceRecordingResponse, params *ListConferenceRecordingParams, recordChannel chan ApiV2010ConferenceRecording, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Recordings
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListConferenceRecordingResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListConferenceRecordingResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListConferenceRecordingResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListConferenceRecordingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateConferenceRecording'
type UpdateConferenceRecordingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resource to update.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	//
	Status *string `json:"Status,omitempty"`
	// Whether to record during a pause. Can be: `skip` or `silence` and the default is `silence`. `skip` does not record during the pause period, while `silence` will replace the actual audio of the call with silence during the pause period. This parameter only applies when setting `status` is set to `paused`.
	PauseBehavior *string `json:"PauseBehavior,omitempty"`
}

func (params *UpdateConferenceRecordingParams) SetPathAccountSid(PathAccountSid string) *UpdateConferenceRecordingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateConferenceRecordingParams) SetStatus(Status string) *UpdateConferenceRecordingParams {
	params.Status = &Status
	return params
}
func (params *UpdateConferenceRecordingParams) SetPauseBehavior(PauseBehavior string) *UpdateConferenceRecordingParams {
	params.PauseBehavior = &PauseBehavior
	return params
}

// Changes the status of the recording to paused, stopped, or in-progress. Note: To use `Twilio.CURRENT`, pass it as recording sid.
func (c *ApiService) UpdateConferenceRecording(ConferenceSid string, Sid string, params *UpdateConferenceRecordingParams) (*ApiV2010ConferenceRecording, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Recordings/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ConferenceSid"+"}", ConferenceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.PauseBehavior != nil {
		data.Set("PauseBehavior", *params.PauseBehavior)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers, queryParams)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010ConferenceRecording{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
