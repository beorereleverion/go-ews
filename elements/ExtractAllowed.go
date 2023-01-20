package elements

// The ExtractAllowed element specifies whether entity extraction is enabled.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/extractallowed
import "encoding/xml"

type ExtractAllowed struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	ExtractAllowedfalse bool = false
	// true
	ExtractAllowedtrue bool = true
)

func (E *ExtractAllowed) SetForMarshal() {
	E.XMLName.Local = "t:ExtractAllowed"
}

func (E *ExtractAllowed) GetSchema() *Schema {
	return &SchemaTypes
}
