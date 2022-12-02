package elements

// The IsPermissionControlled element indicates whether incoming messages must be permission controlled (RMS protected) in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ispermissioncontrolled
type IsPermissionControlled bool

const (
	// false
	IsPermissionControlledfalse IsPermissionControlled = false
	// true
	IsPermissionControlledtrue IsPermissionControlled = true
)
