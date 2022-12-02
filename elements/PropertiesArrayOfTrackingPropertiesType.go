package elements

// The Properties element contains a list of one or more tracking properties.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/properties-arrayoftrackingpropertiestype
type PropertiesArrayOfTrackingPropertiesType struct {
	// The TrackingPropertyType element represents a name and value pair of strings that is used to create properties for message tracking reports.
	TrackingPropertyType *TrackingPropertyType `xml:"m:TrackingPropertyType"`
}
