package elements

// The AddNewTelUriContactToGroupResponse element specifies the result data for the AddNewTelUriContactToGroup WSDL operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/addnewteluricontacttogroupresponse
import "encoding/xml"

type AddNewTelUriContactToGroupResponse struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The Persona element specifies a set of persona data returned by a GetPersona request.
	Persona *Persona `xml:"Persona"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
}

func (A *AddNewTelUriContactToGroupResponse) SetForMarshal() {
	A.XMLName.Local = "m:AddNewTelUriContactToGroupResponse"
}

func (A *AddNewTelUriContactToGroupResponse) GetSchema() *Schema {
	return &SchemaMessages
}
