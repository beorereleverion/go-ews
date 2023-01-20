package elements

// The ToRecipients element specifies a list of recipients to whom the item was sent.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/torecipients-arrayofsmtpaddresstype
import "encoding/xml"

type ToRecipientsArrayOfSmtpAddressType struct {
	XMLName xml.Name

	// The SmtpAddress element represents the Simple Mail Transfer Protocol (SMTP) address of an account to be used for impersonation or a Simple Mail Transfer Protocol (SMTP) recipient address of a calendar or contact sharing request.
	SmtpAddress *SmtpAddress `xml:"SmtpAddress"`
}

func (T *ToRecipientsArrayOfSmtpAddressType) SetForMarshal() {
	T.XMLName.Local = "t:ToRecipients"
}

func (T *ToRecipientsArrayOfSmtpAddressType) GetSchema() *Schema {
	return &SchemaTypes
}
