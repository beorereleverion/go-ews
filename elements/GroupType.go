package elements

// The GroupType element specifies the group class of an instant messaging (IM) group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/grouptype
import "encoding/xml"

type GroupType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (G *GroupType) SetForMarshal() {
	G.XMLName.Local = "t:GroupType"
}

func (G *GroupType) GetSchema() *Schema {
	return &SchemaTypes
}
