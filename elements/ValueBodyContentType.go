package elements

// The Value element specifies the value of a BodyContentAttributedValue element.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/value-bodycontenttype
type ValueBodyContentType struct {
	// The BodyType element identifies how the body text is formatted in the response.
	BodyType *BodyType `xml:"t:BodyType"`
	// The Value element contains the value of an extended property.
	Value *Value `xml:"t:Value"`
}
