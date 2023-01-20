package elements

// The UID element uniquely identifies a calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/uid
import "encoding/xml"

type UID struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (U *UID) SetForMarshal() {
	U.XMLName.Local = "t:UID"
}

func (U *UID) GetSchema() *Schema {
	return &SchemaTypes
}
