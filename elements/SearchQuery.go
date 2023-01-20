package elements

// The SearchQuery element specifies the discovery search query.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchquery
import "encoding/xml"

type SearchQuery struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *SearchQuery) SetForMarshal() {
	S.XMLName.Local = "t:SearchQuery"
}

func (S *SearchQuery) GetSchema() *Schema {
	return &SchemaTypes
}
