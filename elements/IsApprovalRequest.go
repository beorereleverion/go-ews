package elements

// The IsApprovalRequest element indicates whether incoming messages must be approval requests in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isapprovalrequest
import "encoding/xml"

type IsApprovalRequest struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsApprovalRequest) SetForMarshal() {
	I.XMLName.Local = "m:IsApprovalRequest"
}

func (I *IsApprovalRequest) GetSchema() *Schema {
	return &SchemaMessages
}
