package elements

// The ExportAllowed element specifies whether exporting is enabled.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/exportallowed
type ExportAllowed bool

const (
	// false
	ExportAllowedfalse ExportAllowed = false
	// true
	ExportAllowedtrue ExportAllowed = true
)
