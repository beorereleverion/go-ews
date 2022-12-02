package elements

// The UniqueBodyType element specifies whether the unique body is returned in text or HTML format.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/uniquebodytype
type UniqueBodyType string

const (
	// Best
	UniqueBodyTypeBest UniqueBodyType = `Best`
	// HTML
	UniqueBodyTypeHTML UniqueBodyType = `HTML`
	// Text
	UniqueBodyTypeText UniqueBodyType = `Text`
)
