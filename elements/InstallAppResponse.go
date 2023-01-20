package elements

// The InstallAppResponse element specifies the response to an InstallApp request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/installappresponse
import "encoding/xml"

type InstallAppResponse struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// Indicates the class of the response.
	ResponseClass *string `xml:"ResponseClass,attr"`
}

func (I *InstallAppResponse) SetForMarshal() {
	I.XMLName.Local = "m:InstallAppResponse"
}

func (I *InstallAppResponse) GetSchema() *Schema {
	return &SchemaMessages
}
