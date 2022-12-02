package elements

// The IsPermanentFailure element indicates whether a previous attempt to index the item was unsuccessful.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ispermanentfailure
type IsPermanentFailure bool

const (
	// false
	IsPermanentFailurefalse IsPermanentFailure = false
	// true
	IsPermanentFailuretrue IsPermanentFailure = true
)
