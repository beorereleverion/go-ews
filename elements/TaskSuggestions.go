package elements

// The TaskSuggestions element specifies an array of task suggestions extracted from an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/tasksuggestions
import "encoding/xml"

type TaskSuggestions struct {
	XMLName xml.Name

	// The TaskSuggestion element contains a task suggestion that resulted from an entity extracted from an item.
	TaskSuggestion *TaskSuggestion `xml:"TaskSuggestion"`
}

func (T *TaskSuggestions) SetForMarshal() {
	T.XMLName.Local = "t:TaskSuggestions"
}

func (T *TaskSuggestions) GetSchema() *Schema {
	return &SchemaTypes
}
