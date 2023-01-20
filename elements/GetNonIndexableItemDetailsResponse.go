package elements

// The GetNonIndexableItemDetailsResponse element specifies the response to a GetNonIndexableItemDetails request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getnonindexableitemdetailsresponse
import "encoding/xml"

type GetNonIndexableItemDetailsResponse struct {
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

func (G *GetNonIndexableItemDetailsResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetNonIndexableItemDetailsResponse"
}

func (G *GetNonIndexableItemDetailsResponse) GetSchema() *Schema {
	return &SchemaMessages
}
