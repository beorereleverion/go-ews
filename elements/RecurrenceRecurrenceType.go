package elements

// The Recurrence element contains the recurrence pattern for calendar items and meeting requests.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recurrence-recurrencetype
type RecurrenceRecurrenceType struct {
	// The AbsoluteMonthlyRecurrence element represents a monthly recurrence pattern.
	AbsoluteMonthlyRecurrence *AbsoluteMonthlyRecurrence `xml:"t:AbsoluteMonthlyRecurrence"`
	// The AbsoluteYearlyRecurrence element represents a yearly recurrence pattern.
	AbsoluteYearlyRecurrence *AbsoluteYearlyRecurrence `xml:"t:AbsoluteYearlyRecurrence"`
	// The DailyRecurrence element describes the frequency, in days, in which a calendar item or a task recurs.
	DailyRecurrence *DailyRecurrence `xml:"t:DailyRecurrence"`
	// The EndDateRecurrence element describes the start date and the end date of an item recurrence pattern.
	EndDateRecurrence *EndDateRecurrence `xml:"t:EndDateRecurrence"`
	// The NoEndRecurrence element describes the start date of an item recurrence pattern that does not have a defined end date.
	NoEndRecurrence *NoEndRecurrence `xml:"t:NoEndRecurrence"`
	// The NumberedRecurrence element describes the start date and the number of occurrences of a recurring item.
	NumberedRecurrence *NumberedRecurrence `xml:"t:NumberedRecurrence"`
	// The RelativeMonthlyRecurrence element describes a relative monthly recurrence pattern.
	RelativeMonthlyRecurrence *RelativeMonthlyRecurrence `xml:"t:RelativeMonthlyRecurrence"`
	// The RelativeYearlyRecurrence element describes a relative yearly recurrence pattern.
	RelativeYearlyRecurrence *RelativeYearlyRecurrence `xml:"t:RelativeYearlyRecurrence"`
	// The WeeklyRecurrence element describes a weekly recurrence pattern.
	WeeklyRecurrence *WeeklyRecurrence `xml:"t:WeeklyRecurrence"`
}
