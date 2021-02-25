/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// TaskrouterV1WorkspaceWorkflowWorkflowCumulativeStatistics struct for TaskrouterV1WorkspaceWorkflowWorkflowCumulativeStatistics
type TaskrouterV1WorkspaceWorkflowWorkflowCumulativeStatistics struct {
	AccountSid                *string                 `json:"AccountSid,omitempty"`
	AvgTaskAcceptanceTime     *int32                  `json:"AvgTaskAcceptanceTime,omitempty"`
	EndTime                   *time.Time              `json:"EndTime,omitempty"`
	ReservationsAccepted      *int32                  `json:"ReservationsAccepted,omitempty"`
	ReservationsCanceled      *int32                  `json:"ReservationsCanceled,omitempty"`
	ReservationsCreated       *int32                  `json:"ReservationsCreated,omitempty"`
	ReservationsRejected      *int32                  `json:"ReservationsRejected,omitempty"`
	ReservationsRescinded     *int32                  `json:"ReservationsRescinded,omitempty"`
	ReservationsTimedOut      *int32                  `json:"ReservationsTimedOut,omitempty"`
	SplitByWaitTime           *map[string]interface{} `json:"SplitByWaitTime,omitempty"`
	StartTime                 *time.Time              `json:"StartTime,omitempty"`
	TasksCanceled             *int32                  `json:"TasksCanceled,omitempty"`
	TasksCompleted            *int32                  `json:"TasksCompleted,omitempty"`
	TasksDeleted              *int32                  `json:"TasksDeleted,omitempty"`
	TasksEntered              *int32                  `json:"TasksEntered,omitempty"`
	TasksMoved                *int32                  `json:"TasksMoved,omitempty"`
	TasksTimedOutInWorkflow   *int32                  `json:"TasksTimedOutInWorkflow,omitempty"`
	Url                       *string                 `json:"Url,omitempty"`
	WaitDurationUntilAccepted *map[string]interface{} `json:"WaitDurationUntilAccepted,omitempty"`
	WaitDurationUntilCanceled *map[string]interface{} `json:"WaitDurationUntilCanceled,omitempty"`
	WorkflowSid               *string                 `json:"WorkflowSid,omitempty"`
	WorkspaceSid              *string                 `json:"WorkspaceSid,omitempty"`
}
