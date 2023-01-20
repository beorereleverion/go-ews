package elements

// The ApprovalDecision element specifies the decision made on an approval request message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/approvaldecision
import "encoding/xml"

type ApprovalDecision struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

const (
	// 1
	ApprovalDecisionOne int64 = 1
	// 2
	ApprovalDecisionTwo int64 = 2
)

func (A *ApprovalDecision) SetForMarshal() {
	A.XMLName.Local = "t:ApprovalDecision"
}

func (A *ApprovalDecision) GetSchema() *Schema {
	return &SchemaTypes
}
