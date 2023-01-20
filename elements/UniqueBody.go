package elements

// The UniqueBody element represents an HTML fragment or plain text which represents the unique body of this conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/uniquebody
import "encoding/xml"

type UniqueBody struct {
	XMLName xml.Name

	// Describes how the item body is stored in the item.
	BodyType *string `xml:"BodyType,attr"`
}

func (U *UniqueBody) SetForMarshal() {
	U.XMLName.Local = "t:UniqueBody"
}

func (U *UniqueBody) GetSchema() *Schema {
	return &SchemaTypes
}
