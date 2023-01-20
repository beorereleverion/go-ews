package elements

// The WorkFaxes element specifies an array of work fax numbers and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/workfaxes
import "encoding/xml"

type WorkFaxes struct {
	XMLName xml.Name

	// The PhoneNumberAttributedValue element specifies an instance of an array of phone numbers and their associated attributions.
	PhoneNumberAttributedValue *PhoneNumberAttributedValue `xml:"PhoneNumberAttributedValue"`
}

func (W *WorkFaxes) SetForMarshal() {
	W.XMLName.Local = "t:WorkFaxes"
}

func (W *WorkFaxes) GetSchema() *Schema {
	return &SchemaTypes
}
