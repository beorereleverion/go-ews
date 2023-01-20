package elements

// The IsComplete element indicates whether the task has been completed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/iscomplete
import "encoding/xml"

type IsComplete struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsComplete) SetForMarshal() {
	I.XMLName.Local = "t:IsComplete"
}

func (I *IsComplete) GetSchema() *Schema {
	return &SchemaTypes
}
