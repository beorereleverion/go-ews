package elements

// The BodyType element specifies the type of the body of the item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/bodytype-bodytypetype
type BodyTypeBodyTypeType string

const (
	// Indicates that the body is HTML.
	BodyTypeBodyTypeTypeHTML BodyTypeBodyTypeType = `HTML`
	// Indicates that the body is text.
	BodyTypeBodyTypeTypeText BodyTypeBodyTypeType = `Text`
)
