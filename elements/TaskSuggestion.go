package elements

// The TaskSuggestion element contains a task suggestion that resulted from an entity extracted from an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/tasksuggestion
type TaskSuggestion struct {
	// The Assignees element specifies the people to whom a task is assigned.
	Assignees *Assignees `xml:"t:Assignees"`
	// The Position element specifies the position of an entity extracted from a message.
	Position *Position `xml:"t:Position"`
	// The TaskString element contains a suggested task.
	TaskString *TaskString `xml:"t:TaskString"`
}
