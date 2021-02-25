/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateEngagementRequest struct for CreateEngagementRequest
type CreateEngagementRequest struct {
	// The Twilio phone number to send messages or initiate calls from during the Flow Engagement. Available as variable `{{flow.channel.address}}`
	From string `json:"From"`
	// A JSON string we will add to your flow's context and that you can access as variables inside your flow. For example, if you pass in `Parameters={'name':'Zeke'}` then inside a widget you can reference the variable `{{flow.data.name}}` which will return the string 'Zeke'. Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode your JSON string.
	Parameters map[string]interface{} `json:"Parameters,omitempty"`
	// The Contact phone number to start a Studio Flow Engagement, available as variable `{{contact.channel.address}}`.
	To string `json:"To"`
}
