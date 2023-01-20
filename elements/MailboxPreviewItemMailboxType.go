package elements

// The Mailbox element contains the mailbox identifier and the user's primary Simple Mail Transfer Protocol (SMTP) address.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailbox-previewitemmailboxtype
import "encoding/xml"

type MailboxPreviewItemMailboxType struct {
	XMLName xml.Name

	// The MailboxId element specifies an identifier for the mailbox that is accessed by discovery search.
	MailboxId *MailboxId `xml:"MailboxId"`
	// The PrimarySmtpAddress element specifies the primary Simple Mail Transfer Protocol (SMTP) address of the mailbox.
	PrimarySmtpAddress *PrimarySmtpAddressstring `xml:"PrimarySmtpAddress"`
}

func (M *MailboxPreviewItemMailboxType) SetForMarshal() {
	M.XMLName.Local = "t:Mailbox"
}

func (M *MailboxPreviewItemMailboxType) GetSchema() *Schema {
	return &SchemaTypes
}
