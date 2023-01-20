package elements

// The OtherFaxes element specifies an array of fax phone number values and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/otherfaxes
import "encoding/xml"

type OtherFaxes struct {
	XMLName xml.Name

	// The PhoneNumberAttributedValue element specifies an instance of an array of phone numbers and their associated attributions.
	PhoneNumberAttributedValue *PhoneNumberAttributedValue `xml:"PhoneNumberAttributedValue"`
}

func (O *OtherFaxes) SetForMarshal() {
	O.XMLName.Local = "t:OtherFaxes"
}

func (O *OtherFaxes) GetSchema() *Schema {
	return &SchemaTypes
}
