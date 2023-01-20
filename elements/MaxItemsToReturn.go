package elements

// The MaxItemsToReturn element identifies the maximum number of conversations items to return in a GetConversationItems response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/maxitemstoreturn
import "encoding/xml"

type MaxItemsToReturn struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (M *MaxItemsToReturn) SetForMarshal() {
	M.XMLName.Local = "m:MaxItemsToReturn"
}

func (M *MaxItemsToReturn) GetSchema() *Schema {
	return &SchemaMessages
}
