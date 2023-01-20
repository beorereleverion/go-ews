package elements

// The MailboxHoldStatuses element specifies a list of one or more MailboxHoldStatus elements.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxholdstatuses
import "encoding/xml"

type MailboxHoldStatuses struct {
	XMLName xml.Name

	// The MailboxHoldStatus element specifies the hold status of the mailbox.
	MailboxHoldStatus *MailboxHoldStatus `xml:"MailboxHoldStatus"`
}

func (M *MailboxHoldStatuses) SetForMarshal() {
	M.XMLName.Local = "t:MailboxHoldStatuses"
}

func (M *MailboxHoldStatuses) GetSchema() *Schema {
	return &SchemaTypes
}
