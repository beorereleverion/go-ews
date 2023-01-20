package elements

// The AddNewTelUriContactToGroup element specifies the input data for the AddNewTelUriContactToGroup WSDL operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/addnewteluricontacttogroup
import "encoding/xml"

type AddNewTelUriContactToGroup struct {
	XMLName xml.Name

	// The GroupId element uniquely identifies a group.
	GroupId *GroupId `xml:"GroupId"`
	// The ImContactSipUriAddress element contains the SIP URI address of a contact that is added to an instant messaging (IM) group.
	ImContactSipUriAddress *ImContactSipUriAddress `xml:"ImContactSipUriAddress"`
	// The ImTelephoneNumber element represents the telephone number for a contact that is added to an instant messaging (IM) group.
	ImTelephoneNumber *ImTelephoneNumber `xml:"ImTelephoneNumber"`
	// The TelUriAddress element contains the tel Uniform Resource Identifier (URI) for a contact.
	TelUriAddress *TelUriAddress `xml:"TelUriAddress"`
}

func (A *AddNewTelUriContactToGroup) SetForMarshal() {
	A.XMLName.Local = "m:AddNewTelUriContactToGroup"
}

func (A *AddNewTelUriContactToGroup) GetSchema() *Schema {
	return &SchemaMessages
}
