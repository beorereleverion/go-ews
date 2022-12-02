package elements

// The ProgrammaticAccessAllowed element specifies whether programmatic access is enabled for rights managed data.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/programmaticaccessallowed
type ProgrammaticAccessAllowed bool

const (
	// false
	ProgrammaticAccessAllowedfalse ProgrammaticAccessAllowed = false
	// true
	ProgrammaticAccessAllowedtrue ProgrammaticAccessAllowed = true
)
