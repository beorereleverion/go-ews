package elements

// The Path element is the base schema type for all property identifiers. This type is abstract and will never occur directly within instance documents.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/path
import "encoding/xml"

type Path struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (P *Path) SetForMarshal() {
	P.XMLName.Local = "t:Path"
}

func (P *Path) GetSchema() *Schema {
	return &SchemaTypes
}
