package elements

// The SmtpAddress element represents the Simple Mail Transfer Protocol (SMTP) address of an account to be used for impersonation or a Simple Mail Transfer Protocol (SMTP) recipient address of a calendar or contact sharing request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/smtpaddress
import "encoding/xml"

type SmtpAddress struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *SmtpAddress) SetForMarshal() {
	S.XMLName.Local = "t:SmtpAddress"
}

func (S *SmtpAddress) GetSchema() *Schema {
	return &SchemaTypes
}
