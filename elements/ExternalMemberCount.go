package elements

// The ExternalMemberCount element represents the count of external members in a group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/externalmembercount
import "encoding/xml"

type ExternalMemberCount struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (E *ExternalMemberCount) SetForMarshal() {
	E.XMLName.Local = "t:ExternalMemberCount"
}

func (E *ExternalMemberCount) GetSchema() *Schema {
	return &SchemaTypes
}
