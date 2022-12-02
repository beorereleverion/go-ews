package elements

// The TimeZoneContext element is used in the Simple Object Access Protocol (SOAP) header to specify the time zone definition that is to be used as the default when assigning the time zone for the DateTime properties of objects that are created, updated, and retrieved by using Exchange Web Services (EWS).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/timezonecontext
type TimeZoneContext struct {
	// The TimeZoneDefinition element specifies the periods and transitions that define a time zone.
	TimeZoneDefinition *TimeZoneDefinition `xml:"t:TimeZoneDefinition"`
}
