package elements

// The GetSearchableMailboxesResponseMessage element specifies the response message for a GetSearchableMailboxes request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getsearchablemailboxesresponsemessage
import "encoding/xml"

type GetSearchableMailboxesResponseMessage struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// The SearchableMailboxes element contains an array of the mailboxes returned from a GetSearchableMailboxes request.
	SearchableMailboxes *SearchableMailboxes `xml:"SearchableMailboxes"`
	// Indicates the class of the response.
	ResponseClass *string `xml:"ResponseClass,attr"`
}

func (G *GetSearchableMailboxesResponseMessage) SetForMarshal() {
	G.XMLName.Local = "m:GetSearchableMailboxesResponseMessage"
}

func (G *GetSearchableMailboxesResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
