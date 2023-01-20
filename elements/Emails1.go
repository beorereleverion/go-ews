package elements

// The Emails1 element specifies an array of EmailAddressAttributedValue values and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emails1
import "encoding/xml"

type Emails1 struct {
	XMLName xml.Name

	// The EmailAddressAttributedValue element specifies an instance of an array of email addresses and their associated attributions.
	EmailAddressAttributedValue *EmailAddressAttributedValue `xml:"EmailAddressAttributedValue"`
}

func (E *Emails1) SetForMarshal() {
	E.XMLName.Local = "t:Emails1"
}

func (E *Emails1) GetSchema() *Schema {
	return &SchemaTypes
}
