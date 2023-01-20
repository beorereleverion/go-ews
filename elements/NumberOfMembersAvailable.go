package elements

// The NumberOfMembersAvailable element represents the number of distribution list members who are available for a suggested meeting time. This element represents members for whom the status is Free.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/numberofmembersavailable
import "encoding/xml"

type NumberOfMembersAvailable struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (N *NumberOfMembersAvailable) SetForMarshal() {
	N.XMLName.Local = "t:NumberOfMembersAvailable"
}

func (N *NumberOfMembersAvailable) GetSchema() *Schema {
	return &SchemaTypes
}
