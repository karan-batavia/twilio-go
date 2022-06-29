/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Accounts
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

// Optional parameters for the method 'CreateCredentialPublicKey'
type CreateCredentialPublicKeyParams struct {
	// The SID of the Subaccount that this Credential should be associated with. Must be a valid Subaccount of the account issuing the request
	AccountSid *string `json:"AccountSid,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// A URL encoded representation of the public key. For example, `-----BEGIN PUBLIC KEY-----MIIBIjANB.pa9xQIDAQAB-----END PUBLIC KEY-----`
	PublicKey *string `json:"PublicKey,omitempty"`
}

func (params *CreateCredentialPublicKeyParams) SetAccountSid(AccountSid string) *CreateCredentialPublicKeyParams {
	params.AccountSid = &AccountSid
	return params
}
func (params *CreateCredentialPublicKeyParams) SetFriendlyName(FriendlyName string) *CreateCredentialPublicKeyParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateCredentialPublicKeyParams) SetPublicKey(PublicKey string) *CreateCredentialPublicKeyParams {
	params.PublicKey = &PublicKey
	return params
}

// Create a new Public Key Credential
func (c *ApiService) CreateCredentialPublicKey(params *CreateCredentialPublicKeyParams) (*AccountsV1CredentialPublicKey, error) {
	path := "/v1/Credentials/PublicKeys"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AccountSid != nil {
		data.Set("AccountSid", *params.AccountSid)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.PublicKey != nil {
		data.Set("PublicKey", *params.PublicKey)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AccountsV1CredentialPublicKey{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a Credential from your account
func (c *ApiService) DeleteCredentialPublicKey(Sid string) error {
	path := "/v1/Credentials/PublicKeys/{Sid}"
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

// Fetch the public key specified by the provided Credential Sid
func (c *ApiService) FetchCredentialPublicKey(Sid string) (*AccountsV1CredentialPublicKey, error) {
	path := "/v1/Credentials/PublicKeys/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AccountsV1CredentialPublicKey{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListCredentialPublicKey'
type ListCredentialPublicKeyParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListCredentialPublicKeyParams) SetPageSize(PageSize int) *ListCredentialPublicKeyParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListCredentialPublicKeyParams) SetLimit(Limit int) *ListCredentialPublicKeyParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of CredentialPublicKey records from the API. Request is executed immediately.
func (c *ApiService) PageCredentialPublicKey(params *ListCredentialPublicKeyParams, pageToken, pageNumber string) (*ListCredentialPublicKeyResponse, error) {
	path := "/v1/Credentials/PublicKeys"

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

	ps := &ListCredentialPublicKeyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists CredentialPublicKey records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCredentialPublicKey(params *ListCredentialPublicKeyParams) ([]AccountsV1CredentialPublicKey, error) {
	response, errors := c.StreamCredentialPublicKey(params)

	records := make([]AccountsV1CredentialPublicKey, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams CredentialPublicKey records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCredentialPublicKey(params *ListCredentialPublicKeyParams) (chan AccountsV1CredentialPublicKey, chan error) {
	if params == nil {
		params = &ListCredentialPublicKeyParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan AccountsV1CredentialPublicKey, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageCredentialPublicKey(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamCredentialPublicKey(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamCredentialPublicKey(response *ListCredentialPublicKeyResponse, params *ListCredentialPublicKeyParams, recordChannel chan AccountsV1CredentialPublicKey, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Credentials
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListCredentialPublicKeyResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListCredentialPublicKeyResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListCredentialPublicKeyResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCredentialPublicKeyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateCredentialPublicKey'
type UpdateCredentialPublicKeyParams struct {
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateCredentialPublicKeyParams) SetFriendlyName(FriendlyName string) *UpdateCredentialPublicKeyParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Modify the properties of a given Account
func (c *ApiService) UpdateCredentialPublicKey(Sid string, params *UpdateCredentialPublicKeyParams) (*AccountsV1CredentialPublicKey, error) {
	path := "/v1/Credentials/PublicKeys/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AccountsV1CredentialPublicKey{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
