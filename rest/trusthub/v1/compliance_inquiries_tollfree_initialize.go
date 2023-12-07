/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Trusthub
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
)

// Optional parameters for the method 'CreateComplianceTollfreeInquiry'
type CreateComplianceTollfreeInquiryParams struct {
	// The Tollfree phone number to be verified
	TollfreePhoneNumber *string `json:"TollfreePhoneNumber,omitempty"`
	// The notification email to be triggered when verification status is changed
	NotificationEmail *string `json:"NotificationEmail,omitempty"`
}

func (params *CreateComplianceTollfreeInquiryParams) SetTollfreePhoneNumber(TollfreePhoneNumber string) *CreateComplianceTollfreeInquiryParams {
	params.TollfreePhoneNumber = &TollfreePhoneNumber
	return params
}
func (params *CreateComplianceTollfreeInquiryParams) SetNotificationEmail(NotificationEmail string) *CreateComplianceTollfreeInquiryParams {
	params.NotificationEmail = &NotificationEmail
	return params
}

// Create a new Compliance Tollfree Verification Inquiry for the authenticated account. This is necessary to start a new embedded session.
func (c *ApiService) CreateComplianceTollfreeInquiry(params *CreateComplianceTollfreeInquiryParams) (*TrusthubV1ComplianceTollfreeInquiry, error) {
	path := "/v1/ComplianceInquiries/Tollfree/Initialize"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.TollfreePhoneNumber != nil {
		data.Set("TollfreePhoneNumber", *params.TollfreePhoneNumber)
	}
	if params != nil && params.NotificationEmail != nil {
		data.Set("NotificationEmail", *params.NotificationEmail)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1ComplianceTollfreeInquiry{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
