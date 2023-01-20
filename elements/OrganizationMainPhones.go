package elements

// The OrganizationMainPhones element specifies an array of organizational main phone numbers and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/organizationmainphones
import "encoding/xml"

type OrganizationMainPhones struct {
	XMLName xml.Name

	// The PhoneNumberAttributedValue element specifies an instance of an array of phone numbers and their associated attributions.
	PhoneNumberAttributedValue *PhoneNumberAttributedValue `xml:"PhoneNumberAttributedValue"`
}

func (O *OrganizationMainPhones) SetForMarshal() {
	O.XMLName.Local = "t:OrganizationMainPhones"
}

func (O *OrganizationMainPhones) GetSchema() *Schema {
	return &SchemaTypes
}
