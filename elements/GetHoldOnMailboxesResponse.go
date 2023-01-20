package elements

// The GetHoldOnMailboxes element contains the request to get the hold status for a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getholdonmailboxesresponse
import "encoding/xml"

type GetHoldOnMailboxesResponse struct {
	XMLName xml.Name

	// The HoldId element contains the mailbox hold identifier.
	HoldId *HoldId `xml:"HoldId"`
}

func (G *GetHoldOnMailboxesResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetHoldOnMailboxesResponse"
}

func (G *GetHoldOnMailboxesResponse) GetSchema() *Schema {
	return &SchemaMessages
}
