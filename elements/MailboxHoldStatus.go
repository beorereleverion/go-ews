package elements

// The MailboxHoldStatus element specifies the hold status of the mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxholdstatus
import "encoding/xml"

type MailboxHoldStatus struct {
	XMLName xml.Name

	// The AdditionalInfo element specifies additional information about the hold status of a mailbox.
	AdditionalInfo *AdditionalInfo `xml:"AdditionalInfo"`
	// The Mailbox element contains an identifier for a mailbox.
	Mailbox *Mailboxstring `xml:"Mailbox"`
	// The Status element specifies the hold status for a mailbox.
	Status *StatusHoldStatusType `xml:"Status"`
}

func (M *MailboxHoldStatus) SetForMarshal() {
	M.XMLName.Local = "t:MailboxHoldStatus"
}

func (M *MailboxHoldStatus) GetSchema() *Schema {
	return &SchemaTypes
}
