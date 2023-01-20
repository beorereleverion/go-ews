package elements

// The Value element specifies information associated with a postal address.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/value-personapostaladdresstype
import "encoding/xml"

type ValuePersonaPostalAddressType struct {
	XMLName xml.Name

	// The Accuracy element specifies the accuracy of the latitude and longitude of the associated postal address.
	Accuracy *Accuracy `xml:"Accuracy"`
	// The Altitude element specifies the altitude of a postal address.
	Altitude *Altitude `xml:"Altitude"`
	// The AltitudeAccuracy element specifies the accuracy of the altitude property for a postal address.
	AltitudeAccuracy *AltitudeAccuracy `xml:"AltitudeAccuracy"`
	// The City element represents the city name that is associated with a contact.
	City *City `xml:"City"`
	// The Country element identifies a country identifier in a postal address.
	Country *Country `xml:"Country"`
	// The FormattedAddress element specifies the formatted display value of the associated postal address.
	FormattedAddress *FormattedAddress `xml:"FormattedAddress"`
	// The Latitude element specifies the latitude of the location of the associated postal address.
	Latitude *Latitude `xml:"Latitude"`
	// The LocationSource element specifies information about the origin of the associated postal address, for example, a contact or a telephone book.
	LocationSource *LocationSource `xml:"LocationSource"`
	// The LocationUri element contains a string specifying a Uniform Resource Identifier (URI) of the associated postal address.
	LocationUri *LocationUri `xml:"LocationUri"`
	// The Longitude element specifies the longitude of the location of the associated postal address.
	Longitude *Longitude `xml:"Longitude"`
	// The PostOfficeBox element specifies thepost office boxportion of a postal address.
	PostOfficeBox *PostOfficeBox `xml:"PostOfficeBox"`
	// The PostalCode element represents the postal code for a contact item.
	PostalCode *PostalCode `xml:"PostalCode"`
	// The State element represents the state of residence for a contact item.
	State *State `xml:"State"`
	// The Street element represents a street address for a contact item.
	Street *Street `xml:"Street"`
	// The Type element specifies the type of postal address or phone number, for example,HomeorBusiness.
	Type *Typestring `xml:"Type"`
}

func (V *ValuePersonaPostalAddressType) SetForMarshal() {
	V.XMLName.Local = "t:Value"
}

func (V *ValuePersonaPostalAddressType) GetSchema() *Schema {
	return &SchemaTypes
}
