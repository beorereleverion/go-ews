package elements

// The SearchArchiveOnly element indicates whether only the archive mailbox is searched for non-indexable items.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searcharchiveonly
import "encoding/xml"

type SearchArchiveOnly struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	SearchArchiveOnlyfalse bool = false
	// true
	SearchArchiveOnlytrue bool = true
)

func (S *SearchArchiveOnly) SetForMarshal() {
	S.XMLName.Local = "m:SearchArchiveOnly"
}

func (S *SearchArchiveOnly) GetSchema() *Schema {
	return &SchemaMessages
}
