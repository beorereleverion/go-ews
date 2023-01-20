package elements

// The TaskSuggestion element contains a task suggestion that resulted from an entity extracted from an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/tasksuggestion
import "encoding/xml"

type TaskSuggestion struct {
	XMLName xml.Name

	// The Assignees element specifies the people to whom a task is assigned.
	Assignees *Assignees `xml:"Assignees"`
	// The Position element specifies the position of an entity extracted from a message.
	Position *Position `xml:"Position"`
	// The TaskString element contains a suggested task.
	TaskString *TaskString `xml:"TaskString"`
}

func (T *TaskSuggestion) SetForMarshal() {
	T.XMLName.Local = "t:TaskSuggestion"
}

func (T *TaskSuggestion) GetSchema() *Schema {
	return &SchemaTypes
}
