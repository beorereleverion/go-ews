package elements

// The TransitionsGroups element represents an array of time zone transition groups.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/transitionsgroups
type TransitionsGroups struct {
	// The TransitionsGroup element represents an array of time zone transitions.
	TransitionsGroup *TransitionsGroup `xml:"t:TransitionsGroup"`
}
