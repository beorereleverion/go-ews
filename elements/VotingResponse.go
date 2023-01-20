package elements

// The VotingResponse element specifies the submitted vote. This element is only present on responses to voting request messages, not on responses to approvals.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/votingresponse
import "encoding/xml"

type VotingResponse struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (V *VotingResponse) SetForMarshal() {
	V.XMLName.Local = "t:VotingResponse"
}

func (V *VotingResponse) GetSchema() *Schema {
	return &SchemaTypes
}
