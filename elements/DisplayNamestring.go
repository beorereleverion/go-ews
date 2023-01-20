package elements

// The DisplayName element defines the display name of a folder, contact, distribution list, delegate user, location, or rule.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/displayname-string
import "encoding/xml"

type DisplayNamestring struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DisplayNamestring) SetForMarshal() {
	D.XMLName.Local = "t:DisplayName"
}

func (D *DisplayNamestring) GetSchema() *Schema {
	return &SchemaTypes
}
