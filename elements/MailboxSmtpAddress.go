package elements

// The MailboxSmtpAddress element represents the SMTP address of the user whose Inbox rules are to be retrieved or updated; or whose password expiration date is to be retrieved.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxsmtpaddress
import "encoding/xml"

type MailboxSmtpAddress struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (M *MailboxSmtpAddress) SetForMarshal() {
	M.XMLName.Local = "t:MailboxSmtpAddress"
}

func (M *MailboxSmtpAddress) GetSchema() *Schema {
	return &SchemaTypes
}
