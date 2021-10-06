/*
 * Twilio - Autopilot
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
	"net/url"
)

// Optional parameters for the method 'UpdateRestoreAssistant'
type UpdateRestoreAssistantParams struct {
	// The Twilio-provided string that uniquely identifies the Assistant resource to restore.
	Assistant *string `json:"Assistant,omitempty"`
}

func (params *UpdateRestoreAssistantParams) SetAssistant(Assistant string) *UpdateRestoreAssistantParams {
	params.Assistant = &Assistant
	return params
}

func (c *ApiService) UpdateRestoreAssistant(params *UpdateRestoreAssistantParams) (*AutopilotV1RestoreAssistant, error) {
	path := "/v1/Assistants/Restore"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Assistant != nil {
		data.Set("Assistant", *params.Assistant)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1RestoreAssistant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
