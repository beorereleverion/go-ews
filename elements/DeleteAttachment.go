package elements

// The DeleteAttachment element is the root element in a request to delete an attachment from the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteattachment
import "encoding/xml"

type DeleteAttachment struct {
	XMLName xml.Name

	// The AttachmentIds element contains an array of attachment identifiers.
	AttachmentIds *AttachmentIds `xml:"AttachmentIds"`
}

func (D *DeleteAttachment) SetForMarshal() {
	D.XMLName.Local = "m:DeleteAttachment"
}

func (D *DeleteAttachment) GetSchema() *Schema {
	return &SchemaMessages
}
