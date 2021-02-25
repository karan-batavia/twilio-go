/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// DependentHostedNumberOrderStatus the model 'DependentHostedNumberOrderStatus'
type DependentHostedNumberOrderStatus string

// List of dependent_hosted_number_order_status
const (
	DEPENDENTHOSTEDNUMBERORDERSTATUS_RECEIVED             DependentHostedNumberOrderStatus = "received"
	DEPENDENTHOSTEDNUMBERORDERSTATUS_PENDING_VERIFICATION DependentHostedNumberOrderStatus = "pending-verification"
	DEPENDENTHOSTEDNUMBERORDERSTATUS_VERIFIED             DependentHostedNumberOrderStatus = "verified"
	DEPENDENTHOSTEDNUMBERORDERSTATUS_PENDING_LOA          DependentHostedNumberOrderStatus = "pending-loa"
	DEPENDENTHOSTEDNUMBERORDERSTATUS_CARRIER_PROCESSING   DependentHostedNumberOrderStatus = "carrier-processing"
	DEPENDENTHOSTEDNUMBERORDERSTATUS_TESTING              DependentHostedNumberOrderStatus = "testing"
	DEPENDENTHOSTEDNUMBERORDERSTATUS_COMPLETED            DependentHostedNumberOrderStatus = "completed"
	DEPENDENTHOSTEDNUMBERORDERSTATUS_FAILED               DependentHostedNumberOrderStatus = "failed"
	DEPENDENTHOSTEDNUMBERORDERSTATUS_ACTION_REQUIRED      DependentHostedNumberOrderStatus = "action-required"
)
