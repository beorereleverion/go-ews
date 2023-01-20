package elements

// The TextBody element specifies the text body.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/textbody
import "encoding/xml"

type TextBody struct {
	XMLName xml.Name

	// Indicates the body type. The value of Text for the BodyTypeType attribute indicates that the body is in plain text form. The value of HTML for the BodyTypeType attribute indicates that the body is in HTML form. The BodyTypeType attribute is required.
	BodyTypeType *string `xml:"BodyTypeType,attr"`
	// Indicates that the body contents have been truncated. A text value of false for the IsTruncated attribute indicates that the body contents have not been truncated. The normalized body will be truncated if the text body length is longer than the value set in the MaximumBodySize element.
	IsTruncated *string `xml:"IsTruncated,attr"`
}

func (T *TextBody) SetForMarshal() {
	T.XMLName.Local = "t:TextBody"
}

func (T *TextBody) GetSchema() *Schema {
	return &SchemaTypes
}
