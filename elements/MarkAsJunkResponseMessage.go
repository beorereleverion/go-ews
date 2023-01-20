package elements

// The MarkAsJunkResponseMessage element specifies the response message for a MarkAsJunk request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/markasjunkresponsemessage
import "encoding/xml"

type MarkAsJunkResponseMessage struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The MovedItemId element specifies the identifier of the item that was moved by the MarkAsJunk operation.
	MovedItemId *MovedItemId `xml:"MovedItemId"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
}

func (M *MarkAsJunkResponseMessage) SetForMarshal() {
	M.XMLName.Local = "m:MarkAsJunkResponseMessage"
}

func (M *MarkAsJunkResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
