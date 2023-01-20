package elements

// The UnpinTeamMailboxResponse element contains a response to a request to unpin a site mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/unpinteammailboxresponse
import "encoding/xml"

type UnpinTeamMailboxResponse struct {
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

func (U *UnpinTeamMailboxResponse) SetForMarshal() {
	U.XMLName.Local = "m:UnpinTeamMailboxResponse"
}

func (U *UnpinTeamMailboxResponse) GetSchema() *Schema {
	return &SchemaMessages
}
