package elements

// The JobTitle element represents the job title of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/jobtitle
import "encoding/xml"

type JobTitle struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (J *JobTitle) SetForMarshal() {
	J.XMLName.Local = "t:JobTitle"
}

func (J *JobTitle) GetSchema() *Schema {
	return &SchemaTypes
}
