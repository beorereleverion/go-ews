package elements

// The Value element specifies the value of a BodyContentAttributedValue element.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/value-bodycontenttype
import "encoding/xml"

type ValueBodyContentType struct {
	XMLName xml.Name

	// The BodyType element identifies how the body text is formatted in the response.
	BodyType *BodyType `xml:"BodyType"`
	// The Value element contains the value of an extended property.
	Value *Value `xml:"Value"`
}

func (V *ValueBodyContentType) SetForMarshal() {
	V.XMLName.Local = "t:Value"
}

func (V *ValueBodyContentType) GetSchema() *Schema {
	return &SchemaTypes
}
