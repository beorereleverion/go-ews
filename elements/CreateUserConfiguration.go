package elements

// The CreateUserConfiguration element represents a request to create a user configuration object.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createuserconfiguration
type CreateUserConfiguration struct {
	// The UserConfiguration element defines a single user configuration object.
	UserConfiguration *UserConfiguration `xml:"m:UserConfiguration"`
}
