/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.21.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// TaskrouterV1WorkerInstanceStatistics struct for TaskrouterV1WorkerInstanceStatistics
type TaskrouterV1WorkerInstanceStatistics struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// An object that contains the cumulative statistics for the Worker
	Cumulative *map[string]interface{} `json:"cumulative,omitempty"`
	// The absolute URL of the WorkerChannel statistics resource
	Url *string `json:"url,omitempty"`
	// The SID of the Worker that contains the WorkerChannel
	WorkerSid *string `json:"worker_sid,omitempty"`
	// The SID of the Workspace that contains the WorkerChannel
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
}
