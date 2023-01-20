package elements

// The True element specifies a condition that always matches.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/true
import "encoding/xml"

type True struct {
	XMLName xml.Name
	TEXT    struct{} `xml:",chardata"`
}

func (T *True) SetForMarshal() {
	T.XMLName.Local = "t:True"
}

func (T *True) GetSchema() *Schema {
	return &SchemaTypes
}
