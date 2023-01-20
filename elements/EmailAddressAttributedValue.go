package elements

// The EmailAddressAttributedValue element specifies an instance of an array of email addresses and their associated attributions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emailaddressattributedvalue
import "encoding/xml"

type EmailAddressAttributedValue struct {
	XMLName xml.Name

	// The Attributions element specifies an array of attributions for its associated Value element.
	Attributions *AttributionsArrayOfValueAttributionsType `xml:"Attributions"`
	// The Value element specifies the value of an EmailAddress associated with an attributions array.
	Value *ValueEmailAddressType `xml:"Value"`
}

func (E *EmailAddressAttributedValue) SetForMarshal() {
	E.XMLName.Local = "t:EmailAddressAttributedValue"
}

func (E *EmailAddressAttributedValue) GetSchema() *Schema {
	return &SchemaTypes
}
