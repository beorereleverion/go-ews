package elements

// The HasPicture element indicates whether the contact item has a file attachment that represents the contact's picture.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/haspicture
type HasPicture bool

const (
	// false
	HasPicturefalse HasPicture = false
	// true
	HasPicturetrue HasPicture = true
)
