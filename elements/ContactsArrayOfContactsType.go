package elements

// The Contacts element specifies an array of contacts.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/contacts-arrayofcontactstype
type ContactsArrayOfContactsType struct {
	// The Contact element specifies a contact in the Unified Contact Store.
	Contact *ContactContactType `xml:"t:Contact"`
}
