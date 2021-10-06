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

// ListServiceConversationMessageReceiptResponse struct for ListServiceConversationMessageReceiptResponse
type ListServiceConversationMessageReceiptResponse struct {
	DeliveryReceipts []ConversationsV1ServiceConversationMessageReceipt `json:"delivery_receipts,omitempty"`
	Meta             ListConversationResponseMeta                       `json:"meta,omitempty"`
}
