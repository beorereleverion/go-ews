package elements

// The BlockExternalImages element specifies whether external images are blocked in HTML text bodies.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/blockexternalimages
type BlockExternalImages bool

const (
	// false
	BlockExternalImagesfalse BlockExternalImages = false
	// true
	BlockExternalImagestrue BlockExternalImages = true
)
