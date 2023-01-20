package elements

// The EmailAddresses element specifies an array of email address entities.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emailaddresses-arrayofemailaddressentitiestype
import "encoding/xml"

type EmailAddressesArrayOfEmailAddressEntitiesType struct {
	XMLName xml.Name

	// The EmailAddressEntity element specifies a single email address entity.
	EmailAddressEntity *EmailAddressEntity `xml:"EmailAddressEntity"`
}

func (E *EmailAddressesArrayOfEmailAddressEntitiesType) SetForMarshal() {
	E.XMLName.Local = "t:EmailAddresses"
}

func (E *EmailAddressesArrayOfEmailAddressEntitiesType) GetSchema() *Schema {
	return &SchemaTypes
}
