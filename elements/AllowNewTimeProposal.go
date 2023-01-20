package elements

// The AllowNewTimeProposal element indicates whether a new meeting time can be proposed for a meeting by an attendee.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/allownewtimeproposal
import "encoding/xml"

type AllowNewTimeProposal struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (A *AllowNewTimeProposal) SetForMarshal() {
	A.XMLName.Local = "t:AllowNewTimeProposal"
}

func (A *AllowNewTimeProposal) GetSchema() *Schema {
	return &SchemaTypes
}
