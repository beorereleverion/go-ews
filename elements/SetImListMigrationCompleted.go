package elements

// The SetImListMigrationCompleted element represents a request to indicate whether the Exchange store contains the instant messaging items used by instant messaging clients.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setimlistmigrationcompleted
import "encoding/xml"

type SetImListMigrationCompleted struct {
	XMLName xml.Name

	// The ImListMigrationCompleted element indicates whether the Exchange store contains the instant messaging items used by instant messaging clients.
	ImListMigrationCompleted *ImListMigrationCompleted `xml:"ImListMigrationCompleted"`
}

func (S *SetImListMigrationCompleted) SetForMarshal() {
	S.XMLName.Local = "m:SetImListMigrationCompleted"
}

func (S *SetImListMigrationCompleted) GetSchema() *Schema {
	return &SchemaMessages
}
