package elements

// The Name element represents the name of the attachment.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/name-attachmenttype
import "encoding/xml"

type NameAttachmentType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (N *NameAttachmentType) SetForMarshal() {
	N.XMLName.Local = "t:Name"
}

func (N *NameAttachmentType) GetSchema() *Schema {
	return &SchemaTypes
}
