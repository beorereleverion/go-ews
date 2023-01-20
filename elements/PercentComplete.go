package elements

// The PercentComplete element describes the completion status of a task.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/percentcomplete
import "encoding/xml"

type PercentComplete struct {
	XMLName xml.Name
	TEXT    float64 `xml:",chardata"`
}

func (P *PercentComplete) SetForMarshal() {
	P.XMLName.Local = "t:PercentComplete"
}

func (P *PercentComplete) GetSchema() *Schema {
	return &SchemaTypes
}
