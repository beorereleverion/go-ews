package elements

// The DisplayName element contains the display name of a new instant messaging group contact or the display name of a new instant messaging group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/displayname-nonemptystringtype
import "encoding/xml"

type DisplayNameNonEmptyStringType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DisplayNameNonEmptyStringType) SetForMarshal() {
	D.XMLName.Local = "t:DisplayName"
}

func (D *DisplayNameNonEmptyStringType) GetSchema() *Schema {
	return &SchemaTypes
}
