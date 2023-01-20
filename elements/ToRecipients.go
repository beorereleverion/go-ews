package elements

// The ToRecipients element contains an array of recipients of an item. These are the primary recipients of an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/torecipients
import "encoding/xml"

type ToRecipients struct {
	XMLName xml.Name

	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"Mailbox"`
}

func (T *ToRecipients) SetForMarshal() {
	T.XMLName.Local = "t:ToRecipients"
}

func (T *ToRecipients) GetSchema() *Schema {
	return &SchemaTypes
}
