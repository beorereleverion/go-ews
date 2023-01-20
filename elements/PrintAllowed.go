package elements

// The PrintAllowed element specifies whether printing is enabled.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/printallowed
import "encoding/xml"

type PrintAllowed struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	PrintAllowedfalse bool = false
	// true
	PrintAllowedtrue bool = true
)

func (P *PrintAllowed) SetForMarshal() {
	P.XMLName.Local = "t:PrintAllowed"
}

func (P *PrintAllowed) GetSchema() *Schema {
	return &SchemaTypes
}
