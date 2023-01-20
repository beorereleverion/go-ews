package elements

// The RetentionPolicyTag element specifies the retention policy for a mailbox item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/retentionpolicytag
import "encoding/xml"

type RetentionPolicyTag struct {
	XMLName xml.Name

	// The Description element specifies the descriptive text for the retention policy.
	Description *Description `xml:"Description"`
	// The DisplayName element defines the display name of a folder, contact, distribution list, delegate user, location, or rule.
	DisplayName *DisplayNamestring `xml:"DisplayName"`
	// The IsArchive element specifies a Boolean value that indicates whether the mailbox is an archive mailbox.
	IsArchive *IsArchive `xml:"IsArchive"`
	// The IsVisible element indicates whether the retention policy is visible to users.
	IsVisible *IsVisible `xml:"IsVisible"`
	// The OptedInto element specifies a Boolean value that indicates whether the user opted in to the retention policy.
	OptedInto *OptedInto `xml:"OptedInto"`
	// The RetentionAction element specifies the action performed on items with the retention tag.
	RetentionAction *RetentionAction `xml:"RetentionAction"`
	// The RetentionId element specifies the retention tag identifier.
	RetentionId *RetentionId `xml:"RetentionId"`
	// The RetentionPeriod element specifies the number of days that the retention policy is in effect.
	RetentionPeriod *RetentionPeriod `xml:"RetentionPeriod"`
	// The Type element specifies the type of folder used in a retention policy.
	Type *TypeElcFolderType `xml:"Type"`
}

func (R *RetentionPolicyTag) SetForMarshal() {
	R.XMLName.Local = "t:RetentionPolicyTag"
}

func (R *RetentionPolicyTag) GetSchema() *Schema {
	return &SchemaTypes
}
