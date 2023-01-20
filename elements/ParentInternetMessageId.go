package elements

// The ParentInternetMessageId element specifies the Internet message identifier of the parent message in a conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/parentinternetmessageid
import "encoding/xml"

type ParentInternetMessageId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *ParentInternetMessageId) SetForMarshal() {
	P.XMLName.Local = "t:ParentInternetMessageId"
}

func (P *ParentInternetMessageId) GetSchema() *Schema {
	return &SchemaTypes
}
