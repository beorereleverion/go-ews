package elements

// The ContentLocation element contains the Uniform Resource Identifier (URI) that corresponds to the location of the content of an attachment.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/contentlocation
import "encoding/xml"

type ContentLocation struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (C *ContentLocation) SetForMarshal() {
	C.XMLName.Local = "t:ContentLocation"
}

func (C *ContentLocation) GetSchema() *Schema {
	return &SchemaTypes
}
