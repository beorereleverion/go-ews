package elements

// The ReadFlag element indicates the read state to set on items in a folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/readflag
type ReadFlag bool

const (
	// false
	ReadFlagfalse ReadFlag = false
	// true
	ReadFlagtrue ReadFlag = true
)
