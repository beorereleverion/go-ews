package elements

// The HomeFaxes element specifies an array of home fax numbers and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/homefaxes
import "encoding/xml"

type HomeFaxes struct {
	XMLName xml.Name

	// The PhoneNumberAttributedValue element specifies an instance of an array of phone numbers and their associated attributions.
	PhoneNumberAttributedValue *PhoneNumberAttributedValue `xml:"PhoneNumberAttributedValue"`
}

func (H *HomeFaxes) SetForMarshal() {
	H.XMLName.Local = "t:HomeFaxes"
}

func (H *HomeFaxes) GetSchema() *Schema {
	return &SchemaTypes
}
