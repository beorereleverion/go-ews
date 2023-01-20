package elements

// The NormalizedBody element specifies an HTML representation of the Body property of an item as a fragment that can be inserted into another HTML body.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/normalizedbody
import "encoding/xml"

type NormalizedBody struct {
	XMLName xml.Name

	// Indicates the body type. The value of Text for the BodyType attribute indicates that the body is in plain text form. The value of HTML for the BodyType attribute indicates that the body is in HTML form. The BodyType attribute is required.
	BodyType *string `xml:"BodyType,attr"`
	// Indicates that the body contents have been truncated. A text value of false for the IsTruncated attribute indicates that the body contents have not been truncated. The normalized body will be truncated if the normalized body length is longer than the value set in the MaximumBodySize element.
	IsTruncated *string `xml:"IsTruncated,attr"`
}

func (N *NormalizedBody) SetForMarshal() {
	N.XMLName.Local = "t:NormalizedBody"
}

func (N *NormalizedBody) GetSchema() *Schema {
	return &SchemaTypes
}
