package elements

// The IsAssignmentEditable element represents the task type.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isassignmenteditable
import "encoding/xml"

type IsAssignmentEditable struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

const (
	// The default for all task items.
	IsAssignmentEditableZero int64 = 0
	// A task request.
	IsAssignmentEditableOne int64 = 1
	// A task acceptance from a recipient of a task request.
	IsAssignmentEditableTwo int64 = 2
	// A task declination from a recipient of a task request.
	IsAssignmentEditableTree int64 = 3
	// An update to a previous task request.
	IsAssignmentEditableFour int64 = 4
	// Not used.
	IsAssignmentEditable5 int64 = 5
)

func (I *IsAssignmentEditable) SetForMarshal() {
	I.XMLName.Local = "t:IsAssignmentEditable"
}

func (I *IsAssignmentEditable) GetSchema() *Schema {
	return &SchemaTypes
}
