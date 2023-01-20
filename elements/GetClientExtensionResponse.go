package elements

// The GetClientExtensionResponse element contains the response to get configuration information about an app.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getclientextensionresponse
import "encoding/xml"

type GetClientExtensionResponse struct {
	XMLName xml.Name

	// The ClientExtensions element contains an array of user and configuration information about apps.
	ClientExtensions *ClientExtensions `xml:"ClientExtensions"`
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The RawMasterTableXml element is not used.
	RawMasterTableXml *RawMasterTableXml `xml:"RawMasterTableXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
}

func (G *GetClientExtensionResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetClientExtensionResponse"
}

func (G *GetClientExtensionResponse) GetSchema() *Schema {
	return &SchemaMessages
}
