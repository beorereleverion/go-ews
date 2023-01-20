package elements

// The UnknownEntry element represents a single unknown permission entry that cannot be resolved against the Active Directory directory service. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/unknownentry
import "encoding/xml"

type UnknownEntry struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (U *UnknownEntry) SetForMarshal() {
	U.XMLName.Local = "t:UnknownEntry"
}

func (U *UnknownEntry) GetSchema() *Schema {
	return &SchemaTypes
}
