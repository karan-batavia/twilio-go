/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Trunking
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

func (c *ApiService) FetchRecording(TrunkSid string) (*TrunkingV1Recording, error) {
	path := "/v1/Trunks/{TrunkSid}/Recording"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1Recording{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateRecording'
type UpdateRecordingParams struct {
	// The recording mode for the trunk. Can be do-not-record (default), record-from-ringing, record-from-answer, record-from-ringing-dual, or record-from-answer-dual.
	Mode *string `json:"Mode,omitempty"`
	// The recording trim setting for the trunk. Can be do-not-trim (default) or trim-silence.
	Trim *string `json:"Trim,omitempty"`
}

func (params *UpdateRecordingParams) SetMode(Mode string) *UpdateRecordingParams {
	params.Mode = &Mode
	return params
}
func (params *UpdateRecordingParams) SetTrim(Trim string) *UpdateRecordingParams {
	params.Trim = &Trim
	return params
}

func (c *ApiService) UpdateRecording(TrunkSid string, params *UpdateRecordingParams) (*TrunkingV1Recording, error) {
	path := "/v1/Trunks/{TrunkSid}/Recording"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Mode != nil {
		data.Set("Mode", *params.Mode)
	}
	if params != nil && params.Trim != nil {
		data.Set("Trim", *params.Trim)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1Recording{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
