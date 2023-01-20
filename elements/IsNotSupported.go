package elements

// The IsNotSupported element indicates whether the rule cannot be modified by using the managed code APIs.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isnotsupported
import "encoding/xml"

type IsNotSupported struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsNotSupported) SetForMarshal() {
	I.XMLName.Local = "m:IsNotSupported"
}

func (I *IsNotSupported) GetSchema() *Schema {
	return &SchemaMessages
}
