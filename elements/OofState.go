package elements

// The OofState element is used to get or set the user's Out of Office (OOF) state.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/oofstate
type OofState string

const (
	// Disabled

	OofStateDisabled OofState = `Disabled`
	// Enabled

	OofStateEnabled OofState = `Enabled`
	// Scheduled

	OofStateScheduled OofState = `Scheduled`
)
