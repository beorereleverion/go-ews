package elements

// The MustDisplayComment element indicates whether the managed folder comment must be displayed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mustdisplaycomment
import "encoding/xml"

type MustDisplayComment struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (M *MustDisplayComment) SetForMarshal() {
	M.XMLName.Local = "t:MustDisplayComment"
}

func (M *MustDisplayComment) GetSchema() *Schema {
	return &SchemaTypes
}
