package elements

// The EnableAlwaysDelete element specifies a flag that enables deletion for all new items in a conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/enablealwaysdelete
import "encoding/xml"

type EnableAlwaysDelete struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (E *EnableAlwaysDelete) SetForMarshal() {
	E.XMLName.Local = "t:EnableAlwaysDelete"
}

func (E *EnableAlwaysDelete) GetSchema() *Schema {
	return &SchemaTypes
}
