package elements

// The ContactsView element defines a search for contact items based on alphabetical display names.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/contactsview
import "encoding/xml"

type ContactsView struct {
	XMLName xml.Name

	// Defines the last name in the contacts list to return in the response. If the FinalName attribute is omitted, the response will contain all subsequent contacts in the specified sort order. If the specified final name is not in the contacts list, the next alphabetical name as defined by the cultural context will be excluded.  For example, if FinalName="Name", but Name is not in the contacts list, contacts that have display names of Name1 or NAME will not be included.  This attribute is optional.
	FinalName *string `xml:"FinalName,attr"`
	// Defines the first name in the contacts list to return in the response. If the specified initial name is not in the contacts list, the next alphabetical name as defined by the cultural context will be returned, except if the next name comes after FinalName. If the InitialName attribute is omitted, the response will contain a list of contacts that starts with the first name in the contact list. This attribute is optional.
	InitialName *string `xml:"InitialName,attr"`
	// Describes the maximum number of results to return in the FindItem response.
	MaxEntriesReturned *string `xml:"MaxEntriesReturned,attr"`
}

func (C *ContactsView) SetForMarshal() {
	C.XMLName.Local = "m:ContactsView"
}

func (C *ContactsView) GetSchema() *Schema {
	return &SchemaMessages
}
