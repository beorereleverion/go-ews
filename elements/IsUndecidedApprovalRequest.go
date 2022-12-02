package elements

// The IsUndecidedApprovalRequest element specifies whether an approval request message has been acted on.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isundecidedapprovalrequest
type IsUndecidedApprovalRequest bool

const (
	// false
	IsUndecidedApprovalRequestfalse IsUndecidedApprovalRequest = false
	// true
	IsUndecidedApprovalRequesttrue IsUndecidedApprovalRequest = true
)
