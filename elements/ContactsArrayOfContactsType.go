package elements

// The Contacts element specifies an array of contacts.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/contacts-arrayofcontactstype
import "encoding/xml"

type ContactsArrayOfContactsType struct {
	XMLName xml.Name

	// The Contact element specifies a contact in the Unified Contact Store.
	Contact *ContactContactType `xml:"Contact"`
}

func (C *ContactsArrayOfContactsType) SetForMarshal() {
	C.XMLName.Local = "t:Contacts"
}

func (C *ContactsArrayOfContactsType) GetSchema() *Schema {
	return &SchemaTypes
}
