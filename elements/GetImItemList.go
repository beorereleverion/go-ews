package elements

// The GetImItemList element defines a request to get a list of instant messaging groups and contacts.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getimitemlist
import "encoding/xml"

type GetImItemList struct {
	XMLName xml.Name

	// The ExtendedProperties element contains the extended properties used for the Unified Contact Store operations.
	ExtendedProperties *ExtendedPropertiesNonEmptyArrayOfExtendedFieldURIs `xml:"ExtendedProperties"`
}

func (G *GetImItemList) SetForMarshal() {
	G.XMLName.Local = "m:GetImItemList"
}

func (G *GetImItemList) GetSchema() *Schema {
	return &SchemaMessages
}
