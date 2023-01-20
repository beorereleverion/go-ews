package elements

// The PhoneNumbers element specifies an array of phone numbers.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phonenumbers-arrayofphonestype
import "encoding/xml"

type PhoneNumbersArrayOfPhonesType struct {
	XMLName xml.Name

	// The Phone element specifies a single phone number that results from a contact entity extraction.
	Phone *Phone `xml:"Phone"`
}

func (P *PhoneNumbersArrayOfPhonesType) SetForMarshal() {
	P.XMLName.Local = "t:PhoneNumbers"
}

func (P *PhoneNumbersArrayOfPhonesType) GetSchema() *Schema {
	return &SchemaTypes
}
