package elements

// The Senders element specifies an array of Simple Mail Transfer Protocol (SMTP) addresses.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/senders
import "encoding/xml"

type Senders struct {
	XMLName xml.Name

	// The SmtpAddress element represents the Simple Mail Transfer Protocol (SMTP) address of an account to be used for impersonation or a Simple Mail Transfer Protocol (SMTP) recipient address of a calendar or contact sharing request.
	SmtpAddress *SmtpAddress `xml:"SmtpAddress"`
}

func (S *Senders) SetForMarshal() {
	S.XMLName.Local = "t:Senders"
}

func (S *Senders) GetSchema() *Schema {
	return &SchemaTypes
}
