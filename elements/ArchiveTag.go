package elements

// The ArchiveTag element specifies the retention identifier of the archive tag set on an item or folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/archivetag
import "encoding/xml"

type ArchiveTag struct {
	XMLName xml.Name

	// Specifies whether the retention policy is explicitly set on an item or folder or whether it is inherited from a parent folder.
	IsExplicit *string `xml:"IsExplicit,attr"`
}

func (A *ArchiveTag) SetForMarshal() {
	A.XMLName.Local = "t:ArchiveTag"
}

func (A *ArchiveTag) GetSchema() *Schema {
	return &SchemaTypes
}
