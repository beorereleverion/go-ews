package elements

// The GetConversationItemsResponseMessage element specifies the response message for a GetConversationItems request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getconversationitemsresponsemessage
type GetConversationItemsResponseMessage struct {
	// Conversation (ConversationResponseType) represents a single conversation returned in a GetConversationItems response.
	Conversation *ConversationConversationResponseType `xml:"Conversation"`
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"m:MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"m:MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"m:ResponseCode"`
	// Indicates the class of the response.
	ResponseClass *string `xml:"ResponseClass,attr"`
}
