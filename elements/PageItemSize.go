package elements

// The PageItemSize element specifies the number of items to return in a search result pagination.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/pageitemsize
import "encoding/xml"

type PageItemSize struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (P *PageItemSize) SetForMarshal() {
	P.XMLName.Local = "t:PageItemSize"
}

func (P *PageItemSize) GetSchema() *Schema {
	return &SchemaTypes
}
