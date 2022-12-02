package elements

// The EmailUser element specifies an email recipient.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emailuser
type EmailUser struct {
	// The Name element specifies a search refiner name.
	Name *Namestring `xml:"t:Name"`
	// The UserId element specifies the user identifier of an email user.
	UserId *UserIdstring `xml:"t:UserId"`
}
