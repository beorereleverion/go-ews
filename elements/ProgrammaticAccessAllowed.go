package elements

// The ProgrammaticAccessAllowed element specifies whether programmatic access is enabled for rights managed data.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/programmaticaccessallowed
import "encoding/xml"

type ProgrammaticAccessAllowed struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	ProgrammaticAccessAllowedfalse bool = false
	// true
	ProgrammaticAccessAllowedtrue bool = true
)

func (P *ProgrammaticAccessAllowed) SetForMarshal() {
	P.XMLName.Local = "t:ProgrammaticAccessAllowed"
}

func (P *ProgrammaticAccessAllowed) GetSchema() *Schema {
	return &SchemaTypes
}
