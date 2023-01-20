package elements

// The Status element represents the status of a task item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/status
import "encoding/xml"

type Status struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *Status) SetForMarshal() {
	S.XMLName.Local = "t:Status"
}

func (S *Status) GetSchema() *Schema {
	return &SchemaTypes
}
