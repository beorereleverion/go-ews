package elements

// The MimeContent element contains the ASCII MIME stream of an object that is represented in base64Binary format and supports [RFC2045].
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mimecontent
import "encoding/xml"

type MimeContent struct {
	XMLName xml.Name

	// If set, the value for this attribute is ignored by the server.
	CharacterSet *string `xml:"CharacterSet,attr"`
}

func (M *MimeContent) SetForMarshal() {
	M.XMLName.Local = "t:MimeContent"
}

func (M *MimeContent) GetSchema() *Schema {
	return &SchemaTypes
}
