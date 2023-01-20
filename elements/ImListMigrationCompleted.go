package elements

// The ImListMigrationCompleted element indicates whether the Exchange store contains the instant messaging items used by instant messaging clients.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/imlistmigrationcompleted
import "encoding/xml"

type ImListMigrationCompleted struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	ImListMigrationCompletedfalse bool = false
	// true
	ImListMigrationCompletedtrue bool = true
)

func (I *ImListMigrationCompleted) SetForMarshal() {
	I.XMLName.Local = "m:ImListMigrationCompleted"
}

func (I *ImListMigrationCompleted) GetSchema() *Schema {
	return &SchemaMessages
}
