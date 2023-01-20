package elements

// The GetImItems request element defines a request to get information about the specified instant messaging groups and instant messaging contact personas.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getimitems
import "encoding/xml"

type GetImItems struct {
	XMLName xml.Name

	// The ContactIds element contains an array of contact item identifiers.
	ContactIds *ContactIds `xml:"ContactIds"`
	// The ExtendedProperties element contains the extended properties used for the Unified Contact Store operations.
	ExtendedProperties *ExtendedPropertiesNonEmptyArrayOfExtendedFieldURIs `xml:"ExtendedProperties"`
	// The GroupIds element identifies an array of instant messaging group identifiers.
	GroupIds *GroupIds `xml:"GroupIds"`
}

func (G *GetImItems) SetForMarshal() {
	G.XMLName.Local = "m:GetImItems"
}

func (G *GetImItems) GetSchema() *Schema {
	return &SchemaMessages
}
