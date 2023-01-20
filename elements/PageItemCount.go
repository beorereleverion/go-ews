package elements

// The PageItemCount element specifies the number of pages returned in a search result pagination.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/pageitemcount
import "encoding/xml"

type PageItemCount struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (P *PageItemCount) SetForMarshal() {
	P.XMLName.Local = "t:PageItemCount"
}

func (P *PageItemCount) GetSchema() *Schema {
	return &SchemaTypes
}
