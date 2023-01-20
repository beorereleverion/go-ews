package elements

// The SearchableMailboxes element contains an array of the mailboxes returned from a GetSearchableMailboxes request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchablemailboxes
import "encoding/xml"

type SearchableMailboxes struct {
	XMLName xml.Name

	// The SearchableMailbox element specifies a mailbox returned from a GetSearchableMailboxes request.
	SearchableMailbox *SearchableMailbox `xml:"SearchableMailbox"`
}

func (S *SearchableMailboxes) SetForMarshal() {
	S.XMLName.Local = "m:SearchableMailboxes"
}

func (S *SearchableMailboxes) GetSchema() *Schema {
	return &SchemaMessages
}
