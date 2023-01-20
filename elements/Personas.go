package elements

// The Personas element specifies an array of personas returned from the GetImItems and GetImItemList operations.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/personas-ex15websvcsotherref
import "encoding/xml"

type Personas struct {
	XMLName xml.Name

	// The Persona element specifies a set of persona data returned by a GetPersona request.
	Persona *Persona `xml:"Persona"`
}

func (P *Personas) SetForMarshal() {
	P.XMLName.Local = "t:Personas"
}

func (P *Personas) GetSchema() *Schema {
	return &SchemaTypes
}
