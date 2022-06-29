/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Studio
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// StudioV1ExecutionStep struct for StudioV1ExecutionStep
type StudioV1ExecutionStep struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The current state of the flow
	Context *interface{} `json:"context,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The SID of the Execution
	ExecutionSid *string `json:"execution_sid,omitempty"`
	// The SID of the Flow
	FlowSid *string `json:"flow_sid,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The event that caused the Flow to transition to the Step
	Name *string `json:"name,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The Widget that preceded the Widget for the Step
	TransitionedFrom *string `json:"transitioned_from,omitempty"`
	// The Widget that will follow the Widget for the Step
	TransitionedTo *string `json:"transitioned_to,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
