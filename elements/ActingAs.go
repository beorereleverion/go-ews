package elements

// The ActingAs element identifies who the caller is sending as.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/actingas
import "encoding/xml"

type ActingAs struct {
	XMLName xml.Name

	// The EmailAddress element defines the primary SMTP address of a mailbox user.
	EmailAddress *EmailAddressNonEmptyStringType `xml:"EmailAddress"`
	// The RoutingType element represents the routing protocol for the recipient.
	RoutingType *RoutingTypeEmailAddress `xml:"RoutingType"`
}

func (A *ActingAs) SetForMarshal() {
	A.XMLName.Local = "m:ActingAs"
}

func (A *ActingAs) GetSchema() *Schema {
	return &SchemaMessages
}
