package elements

// The ReplyAllAllowed element specifies whether a reply all is allowed for rights managed data.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/replyallallowed
type ReplyAllAllowed bool

const (
	// false
	ReplyAllAllowedfalse ReplyAllAllowed = false
	// true
	ReplyAllAllowedtrue ReplyAllAllowed = true
)
