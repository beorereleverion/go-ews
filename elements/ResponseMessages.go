package elements

// The ResponseMessages element contains the response messages for an Exchange Web Services request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/responsemessages
import "encoding/xml"

type ResponseMessages struct {
	XMLName xml.Name

	// The ApplyConversationActionResponseMessage element contains the status and results of an ApplyConversationAction operation request.
	ApplyConversationActionResponseMessage *ApplyConversationActionResponseMessage `xml:"ApplyConversationActionResponseMessage"`
	// The ConvertIdResponseMessage element contains the status and result of a ConvertId operation request.
	ConvertIdResponseMessage *ConvertIdResponseMessage `xml:"ConvertIdResponseMessage"`
	// The CopyFolderResponseMessage element contains the status and result of a single CopyFolder operation request.
	CopyFolderResponseMessage *CopyFolderResponseMessage `xml:"CopyFolderResponseMessage"`
	// The CopyItemResponseMessage element contains the status and result of a single CopyItem operation request.
	CopyItemResponseMessage *CopyItemResponseMessage `xml:"CopyItemResponseMessage"`
	// The CreateAttachmentResponseMessage element contains the status and result of a single CreateAttachment operation request.
	CreateAttachmentResponseMessage *CreateAttachmentResponseMessage `xml:"CreateAttachmentResponseMessage"`
	// The CreateFolderResponseMessage element contains the status and result of a single CreateFolder operation request.
	CreateFolderResponseMessage *CreateFolderResponseMessage `xml:"CreateFolderResponseMessage"`
	// The CreateItemResponseMessage element contains the status and result of a single CreateItem operation request.
	CreateItemResponseMessage *CreateItemResponseMessage `xml:"CreateItemResponseMessage"`
	// The CreateManagedFolderResponseMessage element contains the status and result of a single CreateManagedFolder operation request.
	CreateManagedFolderResponseMessage *CreateManagedFolderResponseMessage `xml:"CreateManagedFolderResponseMessage"`
	// The CreateUserConfigurationResponseMessage element contains the status and result of a single CreateUserConfiguration request.
	CreateUserConfigurationResponseMessage *CreateUserConfigurationResponseMessage `xml:"CreateUserConfigurationResponseMessage"`
	// The DeleteAttachmentResponseMessage element contains the status and result of a single DeleteAttachment operation request.
	DeleteAttachmentResponseMessage *DeleteAttachmentResponseMessage `xml:"DeleteAttachmentResponseMessage"`
	// The DeleteFolderResponseMessage element contains the status and result of a single DeleteFolder operation request.
	DeleteFolderResponseMessage *DeleteFolderResponseMessage `xml:"DeleteFolderResponseMessage"`
	// The DeleteItemResponseMessage element contains the status and result of a single DeleteItem operation request.
	DeleteItemResponseMessage *DeleteItemResponseMessage `xml:"DeleteItemResponseMessage"`
	// The DeleteUserConfigurationResponseMessage element contains the status and result of a single DeleteUserConfiguration request.
	DeleteUserConfigurationResponseMessage *DeleteUserConfigurationResponseMessage `xml:"DeleteUserConfigurationResponseMessage"`
	// The EmptyFolder element defines a request to empty a folder in a mailbox in the Exchange store. Optionally, subfolders can also be deleted when the folder is emptied.
	EmptyFolder *EmptyFolder `xml:"EmptyFolder"`
	// The EmptyFolderResponseMessage element contains the status and result of a single EmptyFolder request.
	EmptyFolderResponseMessage *EmptyFolderResponseMessage `xml:"EmptyFolderResponseMessage"`
	// The ExpandDLResponseMessage element contains the status and result of a single ExpandDL operation request.
	ExpandDLResponseMessage *ExpandDLResponseMessage `xml:"ExpandDLResponseMessage"`
	// The ExportItemsResponseMessage element contains the status and results of a request to export a single mailbox item.
	ExportItemsResponseMessage *ExportItemsResponseMessage `xml:"ExportItemsResponseMessage"`
	// The FindFolderResponseMessage element contains the status and result of a single FindFolder operation request.
	FindFolderResponseMessage *FindFolderResponseMessage `xml:"FindFolderResponseMessage"`
	// The FindItemResponseMessage element contains the status and result of a single FindItem operation request.
	FindItemResponseMessage *FindItemResponseMessage `xml:"FindItemResponseMessage"`
	// The GetAttachmentResponseMessage element contains the status and result of a single GetAttachment operation request.
	GetAttachmentResponseMessage *GetAttachmentResponseMessage `xml:"GetAttachmentResponseMessage"`
	// The GetEventsResponseMessage element contains the status and result of a single GetEvents operation request.
	GetEventsResponseMessage *GetEventsResponseMessage `xml:"GetEventsResponseMessage"`
	// The GetFolderResponseMessage element contains the status and result of a single GetFolder operation request.
	GetFolderResponseMessage *GetFolderResponseMessage `xml:"GetFolderResponseMessage"`
	// The GetItemResponseMessage element contains the status and result of a single GetItem operation request.
	GetItemResponseMessage *GetItemResponseMessage `xml:"GetItemResponseMessage"`
	// The GetRemindersResponse element specifies the response to a GetReminders request.
	GetRemindersResponse *GetRemindersResponse `xml:"GetRemindersResponse"`
	// The GetRoomListsResponse element defines the response from a GetRoomLists operation request.
	GetRoomListsResponse *GetRoomListsResponse `xml:"GetRoomListsResponse"`
	// The GetRoomsResponse element defines a response to a GetRooms operation request.
	GetRoomsResponse *GetRoomsResponse `xml:"GetRoomsResponse"`
	// The GetServerTimeZonesResponseMessage element contains the status and result of a single GetServerTimeZones operation request.
	GetServerTimeZonesResponseMessage *GetServerTimeZonesResponseMessage `xml:"GetServerTimeZonesResponseMessage"`
	// The GetSharingFolderResponseMessage element contains the status and result of a single GetSharingFolder operation request.
	GetSharingFolderResponseMessage *GetSharingFolderResponseMessage `xml:"GetSharingFolderResponseMessage"`
	// The GetSharingMetadataResponseMessage element contains the status and result of a single GetSharingMetadata operation request.
	GetSharingMetadataResponseMessage *GetSharingMetadataResponseMessage `xml:"GetSharingMetadataResponseMessage"`
	// The GetStreamingEventsResponseMessage element contains the status and result of a single GetStreamingEvents operation request.
	GetStreamingEventsResponseMessage *GetStreamingEventsResponseMessage `xml:"GetStreamingEventsResponseMessage"`
	// The GetUserConfigurationResponseMessage element represents a response that returns a user configuration object.
	GetUserConfigurationResponseMessage *GetUserConfigurationResponseMessage `xml:"GetUserConfigurationResponseMessage"`
	// The MoveFolderResponseMessage element contains the status and result of a single MoveFolder operation request.
	MoveFolderResponseMessage *MoveFolderResponseMessage `xml:"MoveFolderResponseMessage"`
	// The MoveItemResponseMessage element contains the status and result of a single MoveItem operation request.
	MoveItemResponseMessage *MoveItemResponseMessage `xml:"MoveItemResponseMessage"`
	// The PerformReminderActionResponse element specifies the response to a PerformReminderAction request.
	PerformReminderActionResponse *PerformReminderActionResponse `xml:"PerformReminderActionResponse"`
	// The RefreshSharingFolderResponseMessage element contains the status and result of a single RefreshSharingFolder operation request.
	RefreshSharingFolderResponseMessage *RefreshSharingFolderResponseMessage `xml:"RefreshSharingFolderResponseMessage"`
	// The ResolveNamesResponseMessage element contains the status and result of a ResolveNames operation request.
	ResolveNamesResponseMessage *ResolveNamesResponseMessage `xml:"ResolveNamesResponseMessage"`
	// The SendItemResponseMessage element contains the status and result of a single SendItem operation request.
	SendItemResponseMessage *SendItemResponseMessage `xml:"SendItemResponseMessage"`
	// The SendNotificationResponseMessage element contains the status and result of a single SendNotification operation request.
	SendNotificationResponseMessage *SendNotificationResponseMessage `xml:"SendNotificationResponseMessage"`
	// The SubscribeResponseMessage element contains the status and result of a single Subscribe operation request.
	SubscribeResponseMessage *SubscribeResponseMessage `xml:"SubscribeResponseMessage"`
	// The SyncFolderHierarchyResponseMessage element contains the status and result of a single SyncFolderHierarchy operation request.
	SyncFolderHierarchyResponseMessage *SyncFolderHierarchyResponseMessage `xml:"SyncFolderHierarchyResponseMessage"`
	// The SyncFolderItemsResponseMessage element contains the status and result of a single SyncFolderItems operation request.
	SyncFolderItemsResponseMessage *SyncFolderItemsResponseMessage `xml:"SyncFolderItemsResponseMessage"`
	// The UnsubscribeResponseMessage element contains the status and result of a single Unsubscribe operation request.
	UnsubscribeResponseMessage *UnsubscribeResponseMessage `xml:"UnsubscribeResponseMessage"`
	// The UpdateFolderResponseMessage element contains the status and result of updates defined by the FolderChange element of an UpdateFolder operation request.
	UpdateFolderResponseMessage *UpdateFolderResponseMessage `xml:"UpdateFolderResponseMessage"`
	// The UpdateItemResponseMessage element contains the status and result of a single UpdateItem operation request.
	UpdateItemResponseMessage *UpdateItemResponseMessage `xml:"UpdateItemResponseMessage"`
	// The UpdateUserConfigurationResponseMessage element contains the status and result of a single UpdateUserConfiguration operation request.
	UpdateUserConfigurationResponseMessage *UpdateUserConfigurationResponseMessage `xml:"UpdateUserConfigurationResponseMessage"`
	// The UploadItemsResponseMessage element contains the status and results of a request to upload a single mailbox item.
	UploadItemsResponseMessage *UploadItemsResponseMessage `xml:"UploadItemsResponseMessage"`
}

func (R *ResponseMessages) SetForMarshal() {
	R.XMLName.Local = "m:ResponseMessages"
}

func (R *ResponseMessages) GetSchema() *Schema {
	return &SchemaMessages
}
