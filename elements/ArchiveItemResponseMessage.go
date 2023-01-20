package elements

// The ArchiveItemResponseMessage element specifies the response message to an ArchiveItem request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/archiveitemresponsemessage
import "encoding/xml"

type ArchiveItemResponseMessage struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The Items element contains an array of items.
	Items *Items `xml:"Items"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// Indicates the class of the response.
	ResponseClass *string `xml:"ResponseClass,attr"`
}

func (A *ArchiveItemResponseMessage) SetForMarshal() {
	A.XMLName.Local = "m:ArchiveItemResponseMessage"
}

func (A *ArchiveItemResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
