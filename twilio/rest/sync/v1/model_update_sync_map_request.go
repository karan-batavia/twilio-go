/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UpdateSyncMapRequest struct for UpdateSyncMapRequest
type UpdateSyncMapRequest struct {
	// How long, in seconds, before the Sync Map expires (time-to-live) and is deleted. Can be an integer from 0 to 31,536,000 (1 year). The default value is `0`, which means the Sync Map does not expire. The Sync Map will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources's deletion.
	CollectionTtl int32 `json:"CollectionTtl,omitempty"`
	// An alias for `collection_ttl`. If both parameters are provided, this value is ignored.
	Ttl int32 `json:"Ttl,omitempty"`
}
