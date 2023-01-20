package elements

// The ForwardToRecipients element indicates the e-mail addresses to which messages are to be forwarded.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/forwardtorecipients
import "encoding/xml"

type ForwardToRecipients struct {
	XMLName xml.Name

	// The Address element represents a fully resolved e-mail address.
	Address *AddressEmailAddressType `xml:"Address"`
}

func (F *ForwardToRecipients) SetForMarshal() {
	F.XMLName.Local = "m:ForwardToRecipients"
}

func (F *ForwardToRecipients) GetSchema() *Schema {
	return &SchemaMessages
}
