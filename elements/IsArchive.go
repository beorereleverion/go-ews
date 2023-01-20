package elements

// The IsArchive element specifies a Boolean value that indicates whether the mailbox is an archive mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isarchive
import "encoding/xml"

type IsArchive struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsArchivefalse bool = false
	// true
	IsArchivetrue bool = true
)

func (I *IsArchive) SetForMarshal() {
	I.XMLName.Local = "t:IsArchive"
}

func (I *IsArchive) GetSchema() *Schema {
	return &SchemaTypes
}
