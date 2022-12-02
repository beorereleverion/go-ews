package elements

// The RemoveContactFromImList element represents a request to remove an instant messaging contact from all instant messaging groups.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/removecontactfromimlist
type RemoveContactFromImList struct {
	// The ContactId element uniquely identifies a contact.
	ContactId *ContactId `xml:"t:ContactId"`
}
