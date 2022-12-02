package elements

// The TokenType element identifies the type of client access token that will be returned in the GetClientAccessToken response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/tokentype-clientaccesstokentype
type TokenTypeClientAccessTokenType string

const (
	// CallerIdentity
	TokenTypeClientAccessTokenTypeCallerIdentity TokenTypeClientAccessTokenType = `CallerIdentity`
	// ExtensionCallback
	TokenTypeClientAccessTokenTypeExtensionCallback TokenTypeClientAccessTokenType = `ExtensionCallback`
	// ScopedToken
	TokenTypeClientAccessTokenTypeScopedToken TokenTypeClientAccessTokenType = `ScopedToken`
)
