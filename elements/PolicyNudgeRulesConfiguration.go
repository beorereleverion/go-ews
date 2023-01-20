package elements

// The PolicyNudgeRulesConfiguration element contains the data loss prevention configuration data.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/policynudgerulesconfiguration
import "encoding/xml"

type PolicyNudgeRulesConfiguration struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (P *PolicyNudgeRulesConfiguration) SetForMarshal() {
	P.XMLName.Local = "m:PolicyNudgeRulesConfiguration"
}

func (P *PolicyNudgeRulesConfiguration) GetSchema() *Schema {
	return &SchemaMessages
}
