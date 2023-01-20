package elements

// The PageSize element contains the number of items to be returned in a single page for a search result.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/pagesize
import "encoding/xml"

type PageSize struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (P *PageSize) SetForMarshal() {
	P.XMLName.Local = "m:PageSize"
}

func (P *PageSize) GetSchema() *Schema {
	return &SchemaMessages
}
