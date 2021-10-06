/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.21.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSyncMapItemResponse struct for ListSyncMapItemResponse
type ListSyncMapItemResponse struct {
	Items []SyncV1SyncMapItem     `json:"items,omitempty"`
	Meta  ListServiceResponseMeta `json:"meta,omitempty"`
}
