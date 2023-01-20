package elements

// The EmailAddresses element specifies an array of extracted email addresses.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emailaddresses-arrayofextractedemailaddresses
import "encoding/xml"

type EmailAddressesArrayOfExtractedEmailAddresses struct {
	XMLName xml.Name

	// The EmailAddress element specifies a single email address.
	EmailAddress *EmailAddressstring `xml:"EmailAddress"`
}

func (E *EmailAddressesArrayOfExtractedEmailAddresses) SetForMarshal() {
	E.XMLName.Local = "t:EmailAddresses"
}

func (E *EmailAddressesArrayOfExtractedEmailAddresses) GetSchema() *Schema {
	return &SchemaTypes
}
