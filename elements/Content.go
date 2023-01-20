package elements

// The Content element contains the Base64-encoded contents of a file attachment.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/content
import "encoding/xml"

type Content struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (C *Content) SetForMarshal() {
	C.XMLName.Local = "t:Content"
}

func (C *Content) GetSchema() *Schema {
	return &SchemaTypes
}
