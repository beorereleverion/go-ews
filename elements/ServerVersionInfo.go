package elements

// The ServerVersionInfo element represents the Microsoft Exchange Server version number.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/serverversioninfo
type ServerVersionInfo struct {
	// Describes the major build number.
	MajorBuildNumber *string `xml:"MajorBuildNumber,attr"`
	// Describes the major version number.
	MajorVersion *string `xml:"MajorVersion,attr"`
	// Describes the minor build number.
	MinorBuildNumber *string `xml:"MinorBuildNumber,attr"`
	// Describes the minor version number.
	MinorVersion *string `xml:"MinorVersion,attr"`
	// Describes the Exchange Web Services (EWS) schema version.
	Version *string `xml:"Version,attr"`
}
