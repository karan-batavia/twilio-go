/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Insights
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// InsightsV1Metric struct for InsightsV1Metric
type InsightsV1Metric struct {
	// Timestamp of metric sample. Samples are taken every 10 seconds and contain the metrics for the previous 10 seconds.
	Timestamp *string `json:"timestamp,omitempty"`
	// The unique SID identifier of the Call.
	CallSid *string `json:"call_sid,omitempty"`
	// The unique SID identifier of the Account.
	AccountSid *string `json:"account_sid,omitempty"`
	Edge       *string `json:"edge,omitempty"`
	Direction  *string `json:"direction,omitempty"`
	// Contains metrics and properties for the Twilio media gateway of a PSTN call.
	CarrierEdge *interface{} `json:"carrier_edge,omitempty"`
	// Contains metrics and properties for the Twilio media gateway of a SIP Interface or Trunking call.
	SipEdge *interface{} `json:"sip_edge,omitempty"`
	// Contains metrics and properties for the SDK sensor library for Client calls.
	SdkEdge *interface{} `json:"sdk_edge,omitempty"`
	// Contains metrics and properties for the Twilio media gateway of a Client call.
	ClientEdge *interface{} `json:"client_edge,omitempty"`
}
