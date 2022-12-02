package elements

// The IsAssignmentEditable element represents the task type.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isassignmenteditable
type IsAssignmentEditable int64

const (
	// The default for all task items.
	IsAssignmentEditableZero IsAssignmentEditable = 0
	// A task request.
	IsAssignmentEditableOne IsAssignmentEditable = 1
	// A task acceptance from a recipient of a task request.
	IsAssignmentEditableTwo IsAssignmentEditable = 2
	// A task declination from a recipient of a task request.
	IsAssignmentEditableTree IsAssignmentEditable = 3
	// An update to a previous task request.
	IsAssignmentEditableFour IsAssignmentEditable = 4
	// Not used.
	IsAssignmentEditable5 IsAssignmentEditable = 5
)
