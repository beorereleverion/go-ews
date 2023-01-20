package elements

// The UninstallAppResponse element specifies the response to an UninstallApp request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/uninstallappresponse
import "encoding/xml"

type UninstallAppResponse struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
}

func (U *UninstallAppResponse) SetForMarshal() {
	U.XMLName.Local = "m:UninstallAppResponse"
}

func (U *UninstallAppResponse) GetSchema() *Schema {
	return &SchemaMessages
}
