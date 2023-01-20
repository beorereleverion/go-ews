package elements

// The NumberOfMembers element represents the number of users, resources, and rooms in a distribution list.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/numberofmembers
import "encoding/xml"

type NumberOfMembers struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (N *NumberOfMembers) SetForMarshal() {
	N.XMLName.Local = "t:NumberOfMembers"
}

func (N *NumberOfMembers) GetSchema() *Schema {
	return &SchemaTypes
}
