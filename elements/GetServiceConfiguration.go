package elements

// The GetServiceConfiguration element defines a GetServiceConfiguration request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getserviceconfiguration
import "encoding/xml"

type GetServiceConfiguration struct {
	XMLName xml.Name

	// The ActingAs element identifies who the caller is sending as.
	ActingAs *ActingAs `xml:"ActingAs"`
	// The RequestedConfiguration element contains the requested service configurations.
	RequestedConfiguration *RequestedConfiguration `xml:"RequestedConfiguration"`
}

func (G *GetServiceConfiguration) SetForMarshal() {
	G.XMLName.Local = "m:GetServiceConfiguration"
}

func (G *GetServiceConfiguration) GetSchema() *Schema {
	return &SchemaMessages
}
