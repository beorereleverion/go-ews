package elements

// The IsResponseRequested element indicates whether a response to an item is requested.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isresponserequested
import "encoding/xml"

type IsResponseRequested struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsResponseRequested) SetForMarshal() {
	I.XMLName.Local = "t:IsResponseRequested"
}

func (I *IsResponseRequested) GetSchema() *Schema {
	return &SchemaTypes
}
