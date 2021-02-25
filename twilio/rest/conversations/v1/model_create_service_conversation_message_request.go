/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// CreateServiceConversationMessageRequest struct for CreateServiceConversationMessageRequest
type CreateServiceConversationMessageRequest struct {
	// A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \"{}\" will be returned.
	Attributes string `json:"Attributes,omitempty"`
	// The channel specific identifier of the message's author. Defaults to `system`.
	Author string `json:"Author,omitempty"`
	// The content of the message, can be up to 1,600 characters long.
	Body string `json:"Body,omitempty"`
	// The date that this resource was created.
	DateCreated time.Time `json:"DateCreated,omitempty"`
	// The date that this resource was last updated. `null` if the message has not been edited.
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	// The Media SID to be attached to the new Message.
	MediaSid string `json:"MediaSid,omitempty"`
}
