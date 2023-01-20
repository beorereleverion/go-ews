package elements

// The LegacyDN element identifies a mailbox by its legacy distinguished name.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/legacydn
import "encoding/xml"

type LegacyDN struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (L *LegacyDN) SetForMarshal() {
	L.XMLName.Local = "t:LegacyDN"
}

func (L *LegacyDN) GetSchema() *Schema {
	return &SchemaTypes
}
