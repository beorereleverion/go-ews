package elements

// The Range element specifies a range of calendar item occurrences for a repeating calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/range
type Range struct {
	// The text value of true for the CompareOriginalStartTime attribute indicates that the client should compare the original start time with the new start time. A value of false indicates that the client does not need to compare the original start time with the new start time.
	CompareOriginalStartTime *string `xml:"CompareOriginalStartTime,attr"`
	// The text value of the Count attribute is the number of occurrences of the recurring item. This is an integer value.
	Count *string `xml:"Count,attr"`
	// The text value of the End attribute is the end date of the recurring item range. This is a dateTime value.
	End *string `xml:"End,attr"`
	// The text value of the Start attribute is the start date of the recurring item range. This is a dateTime value.
	Start *string `xml:"Start,attr"`
}
