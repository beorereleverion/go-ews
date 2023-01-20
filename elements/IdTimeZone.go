package elements

// The Id element identifies a single time zone definition.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/id-timezone
import "encoding/xml"

type IdTimeZone struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (I *IdTimeZone) SetForMarshal() {
	I.XMLName.Local = "t:Id"
}

func (I *IdTimeZone) GetSchema() *Schema {
	return &SchemaTypes
}
