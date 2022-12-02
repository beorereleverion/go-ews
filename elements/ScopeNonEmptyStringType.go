package elements

// The Scope element specifies the scope of the message tracking report.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/scope-nonemptystringtype
type ScopeNonEmptyStringType string

const (
	// The message tracking scopes spans across a forest.
	ScopeNonEmptyStringTypeForest ScopeNonEmptyStringType = `Forest`
	// The message tracking scopes spans across an organization.
	ScopeNonEmptyStringTypeOrganization ScopeNonEmptyStringType = `Organization`
	// The message tracking scopes spans across a site.
	ScopeNonEmptyStringTypeSite ScopeNonEmptyStringType = `Site`
)
