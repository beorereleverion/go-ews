package elements

// The ApprovalDecisionMaker element specifies the display name of the person who made the approval decision.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/approvaldecisionmaker
import "encoding/xml"

type ApprovalDecisionMaker struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (A *ApprovalDecisionMaker) SetForMarshal() {
	A.XMLName.Local = "t:ApprovalDecisionMaker"
}

func (A *ApprovalDecisionMaker) GetSchema() *Schema {
	return &SchemaTypes
}
