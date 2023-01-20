package elements

// The AddBlankTargetToLinks element specifies that the target attribute in HTML links are set to open a new window.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/addblanktargettolinks
import "encoding/xml"

type AddBlankTargetToLinks struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	AddBlankTargetToLinksfalse bool = false
	// true
	AddBlankTargetToLinkstrue bool = true
)

func (A *AddBlankTargetToLinks) SetForMarshal() {
	A.XMLName.Local = "t:AddBlankTargetToLinks"
}

func (A *AddBlankTargetToLinks) GetSchema() *Schema {
	return &SchemaTypes
}
