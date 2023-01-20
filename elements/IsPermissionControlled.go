package elements

// The IsPermissionControlled element indicates whether incoming messages must be permission controlled (RMS protected) in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ispermissioncontrolled
import "encoding/xml"

type IsPermissionControlled struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsPermissionControlledfalse bool = false
	// true
	IsPermissionControlledtrue bool = true
)

func (I *IsPermissionControlled) SetForMarshal() {
	I.XMLName.Local = "m:IsPermissionControlled"
}

func (I *IsPermissionControlled) GetSchema() *Schema {
	return &SchemaMessages
}
