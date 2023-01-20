package elements

// The OtherTelephones element specifies an array of telephone values and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/othertelephones
import "encoding/xml"

type OtherTelephones struct {
	XMLName xml.Name

	// The PhoneNumberAttributedValue element specifies an instance of an array of phone numbers and their associated attributions.
	PhoneNumberAttributedValue *PhoneNumberAttributedValue `xml:"PhoneNumberAttributedValue"`
}

func (O *OtherTelephones) SetForMarshal() {
	O.XMLName.Local = "t:OtherTelephones"
}

func (O *OtherTelephones) GetSchema() *Schema {
	return &SchemaTypes
}
