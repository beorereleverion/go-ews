package elements

// The UserOptions element specifies the list of voting options on a message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/useroptions
type UserOptions struct {
	// The VotingOptionData element specifies information about each voting option.
	VotingOptionData *VotingOptionData `xml:"t:VotingOptionData"`
}
