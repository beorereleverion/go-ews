package elements

// The AddImContactToGroup element defines a request to add an existing instant messaging contact to an instant messaging group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/addimcontacttogroup
import "encoding/xml"

type AddImContactToGroup struct {
	XMLName xml.Name

	// The ContactId element uniquely identifies a contact.
	ContactId *ContactId `xml:"ContactId"`
	// The GroupId element uniquely identifies a group.
	GroupId *GroupId `xml:"GroupId"`
}

func (A *AddImContactToGroup) SetForMarshal() {
	A.XMLName.Local = "m:AddImContactToGroup"
}

func (A *AddImContactToGroup) GetSchema() *Schema {
	return &SchemaMessages
}
