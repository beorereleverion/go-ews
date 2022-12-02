package elements

// The GetClientAccessTokenResponseMessage element specifies the response message for a GetClientAccessToken request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getclientaccesstokenresponsemessage
type GetClientAccessTokenResponseMessage struct {
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"m:MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"m:MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"m:ResponseCode"`
	// The Token element specifies a client access token.
	Token *TokenClientAccessTokenType `xml:"t:Token"`
	// Indicates the class of the response.
	ResponseClass *string `xml:"ResponseClass,attr"`
}
