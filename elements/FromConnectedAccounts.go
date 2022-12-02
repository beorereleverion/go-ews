package elements

// The FromConnectedAccounts element represents the e-mail account names from which incoming messages have to have been aggregated in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fromconnectedaccounts
type FromConnectedAccounts struct {
	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"t:String"`
}
