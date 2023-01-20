package elements

// The Description element specifies the descriptive text for the retention policy.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/description
import "encoding/xml"

type Description struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *Description) SetForMarshal() {
	D.XMLName.Local = "t:Description"
}

func (D *Description) GetSchema() *Schema {
	return &SchemaTypes
}
