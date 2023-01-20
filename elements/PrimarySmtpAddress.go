package elements

// The PrimarySmtpAddress element represents the primary Simple Mail Transfer Protocol (SMTP) address of an account to be used for server-to-server authorization or delegate access.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/primarysmtpaddress
import "encoding/xml"

type PrimarySmtpAddress struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *PrimarySmtpAddress) SetForMarshal() {
	P.XMLName.Local = "t:PrimarySmtpAddress"
}

func (P *PrimarySmtpAddress) GetSchema() *Schema {
	return &SchemaTypes
}
