package elements

// The SearchScope element specifies the scope of a search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchscope
import "encoding/xml"

type SearchScope struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// All
	SearchScopeAll string = `All`
	// ArchiveOnly
	SearchScopeArchiveOnly string = `ArchiveOnly`
	// PrimaryOnly
	SearchScopePrimaryOnly string = `PrimaryOnly`
)

func (S *SearchScope) SetForMarshal() {
	S.XMLName.Local = "t:SearchScope"
}

func (S *SearchScope) GetSchema() *Schema {
	return &SchemaTypes
}
