/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
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

// Optional parameters for the method 'CreateIncomingPhoneNumberTollFree'
type CreateIncomingPhoneNumberTollFreeParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234.
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
	// The API version to use for incoming calls made to the new phone number. The default is `2010-04-01`.
	ApiVersion *string `json:"ApiVersion,omitempty"`
	// A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, this is a formatted version of the phone number.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The SID of the application that should handle SMS messages sent to the new phone number. If an `sms_application_sid` is present, we ignore all `sms_*_url` values and use those of the application.
	SmsApplicationSid *string `json:"SmsApplicationSid,omitempty"`
	// The HTTP method that we should use to call `sms_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`.
	SmsFallbackMethod *string `json:"SmsFallbackMethod,omitempty"`
	// The URL that we should call when an error occurs while requesting or executing the TwiML defined by `sms_url`.
	SmsFallbackUrl *string `json:"SmsFallbackUrl,omitempty"`
	// The HTTP method that we should use to call `sms_url`. Can be: `GET` or `POST` and defaults to `POST`.
	SmsMethod *string `json:"SmsMethod,omitempty"`
	// The URL we should call when the new phone number receives an incoming SMS message.
	SmsUrl *string `json:"SmsUrl,omitempty"`
	// The URL we should call using the `status_callback_method` to send status information to your application.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
	// The SID of the application we should use to handle calls to the new phone number. If a `voice_application_sid` is present, we ignore all of the voice urls and use those set on the application. Setting a `voice_application_sid` will automatically delete your `trunk_sid` and vice versa.
	VoiceApplicationSid *string `json:"VoiceApplicationSid,omitempty"`
	// Whether to lookup the caller's name from the CNAM database and post it to your app. Can be: `true` or `false` and defaults to `false`.
	VoiceCallerIdLookup *bool `json:"VoiceCallerIdLookup,omitempty"`
	// The HTTP method that we should use to call `voice_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`.
	VoiceFallbackMethod *string `json:"VoiceFallbackMethod,omitempty"`
	// The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`.
	VoiceFallbackUrl *string `json:"VoiceFallbackUrl,omitempty"`
	// The HTTP method that we should use to call `voice_url`. Can be: `GET` or `POST` and defaults to `POST`.
	VoiceMethod *string `json:"VoiceMethod,omitempty"`
	// The URL that we should call to answer a call to the new phone number. The `voice_url` will not be called if a `voice_application_sid` or a `trunk_sid` is set.
	VoiceUrl *string `json:"VoiceUrl,omitempty"`
	// The SID of the Identity resource that we should associate with the new phone number. Some regions require an Identity to meet local regulations.
	IdentitySid *string `json:"IdentitySid,omitempty"`
	// The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations.
	AddressSid *string `json:"AddressSid,omitempty"`
	//
	EmergencyStatus *string `json:"EmergencyStatus,omitempty"`
	// The SID of the emergency address configuration to use for emergency calling from the new phone number.
	EmergencyAddressSid *string `json:"EmergencyAddressSid,omitempty"`
	// The SID of the Trunk we should use to handle calls to the new phone number. If a `trunk_sid` is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a `trunk_sid` will automatically delete your `voice_application_sid` and vice versa.
	TrunkSid *string `json:"TrunkSid,omitempty"`
	//
	VoiceReceiveMode *string `json:"VoiceReceiveMode,omitempty"`
	// The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations.
	BundleSid *string `json:"BundleSid,omitempty"`
}

func (params *CreateIncomingPhoneNumberTollFreeParams) SetPathAccountSid(PathAccountSid string) *CreateIncomingPhoneNumberTollFreeParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetPhoneNumber(PhoneNumber string) *CreateIncomingPhoneNumberTollFreeParams {
	params.PhoneNumber = &PhoneNumber
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetApiVersion(ApiVersion string) *CreateIncomingPhoneNumberTollFreeParams {
	params.ApiVersion = &ApiVersion
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetFriendlyName(FriendlyName string) *CreateIncomingPhoneNumberTollFreeParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetSmsApplicationSid(SmsApplicationSid string) *CreateIncomingPhoneNumberTollFreeParams {
	params.SmsApplicationSid = &SmsApplicationSid
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetSmsFallbackMethod(SmsFallbackMethod string) *CreateIncomingPhoneNumberTollFreeParams {
	params.SmsFallbackMethod = &SmsFallbackMethod
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetSmsFallbackUrl(SmsFallbackUrl string) *CreateIncomingPhoneNumberTollFreeParams {
	params.SmsFallbackUrl = &SmsFallbackUrl
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetSmsMethod(SmsMethod string) *CreateIncomingPhoneNumberTollFreeParams {
	params.SmsMethod = &SmsMethod
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetSmsUrl(SmsUrl string) *CreateIncomingPhoneNumberTollFreeParams {
	params.SmsUrl = &SmsUrl
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetStatusCallback(StatusCallback string) *CreateIncomingPhoneNumberTollFreeParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetStatusCallbackMethod(StatusCallbackMethod string) *CreateIncomingPhoneNumberTollFreeParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetVoiceApplicationSid(VoiceApplicationSid string) *CreateIncomingPhoneNumberTollFreeParams {
	params.VoiceApplicationSid = &VoiceApplicationSid
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetVoiceCallerIdLookup(VoiceCallerIdLookup bool) *CreateIncomingPhoneNumberTollFreeParams {
	params.VoiceCallerIdLookup = &VoiceCallerIdLookup
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetVoiceFallbackMethod(VoiceFallbackMethod string) *CreateIncomingPhoneNumberTollFreeParams {
	params.VoiceFallbackMethod = &VoiceFallbackMethod
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetVoiceFallbackUrl(VoiceFallbackUrl string) *CreateIncomingPhoneNumberTollFreeParams {
	params.VoiceFallbackUrl = &VoiceFallbackUrl
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetVoiceMethod(VoiceMethod string) *CreateIncomingPhoneNumberTollFreeParams {
	params.VoiceMethod = &VoiceMethod
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetVoiceUrl(VoiceUrl string) *CreateIncomingPhoneNumberTollFreeParams {
	params.VoiceUrl = &VoiceUrl
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetIdentitySid(IdentitySid string) *CreateIncomingPhoneNumberTollFreeParams {
	params.IdentitySid = &IdentitySid
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetAddressSid(AddressSid string) *CreateIncomingPhoneNumberTollFreeParams {
	params.AddressSid = &AddressSid
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetEmergencyStatus(EmergencyStatus string) *CreateIncomingPhoneNumberTollFreeParams {
	params.EmergencyStatus = &EmergencyStatus
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetEmergencyAddressSid(EmergencyAddressSid string) *CreateIncomingPhoneNumberTollFreeParams {
	params.EmergencyAddressSid = &EmergencyAddressSid
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetTrunkSid(TrunkSid string) *CreateIncomingPhoneNumberTollFreeParams {
	params.TrunkSid = &TrunkSid
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetVoiceReceiveMode(VoiceReceiveMode string) *CreateIncomingPhoneNumberTollFreeParams {
	params.VoiceReceiveMode = &VoiceReceiveMode
	return params
}
func (params *CreateIncomingPhoneNumberTollFreeParams) SetBundleSid(BundleSid string) *CreateIncomingPhoneNumberTollFreeParams {
	params.BundleSid = &BundleSid
	return params
}

//
func (c *ApiService) CreateIncomingPhoneNumberTollFree(params *CreateIncomingPhoneNumberTollFreeParams) (*ApiV2010IncomingPhoneNumberTollFree, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/TollFree.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PhoneNumber != nil {
		data.Set("PhoneNumber", *params.PhoneNumber)
	}
	if params != nil && params.ApiVersion != nil {
		data.Set("ApiVersion", *params.ApiVersion)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.SmsApplicationSid != nil {
		data.Set("SmsApplicationSid", *params.SmsApplicationSid)
	}
	if params != nil && params.SmsFallbackMethod != nil {
		data.Set("SmsFallbackMethod", *params.SmsFallbackMethod)
	}
	if params != nil && params.SmsFallbackUrl != nil {
		data.Set("SmsFallbackUrl", *params.SmsFallbackUrl)
	}
	if params != nil && params.SmsMethod != nil {
		data.Set("SmsMethod", *params.SmsMethod)
	}
	if params != nil && params.SmsUrl != nil {
		data.Set("SmsUrl", *params.SmsUrl)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}
	if params != nil && params.VoiceApplicationSid != nil {
		data.Set("VoiceApplicationSid", *params.VoiceApplicationSid)
	}
	if params != nil && params.VoiceCallerIdLookup != nil {
		data.Set("VoiceCallerIdLookup", fmt.Sprint(*params.VoiceCallerIdLookup))
	}
	if params != nil && params.VoiceFallbackMethod != nil {
		data.Set("VoiceFallbackMethod", *params.VoiceFallbackMethod)
	}
	if params != nil && params.VoiceFallbackUrl != nil {
		data.Set("VoiceFallbackUrl", *params.VoiceFallbackUrl)
	}
	if params != nil && params.VoiceMethod != nil {
		data.Set("VoiceMethod", *params.VoiceMethod)
	}
	if params != nil && params.VoiceUrl != nil {
		data.Set("VoiceUrl", *params.VoiceUrl)
	}
	if params != nil && params.IdentitySid != nil {
		data.Set("IdentitySid", *params.IdentitySid)
	}
	if params != nil && params.AddressSid != nil {
		data.Set("AddressSid", *params.AddressSid)
	}
	if params != nil && params.EmergencyStatus != nil {
		data.Set("EmergencyStatus", *params.EmergencyStatus)
	}
	if params != nil && params.EmergencyAddressSid != nil {
		data.Set("EmergencyAddressSid", *params.EmergencyAddressSid)
	}
	if params != nil && params.TrunkSid != nil {
		data.Set("TrunkSid", *params.TrunkSid)
	}
	if params != nil && params.VoiceReceiveMode != nil {
		data.Set("VoiceReceiveMode", *params.VoiceReceiveMode)
	}
	if params != nil && params.BundleSid != nil {
		data.Set("BundleSid", *params.BundleSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010IncomingPhoneNumberTollFree{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListIncomingPhoneNumberTollFree'
type ListIncomingPhoneNumberTollFreeParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// Whether to include phone numbers new to the Twilio platform. Can be: `true` or `false` and the default is `true`.
	Beta *bool `json:"Beta,omitempty"`
	// A string that identifies the resources to read.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use '*' as a wildcard for any digit.
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
	// Whether to include phone numbers based on their origin. Can be: `twilio` or `hosted`. By default, phone numbers of all origin are included.
	Origin *string `json:"Origin,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListIncomingPhoneNumberTollFreeParams) SetPathAccountSid(PathAccountSid string) *ListIncomingPhoneNumberTollFreeParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListIncomingPhoneNumberTollFreeParams) SetBeta(Beta bool) *ListIncomingPhoneNumberTollFreeParams {
	params.Beta = &Beta
	return params
}
func (params *ListIncomingPhoneNumberTollFreeParams) SetFriendlyName(FriendlyName string) *ListIncomingPhoneNumberTollFreeParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListIncomingPhoneNumberTollFreeParams) SetPhoneNumber(PhoneNumber string) *ListIncomingPhoneNumberTollFreeParams {
	params.PhoneNumber = &PhoneNumber
	return params
}
func (params *ListIncomingPhoneNumberTollFreeParams) SetOrigin(Origin string) *ListIncomingPhoneNumberTollFreeParams {
	params.Origin = &Origin
	return params
}
func (params *ListIncomingPhoneNumberTollFreeParams) SetPageSize(PageSize int) *ListIncomingPhoneNumberTollFreeParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListIncomingPhoneNumberTollFreeParams) SetLimit(Limit int) *ListIncomingPhoneNumberTollFreeParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of IncomingPhoneNumberTollFree records from the API. Request is executed immediately.
func (c *ApiService) PageIncomingPhoneNumberTollFree(params *ListIncomingPhoneNumberTollFreeParams, pageToken, pageNumber string) (*ListIncomingPhoneNumberTollFree200Response, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/TollFree.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Beta != nil {
		data.Set("Beta", fmt.Sprint(*params.Beta))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.PhoneNumber != nil {
		data.Set("PhoneNumber", *params.PhoneNumber)
	}
	if params != nil && params.Origin != nil {
		data.Set("Origin", *params.Origin)
	}
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

	ps := &ListIncomingPhoneNumberTollFree200Response{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists IncomingPhoneNumberTollFree records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListIncomingPhoneNumberTollFree(params *ListIncomingPhoneNumberTollFreeParams) ([]ApiV2010IncomingPhoneNumberTollFree, error) {
	response, errors := c.StreamIncomingPhoneNumberTollFree(params)

	records := make([]ApiV2010IncomingPhoneNumberTollFree, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams IncomingPhoneNumberTollFree records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamIncomingPhoneNumberTollFree(params *ListIncomingPhoneNumberTollFreeParams) (chan ApiV2010IncomingPhoneNumberTollFree, chan error) {
	if params == nil {
		params = &ListIncomingPhoneNumberTollFreeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ApiV2010IncomingPhoneNumberTollFree, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageIncomingPhoneNumberTollFree(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamIncomingPhoneNumberTollFree(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamIncomingPhoneNumberTollFree(response *ListIncomingPhoneNumberTollFree200Response, params *ListIncomingPhoneNumberTollFreeParams, recordChannel chan ApiV2010IncomingPhoneNumberTollFree, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.IncomingPhoneNumbers
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListIncomingPhoneNumberTollFree200Response)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListIncomingPhoneNumberTollFree200Response)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListIncomingPhoneNumberTollFree200Response(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListIncomingPhoneNumberTollFree200Response{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
