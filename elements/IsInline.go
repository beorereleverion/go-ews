package elements

// The IsInline element represents whether the attachment appears inline within an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isinline
type IsInline bool

const (
	// false
	IsInlinefalse IsInline = false
	// true
	IsInlinetrue IsInline = true
)
