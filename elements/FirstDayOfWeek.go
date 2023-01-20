package elements

// The FirstDayOfWeek element specifies the first day of the week.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/firstdayofweek
import "encoding/xml"

type FirstDayOfWeek struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (F *FirstDayOfWeek) SetForMarshal() {
	F.XMLName.Local = "t:FirstDayOfWeek"
}

func (F *FirstDayOfWeek) GetSchema() *Schema {
	return &SchemaTypes
}
