package elements

// The GetPasswordExpirationDate element defines a request to get the password expiration date for an email account. This element is the base element for the GetPasswordExpirationDate operation operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getpasswordexpirationdate
import "encoding/xml"

type GetPasswordExpirationDate struct {
	XMLName xml.Name

	// The MailboxSmtpAddress element represents the SMTP address of the user whose Inbox rules are to be retrieved or updated; or whose password expiration date is to be retrieved.
	MailboxSmtpAddress *MailboxSmtpAddress `xml:"MailboxSmtpAddress"`
}

func (G *GetPasswordExpirationDate) SetForMarshal() {
	G.XMLName.Local = "m:GetPasswordExpirationDate"
}

func (G *GetPasswordExpirationDate) GetSchema() *Schema {
	return &SchemaMessages
}
