package elements

// The Diagnostics element provides timing and performance information that is used for reporting in a DataCenter.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/diagnostics
import "encoding/xml"

type Diagnostics struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (D *Diagnostics) SetForMarshal() {
	D.XMLName.Local = "m:Diagnostics"
}

func (D *Diagnostics) GetSchema() *Schema {
	return &SchemaMessages
}
