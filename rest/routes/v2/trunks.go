/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Routes
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"
	"strings"
)

// Fetch the Inbound Processing Region assigned to a SIP Trunk.
func (c *ApiService) FetchTrunks(SipTrunkDomain string) (*RoutesV2Trunks, error) {
	path := "/v2/Trunks/{SipTrunkDomain}"
	path = strings.Replace(path, "{"+"SipTrunkDomain"+"}", SipTrunkDomain, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers, queryParams)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &RoutesV2Trunks{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateTrunks'
type UpdateTrunksParams struct {
	// The Inbound Processing Region used for this SIP Trunk for voice
	VoiceRegion *string `json:"VoiceRegion,omitempty"`
	// A human readable description of this resource, up to 64 characters.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateTrunksParams) SetVoiceRegion(VoiceRegion string) *UpdateTrunksParams {
	params.VoiceRegion = &VoiceRegion
	return params
}
func (params *UpdateTrunksParams) SetFriendlyName(FriendlyName string) *UpdateTrunksParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Assign an Inbound Processing Region to a SIP Trunk
func (c *ApiService) UpdateTrunks(SipTrunkDomain string, params *UpdateTrunksParams) (*RoutesV2Trunks, error) {
	path := "/v2/Trunks/{SipTrunkDomain}"
	path = strings.Replace(path, "{"+"SipTrunkDomain"+"}", SipTrunkDomain, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.VoiceRegion != nil {
		data.Set("VoiceRegion", *params.VoiceRegion)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers, queryParams)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &RoutesV2Trunks{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
