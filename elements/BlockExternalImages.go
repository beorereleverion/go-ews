package elements

// The BlockExternalImages element specifies whether external images are blocked in HTML text bodies.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/blockexternalimages
import "encoding/xml"

type BlockExternalImages struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	BlockExternalImagesfalse bool = false
	// true
	BlockExternalImagestrue bool = true
)

func (B *BlockExternalImages) SetForMarshal() {
	B.XMLName.Local = "t:BlockExternalImages"
}

func (B *BlockExternalImages) GetSchema() *Schema {
	return &SchemaTypes
}
