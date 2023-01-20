package elements

// The CompanyNameSortKey element contains the sort key for a company name.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/companynamesortkey
import "encoding/xml"

type CompanyNameSortKey struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (C *CompanyNameSortKey) SetForMarshal() {
	C.XMLName.Local = "t:CompanyNameSortKey"
}

func (C *CompanyNameSortKey) GetSchema() *Schema {
	return &SchemaTypes
}
