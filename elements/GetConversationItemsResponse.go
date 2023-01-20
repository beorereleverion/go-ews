package elements

// The GetConversationItemsResponse element defines a response to a GetConversationItems request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getconversationitemsresponse
import "encoding/xml"

type GetConversationItemsResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (G *GetConversationItemsResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetConversationItemsResponse"
}

func (G *GetConversationItemsResponse) GetSchema() *Schema {
	return &SchemaMessages
}
