package elements

// The UnresolvedEntry element contains the name of a contact or distribution list to resolve.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/unresolvedentry
import "encoding/xml"

type UnresolvedEntry struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (U *UnresolvedEntry) SetForMarshal() {
	U.XMLName.Local = "m:UnresolvedEntry"
}

func (U *UnresolvedEntry) GetSchema() *Schema {
	return &SchemaMessages
}
