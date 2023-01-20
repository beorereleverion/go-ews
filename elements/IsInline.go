package elements

// The IsInline element represents whether the attachment appears inline within an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isinline
import "encoding/xml"

type IsInline struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsInlinefalse bool = false
	// true
	IsInlinetrue bool = true
)

func (I *IsInline) SetForMarshal() {
	I.XMLName.Local = "t:IsInline"
}

func (I *IsInline) GetSchema() *Schema {
	return &SchemaTypes
}
