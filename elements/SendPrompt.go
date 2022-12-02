package elements

// The SendPrompt element specifies the type of action allowed for a voting option.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sendprompt
type SendPrompt string

const (
	// None
	SendPromptNone SendPrompt = `None`
	// Send
	SendPromptSend SendPrompt = `Send`
	// VotingOption
	SendPromptVotingOption SendPrompt = `VotingOption`
)
