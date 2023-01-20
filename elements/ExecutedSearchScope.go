package elements

// The ExecutedSearchScope element contains the scope of the search that was performed to get the search results.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/executedsearchscope
import "encoding/xml"

type ExecutedSearchScope struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (E *ExecutedSearchScope) SetForMarshal() {
	E.XMLName.Local = "t:ExecutedSearchScope"
}

func (E *ExecutedSearchScope) GetSchema() *Schema {
	return &SchemaTypes
}
