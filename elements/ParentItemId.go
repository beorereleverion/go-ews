package elements

// The ParentItemId element identifies the parent item that links to an associated attachment.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/parentitemid
import "encoding/xml"

type ParentItemId struct {
	XMLName xml.Name

	// Identifies an unspecified version of an item that is identified by the Id attribute in the Exchange store. This is used to make sure that a current item is used when it is updated with an attachment. This value is a string. This attribute is optional.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// Identifies a single item in the Exchange store to associate with an attachment. This value is a string. This attribute is required.
	Id *string `xml:"Id,attr"`
}

func (P *ParentItemId) SetForMarshal() {
	P.XMLName.Local = "m:ParentItemId"
}

func (P *ParentItemId) GetSchema() *Schema {
	return &SchemaMessages
}
