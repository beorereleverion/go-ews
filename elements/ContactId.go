package elements

// The ContactId element uniquely identifies a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/contactid
import "encoding/xml"

type ContactId struct {
	XMLName xml.Name

	// The text value of the ChangeKey attribute is the change key of the contact item.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// The text value of the Id attribute is the identifier of the contact item.
	Id *string `xml:"Id,attr"`
}

func (C *ContactId) SetForMarshal() {
	C.XMLName.Local = "t:ContactId"
}

func (C *ContactId) GetSchema() *Schema {
	return &SchemaTypes
}
