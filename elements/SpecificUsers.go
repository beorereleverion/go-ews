package elements

// The SpecificUsers element specifies the email accounts that can access the app.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/specificusers
type SpecificUsers struct {
	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"t:String"`
}
