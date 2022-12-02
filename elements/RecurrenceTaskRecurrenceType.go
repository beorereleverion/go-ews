package elements

// The Recurrence element contains recurrence information for recurring tasks.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recurrence-taskrecurrencetype
type RecurrenceTaskRecurrenceType struct {
	// The AbsoluteMonthlyRecurrence element represents a monthly recurrence pattern.
	AbsoluteMonthlyRecurrence *AbsoluteMonthlyRecurrence `xml:"t:AbsoluteMonthlyRecurrence"`
	// The AbsoluteYearlyRecurrence element represents a yearly recurrence pattern.
	AbsoluteYearlyRecurrence *AbsoluteYearlyRecurrence `xml:"t:AbsoluteYearlyRecurrence"`
	// The DailyRecurrence element describes the frequency, in days, in which a calendar item or a task recurs.
	DailyRecurrence *DailyRecurrence `xml:"t:DailyRecurrence"`
	// The DailyRegeneration element describes the frequency, in days, in which a task is regenerated.
	DailyRegeneration *DailyRegeneration `xml:"t:DailyRegeneration"`
	// The EndDateRecurrence element describes the start date and the end date of an item recurrence pattern.
	EndDateRecurrence *EndDateRecurrence `xml:"t:EndDateRecurrence"`
	// The MonthlyRegeneration element describes the frequency, in months, of which task is regenerated.
	MonthlyRegeneration *MonthlyRegeneration `xml:"t:MonthlyRegeneration"`
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
	// The WeeklyRegeneration element describes the frequency, in weeks, in which a task is regenerated.
	WeeklyRegeneration *WeeklyRegeneration `xml:"t:WeeklyRegeneration"`
	// The YearlyRegeneration element describes the frequency, in years, in which a task is regenerated.
	YearlyRegeneration *YearlyRegeneration `xml:"t:YearlyRegeneration"`
}
