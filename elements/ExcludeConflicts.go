package elements

// The ExcludeConflicts element specifies whether to return suggested times for calendar times that conflict among the attendees.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/excludeconflicts
type ExcludeConflicts bool

const (
	// false
	ExcludeConflictsfalse ExcludeConflicts = false
	// true
	ExcludeConflictstrue ExcludeConflicts = true
)
