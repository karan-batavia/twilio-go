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

// ListServiceConversationParticipantResponse struct for ListServiceConversationParticipantResponse
type ListServiceConversationParticipantResponse struct {
	Meta         ListConversationResponseMeta                                              `json:"Meta,omitempty"`
	Participants []ConversationsV1ServiceServiceConversationServiceConversationParticipant `json:"Participants,omitempty"`
}
