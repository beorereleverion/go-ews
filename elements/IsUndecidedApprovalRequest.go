package elements

// The IsUndecidedApprovalRequest element specifies whether an approval request message has been acted on.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isundecidedapprovalrequest
import "encoding/xml"

type IsUndecidedApprovalRequest struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsUndecidedApprovalRequestfalse bool = false
	// true
	IsUndecidedApprovalRequesttrue bool = true
)

func (I *IsUndecidedApprovalRequest) SetForMarshal() {
	I.XMLName.Local = "t:IsUndecidedApprovalRequest"
}

func (I *IsUndecidedApprovalRequest) GetSchema() *Schema {
	return &SchemaTypes
}
