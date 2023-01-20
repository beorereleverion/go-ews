package elements

// The SearchId element specifies the identifier of a discovery search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchid
import "encoding/xml"

type SearchId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *SearchId) SetForMarshal() {
	S.XMLName.Local = "t:SearchId"
}

func (S *SearchId) GetSchema() *Schema {
	return &SchemaTypes
}
