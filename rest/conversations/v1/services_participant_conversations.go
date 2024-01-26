/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Conversations
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

// Optional parameters for the method 'ListServiceParticipantConversation'
type ListServiceParticipantConversationParams struct {
	// A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversations SDK to communicate. Limited to 256 characters.
	Identity *string `json:"Identity,omitempty"`
	// A unique string identifier for the conversation participant who's not a Conversation User. This parameter could be found in messaging_binding.address field of Participant resource. It should be url-encoded.
	Address *string `json:"Address,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListServiceParticipantConversationParams) SetIdentity(Identity string) *ListServiceParticipantConversationParams {
	params.Identity = &Identity
	return params
}
func (params *ListServiceParticipantConversationParams) SetAddress(Address string) *ListServiceParticipantConversationParams {
	params.Address = &Address
	return params
}
func (params *ListServiceParticipantConversationParams) SetPageSize(PageSize int) *ListServiceParticipantConversationParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListServiceParticipantConversationParams) SetLimit(Limit int) *ListServiceParticipantConversationParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ServiceParticipantConversation records from the API. Request is executed immediately.
func (c *ApiService) PageServiceParticipantConversation(ChatServiceSid string, params *ListServiceParticipantConversationParams, pageToken, pageNumber string) (*ListServiceParticipantConversationResponse, error) {
	path := "/v1/Services/{ChatServiceSid}/ParticipantConversations"

	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Identity != nil {
		queryParams.Set("Identity", *params.Identity)
	}
	if params != nil && params.Address != nil {
		queryParams.Set("Address", *params.Address)
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

	ps := &ListServiceParticipantConversationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ServiceParticipantConversation records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListServiceParticipantConversation(ChatServiceSid string, params *ListServiceParticipantConversationParams) ([]ConversationsV1ServiceParticipantConversation, error) {
	response, errors := c.StreamServiceParticipantConversation(ChatServiceSid, params)

	records := make([]ConversationsV1ServiceParticipantConversation, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams ServiceParticipantConversation records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamServiceParticipantConversation(ChatServiceSid string, params *ListServiceParticipantConversationParams) (chan ConversationsV1ServiceParticipantConversation, chan error) {
	if params == nil {
		params = &ListServiceParticipantConversationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ConversationsV1ServiceParticipantConversation, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageServiceParticipantConversation(ChatServiceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamServiceParticipantConversation(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamServiceParticipantConversation(response *ListServiceParticipantConversationResponse, params *ListServiceParticipantConversationParams, recordChannel chan ConversationsV1ServiceParticipantConversation, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Conversations
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListServiceParticipantConversationResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListServiceParticipantConversationResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListServiceParticipantConversationResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceParticipantConversationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
