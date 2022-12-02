package elements

// The UpdateUserConfiguration element represents a request to update a user configuration object.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updateuserconfiguration
type UpdateUserConfiguration struct {
	// The UserConfiguration element defines a single user configuration object.
	UserConfiguration *UserConfiguration `xml:"m:UserConfiguration"`
}
