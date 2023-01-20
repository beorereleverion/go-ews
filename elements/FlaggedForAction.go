package elements

// The FlaggedForAction element specifies the flag for action value that must appear on incoming messages in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/flaggedforaction
import "encoding/xml"

type FlaggedForAction struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (F *FlaggedForAction) SetForMarshal() {
	F.XMLName.Local = "m:FlaggedForAction"
}

func (F *FlaggedForAction) GetSchema() *Schema {
	return &SchemaMessages
}
