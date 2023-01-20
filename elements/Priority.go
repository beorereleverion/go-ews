package elements

// The Priority element indicates the order in which a rule is to be run.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/priority
import "encoding/xml"

type Priority struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (P *Priority) SetForMarshal() {
	P.XMLName.Local = "m:Priority"
}

func (P *Priority) GetSchema() *Schema {
	return &SchemaMessages
}
