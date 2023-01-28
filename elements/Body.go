package elements

// The Body element specifies the body of an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/body
import "encoding/xml"

type Body struct {
	XMLName xml.Name

	// Specifies the type of the body.
	BodyType *string `xml:"BodyType,attr"`
	// Boolean value that indicates whether the body is truncated.
	IsTruncated *string `xml:"IsTruncated,attr"`
	TEXT    string `xml:",chardata"`
}

func (B *Body) SetForMarshal() {
	B.XMLName.Local = "t:Body"
}

func (B *Body) GetSchema() *Schema {
	return &SchemaTypes
}
