package elements

// The ReturnNewItemIds element indicates whether the item identifiers of new items are returned in the response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/returnnewitemids
import "encoding/xml"

type ReturnNewItemIds struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (R *ReturnNewItemIds) SetForMarshal() {
	R.XMLName.Local = "m:ReturnNewItemIds"
}

func (R *ReturnNewItemIds) GetSchema() *Schema {
	return &SchemaMessages
}
