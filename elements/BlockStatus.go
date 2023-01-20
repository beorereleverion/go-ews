package elements

// The BlockStatus element specifies the block status of an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/blockstatus
import "encoding/xml"

type BlockStatus struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	BlockStatusfalse bool = false
	// true
	BlockStatustrue bool = true
)

func (B *BlockStatus) SetForMarshal() {
	B.XMLName.Local = "t:BlockStatus"
}

func (B *BlockStatus) GetSchema() *Schema {
	return &SchemaTypes
}
