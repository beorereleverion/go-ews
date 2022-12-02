package elements

// The GetPersona element contains the request to get a persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getpersona
type GetPersona struct {
	// The PersonaId element specifies the persona identifier for the associated persona.
	PersonaId *PersonaId `xml:"m:PersonaId"`
}
