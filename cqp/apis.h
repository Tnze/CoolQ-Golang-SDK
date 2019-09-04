// +build windows,386,cgo

#include <windows.h>
#include <stdint.h>
#include <stdlib.h>

// #define __stdcall __attribute__((__stdcall__))
#define cq_bool_t int32_t

#define CQAPI(RetType, Name, ...)                                         \
      typedef RetType(__stdcall *Name##_Type)(int32_t ac, ##__VA_ARGS__); \
      Name##_Type Name##_Ptr;                                             \
      RetType Name(__VA_ARGS__);

extern char *__stdcall AppInfo();

int32_t ac; //AccessCode

// Message
CQAPI(int32_t, CQ_sendPrivateMsg, int64_t qq, char *msg)
CQAPI(int32_t, CQ_sendGroupMsg, int64_t group_id, char *msg)
CQAPI(int32_t, CQ_sendDiscussMsg, int64_t discuss_id, char *msg)
CQAPI(int32_t, CQ_deleteMsg, int64_t msg_id)

// Send Like
CQAPI(int32_t, CQ_sendLike, int64_t qq)
CQAPI(int32_t, CQ_sendLikeV2, int64_t qq, int32_t times)

// Group & Discuss Operation
CQAPI(int32_t, CQ_setGroupKick, int64_t group_id, int64_t qq, cq_bool_t reject_add_request)
CQAPI(int32_t, CQ_setGroupBan, int64_t group_id, int64_t qq, int64_t duration)
CQAPI(int32_t, CQ_setGroupAnonymousBan, int64_t group_id, char *anonymous, int64_t duration)
CQAPI(int32_t, CQ_setGroupWholeBan, int64_t group_id, cq_bool_t enable)
CQAPI(int32_t, CQ_setGroupAdmin, int64_t group_id, int64_t qq, cq_bool_t set)
CQAPI(int32_t, CQ_setGroupAnonymous, int64_t group_id, cq_bool_t enable)
CQAPI(int32_t, CQ_setGroupCard, int64_t group_id, int64_t qq, char *new_card)
CQAPI(int32_t, CQ_setGroupLeave, int64_t group_id, cq_bool_t is_dismiss)
CQAPI(int32_t, CQ_setGroupSpecialTitle, int64_t group_id, int64_t qq, char *new_special_title,
      int64_t duration)
CQAPI(int32_t, CQ_setDiscussLeave, int64_t discuss_id)

// Request Operation
CQAPI(int32_t, CQ_setFriendAddRequest, char *response_flag, int32_t response_operation,
      char *remark)
CQAPI(int32_t, CQ_setGroupAddRequest, char *response_flag, int32_t request_type,
      int32_t response_operation)
CQAPI(int32_t, CQ_setGroupAddRequestV2, char *response_flag, int32_t request_type,
      int32_t response_operation, char *reason)

// Get QQ Information
CQAPI(int64_t, CQ_getLoginQQ)
CQAPI(char *, CQ_getLoginNick)
CQAPI(char *, CQ_getStrangerInfo, int64_t qq, cq_bool_t no_cache)
CQAPI(char *, CQ_getGroupList)
CQAPI(char *, CQ_getGroupMemberList, int64_t group_id)
CQAPI(char *, CQ_getGroupMemberInfoV2, int64_t group_id, int64_t qq, cq_bool_t no_cache)

// Get CoolQ Information
CQAPI(char *, CQ_getCookies)
CQAPI(char *, CQ_getCookiesV2, char *domain)
CQAPI(int32_t, CQ_getCsrfToken)
CQAPI(char *, CQ_getAppDirectory)
CQAPI(char *, CQ_getRecord, char *file, char *out_format)
CQAPI(char *, CQ_getRecordV2, char *file, char *out_format)
CQAPI(char *, CQ_getImage, char *file)
CQAPI(int32_t, CQ_canSendImage)
CQAPI(int32_t, CQ_canSendRecord)

CQAPI(int32_t, CQ_addLog, int32_t log_level, char *category, char *log_msg)
CQAPI(int32_t, CQ_setFatal, char *error_info)
CQAPI(int32_t, CQ_setRestart) // currently ineffective