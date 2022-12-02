package elements

// The BodyType element identifies how the body text is formatted in the response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/bodytype
type BodyType string

const (
	// The response will return the richest available content of body text. This is useful if it is unknown whether the content is text or HTML. The returned body will be text if the stored body is plain text. Otherwise, the response will return HTML if the stored body is in either HTML or RTF format. This is the default value.
	BodyTypeBest BodyType = `Best`
	// The response will return an item body as HTML.
	BodyTypeHTML BodyType = `HTML`
	// The response will return an item body as plain text.
	BodyTypeText BodyType = `Text`
)
