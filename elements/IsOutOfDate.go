package elements

// The IsOutOfDate element indicates whether a meeting message, request, response, or cancellation is out-of-date.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isoutofdate
import "encoding/xml"

type IsOutOfDate struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsOutOfDate) SetForMarshal() {
	I.XMLName.Local = "t:IsOutOfDate"
}

func (I *IsOutOfDate) GetSchema() *Schema {
	return &SchemaTypes
}
