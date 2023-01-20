package elements

// The GetHoldOnMailboxesResponseMessage element specifies the response message for a GetHoldOnMailboxes request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getholdonmailboxesresponsemessage
import "encoding/xml"

type GetHoldOnMailboxesResponseMessage struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MailboxHoldResult element contains the result of the GetHoldOnMailboxes request.
	MailboxHoldResult *MailboxHoldResult `xml:"MailboxHoldResult"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// Indicates the class of the response.
	ResponseClass *string `xml:"ResponseClass,attr"`
}

func (G *GetHoldOnMailboxesResponseMessage) SetForMarshal() {
	G.XMLName.Local = "m:GetHoldOnMailboxesResponseMessage"
}

func (G *GetHoldOnMailboxesResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
