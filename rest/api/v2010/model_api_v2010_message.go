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

// ApiV2010Message struct for ApiV2010Message
type ApiV2010Message struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The API version used to process the message
	ApiVersion *string `json:"api_version,omitempty"`
	// The message text
	Body *string `json:"body,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the message was sent
	DateSent *string `json:"date_sent,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// The direction of the message
	Direction *string `json:"direction,omitempty"`
	// The error code associated with the message
	ErrorCode *int `json:"error_code,omitempty"`
	// The description of the error_code
	ErrorMessage *string `json:"error_message,omitempty"`
	// The phone number that initiated the message
	From *string `json:"from,omitempty"`
	// The SID of the Messaging Service used with the message.
	MessagingServiceSid *string `json:"messaging_service_sid,omitempty"`
	// The number of media files associated with the message
	NumMedia *string `json:"num_media,omitempty"`
	// The number of messages used to deliver the message body
	NumSegments *string `json:"num_segments,omitempty"`
	// The amount billed for the message
	Price *string `json:"price,omitempty"`
	// The currency in which price is measured
	PriceUnit *string `json:"price_unit,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The status of the message
	Status *string `json:"status,omitempty"`
	// A list of related resources identified by their relative URIs
	SubresourceUris *map[string]interface{} `json:"subresource_uris,omitempty"`
	// The phone number that received the message
	To *string `json:"to,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
}
