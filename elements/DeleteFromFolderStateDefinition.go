package elements

// The DeleteFromFolderStateDefinition element specifies the state when an item is deleted from a folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deletefromfolderstatedefinition
type DeleteFromFolderStateDefinition struct {
	// The IsOccurrencePresent element is intended for internal use only.
	IsOccurrencePresent *IsOccurrencePresent `xml:"IsOccurrencePresent"`
	// The Occurrence element represents the occurrence of the day of the week in the month that the time zone transition occurs.
	Occurrence *OccurrenceTimeZoneTransition `xml:"t:Occurrence"`
}
