package elements

// The Entry element describes a single physical address for a contact item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/entry-physicaladdress
import "encoding/xml"

type EntryPhysicalAddress struct {
	XMLName xml.Name

	// The City element represents the city name that is associated with a contact.
	City *City `xml:"City"`
	// The Country element represents the country or region for a given physical address.
	CountryOrRegion *CountryOrRegion `xml:"CountryOrRegion"`
	// The PostalCode element represents the postal code for a contact item.
	PostalCode *PostalCode `xml:"PostalCode"`
	// The State element represents the state of residence for a contact item.
	State *State `xml:"State"`
	// The Street element represents a street address for a contact item.
	Street *Street `xml:"Street"`
	// Identifies a physical address. The following are the possible values for this attribute:  - Business  - Home  - Other
	Key *string `xml:"Key,attr"`
}

func (E *EntryPhysicalAddress) SetForMarshal() {
	E.XMLName.Local = "t:Entry"
}

func (E *EntryPhysicalAddress) GetSchema() *Schema {
	return &SchemaTypes
}
