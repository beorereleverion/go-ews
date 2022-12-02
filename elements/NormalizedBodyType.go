package elements

// The NormalizedBodyType element specifies whether the normalized body is returned in text or HTML format.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/normalizedbodytype
type NormalizedBodyType string

const (
	// Best
	NormalizedBodyTypeBest NormalizedBodyType = `Best`
	// HTML
	NormalizedBodyTypeHTML NormalizedBodyType = `HTML`
	// Text
	NormalizedBodyTypeText NormalizedBodyType = `Text`
)
