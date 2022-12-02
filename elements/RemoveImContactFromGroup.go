package elements

// The RemoveImContactFromGroup element defines a request to remove an instant messaging contact from an instant messaging group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/removeimcontactfromgroup
type RemoveImContactFromGroup struct {
	// The ContactId element uniquely identifies a contact.
	ContactId *ContactId `xml:"t:ContactId"`
	// The GroupId element uniquely identifies a group.
	GroupId *GroupId `xml:"m:GroupId"`
}
