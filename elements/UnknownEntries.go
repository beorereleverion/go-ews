package elements

// The UnknownEntries element contains an array of unknown permission entries that cannot be resolved against the Active Directory directory service. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/unknownentries
import "encoding/xml"

type UnknownEntries struct {
	XMLName xml.Name

	// The UnknownEntry element represents a single unknown permission entry that cannot be resolved against the Active Directory directory service. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	UnknownEntry *UnknownEntry `xml:"UnknownEntry"`
}

func (U *UnknownEntries) SetForMarshal() {
	U.XMLName.Local = "t:UnknownEntries"
}

func (U *UnknownEntries) GetSchema() *Schema {
	return &SchemaTypes
}
