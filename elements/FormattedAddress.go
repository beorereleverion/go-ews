package elements

// The FormattedAddress element specifies the formatted display value of the associated postal address.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/formattedaddress
import "encoding/xml"

type FormattedAddress struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (F *FormattedAddress) SetForMarshal() {
	F.XMLName.Local = "t:FormattedAddress"
}

func (F *FormattedAddress) GetSchema() *Schema {
	return &SchemaTypes
}
