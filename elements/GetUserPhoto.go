package elements

// The GetUserPhoto element contains the request to get a user's photo.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuserphoto
type GetUserPhoto struct {
	// The Email element identifies the email address of the user whose photo is requested in the GetUserPhoto operation.
	Email *EmailString `xml:"t:Email"`
	// The SizeRequested element contains the requested photo size for a GetUserPhoto operation.
	SizeRequested *SizeRequested `xml:"m:SizeRequested"`
}
