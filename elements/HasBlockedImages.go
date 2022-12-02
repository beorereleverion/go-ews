package elements

// The HasBlockedImages element specifies a Boolean value that indicates whether the item has blocked images.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/hasblockedimages
type HasBlockedImages bool

const (
	// false
	HasBlockedImagesfalse HasBlockedImages = false
	// true
	HasBlockedImagestrue HasBlockedImages = true
)
