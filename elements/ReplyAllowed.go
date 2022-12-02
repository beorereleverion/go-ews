package elements

// The ReplyAllowed element specifies whether a reply is allowed for rights managed data.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/replyallowed
type ReplyAllowed bool

const (
	// false
	ReplyAllowedfalse ReplyAllowed = false
	// true
	ReplyAllowedtrue ReplyAllowed = true
)
