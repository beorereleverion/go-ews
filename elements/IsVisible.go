package elements

// The IsVisible element indicates whether the retention policy is visible to users.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isvisible
type IsVisible bool

const (
	// false
	IsVisiblefalse IsVisible = false
	// true
	IsVisibletrue IsVisible = true
)
