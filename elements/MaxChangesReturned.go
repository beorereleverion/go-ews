package elements

// The MaxChangesReturned element describes the maximum number of changes that can be returned in a synchronization response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/maxchangesreturned
import "encoding/xml"

type MaxChangesReturned struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (M *MaxChangesReturned) SetForMarshal() {
	M.XMLName.Local = "t:MaxChangesReturned"
}

func (M *MaxChangesReturned) GetSchema() *Schema {
	return &SchemaTypes
}
