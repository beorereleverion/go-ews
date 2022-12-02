package elements

// The MimeContentUTF8 element contains the UTF-8 MIME stream of an object that is represented in base64Binary format and supports email address internationalization and [RFC6530].
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mimecontentutf8
type MimeContentUTF8 struct {
	// If set, the value for this attribute is ignored by the server.
	CharacterSet *string `xml:"CharacterSet,attr"`
}
