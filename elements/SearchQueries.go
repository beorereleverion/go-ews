package elements

// The SearchQueries element contains a list of mailboxes and associated queries for discovery search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchqueries
import "encoding/xml"

type SearchQueries struct {
	XMLName xml.Name

	// The MailboxQuery element specifies a query and the scope of a discovery search.
	MailboxQuery *MailboxQuery `xml:"MailboxQuery"`
}

func (S *SearchQueries) SetForMarshal() {
	S.XMLName.Local = "m:SearchQueries"
}

func (S *SearchQueries) GetSchema() *Schema {
	return &SchemaMessages
}
