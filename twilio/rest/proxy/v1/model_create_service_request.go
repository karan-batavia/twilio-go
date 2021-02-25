/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateServiceRequest struct for CreateServiceRequest
type CreateServiceRequest struct {
	// The URL we should call when the interaction status changes.
	CallbackUrl string `json:"CallbackUrl,omitempty"`
	// The SID of the Chat Service Instance managed by Proxy Service. The Chat Service enables Proxy to forward SMS and channel messages to this chat instance. This is a one-to-one relationship.
	ChatInstanceSid string `json:"ChatInstanceSid,omitempty"`
	// The default `ttl` value to set for Sessions created in the Service. The TTL (time to live) is measured in seconds after the Session's last create or last Interaction. The default value of `0` indicates an unlimited Session length. You can override a Session's default TTL value by setting its `ttl` value.
	DefaultTtl int32 `json:"DefaultTtl,omitempty"`
	// Where a proxy number must be located relative to the participant identifier. Can be: `country`, `area-code`, or `extended-area-code`. The default value is `country` and more specific areas than `country` are only available in North America.
	GeoMatchLevel string `json:"GeoMatchLevel,omitempty"`
	// The URL we call on each interaction. If we receive a 403 status, we block the interaction; otherwise the interaction continues.
	InterceptCallbackUrl string `json:"InterceptCallbackUrl,omitempty"`
	// The preference for Proxy Number selection in the Service instance. Can be: `prefer-sticky` or `avoid-sticky` and the default is `prefer-sticky`. `prefer-sticky` means that we will try and select the same Proxy Number for a given participant if they have previous [Sessions](https://www.twilio.com/docs/proxy/api/session), but we will not fail if that Proxy Number cannot be used.  `avoid-sticky` means that we will try to use different Proxy Numbers as long as that is possible within a given pool rather than try and use a previously assigned number.
	NumberSelectionBehavior string `json:"NumberSelectionBehavior,omitempty"`
	// The URL we should call when an inbound call or SMS action occurs on a closed or non-existent Session. If your server (or a Twilio [function](https://www.twilio.com/functions)) responds with valid [TwiML](https://www.twilio.com/docs/voice/twiml), we will process it. This means it is possible, for example, to play a message for a call, send an automated text message response, or redirect a call to another Phone Number. See [Out-of-Session Callback Response Guide](https://www.twilio.com/docs/proxy/out-session-callback-response-guide) for more information.
	OutOfSessionCallbackUrl string `json:"OutOfSessionCallbackUrl,omitempty"`
	// An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. **This value should not have PII.**
	UniqueName string `json:"UniqueName"`
}
