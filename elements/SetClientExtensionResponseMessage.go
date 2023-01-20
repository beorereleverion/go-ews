package elements

// The SetClientExtensionResponseMessage element specifies the response message for a SetClientExtension request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setclientextensionresponsemessage
import "encoding/xml"

type SetClientExtensionResponseMessage struct {
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

func (S *SetClientExtensionResponseMessage) SetForMarshal() {
	S.XMLName.Local = "t:SetClientExtensionResponseMessage"
}

func (S *SetClientExtensionResponseMessage) GetSchema() *Schema {
	return &SchemaTypes
}
