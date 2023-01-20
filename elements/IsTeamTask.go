package elements

// The IsTeamTask element indicates whether the task is owned by a team.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isteamtask
import "encoding/xml"

type IsTeamTask struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsTeamTask) SetForMarshal() {
	I.XMLName.Local = "t:IsTeamTask"
}

func (I *IsTeamTask) GetSchema() *Schema {
	return &SchemaTypes
}
