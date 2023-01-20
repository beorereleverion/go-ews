package elements

// The AttemptCount element represents the number of attempts that have been made to index the item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attemptcount
import "encoding/xml"

type AttemptCount struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (A *AttemptCount) SetForMarshal() {
	A.XMLName.Local = "t:AttemptCount"
}

func (A *AttemptCount) GetSchema() *Schema {
	return &SchemaTypes
}
