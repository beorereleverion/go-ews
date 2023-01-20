package elements

// The PageDirection element contains the direction for pagination in the search result. The value is Previous or Next.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/pagedirection
import "encoding/xml"

type PageDirection struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Next
	PageDirectionNext string = `Next`
	// Previous
	PageDirectionPrevious string = `Previous`
)

func (P *PageDirection) SetForMarshal() {
	P.XMLName.Local = "m:PageDirection"
}

func (P *PageDirection) GetSchema() *Schema {
	return &SchemaMessages
}
