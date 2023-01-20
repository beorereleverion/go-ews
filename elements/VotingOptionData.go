package elements

// The VotingOptionData element specifies information about each voting option.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/votingoptiondata
import "encoding/xml"

type VotingOptionData struct {
	XMLName xml.Name

	// The DisplayName (VotingOptionDataType) element specifies the display name of a voting option.
	DisplayName *DisplayNameVotingOptionDataType `xml:"DisplayName"`
	// The SendPrompt element specifies the type of action allowed for a voting option.
	SendPrompt *SendPrompt `xml:"SendPrompt"`
}

func (V *VotingOptionData) SetForMarshal() {
	V.XMLName.Local = "t:VotingOptionData"
}

func (V *VotingOptionData) GetSchema() *Schema {
	return &SchemaTypes
}
