package elements

// The Recurrence element contains the recurrence pattern for calendar items and meeting requests.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recurrence-recurrencetype
import "encoding/xml"

type RecurrenceRecurrenceType struct {
	XMLName xml.Name

	// The AbsoluteMonthlyRecurrence element represents a monthly recurrence pattern.
	AbsoluteMonthlyRecurrence *AbsoluteMonthlyRecurrence `xml:"AbsoluteMonthlyRecurrence"`
	// The AbsoluteYearlyRecurrence element represents a yearly recurrence pattern.
	AbsoluteYearlyRecurrence *AbsoluteYearlyRecurrence `xml:"AbsoluteYearlyRecurrence"`
	// The DailyRecurrence element describes the frequency, in days, in which a calendar item or a task recurs.
	DailyRecurrence *DailyRecurrence `xml:"DailyRecurrence"`
	// The EndDateRecurrence element describes the start date and the end date of an item recurrence pattern.
	EndDateRecurrence *EndDateRecurrence `xml:"EndDateRecurrence"`
	// The NoEndRecurrence element describes the start date of an item recurrence pattern that does not have a defined end date.
	NoEndRecurrence *NoEndRecurrence `xml:"NoEndRecurrence"`
	// The NumberedRecurrence element describes the start date and the number of occurrences of a recurring item.
	NumberedRecurrence *NumberedRecurrence `xml:"NumberedRecurrence"`
	// The RelativeMonthlyRecurrence element describes a relative monthly recurrence pattern.
	RelativeMonthlyRecurrence *RelativeMonthlyRecurrence `xml:"RelativeMonthlyRecurrence"`
	// The RelativeYearlyRecurrence element describes a relative yearly recurrence pattern.
	RelativeYearlyRecurrence *RelativeYearlyRecurrence `xml:"RelativeYearlyRecurrence"`
	// The WeeklyRecurrence element describes a weekly recurrence pattern.
	WeeklyRecurrence *WeeklyRecurrence `xml:"WeeklyRecurrence"`
}

func (R *RecurrenceRecurrenceType) SetForMarshal() {
	R.XMLName.Local = "t:Recurrence"
}

func (R *RecurrenceRecurrenceType) GetSchema() *Schema {
	return &SchemaTypes
}
