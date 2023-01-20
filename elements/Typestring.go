package elements

// The Type element specifies the type of postal address or phone number, for example,HomeorBusiness.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/type-string
import "encoding/xml"

type Typestring struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (T *Typestring) SetForMarshal() {
	T.XMLName.Local = "t:Type"
}

func (T *Typestring) GetSchema() *Schema {
	return &SchemaTypes
}
