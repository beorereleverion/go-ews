package elements

// The Emails2 element contains an array of EmailAddressAttributedValue values and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emails2
import "encoding/xml"

type Emails2 struct {
	XMLName xml.Name

	// The EmailAddressAttributedValue element specifies an instance of an array of email addresses and their associated attributions.
	EmailAddressAttributedValue *EmailAddressAttributedValue `xml:"EmailAddressAttributedValue"`
}

func (E *Emails2) SetForMarshal() {
	E.XMLName.Local = "t:Emails2"
}

func (E *Emails2) GetSchema() *Schema {
	return &SchemaTypes
}
