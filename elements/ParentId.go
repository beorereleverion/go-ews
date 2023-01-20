package elements

// The ParentId element specifies the identifier of the parent item in a search preview.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/parentid
import "encoding/xml"

type ParentId struct {
	XMLName xml.Name

	// The text value of the ChangeKey attribute is the change key of the parent item.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// The text value of the Id attribute is the identifier of the parent item.
	Id *string `xml:"Id,attr"`
}

func (P *ParentId) SetForMarshal() {
	P.XMLName.Local = "t:ParentId"
}

func (P *ParentId) GetSchema() *Schema {
	return &SchemaTypes
}
