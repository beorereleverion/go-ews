package elements

// The NumberOfMembersWithNoData element represents the number of distribution list members who do not have published free/busy data to compare to a suggested meeting time. This element represents members of a distribution list that is too large or members who have No Data status.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/numberofmemberswithnodata
import "encoding/xml"

type NumberOfMembersWithNoData struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (N *NumberOfMembersWithNoData) SetForMarshal() {
	N.XMLName.Local = "t:NumberOfMembersWithNoData"
}

func (N *NumberOfMembersWithNoData) GetSchema() *Schema {
	return &SchemaTypes
}
