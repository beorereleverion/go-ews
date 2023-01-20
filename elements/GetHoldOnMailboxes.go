package elements

// The GetHoldOnMailboxes element indicates the beginning of the GetHoldOnMailboxes request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getholdonmailboxes
import "encoding/xml"

type GetHoldOnMailboxes struct {
	XMLName xml.Name

	// The HoldId element contains the mailbox hold identifier.
	HoldId *HoldId `xml:"HoldId"`
}

func (G *GetHoldOnMailboxes) SetForMarshal() {
	G.XMLName.Local = "m:GetHoldOnMailboxes"
}

func (G *GetHoldOnMailboxes) GetSchema() *Schema {
	return &SchemaMessages
}
