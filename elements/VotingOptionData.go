package elements

// The VotingOptionData element specifies information about each voting option.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/votingoptiondata
type VotingOptionData struct {
	// The DisplayName (VotingOptionDataType) element specifies the display name of a voting option.
	DisplayName *DisplayNameVotingOptionDataType `xml:"t:DisplayName"`
	// The SendPrompt element specifies the type of action allowed for a voting option.
	SendPrompt *SendPrompt `xml:"t:SendPrompt"`
}
