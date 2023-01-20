package elements

// The IsEnabled element indicates whether the rule is enabled.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isenabled
import "encoding/xml"

type IsEnabled struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsEnabled) SetForMarshal() {
	I.XMLName.Local = "m:IsEnabled"
}

func (I *IsEnabled) GetSchema() *Schema {
	return &SchemaMessages
}
