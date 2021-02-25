/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UpdateUserRequest struct for UpdateUserRequest
type UpdateUserRequest struct {
	// The JSON Object string that stores application-specific data. If attributes have not been set, `{}` is returned.
	Attributes string `json:"Attributes,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user.
	RoleSid string `json:"RoleSid,omitempty"`
}
