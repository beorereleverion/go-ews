package elements

// The UserParameters element contains a list of enabled and disabled client extensions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/userparameters
type UserParameters struct {
	// The UserDisabledExtensions element lists the disabled apps.
	UserDisabledExtensions *UserDisabledExtensions `xml:"t:UserDisabledExtensions"`
	// The UserEnabledExtensions element lists the enabled apps.
	UserEnabledExtensions *UserEnabledExtensions `xml:"t:UserEnabledExtensions"`
	// The text value of the EnabledOnly indicates whether the response only contains the enabled extensions.
	EnabledOnly *string `xml:"EnabledOnly,attr"`
	// The text value of the UserId attribute is the identifier of the user.
	UserId *string `xml:"UserId,attr"`
}
