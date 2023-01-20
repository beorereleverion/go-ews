package elements

// The IncludePersonalArchive element specifies whether to include the personal archive in the search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/includepersonalarchive
import "encoding/xml"

type IncludePersonalArchive struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IncludePersonalArchivefalse bool = false
	// true
	IncludePersonalArchivetrue bool = true
)

func (I *IncludePersonalArchive) SetForMarshal() {
	I.XMLName.Local = "m:IncludePersonalArchive"
}

func (I *IncludePersonalArchive) GetSchema() *Schema {
	return &SchemaMessages
}
