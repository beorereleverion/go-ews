package elements

// The GetClientAccessTokenResponseMessage element specifies the response message for a GetClientAccessToken request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getclientaccesstokenresponsemessage
import "encoding/xml"

type GetClientAccessTokenResponseMessage struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// The Token element specifies a client access token.
	Token *TokenClientAccessTokenType `xml:"Token"`
	// Indicates the class of the response.
	ResponseClass *string `xml:"ResponseClass,attr"`
}

func (G *GetClientAccessTokenResponseMessage) SetForMarshal() {
	G.XMLName.Local = "m:GetClientAccessTokenResponseMessage"
}

func (G *GetClientAccessTokenResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
