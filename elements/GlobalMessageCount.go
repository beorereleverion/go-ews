package elements

// The GlobalMessageCount element contains the total number of conversation items in the mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globalmessagecount
import "encoding/xml"

type GlobalMessageCount struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (G *GlobalMessageCount) SetForMarshal() {
	G.XMLName.Local = "t:GlobalMessageCount"
}

func (G *GlobalMessageCount) GetSchema() *Schema {
	return &SchemaTypes
}
