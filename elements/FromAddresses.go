package elements

// The FromAddresses element indicates the e-mail addresses from which incoming messages must be sent in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fromaddresses
import "encoding/xml"

type FromAddresses struct {
	XMLName xml.Name

	// The Address element represents a fully resolved e-mail address.
	Address *AddressEmailAddressType `xml:"Address"`
}

func (F *FromAddresses) SetForMarshal() {
	F.XMLName.Local = "m:FromAddresses"
}

func (F *FromAddresses) GetSchema() *Schema {
	return &SchemaMessages
}
