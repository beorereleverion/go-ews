package elements

// The InternetMessageId element represents the Internet message identifier of an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/internetmessageid
import "encoding/xml"

type InternetMessageId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (I *InternetMessageId) SetForMarshal() {
	I.XMLName.Local = "t:InternetMessageId"
}

func (I *InternetMessageId) GetSchema() *Schema {
	return &SchemaTypes
}
