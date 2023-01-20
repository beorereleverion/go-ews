package elements

// The InPlaceHoldIdentity element specifies the identity of a hold that preserves the mailbox items.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/inplaceholdidentity
import "encoding/xml"

type InPlaceHoldIdentity struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (I *InPlaceHoldIdentity) SetForMarshal() {
	I.XMLName.Local = "m:InPlaceHoldIdentity"
}

func (I *InPlaceHoldIdentity) GetSchema() *Schema {
	return &SchemaMessages
}
