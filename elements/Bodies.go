package elements

// The Bodies element specifies an array of BodyContentAttributedValue elements.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/bodies
import "encoding/xml"

type Bodies struct {
	XMLName xml.Name

	// The BodyContentAttributedValue element specifies the body content of an item.
	BodyContentAttributedValue *BodyContentAttributedValue `xml:"BodyContentAttributedValue"`
}

func (B *Bodies) SetForMarshal() {
	B.XMLName.Local = "t:Bodies"
}

func (B *Bodies) GetSchema() *Schema {
	return &SchemaTypes
}
