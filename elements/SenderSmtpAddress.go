package elements

// The SenderSmtpAddress element represents the SMTP e-mail address corresponding to the mailbox that contains the folder that will be shared.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sendersmtpaddress
import "encoding/xml"

type SenderSmtpAddress struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *SenderSmtpAddress) SetForMarshal() {
	S.XMLName.Local = "m:SenderSmtpAddress"
}

func (S *SenderSmtpAddress) GetSchema() *Schema {
	return &SchemaMessages
}
