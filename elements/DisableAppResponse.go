package elements

// The DisableAppResponse element specifies the response to a DisableApp request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/disableappresponse
import "encoding/xml"

type DisableAppResponse struct {
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

func (D *DisableAppResponse) SetForMarshal() {
	D.XMLName.Local = "m:DisableAppResponse"
}

func (D *DisableAppResponse) GetSchema() *Schema {
	return &SchemaMessages
}
