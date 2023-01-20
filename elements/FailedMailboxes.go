package elements

// The FailedMailboxes element specifies an array of mailboxes that failed on search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/failedmailboxes
import "encoding/xml"

type FailedMailboxes struct {
	XMLName xml.Name

	// The FailedMailbox element specifies the error message for a mailbox that failed on search.
	FailedMailbox *FailedMailbox `xml:"FailedMailbox"`
}

func (F *FailedMailboxes) SetForMarshal() {
	F.XMLName.Local = "t:FailedMailboxes"
}

func (F *FailedMailboxes) GetSchema() *Schema {
	return &SchemaTypes
}
