package elements

// The HasAttachment element specifies a Boolean value to indicate whether the item has attachments.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/hasattachment
import "encoding/xml"

type HasAttachment struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	HasAttachmentfalse bool = false
	// true
	HasAttachmenttrue bool = true
)

func (H *HasAttachment) SetForMarshal() {
	H.XMLName.Local = "t:HasAttachment"
}

func (H *HasAttachment) GetSchema() *Schema {
	return &SchemaTypes
}
