package elements

// The ResponseMessages element contains the response messages for an Exchange Web Services request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/responsemessages
type ResponseMessages struct {
	// The ApplyConversationActionResponseMessage element contains the status and results of an ApplyConversationAction operation request.
	ApplyConversationActionResponseMessage *ApplyConversationActionResponseMessage `xml:"m:ApplyConversationActionResponseMessage"`
	// The ConvertIdResponseMessage element contains the status and result of a ConvertId operation request.
	ConvertIdResponseMessage *ConvertIdResponseMessage `xml:"m:ConvertIdResponseMessage"`
	// The CopyFolderResponseMessage element contains the status and result of a single CopyFolder operation request.
	CopyFolderResponseMessage *CopyFolderResponseMessage `xml:"m:CopyFolderResponseMessage"`
	// The CopyItemResponseMessage element contains the status and result of a single CopyItem operation request.
	CopyItemResponseMessage *CopyItemResponseMessage `xml:"m:CopyItemResponseMessage"`
	// The CreateAttachmentResponseMessage element contains the status and result of a single CreateAttachment operation request.
	CreateAttachmentResponseMessage *CreateAttachmentResponseMessage `xml:"m:CreateAttachmentResponseMessage"`
	// The CreateFolderResponseMessage element contains the status and result of a single CreateFolder operation request.
	CreateFolderResponseMessage *CreateFolderResponseMessage `xml:"m:CreateFolderResponseMessage"`
	// The CreateItemResponseMessage element contains the status and result of a single CreateItem operation request.
	CreateItemResponseMessage *CreateItemResponseMessage `xml:"m:CreateItemResponseMessage"`
	// The CreateManagedFolderResponseMessage element contains the status and result of a single CreateManagedFolder operation request.
	CreateManagedFolderResponseMessage *CreateManagedFolderResponseMessage `xml:"m:CreateManagedFolderResponseMessage"`
	// The CreateUserConfigurationResponseMessage element contains the status and result of a single CreateUserConfiguration request.
	CreateUserConfigurationResponseMessage *CreateUserConfigurationResponseMessage `xml:"m:CreateUserConfigurationResponseMessage"`
	// The DeleteAttachmentResponseMessage element contains the status and result of a single DeleteAttachment operation request.
	DeleteAttachmentResponseMessage *DeleteAttachmentResponseMessage `xml:"m:DeleteAttachmentResponseMessage"`
	// The DeleteFolderResponseMessage element contains the status and result of a single DeleteFolder operation request.
	DeleteFolderResponseMessage *DeleteFolderResponseMessage `xml:"m:DeleteFolderResponseMessage"`
	// The DeleteItemResponseMessage element contains the status and result of a single DeleteItem operation request.
	DeleteItemResponseMessage *DeleteItemResponseMessage `xml:"m:DeleteItemResponseMessage"`
	// The DeleteUserConfigurationResponseMessage element contains the status and result of a single DeleteUserConfiguration request.
	DeleteUserConfigurationResponseMessage *DeleteUserConfigurationResponseMessage `xml:"m:DeleteUserConfigurationResponseMessage"`
	// The EmptyFolder element defines a request to empty a folder in a mailbox in the Exchange store. Optionally, subfolders can also be deleted when the folder is emptied.
	EmptyFolder *EmptyFolder `xml:"m:EmptyFolder"`
	// The EmptyFolderResponseMessage element contains the status and result of a single EmptyFolder request.
	EmptyFolderResponseMessage *EmptyFolderResponseMessage `xml:"m:EmptyFolderResponseMessage"`
	// The ExpandDLResponseMessage element contains the status and result of a single ExpandDL operation request.
	ExpandDLResponseMessage *ExpandDLResponseMessage `xml:"t:ExpandDLResponseMessage"`
	// The ExportItemsResponseMessage element contains the status and results of a request to export a single mailbox item.
	ExportItemsResponseMessage *ExportItemsResponseMessage `xml:"m:ExportItemsResponseMessage"`
	// The FindFolderResponseMessage element contains the status and result of a single FindFolder operation request.
	FindFolderResponseMessage *FindFolderResponseMessage `xml:"m:FindFolderResponseMessage"`
	// The FindItemResponseMessage element contains the status and result of a single FindItem operation request.
	FindItemResponseMessage *FindItemResponseMessage `xml:"m:FindItemResponseMessage"`
	// The GetAttachmentResponseMessage element contains the status and result of a single GetAttachment operation request.
	GetAttachmentResponseMessage *GetAttachmentResponseMessage `xml:"m:GetAttachmentResponseMessage"`
	// The GetEventsResponseMessage element contains the status and result of a single GetEvents operation request.
	GetEventsResponseMessage *GetEventsResponseMessage `xml:"m:GetEventsResponseMessage"`
	// The GetFolderResponseMessage element contains the status and result of a single GetFolder operation request.
	GetFolderResponseMessage *GetFolderResponseMessage `xml:"m:GetFolderResponseMessage"`
	// The GetItemResponseMessage element contains the status and result of a single GetItem operation request.
	GetItemResponseMessage *GetItemResponseMessage `xml:"m:GetItemResponseMessage"`
	// The GetRemindersResponse element specifies the response to a GetReminders request.
	GetRemindersResponse *GetRemindersResponse `xml:"m:GetRemindersResponse"`
	// The GetRoomListsResponse element defines the response from a GetRoomLists operation request.
	GetRoomListsResponse *GetRoomListsResponse `xml:"m:GetRoomListsResponse"`
	// The GetRoomsResponse element defines a response to a GetRooms operation request.
	GetRoomsResponse *GetRoomsResponse `xml:"m:GetRoomsResponse"`
	// The GetServerTimeZonesResponseMessage element contains the status and result of a single GetServerTimeZones operation request.
	GetServerTimeZonesResponseMessage *GetServerTimeZonesResponseMessage `xml:"m:GetServerTimeZonesResponseMessage"`
	// The GetSharingFolderResponseMessage element contains the status and result of a single GetSharingFolder operation request.
	GetSharingFolderResponseMessage *GetSharingFolderResponseMessage `xml:"m:GetSharingFolderResponseMessage"`
	// The GetSharingMetadataResponseMessage element contains the status and result of a single GetSharingMetadata operation request.
	GetSharingMetadataResponseMessage *GetSharingMetadataResponseMessage `xml:"m:GetSharingMetadataResponseMessage"`
	// The GetStreamingEventsResponseMessage element contains the status and result of a single GetStreamingEvents operation request.
	GetStreamingEventsResponseMessage *GetStreamingEventsResponseMessage `xml:"m:GetStreamingEventsResponseMessage"`
	// The GetUserConfigurationResponseMessage element represents a response that returns a user configuration object.
	GetUserConfigurationResponseMessage *GetUserConfigurationResponseMessage `xml:"m:GetUserConfigurationResponseMessage"`
	// The MoveFolderResponseMessage element contains the status and result of a single MoveFolder operation request.
	MoveFolderResponseMessage *MoveFolderResponseMessage `xml:"m:MoveFolderResponseMessage"`
	// The MoveItemResponseMessage element contains the status and result of a single MoveItem operation request.
	MoveItemResponseMessage *MoveItemResponseMessage `xml:"m:MoveItemResponseMessage"`
	// The PerformReminderActionResponse element specifies the response to a PerformReminderAction request.
	PerformReminderActionResponse *PerformReminderActionResponse `xml:"m:PerformReminderActionResponse"`
	// The RefreshSharingFolderResponseMessage element contains the status and result of a single RefreshSharingFolder operation request.
	RefreshSharingFolderResponseMessage *RefreshSharingFolderResponseMessage `xml:"m:RefreshSharingFolderResponseMessage"`
	// The ResolveNamesResponseMessage element contains the status and result of a ResolveNames operation request.
	ResolveNamesResponseMessage *ResolveNamesResponseMessage `xml:"m:ResolveNamesResponseMessage"`
	// The SendItemResponseMessage element contains the status and result of a single SendItem operation request.
	SendItemResponseMessage *SendItemResponseMessage `xml:"m:SendItemResponseMessage"`
	// The SendNotificationResponseMessage element contains the status and result of a single SendNotification operation request.
	SendNotificationResponseMessage *SendNotificationResponseMessage `xml:"m:SendNotificationResponseMessage"`
	// The SubscribeResponseMessage element contains the status and result of a single Subscribe operation request.
	SubscribeResponseMessage *SubscribeResponseMessage `xml:"m:SubscribeResponseMessage"`
	// The SyncFolderHierarchyResponseMessage element contains the status and result of a single SyncFolderHierarchy operation request.
	SyncFolderHierarchyResponseMessage *SyncFolderHierarchyResponseMessage `xml:"m:SyncFolderHierarchyResponseMessage"`
	// The SyncFolderItemsResponseMessage element contains the status and result of a single SyncFolderItems operation request.
	SyncFolderItemsResponseMessage *SyncFolderItemsResponseMessage `xml:"m:SyncFolderItemsResponseMessage"`
	// The UnsubscribeResponseMessage element contains the status and result of a single Unsubscribe operation request.
	UnsubscribeResponseMessage *UnsubscribeResponseMessage `xml:"m:UnsubscribeResponseMessage"`
	// The UpdateFolderResponseMessage element contains the status and result of updates defined by the FolderChange element of an UpdateFolder operation request.
	UpdateFolderResponseMessage *UpdateFolderResponseMessage `xml:"m:UpdateFolderResponseMessage"`
	// The UpdateItemResponseMessage element contains the status and result of a single UpdateItem operation request.
	UpdateItemResponseMessage *UpdateItemResponseMessage `xml:"m:UpdateItemResponseMessage"`
	// The UpdateUserConfigurationResponseMessage element contains the status and result of a single UpdateUserConfiguration operation request.
	UpdateUserConfigurationResponseMessage *UpdateUserConfigurationResponseMessage `xml:"m:UpdateUserConfigurationResponseMessage"`
	// The UploadItemsResponseMessage element contains the status and results of a request to upload a single mailbox item.
	UploadItemsResponseMessage *UploadItemsResponseMessage `xml:"m:UploadItemsResponseMessage"`
}
