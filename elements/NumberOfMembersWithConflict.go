package elements

// The NumberOfMembersWithConflict element represents the number of distribution list members who have a conflict with a suggested meeting time. This element represents members who have a status of Busy, OOF, or Tentative.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/numberofmemberswithconflict
import "encoding/xml"

type NumberOfMembersWithConflict struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (N *NumberOfMembersWithConflict) SetForMarshal() {
	N.XMLName.Local = "t:NumberOfMembersWithConflict"
}

func (N *NumberOfMembersWithConflict) GetSchema() *Schema {
	return &SchemaTypes
}
