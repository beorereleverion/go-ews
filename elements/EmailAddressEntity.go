package elements

// The EmailAddressEntity element specifies a single email address entity.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emailaddressentity
import "encoding/xml"

type EmailAddressEntity struct {
	XMLName xml.Name

	// The EmailAddress element specifies a single email address.
	EmailAddress *EmailAddressstring `xml:"EmailAddress"`
}

func (E *EmailAddressEntity) SetForMarshal() {
	E.XMLName.Local = "t:EmailAddressEntity"
}

func (E *EmailAddressEntity) GetSchema() *Schema {
	return &SchemaTypes
}
