/*
 * Twilio - Fax
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UpdateFaxRequest struct for UpdateFaxRequest
type UpdateFaxRequest struct {
	// The new [status](https://www.twilio.com/docs/fax/api/fax-resource#fax-status-values) of the resource. Can be only `canceled`. This may fail if transmission has already started.
	Status string `json:"Status,omitempty"`
}
