package elements

// The IncludeMimeContent element specifies whether the Multipurpose Internet Mail Extensions (MIME) content of an item or attachment is returned in the response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/includemimecontent
type IncludeMimeContent bool

const (
	// false
	IncludeMimeContentfalse IncludeMimeContent = false
	// true
	IncludeMimeContenttrue IncludeMimeContent = true
)
