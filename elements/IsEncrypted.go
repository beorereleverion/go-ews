package elements

// The IsEncrypted element indicates whether incoming messages must be S/MIME encrypted in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isencrypted
type IsEncrypted bool

const (
	// false
	IsEncryptedfalse IsEncrypted = false
	// true
	IsEncryptedtrue IsEncrypted = true
)
