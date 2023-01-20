package elements

// The TotalMemberCount element represents the count of all members in a group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/totalmembercount
import "encoding/xml"

type TotalMemberCount struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (T *TotalMemberCount) SetForMarshal() {
	T.XMLName.Local = "t:TotalMemberCount"
}

func (T *TotalMemberCount) GetSchema() *Schema {
	return &SchemaTypes
}
