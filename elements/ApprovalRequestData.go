package elements

// The ApprovalRequestData element specifies the approval state of an approval request message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/approvalrequestdata
import "encoding/xml"

type ApprovalRequestData struct {
	XMLName xml.Name

	// The ApprovalDecision element specifies the decision made on an approval request message.
	ApprovalDecision *ApprovalDecision `xml:"ApprovalDecision"`
	// The ApprovalDecisionMaker element specifies the display name of the person who made the approval decision.
	ApprovalDecisionMaker *ApprovalDecisionMaker `xml:"ApprovalDecisionMaker"`
	// The ApprovalDecisionTime element specifies the time at which the approval decision was made.
	ApprovalDecisionTime *ApprovalDecisionTime `xml:"ApprovalDecisionTime"`
	// The IsUndecidedApprovalRequest element specifies whether an approval request message has been acted on.
	IsUndecidedApprovalRequest *IsUndecidedApprovalRequest `xml:"IsUndecidedApprovalRequest"`
}

func (A *ApprovalRequestData) SetForMarshal() {
	A.XMLName.Local = "t:ApprovalRequestData"
}

func (A *ApprovalRequestData) GetSchema() *Schema {
	return &SchemaTypes
}
