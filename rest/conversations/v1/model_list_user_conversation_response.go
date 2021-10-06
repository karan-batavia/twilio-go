/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.21.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListUserConversationResponse struct for ListUserConversationResponse
type ListUserConversationResponse struct {
	Conversations []ConversationsV1UserConversation `json:"conversations,omitempty"`
	Meta          ListConversationResponseMeta      `json:"meta,omitempty"`
}
