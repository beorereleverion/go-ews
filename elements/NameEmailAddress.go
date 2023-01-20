package elements

// The Name element represents the display name of the mailbox user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/name-emailaddress
import "encoding/xml"

type NameEmailAddress struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (N *NameEmailAddress) SetForMarshal() {
	N.XMLName.Local = "t:Name"
}

func (N *NameEmailAddress) GetSchema() *Schema {
	return &SchemaTypes
}
