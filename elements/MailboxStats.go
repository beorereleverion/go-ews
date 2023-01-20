package elements

// The MailboxStats element specifies a list of one or more MailboxStat elements.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxstats
import "encoding/xml"

type MailboxStats struct {
	XMLName xml.Name

	// The MailboxStat element specifies statistics for a mailbox searched by discovery search.
	MailboxStat *MailboxStat `xml:"MailboxStat"`
}

func (M *MailboxStats) SetForMarshal() {
	M.XMLName.Local = "t:MailboxStats"
}

func (M *MailboxStats) GetSchema() *Schema {
	return &SchemaTypes
}
