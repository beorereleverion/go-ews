package elements

// The ActualWork element represents the actual amount of time that is spent on a task.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/actualwork
import "encoding/xml"

type ActualWork struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (A *ActualWork) SetForMarshal() {
	A.XMLName.Local = "t:ActualWork"
}

func (A *ActualWork) GetSchema() *Schema {
	return &SchemaTypes
}
