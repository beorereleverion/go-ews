package elements

// The ApplyConversationActionResponse element defines a response to an ApplyConversationAction operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/applyconversationactionresponse
import "encoding/xml"

type ApplyConversationActionResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (A *ApplyConversationActionResponse) SetForMarshal() {
	A.XMLName.Local = "m:ApplyConversationActionResponse"
}

func (A *ApplyConversationActionResponse) GetSchema() *Schema {
	return &SchemaMessages
}
