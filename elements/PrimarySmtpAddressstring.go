package elements

// The PrimarySmtpAddress element specifies the primary Simple Mail Transfer Protocol (SMTP) address of the mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/primarysmtpaddress-string
import "encoding/xml"

type PrimarySmtpAddressstring struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *PrimarySmtpAddressstring) SetForMarshal() {
	P.XMLName.Local = "t:PrimarySmtpAddress"
}

func (P *PrimarySmtpAddressstring) GetSchema() *Schema {
	return &SchemaTypes
}
