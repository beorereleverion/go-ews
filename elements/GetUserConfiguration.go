package elements

// The GetUserConfiguration element represent a request to get a user configuration object.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuserconfiguration
type GetUserConfiguration struct {
	// The UserConfigurationName element represents the name of a user configuration object. The user configuration object name is the identifier for a user configuration object.
	UserConfigurationName *UserConfigurationName `xml:"t:UserConfigurationName"`
	// The UserConfigurationProperties element specifies the property types to get in a GetUserConfiguration operation.
	UserConfigurationProperties *UserConfigurationProperties `xml:"m:UserConfigurationProperties"`
}
