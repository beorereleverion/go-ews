package elements

// The SearchDumpster element specifies whether to search in the Exchange Dumpster.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchdumpster
import "encoding/xml"

type SearchDumpster struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// false
	SearchDumpsterfalse string = `false`
	// true
	SearchDumpstertrue string = `true`
)

func (S *SearchDumpster) SetForMarshal() {
	S.XMLName.Local = "t:SearchDumpster"
}

func (S *SearchDumpster) GetSchema() *Schema {
	return &SchemaTypes
}
