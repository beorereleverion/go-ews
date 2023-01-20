package elements

// The GlobalHasIrm element specifies whether at least one message in the conversation and across all folders is an IRM protected message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globalhasirm
import "encoding/xml"

type GlobalHasIrm struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	GlobalHasIrmfalse bool = false
	// true
	GlobalHasIrmtrue bool = true
)

func (G *GlobalHasIrm) SetForMarshal() {
	G.XMLName.Local = "t:GlobalHasIrm"
}

func (G *GlobalHasIrm) GetSchema() *Schema {
	return &SchemaTypes
}
