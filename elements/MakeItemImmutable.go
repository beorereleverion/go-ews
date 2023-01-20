package elements

// The MakeItemImmutable element specifies a Boolean value that indicates whether an item should be made read-only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/makeitemimmutable
import "encoding/xml"

type MakeItemImmutable struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	MakeItemImmutablefalse bool = false
	// true
	MakeItemImmutabletrue bool = true
)

func (M *MakeItemImmutable) SetForMarshal() {
	M.XMLName.Local = "m:MakeItemImmutable"
}

func (M *MakeItemImmutable) GetSchema() *Schema {
	return &SchemaMessages
}
