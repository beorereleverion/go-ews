package elements

// The RemoveContactFromImList element represents a request to remove an instant messaging contact from all instant messaging groups.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/removecontactfromimlist
import "encoding/xml"

type RemoveContactFromImList struct {
	XMLName xml.Name

	// The ContactId element uniquely identifies a contact.
	ContactId *ContactId `xml:"ContactId"`
}

func (R *RemoveContactFromImList) SetForMarshal() {
	R.XMLName.Local = "m:RemoveContactFromImList"
}

func (R *RemoveContactFromImList) GetSchema() *Schema {
	return &SchemaMessages
}
