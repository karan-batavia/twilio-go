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

// Optional parameters for the method 'CreateConfigurationAddress'
type CreateConfigurationAddressParams struct {
	//
	Type *string `json:"Type,omitempty"`
	// The unique address to be configured. The address can be a whatsapp address or phone number
	Address *string `json:"Address,omitempty"`
	// The human-readable name of this configuration, limited to 256 characters. Optional.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Enable/Disable auto-creating conversations for messages to this address
	AutoCreationEnabled *bool `json:"AutoCreation.Enabled,omitempty"`
	//
	AutoCreationType *string `json:"AutoCreation.Type,omitempty"`
	// Conversation Service for the auto-created conversation. If not set, the conversation is created in the default service.
	AutoCreationConversationServiceSid *string `json:"AutoCreation.ConversationServiceSid,omitempty"`
	// For type `webhook`, the url for the webhook request.
	AutoCreationWebhookUrl *string `json:"AutoCreation.WebhookUrl,omitempty"`
	//
	AutoCreationWebhookMethod *string `json:"AutoCreation.WebhookMethod,omitempty"`
	// The list of events, firing webhook event for this Conversation. Values can be any of the following: `onMessageAdded`, `onMessageUpdated`, `onMessageRemoved`, `onConversationUpdated`, `onConversationStateUpdated`, `onConversationRemoved`, `onParticipantAdded`, `onParticipantUpdated`, `onParticipantRemoved`, `onDeliveryUpdated`
	AutoCreationWebhookFilters *[]string `json:"AutoCreation.WebhookFilters,omitempty"`
	// For type `studio`, the studio flow SID where the webhook should be sent to.
	AutoCreationStudioFlowSid *string `json:"AutoCreation.StudioFlowSid,omitempty"`
	// For type `studio`, number of times to retry the webhook request
	AutoCreationStudioRetryCount *int `json:"AutoCreation.StudioRetryCount,omitempty"`
	// An ISO 3166-1 alpha-2n country code which the address belongs to. This is currently only applicable to short code addresses.
	AddressCountry *string `json:"AddressCountry,omitempty"`
}

func (params *CreateConfigurationAddressParams) SetType(Type string) *CreateConfigurationAddressParams {
	params.Type = &Type
	return params
}
func (params *CreateConfigurationAddressParams) SetAddress(Address string) *CreateConfigurationAddressParams {
	params.Address = &Address
	return params
}
func (params *CreateConfigurationAddressParams) SetFriendlyName(FriendlyName string) *CreateConfigurationAddressParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateConfigurationAddressParams) SetAutoCreationEnabled(AutoCreationEnabled bool) *CreateConfigurationAddressParams {
	params.AutoCreationEnabled = &AutoCreationEnabled
	return params
}
func (params *CreateConfigurationAddressParams) SetAutoCreationType(AutoCreationType string) *CreateConfigurationAddressParams {
	params.AutoCreationType = &AutoCreationType
	return params
}
func (params *CreateConfigurationAddressParams) SetAutoCreationConversationServiceSid(AutoCreationConversationServiceSid string) *CreateConfigurationAddressParams {
	params.AutoCreationConversationServiceSid = &AutoCreationConversationServiceSid
	return params
}
func (params *CreateConfigurationAddressParams) SetAutoCreationWebhookUrl(AutoCreationWebhookUrl string) *CreateConfigurationAddressParams {
	params.AutoCreationWebhookUrl = &AutoCreationWebhookUrl
	return params
}
func (params *CreateConfigurationAddressParams) SetAutoCreationWebhookMethod(AutoCreationWebhookMethod string) *CreateConfigurationAddressParams {
	params.AutoCreationWebhookMethod = &AutoCreationWebhookMethod
	return params
}
func (params *CreateConfigurationAddressParams) SetAutoCreationWebhookFilters(AutoCreationWebhookFilters []string) *CreateConfigurationAddressParams {
	params.AutoCreationWebhookFilters = &AutoCreationWebhookFilters
	return params
}
func (params *CreateConfigurationAddressParams) SetAutoCreationStudioFlowSid(AutoCreationStudioFlowSid string) *CreateConfigurationAddressParams {
	params.AutoCreationStudioFlowSid = &AutoCreationStudioFlowSid
	return params
}
func (params *CreateConfigurationAddressParams) SetAutoCreationStudioRetryCount(AutoCreationStudioRetryCount int) *CreateConfigurationAddressParams {
	params.AutoCreationStudioRetryCount = &AutoCreationStudioRetryCount
	return params
}
func (params *CreateConfigurationAddressParams) SetAddressCountry(AddressCountry string) *CreateConfigurationAddressParams {
	params.AddressCountry = &AddressCountry
	return params
}

// Create a new address configuration
func (c *ApiService) CreateConfigurationAddress(params *CreateConfigurationAddressParams) (*ConversationsV1ConfigurationAddress, error) {
	path := "/v1/Configuration/Addresses"

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Type != nil {
		data.Set("Type", *params.Type)
	}
	if params != nil && params.Address != nil {
		data.Set("Address", *params.Address)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.AutoCreationEnabled != nil {
		data.Set("AutoCreation.Enabled", fmt.Sprint(*params.AutoCreationEnabled))
	}
	if params != nil && params.AutoCreationType != nil {
		data.Set("AutoCreation.Type", *params.AutoCreationType)
	}
	if params != nil && params.AutoCreationConversationServiceSid != nil {
		data.Set("AutoCreation.ConversationServiceSid", *params.AutoCreationConversationServiceSid)
	}
	if params != nil && params.AutoCreationWebhookUrl != nil {
		data.Set("AutoCreation.WebhookUrl", *params.AutoCreationWebhookUrl)
	}
	if params != nil && params.AutoCreationWebhookMethod != nil {
		data.Set("AutoCreation.WebhookMethod", *params.AutoCreationWebhookMethod)
	}
	if params != nil && params.AutoCreationWebhookFilters != nil {
		for _, item := range *params.AutoCreationWebhookFilters {
			data.Add("AutoCreation.WebhookFilters", item)
		}
	}
	if params != nil && params.AutoCreationStudioFlowSid != nil {
		data.Set("AutoCreation.StudioFlowSid", *params.AutoCreationStudioFlowSid)
	}
	if params != nil && params.AutoCreationStudioRetryCount != nil {
		data.Set("AutoCreation.StudioRetryCount", fmt.Sprint(*params.AutoCreationStudioRetryCount))
	}
	if params != nil && params.AddressCountry != nil {
		data.Set("AddressCountry", *params.AddressCountry)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers, queryParams)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConfigurationAddress{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an existing address configuration
func (c *ApiService) DeleteConfigurationAddress(Sid string) error {
	path := "/v1/Configuration/Addresses/{Sid}"
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

// Fetch an address configuration
func (c *ApiService) FetchConfigurationAddress(Sid string) (*ConversationsV1ConfigurationAddress, error) {
	path := "/v1/Configuration/Addresses/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers, queryParams)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConfigurationAddress{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListConfigurationAddress'
type ListConfigurationAddressParams struct {
	// Filter the address configurations by its type. This value can be one of: `whatsapp`, `sms`.
	Type *string `json:"Type,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListConfigurationAddressParams) SetType(Type string) *ListConfigurationAddressParams {
	params.Type = &Type
	return params
}
func (params *ListConfigurationAddressParams) SetPageSize(PageSize int) *ListConfigurationAddressParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListConfigurationAddressParams) SetLimit(Limit int) *ListConfigurationAddressParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ConfigurationAddress records from the API. Request is executed immediately.
func (c *ApiService) PageConfigurationAddress(params *ListConfigurationAddressParams, pageToken, pageNumber string) (*ListConfigurationAddressResponse, error) {
	path := "/v1/Configuration/Addresses"

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Type != nil {
		queryParams.Set("Type", *params.Type)
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

	ps := &ListConfigurationAddressResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ConfigurationAddress records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListConfigurationAddress(params *ListConfigurationAddressParams) ([]ConversationsV1ConfigurationAddress, error) {
	response, errors := c.StreamConfigurationAddress(params)

	records := make([]ConversationsV1ConfigurationAddress, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams ConfigurationAddress records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamConfigurationAddress(params *ListConfigurationAddressParams) (chan ConversationsV1ConfigurationAddress, chan error) {
	if params == nil {
		params = &ListConfigurationAddressParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ConversationsV1ConfigurationAddress, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageConfigurationAddress(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamConfigurationAddress(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamConfigurationAddress(response *ListConfigurationAddressResponse, params *ListConfigurationAddressParams, recordChannel chan ConversationsV1ConfigurationAddress, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.AddressConfigurations
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListConfigurationAddressResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListConfigurationAddressResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListConfigurationAddressResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListConfigurationAddressResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateConfigurationAddress'
type UpdateConfigurationAddressParams struct {
	// The human-readable name of this configuration, limited to 256 characters. Optional.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Enable/Disable auto-creating conversations for messages to this address
	AutoCreationEnabled *bool `json:"AutoCreation.Enabled,omitempty"`
	//
	AutoCreationType *string `json:"AutoCreation.Type,omitempty"`
	// Conversation Service for the auto-created conversation. If not set, the conversation is created in the default service.
	AutoCreationConversationServiceSid *string `json:"AutoCreation.ConversationServiceSid,omitempty"`
	// For type `webhook`, the url for the webhook request.
	AutoCreationWebhookUrl *string `json:"AutoCreation.WebhookUrl,omitempty"`
	//
	AutoCreationWebhookMethod *string `json:"AutoCreation.WebhookMethod,omitempty"`
	// The list of events, firing webhook event for this Conversation. Values can be any of the following: `onMessageAdded`, `onMessageUpdated`, `onMessageRemoved`, `onConversationUpdated`, `onConversationStateUpdated`, `onConversationRemoved`, `onParticipantAdded`, `onParticipantUpdated`, `onParticipantRemoved`, `onDeliveryUpdated`
	AutoCreationWebhookFilters *[]string `json:"AutoCreation.WebhookFilters,omitempty"`
	// For type `studio`, the studio flow SID where the webhook should be sent to.
	AutoCreationStudioFlowSid *string `json:"AutoCreation.StudioFlowSid,omitempty"`
	// For type `studio`, number of times to retry the webhook request
	AutoCreationStudioRetryCount *int `json:"AutoCreation.StudioRetryCount,omitempty"`
}

func (params *UpdateConfigurationAddressParams) SetFriendlyName(FriendlyName string) *UpdateConfigurationAddressParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateConfigurationAddressParams) SetAutoCreationEnabled(AutoCreationEnabled bool) *UpdateConfigurationAddressParams {
	params.AutoCreationEnabled = &AutoCreationEnabled
	return params
}
func (params *UpdateConfigurationAddressParams) SetAutoCreationType(AutoCreationType string) *UpdateConfigurationAddressParams {
	params.AutoCreationType = &AutoCreationType
	return params
}
func (params *UpdateConfigurationAddressParams) SetAutoCreationConversationServiceSid(AutoCreationConversationServiceSid string) *UpdateConfigurationAddressParams {
	params.AutoCreationConversationServiceSid = &AutoCreationConversationServiceSid
	return params
}
func (params *UpdateConfigurationAddressParams) SetAutoCreationWebhookUrl(AutoCreationWebhookUrl string) *UpdateConfigurationAddressParams {
	params.AutoCreationWebhookUrl = &AutoCreationWebhookUrl
	return params
}
func (params *UpdateConfigurationAddressParams) SetAutoCreationWebhookMethod(AutoCreationWebhookMethod string) *UpdateConfigurationAddressParams {
	params.AutoCreationWebhookMethod = &AutoCreationWebhookMethod
	return params
}
func (params *UpdateConfigurationAddressParams) SetAutoCreationWebhookFilters(AutoCreationWebhookFilters []string) *UpdateConfigurationAddressParams {
	params.AutoCreationWebhookFilters = &AutoCreationWebhookFilters
	return params
}
func (params *UpdateConfigurationAddressParams) SetAutoCreationStudioFlowSid(AutoCreationStudioFlowSid string) *UpdateConfigurationAddressParams {
	params.AutoCreationStudioFlowSid = &AutoCreationStudioFlowSid
	return params
}
func (params *UpdateConfigurationAddressParams) SetAutoCreationStudioRetryCount(AutoCreationStudioRetryCount int) *UpdateConfigurationAddressParams {
	params.AutoCreationStudioRetryCount = &AutoCreationStudioRetryCount
	return params
}

// Update an existing address configuration
func (c *ApiService) UpdateConfigurationAddress(Sid string, params *UpdateConfigurationAddressParams) (*ConversationsV1ConfigurationAddress, error) {
	path := "/v1/Configuration/Addresses/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.AutoCreationEnabled != nil {
		data.Set("AutoCreation.Enabled", fmt.Sprint(*params.AutoCreationEnabled))
	}
	if params != nil && params.AutoCreationType != nil {
		data.Set("AutoCreation.Type", *params.AutoCreationType)
	}
	if params != nil && params.AutoCreationConversationServiceSid != nil {
		data.Set("AutoCreation.ConversationServiceSid", *params.AutoCreationConversationServiceSid)
	}
	if params != nil && params.AutoCreationWebhookUrl != nil {
		data.Set("AutoCreation.WebhookUrl", *params.AutoCreationWebhookUrl)
	}
	if params != nil && params.AutoCreationWebhookMethod != nil {
		data.Set("AutoCreation.WebhookMethod", *params.AutoCreationWebhookMethod)
	}
	if params != nil && params.AutoCreationWebhookFilters != nil {
		for _, item := range *params.AutoCreationWebhookFilters {
			data.Add("AutoCreation.WebhookFilters", item)
		}
	}
	if params != nil && params.AutoCreationStudioFlowSid != nil {
		data.Set("AutoCreation.StudioFlowSid", *params.AutoCreationStudioFlowSid)
	}
	if params != nil && params.AutoCreationStudioRetryCount != nil {
		data.Set("AutoCreation.StudioRetryCount", fmt.Sprint(*params.AutoCreationStudioRetryCount))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers, queryParams)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConfigurationAddress{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
