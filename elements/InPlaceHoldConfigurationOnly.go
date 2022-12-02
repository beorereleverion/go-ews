package elements

// The InPlaceHoldConfigurationOnly element specifies whether to include the in-place hold configuration.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/inplaceholdconfigurationonly
type InPlaceHoldConfigurationOnly bool

const (
	// false
	InPlaceHoldConfigurationOnlyfalse InPlaceHoldConfigurationOnly = false
	// true
	InPlaceHoldConfigurationOnlytrue InPlaceHoldConfigurationOnly = true
)
