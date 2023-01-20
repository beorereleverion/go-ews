package elements

// The OriginalRecipients element represents a list of e-mail addresses of the first message recipients.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/originalrecipients
import "encoding/xml"

type OriginalRecipients struct {
	XMLName xml.Name

	// The Address element represents a fully resolved e-mail address.
	Address *AddressEmailAddressType `xml:"Address"`
}

func (O *OriginalRecipients) SetForMarshal() {
	O.XMLName.Local = "t:OriginalRecipients"
}

func (O *OriginalRecipients) GetSchema() *Schema {
	return &SchemaTypes
}
