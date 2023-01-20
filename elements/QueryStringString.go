package elements

// The QueryString element specifies a set of values to be returned that match the query string in a FindPeople operation request. A search with no QueryString specified returns all items from the specified folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/querystring-string
import "encoding/xml"

type QueryStringString struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (Q *QueryStringString) SetForMarshal() {
	Q.XMLName.Local = "m:QueryString"
}

func (Q *QueryStringString) GetSchema() *Schema {
	return &SchemaMessages
}
