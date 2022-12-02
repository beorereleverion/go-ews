package elements

// The TokenType element specifies the type of token.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/tokentype
type TokenType string

const (
	// CallerIdentity
	TokenTypeCallerIdentity TokenType = `CallerIdentity`
	// ExtensionCallback
	TokenTypeExtensionCallback TokenType = `ExtensionCallback`
	// ScopedToken
	TokenTypeScopedToken TokenType = `ScopedToken`
)
