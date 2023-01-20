package elements

// The String element represents a string that is used by items, contacts, tasks, and conversations.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/string
import "encoding/xml"

type String struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *String) SetForMarshal() {
	S.XMLName.Local = "t:String"
}

func (S *String) GetSchema() *Schema {
	return &SchemaTypes
}
