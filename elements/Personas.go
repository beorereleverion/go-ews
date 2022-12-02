package elements

// The Personas element specifies an array of personas returned from the GetImItems and GetImItemList operations.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/personas-ex15websvcsotherref
type Personas struct {
	// The Persona element specifies a set of persona data returned by a GetPersona request.
	Persona *Persona `xml:"m:Persona"`
}
