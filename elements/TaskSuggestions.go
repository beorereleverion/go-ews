package elements

// The TaskSuggestions element specifies an array of task suggestions extracted from an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/tasksuggestions
type TaskSuggestions struct {
	// The TaskSuggestion element contains a task suggestion that resulted from an entity extracted from an item.
	TaskSuggestion *TaskSuggestion `xml:"t:TaskSuggestion"`
}
