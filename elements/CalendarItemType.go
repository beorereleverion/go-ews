package elements

// The CalendarItemType element represents the type of a calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendaritemtype
type CalendarItemType string

const (
	// Exception The item is an exception to a recurring calendar item.

	CalendarItemTypeException CalendarItemType = `Exception`
	// Occurrence The item is an occurrence of a recurring calendar item.

	CalendarItemTypeOccurrence CalendarItemType = `Occurrence`
	// RecurringMaster The item is master for a set of recurring calendar items.

	CalendarItemTypeRecurringMaster CalendarItemType = `RecurringMaster`
	// Single The item is not associated with a recurring calendar item.

	CalendarItemTypeSingle CalendarItemType = `Single`
)
