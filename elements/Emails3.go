package elements

// The Emails3 element specifies an array of EmailAddressAttributedValue values and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emails3
import "encoding/xml"

type Emails3 struct {
	XMLName xml.Name

	// The EmailAddressAttributedValue element specifies an instance of an array of email addresses and their associated attributions.
	EmailAddressAttributedValue *EmailAddressAttributedValue `xml:"EmailAddressAttributedValue"`
}

func (E *Emails3) SetForMarshal() {
	E.XMLName.Local = "t:Emails3"
}

func (E *Emails3) GetSchema() *Schema {
	return &SchemaTypes
}
