package elements

// The DisplayName (VotingOptionDataType) element specifies the display name of a voting option.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/displayname-votingoptiondatatype
import "encoding/xml"

type DisplayNameVotingOptionDataType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DisplayNameVotingOptionDataType) SetForMarshal() {
	D.XMLName.Local = "t:DisplayName"
}

func (D *DisplayNameVotingOptionDataType) GetSchema() *Schema {
	return &SchemaTypes
}
