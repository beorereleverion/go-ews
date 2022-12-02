package elements

// The VotingInformation element specifies voting information on a voting message and approval request message whereApproveandRejectare the voting options.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/votinginformation
type VotingInformation struct {
	// The UserOptions element specifies the list of voting options on a message.
	UserOptions *UserOptions `xml:"t:UserOptions"`
	// The VotingResponse element specifies the submitted vote. This element is only present on responses to voting request messages, not on responses to approvals.
	VotingResponse *VotingResponse `xml:"t:VotingResponse"`
}
