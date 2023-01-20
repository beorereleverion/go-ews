package elements

// The AddNewImContactToGroupResponse element defines a response to an AddNewImContactToGroup request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/addnewimcontacttogroupresponse
import "encoding/xml"

type AddNewImContactToGroupResponse struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The Persona element specifies a set of persona data returned by a GetPersona request.
	Persona *Persona `xml:"Persona"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
}

func (A *AddNewImContactToGroupResponse) SetForMarshal() {
	A.XMLName.Local = "m:AddNewImContactToGroupResponse"
}

func (A *AddNewImContactToGroupResponse) GetSchema() *Schema {
	return &SchemaMessages
}
