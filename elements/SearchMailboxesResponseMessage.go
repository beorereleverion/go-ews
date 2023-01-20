package elements

// The SearchMailboxesResponseMessage element specifies the response message for a SearchMailboxes request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchmailboxesresponsemessage
import "encoding/xml"

type SearchMailboxesResponseMessage struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// The SearchMailboxesResult element contains the result of the SearchMailboxes request.
	SearchMailboxesResult *SearchMailboxesResult `xml:"SearchMailboxesResult"`
}

func (S *SearchMailboxesResponseMessage) SetForMarshal() {
	S.XMLName.Local = "m:SearchMailboxesResponseMessage"
}

func (S *SearchMailboxesResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
