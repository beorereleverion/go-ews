package elements

// The IsWritable element specifies whether the underlying contact or Active Directory recipient can be written to.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/iswritable
type IsWritable bool

const (
	// false
	IsWritablefalse IsWritable = false
	// true
	IsWritabletrue IsWritable = true
)
