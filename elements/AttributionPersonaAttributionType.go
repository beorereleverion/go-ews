package elements

// The Attribution element specifies an instance in an array of attributes for a PersonaType element.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attribution-personaattributiontype
import "encoding/xml"

type AttributionPersonaAttributionType struct {
	XMLName xml.Name

	// The DisplayName element defines the display name of a folder, contact, distribution list, delegate user, location, or rule.
	DisplayName *DisplayNamestring `xml:"DisplayName"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"FolderId"`
	// The ID element specifies the identifier of an app.
	ID *IDString `xml:"ID"`
	// The IsHidden element contains a Boolean value that indicates whether the underlying contact should be hidden or displayed as part of the persona.
	IsHidden *IsHidden `xml:"IsHidden"`
	// The IsQuickContact element specifies a Boolean value that indicates whether the underlying contact is a quick contact.
	IsQuickContact *IsQuickContact `xml:"IsQuickContact"`
	// The IsWritable element specifies whether the underlying contact or Active Directory recipient can be written to.
	IsWritable *IsWritable `xml:"IsWritable"`
	// The SourceId element specifies the identifier of the attributed contact in a persona.
	SourceId *SourceId `xml:"SourceId"`
}

func (A *AttributionPersonaAttributionType) SetForMarshal() {
	A.XMLName.Local = "t:Attribution"
}

func (A *AttributionPersonaAttributionType) GetSchema() *Schema {
	return &SchemaTypes
}
