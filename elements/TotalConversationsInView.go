package elements

// The TotalConversationsInView element contains the count of conversations returned in a FindConversation response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/totalconversationsinview
import "encoding/xml"

type TotalConversationsInView struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (T *TotalConversationsInView) SetForMarshal() {
	T.XMLName.Local = "m:TotalConversationsInView"
}

func (T *TotalConversationsInView) GetSchema() *Schema {
	return &SchemaMessages
}
