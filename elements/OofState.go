package elements

// The OofState element is used to get or set the user's Out of Office (OOF) state.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/oofstate
import "encoding/xml"

type OofState struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Disabled

	OofStateDisabled string = `Disabled`
	// Enabled

	OofStateEnabled string = `Enabled`
	// Scheduled

	OofStateScheduled string = `Scheduled`
)

func (O *OofState) SetForMarshal() {
	O.XMLName.Local = "t:OofState"
}

func (O *OofState) GetSchema() *Schema {
	return &SchemaTypes
}
