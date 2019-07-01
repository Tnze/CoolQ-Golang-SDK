// +build windows,386,cgo

#include <stdint.h>

#define CQEVENT extern int32_t __stdcall

//events
CQEVENT Initialize(int32_t p0);
CQEVENT EVENT_ON_ENABLE();
CQEVENT EVENT_ON_DISABLE();
CQEVENT EVENT_ON_PRIVATE_MSG(int32_t subType, int32_t msgId, int64_t fromQQ, char *msg, int32_t font);

extern char *_appinfo();
extern int32_t _on_enable();
extern int32_t _on_disable();
extern int32_t _on_private_msg(int32_t subType, int32_t msgId, int64_t fromQQ, char *msg, int32_t font);
extern int32_t _on_group_msg(int32_t subType, int32_t msgId, int64_t fromGroup, int64_t fromQQ, char *fromAnonymous, char *msg, int32_t font);

// extern int32_t PrivateMessageEvent &)> on_private_msg;
// extern int32_t GroupMessageEvent &)> on_group_msg;
// extern int32_t DiscussMessageEvent &)> on_discuss_msg;
// extern int32_t GroupUploadEvent &)> on_group_upload;
// extern int32_t GroupAdminEvent &)> on_group_admin;
// extern int32_t GroupMemberDecreaseEvent &)> on_group_member_decrease;
// extern int32_t GroupMemberIncreaseEvent &)> on_group_member_increase;
// extern int32_t FriendAddEvent &)> on_friend_add;
// extern int32_t FriendRequestEvent &)> on_friend_request;
// extern int32_t GroupRequestEvent &)> on_group_request;