package elements

// The NewBodyContent element represents the new body content of a message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/newbodycontent
import "encoding/xml"

type NewBodyContent struct {
	XMLName xml.Name

	// Represents the actual body content of a message.
	BodyType *string `xml:"BodyType,attr"`
}

func (N *NewBodyContent) SetForMarshal() {
	N.XMLName.Local = "t:NewBodyContent"
}

func (N *NewBodyContent) GetSchema() *Schema {
	return &SchemaTypes
}
