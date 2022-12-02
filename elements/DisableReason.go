package elements

// The DisableReason element specifies the reason for disabling an app.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/disablereason
type DisableReason string

const (
	// To improve mobile client performance.
	DisableReasonMobileClientPerformance DisableReason = `MobileClientPerformance`
	// No reason given
	DisableReasonNoReason DisableReason = `NoReason`
	// To improve Web app client performance.
	DisableReasonOWAClientPerformance DisableReason = `OWAClientPerformance`
	// To improve email client performance.
	DisableReasonOutlookClientPerformance DisableReason = `OutlookClientPerformance`
)
