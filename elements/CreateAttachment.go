package elements

// The CreateAttachment element defines a request to create an attachment to an item in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createattachment
import "encoding/xml"

type CreateAttachment struct {
	XMLName xml.Name

	// The Attachments element contains the items or files that are attached to an item in the Exchange store.
	Attachments *Attachments `xml:"Attachments"`
	// The ParentItemId element identifies the parent item that links to an associated attachment.
	ParentItemId *ParentItemId `xml:"ParentItemId"`
}

func (C *CreateAttachment) SetForMarshal() {
	C.XMLName.Local = "m:CreateAttachment"
}

func (C *CreateAttachment) GetSchema() *Schema {
	return &SchemaMessages
}
