package elements

// The GroupId element uniquely identifies a group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/groupid
import "encoding/xml"

type GroupId struct {
	XMLName xml.Name

	// The text value of the ChangeKey attribute is the change key of the group.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// The text value of the Id attribute is the identifier of the group.
	Id *string `xml:"Id,attr"`
}

func (G *GroupId) SetForMarshal() {
	G.XMLName.Local = "m:GroupId"
}

func (G *GroupId) GetSchema() *Schema {
	return &SchemaMessages
}
