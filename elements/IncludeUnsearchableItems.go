package elements

// The IncludeUnsearchableItems element specifies whether to include items that cannot be searched.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/includeunsearchableitems
import "encoding/xml"

type IncludeUnsearchableItems struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IncludeUnsearchableItemsfalse bool = false
	// true
	IncludeUnsearchableItemstrue bool = true
)

func (I *IncludeUnsearchableItems) SetForMarshal() {
	I.XMLName.Local = "m:IncludeUnsearchableItems"
}

func (I *IncludeUnsearchableItems) GetSchema() *Schema {
	return &SchemaMessages
}
