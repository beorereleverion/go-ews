package elements

// The TimeZoneDefinitions element represents an array of time zone definitions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/timezonedefinitions
type TimeZoneDefinitions struct {
	// The TimeZoneDefinition element specifies the periods and transitions that define a time zone.
	TimeZoneDefinition *TimeZoneDefinition `xml:"t:TimeZoneDefinition"`
}
