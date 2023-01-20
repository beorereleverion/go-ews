package elements

// The RemoveImContactFromGroup element defines a request to remove an instant messaging contact from an instant messaging group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/removeimcontactfromgroup
import "encoding/xml"

type RemoveImContactFromGroup struct {
	XMLName xml.Name

	// The ContactId element uniquely identifies a contact.
	ContactId *ContactId `xml:"ContactId"`
	// The GroupId element uniquely identifies a group.
	GroupId *GroupId `xml:"GroupId"`
}

func (R *RemoveImContactFromGroup) SetForMarshal() {
	R.XMLName.Local = "m:RemoveImContactFromGroup"
}

func (R *RemoveImContactFromGroup) GetSchema() *Schema {
	return &SchemaMessages
}
