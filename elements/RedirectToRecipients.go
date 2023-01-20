package elements

// The RedirectToRecipients element indicates the e-mail addresses to which messages are to be redirected.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/redirecttorecipients
import "encoding/xml"

type RedirectToRecipients struct {
	XMLName xml.Name

	// The Address element represents a fully resolved e-mail address.
	Address *AddressEmailAddressType `xml:"Address"`
}

func (R *RedirectToRecipients) SetForMarshal() {
	R.XMLName.Local = "m:RedirectToRecipients"
}

func (R *RedirectToRecipients) GetSchema() *Schema {
	return &SchemaMessages
}
