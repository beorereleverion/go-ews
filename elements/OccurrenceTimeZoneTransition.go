package elements

// The Occurrence element represents the occurrence of the day of the week in the month that the time zone transition occurs.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/occurrence-time-zone-transition
type OccurrenceTimeZoneTransition int64

const (
	// The first occurrence of the specified day of the week from the end of the month.
	OccurrenceTimeZoneTransitionMinusOne OccurrenceTimeZoneTransition = -1
	// The second occurrence of the specified day of the week from the end of the month.
	OccurrenceTimeZoneTransitionMinusTwo OccurrenceTimeZoneTransition = -2
	// The third occurrence of the specified day of the week from the end of the month.
	OccurrenceTimeZoneTransitionMinusTree OccurrenceTimeZoneTransition = -3
	// The fourth occurrence of the specified day of the week from the end of the month.
	OccurrenceTimeZoneTransitionMinusFour OccurrenceTimeZoneTransition = -4
	// The first occurrence of the specified day of the week from the beginning of the month.
	OccurrenceTimeZoneTransitionOne OccurrenceTimeZoneTransition = 1
	// The second occurrence of the specified day of the week from the beginning of the month.
	OccurrenceTimeZoneTransitionTwo OccurrenceTimeZoneTransition = 2
	// The third occurrence of the specified day of the week from the beginning of the month.
	OccurrenceTimeZoneTransitionTree OccurrenceTimeZoneTransition = 3
	// The fourth occurrence of the specified day of the week from the beginning of the month.
	OccurrenceTimeZoneTransitionFour OccurrenceTimeZoneTransition = 4
)
