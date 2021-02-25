/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UpdateQueryRequest struct for UpdateQueryRequest
type UpdateQueryRequest struct {
	// The SID of an optional reference to the [Sample](https://www.twilio.com/docs/autopilot/api/task-sample) created from the query.
	SampleSid string `json:"SampleSid,omitempty"`
	// The new status of the resource. Can be: `pending-review`, `reviewed`, or `discarded`
	Status string `json:"Status,omitempty"`
}
