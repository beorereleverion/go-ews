package elements

// The InternalId element represents an integer value for the event identification.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/internalid
import "encoding/xml"

type InternalId struct {
	XMLName xml.Name
	TEXT    uint32 `xml:",chardata"`
}

func (I *InternalId) SetForMarshal() {
	I.XMLName.Local = "t:InternalId"
}

func (I *InternalId) GetSchema() *Schema {
	return &SchemaTypes
}
