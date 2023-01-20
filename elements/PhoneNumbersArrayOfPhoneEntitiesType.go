package elements

// The PhoneNumbers element specifies an array of extracted phone numbers.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phonenumbers-arrayofphoneentitiestype
import "encoding/xml"

type PhoneNumbersArrayOfPhoneEntitiesType struct {
	XMLName xml.Name

	// The Phone element specifies a single phone number that results from a phone number entity extraction.
	Phone *PhonePhoneEntityType `xml:"Phone"`
}

func (P *PhoneNumbersArrayOfPhoneEntitiesType) SetForMarshal() {
	P.XMLName.Local = "t:PhoneNumbers"
}

func (P *PhoneNumbersArrayOfPhoneEntitiesType) GetSchema() *Schema {
	return &SchemaTypes
}
