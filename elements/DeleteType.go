package elements

// The DeleteType element indicates how items in a conversation are deleted.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deletetype
import "encoding/xml"

type DeleteType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DeleteType) SetForMarshal() {
	D.XMLName.Local = "t:DeleteType"
}

func (D *DeleteType) GetSchema() *Schema {
	return &SchemaTypes
}
