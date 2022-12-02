package elements

// The AddImContactToGroup element defines a request to add an existing instant messaging contact to an instant messaging group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/addimcontacttogroup
type AddImContactToGroup struct {
	// The ContactId element uniquely identifies a contact.
	ContactId *ContactId `xml:"t:ContactId"`
	// The GroupId element uniquely identifies a group.
	GroupId *GroupId `xml:"m:GroupId"`
}
