package elements

// The Sender element specifies the e-mail address of the person who sent an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sender-string
import "encoding/xml"

type Senderstring struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *Senderstring) SetForMarshal() {
	S.XMLName.Local = "t:Sender"
}

func (S *Senderstring) GetSchema() *Schema {
	return &SchemaTypes
}
