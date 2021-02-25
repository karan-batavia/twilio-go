/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UpdateBucketRequest struct for UpdateBucketRequest
type UpdateBucketRequest struct {
	// Number of seconds that the rate limit will be enforced over.
	Interval int32 `json:"Interval,omitempty"`
	// Maximum number of requests permitted in during the interval.
	Max int32 `json:"Max,omitempty"`
}
