package elements

// The ApprovalRequestData element specifies the approval state of an approval request message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/approvalrequestdata
type ApprovalRequestData struct {
	// The ApprovalDecision element specifies the decision made on an approval request message.
	ApprovalDecision *ApprovalDecision `xml:"t:ApprovalDecision"`
	// The ApprovalDecisionMaker element specifies the display name of the person who made the approval decision.
	ApprovalDecisionMaker *ApprovalDecisionMaker `xml:"t:ApprovalDecisionMaker"`
	// The ApprovalDecisionTime element specifies the time at which the approval decision was made.
	ApprovalDecisionTime *ApprovalDecisionTime `xml:"t:ApprovalDecisionTime"`
	// The IsUndecidedApprovalRequest element specifies whether an approval request message has been acted on.
	IsUndecidedApprovalRequest *IsUndecidedApprovalRequest `xml:"t:IsUndecidedApprovalRequest"`
}
