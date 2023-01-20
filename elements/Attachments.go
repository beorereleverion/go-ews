package elements

// The Attachments element contains the items or files that are attached to an item in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attachments-ex15websvcsotherref
import "encoding/xml"

type Attachments struct {
	XMLName xml.Name

	// The FileAttachment element represents a file that is attached to an item in the Exchange store.
	FileAttachment *FileAttachment `xml:"FileAttachment"`
	// The ItemAttachment element represents an Exchange item that is attached to another Exchange item.
	ItemAttachment *ItemAttachment `xml:"ItemAttachment"`
}

func (A *Attachments) SetForMarshal() {
	A.XMLName.Local = "t:Attachments"
}

func (A *Attachments) GetSchema() *Schema {
	return &SchemaTypes
}
