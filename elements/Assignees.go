package elements

// The Assignees element specifies the people to whom a task is assigned.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/assignees
type Assignees struct {
	// The Name element represents the display name of the mailbox user.
	Name *NameEmailAddress `xml:"t:Name"`
	// The UserId element specifies the user identifier of an email user.
	UserId *UserIdstring `xml:"t:UserId"`
}
