package elements

// The ExceptionFieldURI element identifies particular errors in a request. This element is only used as part of an error response in the MessageXml node.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/exceptionfielduri
import "encoding/xml"

type ExceptionFieldURI struct {
	XMLName xml.Name

	// Identifies a property of an occurrence of a recurring item. This attribute is required.
	FieldURI *string `xml:"FieldURI,attr"`
}

func (E *ExceptionFieldURI) SetForMarshal() {
	E.XMLName.Local = "t:ExceptionFieldURI"
}

func (E *ExceptionFieldURI) GetSchema() *Schema {
	return &SchemaTypes
}
