package elements

// The GetNonIndexableItemDetailsResponseMessage element specifies the response message for a GetNonIndexableItemDetails request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getnonindexableitemdetailsresponsemessage
import "encoding/xml"

type GetNonIndexableItemDetailsResponseMessage struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The NonIndexableItemDetailsResult element specifies the results of the GetNonIndexableItemDetails WSDL operation.
	NonIndexableItemDetailsResult *NonIndexableItemDetailsResult `xml:"NonIndexableItemDetailsResult"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
}

func (G *GetNonIndexableItemDetailsResponseMessage) SetForMarshal() {
	G.XMLName.Local = "m:GetNonIndexableItemDetailsResponseMessage"
}

func (G *GetNonIndexableItemDetailsResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
