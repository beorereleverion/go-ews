package elements

// The MailboxStat element specifies statistics for a mailbox searched by discovery search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxstat
import "encoding/xml"

type MailboxStat struct {
	XMLName xml.Name

	// The DisplayName element defines the display name of a folder, contact, distribution list, delegate user, location, or rule.
	DisplayName *DisplayNamestring `xml:"DisplayName"`
	// The ItemCount element specifies the total number of items in a search result.
	ItemCount *ItemCount `xml:"ItemCount"`
	// The MailboxId element specifies an identifier for the mailbox that is accessed by discovery search.
	MailboxId *MailboxId `xml:"MailboxId"`
	// The Size element specifies the total size of one or more mailbox items.
	Size *Sizelong `xml:"Size"`
}

func (M *MailboxStat) SetForMarshal() {
	M.XMLName.Local = "t:MailboxStat"
}

func (M *MailboxStat) GetSchema() *Schema {
	return &SchemaTypes
}
