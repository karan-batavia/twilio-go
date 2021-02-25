/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListTranscriptionResponse struct for ListTranscriptionResponse
type ListTranscriptionResponse struct {
	End             int32                          `json:"End,omitempty"`
	FirstPageUri    string                         `json:"FirstPageUri,omitempty"`
	NextPageUri     string                         `json:"NextPageUri,omitempty"`
	Page            int32                          `json:"Page,omitempty"`
	PageSize        int32                          `json:"PageSize,omitempty"`
	PreviousPageUri string                         `json:"PreviousPageUri,omitempty"`
	Start           int32                          `json:"Start,omitempty"`
	Transcriptions  []ApiV2010AccountTranscription `json:"Transcriptions,omitempty"`
	Uri             string                         `json:"Uri,omitempty"`
}
