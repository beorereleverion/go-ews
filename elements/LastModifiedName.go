package elements

// The LastModifiedName element contains the display name of the last user to modify an item. This element is read-only. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/lastmodifiedname
import "encoding/xml"

type LastModifiedName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (L *LastModifiedName) SetForMarshal() {
	L.XMLName.Local = "t:LastModifiedName"
}

func (L *LastModifiedName) GetSchema() *Schema {
	return &SchemaTypes
}
