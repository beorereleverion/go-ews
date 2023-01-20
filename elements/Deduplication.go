package elements

// The Deduplication element indicates whether the search result should remove duplicate items.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deduplication
import "encoding/xml"

type Deduplication struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	Deduplicationfalse bool = false
	// true
	Deduplicationtrue bool = true
)

func (D *Deduplication) SetForMarshal() {
	D.XMLName.Local = "t:Deduplication"
}

func (D *Deduplication) GetSchema() *Schema {
	return &SchemaTypes
}
