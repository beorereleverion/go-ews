package elements

// The Period element defines the name, time offset, and unique identifier for a specific stage of the time zone.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/period
type Period struct {
	// An xs:duration value that represents the time offset from Coordinated Universal Time (UTC) for the period.
	Bias *string `xml:"Bias,attr"`
	// A string value that represents the identifier for the period.
	Id *string `xml:"Id,attr"`
	// A string value that represents the descriptive name of the period.
	Name *string `xml:"Name,attr"`
}
