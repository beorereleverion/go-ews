package elements

// The SetTeamMailboxResponse element represents a response to a SetTeamMailbox request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setteammailboxresponse
import "encoding/xml"

type SetTeamMailboxResponse struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
}

func (S *SetTeamMailboxResponse) SetForMarshal() {
	S.XMLName.Local = "m:SetTeamMailboxResponse"
}

func (S *SetTeamMailboxResponse) GetSchema() *Schema {
	return &SchemaMessages
}
