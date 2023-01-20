package elements

// The PageItemReference element specifies the reference for a page item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/pageitemreference
import "encoding/xml"

type PageItemReference struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *PageItemReference) SetForMarshal() {
	P.XMLName.Local = "m:PageItemReference"
}

func (P *PageItemReference) GetSchema() *Schema {
	return &SchemaMessages
}
