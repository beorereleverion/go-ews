package elements

// The ExportAllowed element specifies whether exporting is enabled.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/exportallowed
import "encoding/xml"

type ExportAllowed struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	ExportAllowedfalse bool = false
	// true
	ExportAllowedtrue bool = true
)

func (E *ExportAllowed) SetForMarshal() {
	E.XMLName.Local = "t:ExportAllowed"
}

func (E *ExportAllowed) GetSchema() *Schema {
	return &SchemaTypes
}
