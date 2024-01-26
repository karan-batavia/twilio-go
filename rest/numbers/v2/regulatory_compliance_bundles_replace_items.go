/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Numbers
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

// Optional parameters for the method 'CreateReplaceItems'
type CreateReplaceItemsParams struct {
	// The source bundle sid to copy the item assignments from.
	FromBundleSid *string `json:"FromBundleSid,omitempty"`
}

func (params *CreateReplaceItemsParams) SetFromBundleSid(FromBundleSid string) *CreateReplaceItemsParams {
	params.FromBundleSid = &FromBundleSid
	return params
}

// Replaces all bundle items in the target bundle (specified in the path) with all the bundle items of the source bundle (specified by the from_bundle_sid body param)
func (c *ApiService) CreateReplaceItems(BundleSid string, params *CreateReplaceItemsParams) (*NumbersV2ReplaceItems, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/ReplaceItems"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FromBundleSid != nil {
		data.Set("FromBundleSid", *params.FromBundleSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers, queryParams)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2ReplaceItems{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
