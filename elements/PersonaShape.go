package elements

// The PersonaShape element specifies the set of persona properties to be returned from a FindPeople request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/personashape
type PersonaShape struct {
	// The AdditionalProperties element identifies additional properties for use in GetItem, UpdateItem, CreateItem, FindItem, or FindFolder requests.
	AdditionalProperties *AdditionalProperties `xml:"t:AdditionalProperties"`
	// The BaseShape element identifies the set of properties to return in an item or folder response.
	BaseShape *BaseShape `xml:"t:BaseShape"`
}
