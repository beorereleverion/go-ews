package elements

// The BaseShape element identifies the set of properties to return in an item or folder response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/baseshape
type BaseShape string

const (
	// Returns only the item or folder ID.
	BaseShapeIdOnly BaseShape = `IdOnly`
	// Returns a set of properties that are defined as the default for the item or folder.
	BaseShapeDefault BaseShape = `Default`
	// Returns all the properties used by the Exchange Business Logic layer to construct a folder.
	BaseShapeAllProperties BaseShape = `AllProperties`
)
