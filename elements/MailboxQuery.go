package elements

// The MailboxQuery element specifies a query and the scope of a discovery search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxquery
import "encoding/xml"

type MailboxQuery struct {
	XMLName xml.Name

	// The MailboxSearchScopes element specifies a list of one or more mailboxes and associated search scopes for a discovery search.
	MailboxSearchScopes *MailboxSearchScopes `xml:"MailboxSearchScopes"`
	// The Query element contains the search query for the hold.
	Query *Query `xml:"Query"`
}

func (M *MailboxQuery) SetForMarshal() {
	M.XMLName.Local = "t:MailboxQuery"
}

func (M *MailboxQuery) GetSchema() *Schema {
	return &SchemaTypes
}
