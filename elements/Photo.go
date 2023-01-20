package elements

// The Photo element contains a value that encodes the photo of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/photo
import "encoding/xml"

type Photo struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *Photo) SetForMarshal() {
	P.XMLName.Local = "t:Photo"
}

func (P *Photo) GetSchema() *Schema {
	return &SchemaTypes
}
