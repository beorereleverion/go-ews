package elements

// The FilterHtmlContent element specifies whether potentially unsafe HTML content is filtered from an item or attachment.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/filterhtmlcontent
type FilterHtmlContent bool

const (
	// false
	FilterHtmlContentfalse FilterHtmlContent = false
	// true
	FilterHtmlContenttrue FilterHtmlContent = true
)
