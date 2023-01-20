package elements

// The MailboxSearchScopes element specifies a list of one or more mailboxes and associated search scopes for a discovery search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxsearchscopes
import "encoding/xml"

type MailboxSearchScopes struct {
	XMLName xml.Name

	// The MailboxSearchScope element specifies a mailbox and a search scope for a discovery search.
	MailboxSearchScope *MailboxSearchScope `xml:"MailboxSearchScope"`
}

func (M *MailboxSearchScopes) SetForMarshal() {
	M.XMLName.Local = "t:MailboxSearchScopes"
}

func (M *MailboxSearchScopes) GetSchema() *Schema {
	return &SchemaTypes
}
