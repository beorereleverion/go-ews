package elements

// The People element specifies an array of persona data returned as the result of a FindPeople request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/people
type People struct {
	// The Persona element specifies a set of persona data returned by a GetPersona request.
	Persona *Persona `xml:"m:Persona"`
}
