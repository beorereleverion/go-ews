package elements

// The HasBeenProcessed element indicates whether a meeting message item has been processed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/hasbeenprocessed
import "encoding/xml"

type HasBeenProcessed struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (H *HasBeenProcessed) SetForMarshal() {
	H.XMLName.Local = "t:HasBeenProcessed"
}

func (H *HasBeenProcessed) GetSchema() *Schema {
	return &SchemaTypes
}
