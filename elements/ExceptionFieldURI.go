package elements

// The ExceptionFieldURI element identifies particular errors in a request. This element is only used as part of an error response in the MessageXml node.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/exceptionfielduri
type ExceptionFieldURI struct {
	// Identifies a property of an occurrence of a recurring item. This attribute is required.
	FieldURI *string `xml:"FieldURI,attr"`
}
