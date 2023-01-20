package elements

// The SearchFilter element contains the query string to filter the mailboxes to be returned from a GetSearchableMailboxes request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchfilter
import "encoding/xml"

type SearchFilter struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *SearchFilter) SetForMarshal() {
	S.XMLName.Local = "m:SearchFilter"
}

func (S *SearchFilter) GetSchema() *Schema {
	return &SchemaMessages
}
