package elements

// The BaseShape element identifies the set of properties to return in an item or folder response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/baseshape
type BaseShape string

const (
	// AllProperties
	BaseShapeAllProperties BaseShape = `AllProperties`
	// Default
	BaseShapeDefault BaseShape = `Default`
	// IdOnly
	BaseShapeIdOnly BaseShape = `IdOnly`
)
