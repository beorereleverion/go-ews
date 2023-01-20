package elements

// The GetConversationItemsResponseMessage element specifies the response message for a GetConversationItems request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getconversationitemsresponsemessage
import "encoding/xml"

type GetConversationItemsResponseMessage struct {
	XMLName xml.Name

	// Conversation (ConversationResponseType) represents a single conversation returned in a GetConversationItems response.
	Conversation *ConversationConversationResponseType `xml:"Conversation"`
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// Indicates the class of the response.
	ResponseClass *string `xml:"ResponseClass,attr"`
}

func (G *GetConversationItemsResponseMessage) SetForMarshal() {
	G.XMLName.Local = "m:GetConversationItemsResponseMessage"
}

func (G *GetConversationItemsResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
