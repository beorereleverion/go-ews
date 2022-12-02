package elements

// The BlockStatus element specifies the block status of an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/blockstatus
type BlockStatus bool

const (
	// false
	BlockStatusfalse BlockStatus = false
	// true
	BlockStatustrue BlockStatus = true
)
