package elements

// The ConfigurationName element specifies the requested service configurations by name.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/configurationname
type ConfigurationName string

const (
	// MailTips
	ConfigurationNameMailTips ConfigurationName = `MailTips`
	// ProtectionRules
	ConfigurationNameProtectionRules ConfigurationName = `ProtectionRules`
	// UnifiedMessagingConfiguration
	ConfigurationNameUnifiedMessagingConfiguration ConfigurationName = `UnifiedMessagingConfiguration`
)
