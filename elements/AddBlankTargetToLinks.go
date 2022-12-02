package elements

// The AddBlankTargetToLinks element specifies that the target attribute in HTML links are set to open a new window.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/addblanktargettolinks
type AddBlankTargetToLinks bool

const (
	// false
	AddBlankTargetToLinksfalse AddBlankTargetToLinks = false
	// true
	AddBlankTargetToLinkstrue AddBlankTargetToLinks = true
)
