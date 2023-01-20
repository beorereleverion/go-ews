package elements

// The MimeContentUTF8 element contains the UTF-8 MIME stream of an object that is represented in base64Binary format and supports email address internationalization and [RFC6530].
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mimecontentutf8
import "encoding/xml"

type MimeContentUTF8 struct {
	XMLName xml.Name

	// If set, the value for this attribute is ignored by the server.
	CharacterSet *string `xml:"CharacterSet,attr"`
}

func (M *MimeContentUTF8) SetForMarshal() {
	M.XMLName.Local = "t:MimeContentUTF8"
}

func (M *MimeContentUTF8) GetSchema() *Schema {
	return &SchemaTypes
}
