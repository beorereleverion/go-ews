package elements

// The Language (DiscoverySearchConfigurationType) element identifies the culture to be used for the culture-specific format of date ranges. It also specifies the language used in a search query.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/language-discoverysearchconfigurationtype
import "encoding/xml"

type LanguageDiscoverySearchConfigurationType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (L *LanguageDiscoverySearchConfigurationType) SetForMarshal() {
	L.XMLName.Local = "t:Language"
}

func (L *LanguageDiscoverySearchConfigurationType) GetSchema() *Schema {
	return &SchemaTypes
}
