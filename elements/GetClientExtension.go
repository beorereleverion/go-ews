package elements

// The GetClientExtension element represents a request to get a client extension.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getclientextension
import "encoding/xml"

type GetClientExtension struct {
	XMLName xml.Name

	// The IsDebug element is not used.
	IsDebug *IsDebug `xml:"IsDebug"`
	// The RequestedExtensionIds element contains an array of extension identifiers.
	RequestedExtensionIds *RequestedExtensionIds `xml:"RequestedExtensionIds"`
	// The UserParameters element contains a list of enabled and disabled client extensions.
	UserParameters *UserParameters `xml:"UserParameters"`
}

func (G *GetClientExtension) SetForMarshal() {
	G.XMLName.Local = "m:GetClientExtension"
}

func (G *GetClientExtension) GetSchema() *Schema {
	return &SchemaMessages
}
