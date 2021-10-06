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

	"github.com/twilio/twilio-go/client"
)

// Returns a single Track resource represented by TrackName or SID.
func (c *ApiService) FetchRoomParticipantPublishedTrack(RoomSid string, ParticipantSid string, Sid string) (*VideoV1RoomParticipantPublishedTrack, error) {
	path := "/v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/PublishedTracks/{Sid}"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
	path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1RoomParticipantPublishedTrack{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRoomParticipantPublishedTrack'
type ListRoomParticipantPublishedTrackParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListRoomParticipantPublishedTrackParams) SetPageSize(PageSize int) *ListRoomParticipantPublishedTrackParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListRoomParticipantPublishedTrackParams) SetLimit(Limit int) *ListRoomParticipantPublishedTrackParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of RoomParticipantPublishedTrack records from the API. Request is executed immediately.
func (c *ApiService) PageRoomParticipantPublishedTrack(RoomSid string, ParticipantSid string, params *ListRoomParticipantPublishedTrackParams, pageToken, pageNumber string) (*ListRoomParticipantPublishedTrackResponse, error) {
	path := "/v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/PublishedTracks"

	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
	path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)

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

	ps := &ListRoomParticipantPublishedTrackResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists RoomParticipantPublishedTrack records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRoomParticipantPublishedTrack(RoomSid string, ParticipantSid string, params *ListRoomParticipantPublishedTrackParams) ([]VideoV1RoomParticipantPublishedTrack, error) {
	if params == nil {
		params = &ListRoomParticipantPublishedTrackParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageRoomParticipantPublishedTrack(RoomSid, ParticipantSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []VideoV1RoomParticipantPublishedTrack

	for response != nil {
		records = append(records, response.PublishedTracks...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListRoomParticipantPublishedTrackResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListRoomParticipantPublishedTrackResponse)
	}

	return records, err
}

// Streams RoomParticipantPublishedTrack records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRoomParticipantPublishedTrack(RoomSid string, ParticipantSid string, params *ListRoomParticipantPublishedTrackParams) (chan VideoV1RoomParticipantPublishedTrack, error) {
	if params == nil {
		params = &ListRoomParticipantPublishedTrackParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageRoomParticipantPublishedTrack(RoomSid, ParticipantSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan VideoV1RoomParticipantPublishedTrack, 1)

	go func() {
		for response != nil {
			for item := range response.PublishedTracks {
				channel <- response.PublishedTracks[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListRoomParticipantPublishedTrackResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListRoomParticipantPublishedTrackResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListRoomParticipantPublishedTrackResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRoomParticipantPublishedTrackResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
