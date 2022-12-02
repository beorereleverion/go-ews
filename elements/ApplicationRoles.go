package elements

// The ApplicationRoles element specifies the application roles that the calling partner application uses for the current call.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/applicationroles
type ApplicationRoles struct {
	// The Role element specifies a string that represents a management role.
	Role *Role `xml:"t:Role"`
}
