package elements

// The Status element provides information about the status of a distribution list member on the server.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/status-memberstatustype
import "encoding/xml"

type StatusMemberStatusType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Demoted
	StatusMemberStatusTypeDemoted string = `Demoted`
	// Normal
	StatusMemberStatusTypeNormal string = `Normal`
	// Unrecognized
	StatusMemberStatusTypeUnrecognized string = `Unrecognized`
)

func (S *StatusMemberStatusType) SetForMarshal() {
	S.XMLName.Local = "t:Status"
}

func (S *StatusMemberStatusType) GetSchema() *Schema {
	return &SchemaTypes
}
