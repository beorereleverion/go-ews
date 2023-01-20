package elements

// The MessageTypes element specifies an array of messages to search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/messagetypes
import "encoding/xml"

type MessageTypes struct {
	XMLName xml.Name

	// The SearchItemKind element indicates the type of items that are searched for a FindMailboxStatisticsByKeyword operation.
	SearchItemKind *SearchItemKind `xml:"SearchItemKind"`
}

func (M *MessageTypes) SetForMarshal() {
	M.XMLName.Local = "m:MessageTypes"
}

func (M *MessageTypes) GetSchema() *Schema {
	return &SchemaMessages
}
