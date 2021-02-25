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

// TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics struct for TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics
type TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics struct {
	AccountSid   *string                 `json:"AccountSid,omitempty"`
	Cumulative   *map[string]interface{} `json:"Cumulative,omitempty"`
	Realtime     *map[string]interface{} `json:"Realtime,omitempty"`
	TaskQueueSid *string                 `json:"TaskQueueSid,omitempty"`
	WorkspaceSid *string                 `json:"WorkspaceSid,omitempty"`
}
