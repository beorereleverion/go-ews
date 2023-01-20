package elements

// The DeleteItemResponseMessage element contains the status and result of a single DeleteItem operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteitemresponsemessage
import "encoding/xml"

type DeleteItemResponseMessage struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// Describes the status of a DeleteItem operation response.The following values are valid for this attribute:- Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}

func (D *DeleteItemResponseMessage) SetForMarshal() {
	D.XMLName.Local = "m:DeleteItemResponseMessage"
}

func (D *DeleteItemResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
