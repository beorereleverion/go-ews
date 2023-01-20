package elements

// The ContentType element describes the Multipurpose Internet Mail Extensions (MIME) type of the attachment content.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/contenttype
import "encoding/xml"

type ContentType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (C *ContentType) SetForMarshal() {
	C.XMLName.Local = "t:ContentType"
}

func (C *ContentType) GetSchema() *Schema {
	return &SchemaTypes
}
