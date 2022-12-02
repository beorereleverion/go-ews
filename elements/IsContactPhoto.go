package elements

// The IsContactPhoto element indicates whether the file attachment is a contact picture.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/iscontactphoto
type IsContactPhoto bool

const (
	// false
	IsContactPhotofalse IsContactPhoto = false
	// true
	IsContactPhototrue IsContactPhoto = true
)
