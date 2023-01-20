package elements

// The Ids element contains an array of time zone definition identifiers.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ids
import "encoding/xml"

type Ids struct {
	XMLName xml.Name

	// The Id element identifies a single time zone definition.
	Id *IdTimeZone `xml:"Id"`
}

func (I *Ids) SetForMarshal() {
	I.XMLName.Local = "m:Ids"
}

func (I *Ids) GetSchema() *Schema {
	return &SchemaMessages
}
