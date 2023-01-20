package elements

// The SetHoldOnMailboxesResponse element represents a response to a SetHoldOnMailboxes request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setholdonmailboxesresponse
import "encoding/xml"

type SetHoldOnMailboxesResponse struct {
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
}

func (S *SetHoldOnMailboxesResponse) SetForMarshal() {
	S.XMLName.Local = "m:SetHoldOnMailboxesResponse"
}

func (S *SetHoldOnMailboxesResponse) GetSchema() *Schema {
	return &SchemaMessages
}
