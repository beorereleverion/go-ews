package elements

// The IsResend element indicates whether the item had previously been sent.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isresend
import "encoding/xml"

type IsResend struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsResend) SetForMarshal() {
	I.XMLName.Local = "t:IsResend"
}

func (I *IsResend) GetSchema() *Schema {
	return &SchemaTypes
}
