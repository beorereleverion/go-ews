package elements

// The TotalNumberOfPeopleInView element specifies the total number of personas returned in a FindPeople response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/totalnumberofpeopleinview
import "encoding/xml"

type TotalNumberOfPeopleInView struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (T *TotalNumberOfPeopleInView) SetForMarshal() {
	T.XMLName.Local = "m:TotalNumberOfPeopleInView"
}

func (T *TotalNumberOfPeopleInView) GetSchema() *Schema {
	return &SchemaMessages
}
