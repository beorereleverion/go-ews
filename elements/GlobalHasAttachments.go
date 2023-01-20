package elements

// The GlobalHasAttachments element contains a value that indicates whether at least one conversation item in a mailbox has an attachment.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globalhasattachments
import "encoding/xml"

type GlobalHasAttachments struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (G *GlobalHasAttachments) SetForMarshal() {
	G.XMLName.Local = "t:GlobalHasAttachments"
}

func (G *GlobalHasAttachments) GetSchema() *Schema {
	return &SchemaTypes
}
