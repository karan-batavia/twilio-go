/*
 * Twilio - Video
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
	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateRoom'
type CreateRoomParams struct {
	// Deprecated, now always considered to be true.
	EnableTurn *bool `json:"EnableTurn,omitempty"`
	// The maximum number of concurrent Participants allowed in the room. Peer-to-peer rooms can have up to 10 Participants. Small Group rooms can have up to 4 Participants. Group rooms can have up to 50 Participants.
	MaxParticipants *int `json:"MaxParticipants,omitempty"`
	// The region for the media server in Group Rooms.  Can be: one of the [available Media Regions](https://www.twilio.com/docs/video/ip-address-whitelisting#group-rooms-media-servers). ***This feature is not available in `peer-to-peer` rooms.***
	MediaRegion *string `json:"MediaRegion,omitempty"`
	// Whether to start recording when Participants connect. ***This feature is not available in `peer-to-peer` rooms.***
	RecordParticipantsOnConnect *bool `json:"RecordParticipantsOnConnect,omitempty"`
	// A collection of Recording Rules that describe how to include or exclude matching tracks for recording
	RecordingRules *map[string]interface{} `json:"RecordingRules,omitempty"`
	// The URL we should call using the `status_callback_method` to send status information to your application on every room event. See [Status Callbacks](https://www.twilio.com/docs/video/api/status-callbacks) for more info.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The HTTP method we should use to call `status_callback`. Can be `POST` or `GET`.
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
	// The type of room. Can be: `go`, `peer-to-peer`, `group-small`, or `group`. The default value is `group`.
	Type *string `json:"Type,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used as a `room_sid` in place of the resource's `sid` in the URL to address the resource. This value is unique for `in-progress` rooms. SDK clients can use this name to connect to the room. REST API clients can use this name in place of the Room SID to interact with the room as long as the room is `in-progress`.
	UniqueName *string `json:"UniqueName,omitempty"`
	// An array of the video codecs that are supported when publishing a track in the room.  Can be: `VP8` and `H264`.  ***This feature is not available in `peer-to-peer` rooms***
	VideoCodecs *[]string `json:"VideoCodecs,omitempty"`
}

func (params *CreateRoomParams) SetEnableTurn(EnableTurn bool) *CreateRoomParams {
	params.EnableTurn = &EnableTurn
	return params
}
func (params *CreateRoomParams) SetMaxParticipants(MaxParticipants int) *CreateRoomParams {
	params.MaxParticipants = &MaxParticipants
	return params
}
func (params *CreateRoomParams) SetMediaRegion(MediaRegion string) *CreateRoomParams {
	params.MediaRegion = &MediaRegion
	return params
}
func (params *CreateRoomParams) SetRecordParticipantsOnConnect(RecordParticipantsOnConnect bool) *CreateRoomParams {
	params.RecordParticipantsOnConnect = &RecordParticipantsOnConnect
	return params
}
func (params *CreateRoomParams) SetRecordingRules(RecordingRules map[string]interface{}) *CreateRoomParams {
	params.RecordingRules = &RecordingRules
	return params
}
func (params *CreateRoomParams) SetStatusCallback(StatusCallback string) *CreateRoomParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *CreateRoomParams) SetStatusCallbackMethod(StatusCallbackMethod string) *CreateRoomParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}
func (params *CreateRoomParams) SetType(Type string) *CreateRoomParams {
	params.Type = &Type
	return params
}
func (params *CreateRoomParams) SetUniqueName(UniqueName string) *CreateRoomParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *CreateRoomParams) SetVideoCodecs(VideoCodecs []string) *CreateRoomParams {
	params.VideoCodecs = &VideoCodecs
	return params
}

func (c *ApiService) CreateRoom(params *CreateRoomParams) (*VideoV1Room, error) {
	path := "/v1/Rooms"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.EnableTurn != nil {
		data.Set("EnableTurn", fmt.Sprint(*params.EnableTurn))
	}
	if params != nil && params.MaxParticipants != nil {
		data.Set("MaxParticipants", fmt.Sprint(*params.MaxParticipants))
	}
	if params != nil && params.MediaRegion != nil {
		data.Set("MediaRegion", *params.MediaRegion)
	}
	if params != nil && params.RecordParticipantsOnConnect != nil {
		data.Set("RecordParticipantsOnConnect", fmt.Sprint(*params.RecordParticipantsOnConnect))
	}
	if params != nil && params.RecordingRules != nil {
		v, err := json.Marshal(params.RecordingRules)

		if err != nil {
			return nil, err
		}

		data.Set("RecordingRules", string(v))
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}
	if params != nil && params.Type != nil {
		data.Set("Type", *params.Type)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.VideoCodecs != nil {
		for _, item := range *params.VideoCodecs {
			data.Add("VideoCodecs", item)
		}
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1Room{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) FetchRoom(Sid string) (*VideoV1Room, error) {
	path := "/v1/Rooms/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1Room{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRoom'
type ListRoomParams struct {
	// Read only the rooms with this status. Can be: `in-progress` (default) or `completed`
	Status *string `json:"Status,omitempty"`
	// Read only rooms with the this `unique_name`.
	UniqueName *string `json:"UniqueName,omitempty"`
	// Read only rooms that started on or after this date, given as `YYYY-MM-DD`.
	DateCreatedAfter *time.Time `json:"DateCreatedAfter,omitempty"`
	// Read only rooms that started before this date, given as `YYYY-MM-DD`.
	DateCreatedBefore *time.Time `json:"DateCreatedBefore,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListRoomParams) SetStatus(Status string) *ListRoomParams {
	params.Status = &Status
	return params
}
func (params *ListRoomParams) SetUniqueName(UniqueName string) *ListRoomParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *ListRoomParams) SetDateCreatedAfter(DateCreatedAfter time.Time) *ListRoomParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *ListRoomParams) SetDateCreatedBefore(DateCreatedBefore time.Time) *ListRoomParams {
	params.DateCreatedBefore = &DateCreatedBefore
	return params
}
func (params *ListRoomParams) SetPageSize(PageSize int) *ListRoomParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListRoomParams) SetLimit(Limit int) *ListRoomParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Room records from the API. Request is executed immediately.
func (c *ApiService) PageRoom(params *ListRoomParams, pageToken, pageNumber string) (*ListRoomResponse, error) {
	path := "/v1/Rooms"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreatedAfter", fmt.Sprint((*params.DateCreatedAfter).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedBefore != nil {
		data.Set("DateCreatedBefore", fmt.Sprint((*params.DateCreatedBefore).Format(time.RFC3339)))
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

	ps := &ListRoomResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Room records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRoom(params *ListRoomParams) ([]VideoV1Room, error) {
	if params == nil {
		params = &ListRoomParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageRoom(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []VideoV1Room

	for response != nil {
		records = append(records, response.Rooms...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListRoomResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListRoomResponse)
	}

	return records, err
}

// Streams Room records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRoom(params *ListRoomParams) (chan VideoV1Room, error) {
	if params == nil {
		params = &ListRoomParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageRoom(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan VideoV1Room, 1)

	go func() {
		for response != nil {
			for item := range response.Rooms {
				channel <- response.Rooms[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListRoomResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListRoomResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListRoomResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRoomResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateRoom'
type UpdateRoomParams struct {
	// The new status of the resource. Set to `completed` to end the room.
	Status *string `json:"Status,omitempty"`
}

func (params *UpdateRoomParams) SetStatus(Status string) *UpdateRoomParams {
	params.Status = &Status
	return params
}

func (c *ApiService) UpdateRoom(Sid string, params *UpdateRoomParams) (*VideoV1Room, error) {
	path := "/v1/Rooms/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1Room{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
