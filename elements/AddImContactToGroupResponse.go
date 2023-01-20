package elements

// The AddImContactToGroupResponse element defines a response to an AddImContactToGroup request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/addimcontacttogroupresponse
import "encoding/xml"

type AddImContactToGroupResponse struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
}

func (A *AddImContactToGroupResponse) SetForMarshal() {
	A.XMLName.Local = "m:AddImContactToGroupResponse"
}

func (A *AddImContactToGroupResponse) GetSchema() *Schema {
	return &SchemaMessages
}
