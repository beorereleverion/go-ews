package elements

// The VotingInformation element specifies voting information on a voting message and approval request message whereApproveandRejectare the voting options.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/votinginformation
import "encoding/xml"

type VotingInformation struct {
	XMLName xml.Name

	// The UserOptions element specifies the list of voting options on a message.
	UserOptions *UserOptions `xml:"UserOptions"`
	// The VotingResponse element specifies the submitted vote. This element is only present on responses to voting request messages, not on responses to approvals.
	VotingResponse *VotingResponse `xml:"VotingResponse"`
}

func (V *VotingInformation) SetForMarshal() {
	V.XMLName.Local = "t:VotingInformation"
}

func (V *VotingInformation) GetSchema() *Schema {
	return &SchemaTypes
}
