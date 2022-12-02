package elements

// The ApprovalDecision element specifies the decision made on an approval request message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/approvaldecision
type ApprovalDecision int64

const (
	// 1
	ApprovalDecisionOne ApprovalDecision = 1
	// 2
	ApprovalDecisionTwo ApprovalDecision = 2
)
