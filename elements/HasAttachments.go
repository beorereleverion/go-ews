package elements

// The HasAttachments element represents a property that is set to true if an item has at least one visible attachment or if a conversation contains at least one item that has an attachment. This property is read-only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/hasattachments
import "encoding/xml"

type HasAttachments struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (H *HasAttachments) SetForMarshal() {
	H.XMLName.Local = "t:HasAttachments"
}

func (H *HasAttachments) GetSchema() *Schema {
	return &SchemaTypes
}
