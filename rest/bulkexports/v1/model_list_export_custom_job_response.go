/*
 * Twilio - Bulkexports
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.21.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListExportCustomJobResponse struct for ListExportCustomJobResponse
type ListExportCustomJobResponse struct {
	Jobs []BulkexportsV1ExportCustomJob `json:"jobs,omitempty"`
	Meta ListDayResponseMeta            `json:"meta,omitempty"`
}
