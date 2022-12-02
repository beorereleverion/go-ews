package elements

// The Presenters element specifies the presenters for an online meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/presenters
type Presenters string

const (
	// Disabled
	PresentersDisabled Presenters = `Disabled`
	// Everyone
	PresentersEveryone Presenters = `Everyone`
	// Internal
	PresentersInternal Presenters = `Internal`
)
