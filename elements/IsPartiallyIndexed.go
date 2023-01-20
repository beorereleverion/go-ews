package elements

// The IsPartiallyIndexed element indicates whether the item is partially indexed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ispartiallyindexed
import "encoding/xml"

type IsPartiallyIndexed struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsPartiallyIndexedfalse bool = false
	// true
	IsPartiallyIndexedtrue bool = true
)

func (I *IsPartiallyIndexed) SetForMarshal() {
	I.XMLName.Local = "t:IsPartiallyIndexed"
}

func (I *IsPartiallyIndexed) GetSchema() *Schema {
	return &SchemaTypes
}
