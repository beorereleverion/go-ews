package elements

// The MobilePhones2 element specifies an array of MobilePhone values and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mobilephones2
import "encoding/xml"

type MobilePhones2 struct {
	XMLName xml.Name

	// The PhoneNumberAttributedValue element specifies an instance of an array of phone numbers and their associated attributions.
	PhoneNumberAttributedValue *PhoneNumberAttributedValue `xml:"PhoneNumberAttributedValue"`
}

func (M *MobilePhones2) SetForMarshal() {
	M.XMLName.Local = "t:MobilePhones2"
}

func (M *MobilePhones2) GetSchema() *Schema {
	return &SchemaTypes
}
