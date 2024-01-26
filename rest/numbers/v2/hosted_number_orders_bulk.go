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

// Optional parameters for the method 'FetchBulkHostedNumberOrder'
type FetchBulkHostedNumberOrderParams struct {
	// Order status can be used for filtering on Hosted Number Order status values. To see a complete list of order statuses, please check 'https://www.twilio.com/docs/phone-numbers/hosted-numbers/hosted-numbers-api/hosted-number-order-resource#status-values'.
	OrderStatus *string `json:"OrderStatus,omitempty"`
}

func (params *FetchBulkHostedNumberOrderParams) SetOrderStatus(OrderStatus string) *FetchBulkHostedNumberOrderParams {
	params.OrderStatus = &OrderStatus
	return params
}

// Fetch a specific BulkHostedNumberOrder.
func (c *ApiService) FetchBulkHostedNumberOrder(BulkHostingSid string, params *FetchBulkHostedNumberOrderParams) (*NumbersV2BulkHostedNumberOrder, error) {
	path := "/v2/HostedNumber/Orders/Bulk/{BulkHostingSid}"
	path = strings.Replace(path, "{"+"BulkHostingSid"+"}", BulkHostingSid, -1)

	data := url.Values{}
	queryParams := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.OrderStatus != nil {
		queryParams.Set("OrderStatus", *params.OrderStatus)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers, queryParams)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2BulkHostedNumberOrder{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
