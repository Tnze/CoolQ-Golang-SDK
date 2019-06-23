// +build windows,386,cgo

#include <windows.h>
#include "cq.h"
#include "events.h"

#define LoadAPI(Name) Name##_Ptr = (Name##_Type)GetProcAddress(hmod, #Name)

extern char *__stdcall AppInfo()
{
    return _appinfo();
}

CQEVENT Initialize(int32_t access_code)
{
    ac = access_code;
    HMODULE hmod = LoadLibrary("CQP.dll");

    LoadAPI(CQ_addLog);
    LoadAPI(CQ_sendPrivateMsg);
    LoadAPI(CQ_sendGroupMsg);
    LoadAPI(CQ_sendDiscussMsg);
    LoadAPI(CQ_sendLike);
    LoadAPI(CQ_sendLikeV2);
    LoadAPI(CQ_getCookies);
    LoadAPI(CQ_getRecord);
    LoadAPI(CQ_getCsrfToken);
    LoadAPI(CQ_getAppDirectory);
    LoadAPI(CQ_getLoginQQ);
    LoadAPI(CQ_getLoginNick);
    LoadAPI(CQ_setGroupKick);
    LoadAPI(CQ_setGroupBan);
    LoadAPI(CQ_setGroupAdmin);
    LoadAPI(CQ_setGroupSpecialTitle);
    LoadAPI(CQ_setGroupWholeBan);
    LoadAPI(CQ_setGroupAnonymousBan);
    LoadAPI(CQ_setGroupAnonymous);
    LoadAPI(CQ_setGroupCard);
    LoadAPI(CQ_setGroupLeave);
    LoadAPI(CQ_setDiscussLeave);
    LoadAPI(CQ_setFriendAddRequest);
    LoadAPI(CQ_setGroupAddRequest);
    LoadAPI(CQ_setGroupAddRequestV2);
    LoadAPI(CQ_setFatal);
    LoadAPI(CQ_getGroupMemberInfo);
    LoadAPI(CQ_getGroupMemberInfoV2);
    LoadAPI(CQ_getStrangerInfo);
    LoadAPI(CQ_getGroupMemberList);
    LoadAPI(CQ_getGroupList);
    LoadAPI(CQ_deleteMsg);

    return 0;
}

CQEVENT EVENT_ON_ENABLE() { return _on_enable(); }
CQEVENT EVENT_ON_DISABLE() { return _on_disable(); }
CQEVENT EVENT_ON_PRIVATE_MSG(int32_t subType, int32_t msgId, int64_t fromQQ, char *msg, int32_t font) { return _on_private_msg(subType, msgId, fromQQ, msg, font); }

//apis，由于Go语言CGO不支持直接调用C函数指针，我们只能用C函数再包一层。顺便做一些内存的释放工作
int32_t CQ_addLog(int32_t priority, char *type, char *content)
{
    int32_t ret = CQ_addLog_Ptr(ac, priority, type, content);
    free(type);
    free(content);
    return ret;
}
int32_t CQ_sendPrivateMsg(int64_t QQ, char *msg)
{
    int32_t ret = CQ_sendPrivateMsg_Ptr(ac, QQ, msg);
    free(msg);
    return ret;
}
int32_t CQ_sendGroupMsg(int64_t GroupNum, char *msg)
{
    int32_t ret = CQ_sendGroupMsg_Ptr(ac, GroupNum, msg);
    free(msg);
    return ret;
}
int32_t CQ_sendDiscussMsg(int64_t DiscussNum, char *msg)
{
    int32_t ret = CQ_sendDiscussMsg_Ptr(ac, DiscussNum, msg);
    free(msg);
    return ret;
}
int32_t CQ_sendLike(int64_t QQ)
{
    return CQ_sendLike_Ptr(ac, QQ);
}
int32_t CQ_sendLikeV2(int64_t QQ, int32_t times)
{
    return CQ_sendLikeV2_Ptr(ac, QQ, times);
}
char *CQ_getCookies()
{
    return CQ_getCookies_Ptr(ac);
}
char *CQ_getRecord(char *file, char *outformat)
{
    char *ret = CQ_getRecord_Ptr(ac, file, outformat);
    free(file);
    free(outformat);
    return ret;
}
int32_t CQ_getCsrfToken()
{
    return CQ_getCsrfToken_Ptr(ac);
}
char *CQ_getAppDirectory()
{
    return CQ_getAppDirectory_Ptr(ac);
}
int64_t CQ_getLoginQQ()
{
    return CQ_getLoginQQ_Ptr(ac);
}
char *CQ_getLoginNick()
{
    return CQ_getLoginNick_Ptr(ac);
}
int32_t CQ_setGroupKick(int64_t GroupNum, int64_t QQ, int32_t RejectNextTime)
{
    return CQ_setGroupKick_Ptr(ac, GroupNum, QQ, RejectNextTime);
}
int32_t CQ_setGroupBan(int64_t GroupNum, int64_t QQ, int64_t BanTime)
{
    return CQ_setGroupBan_Ptr(ac, GroupNum, QQ, BanTime);
}
int32_t CQ_setGroupAdmin(int64_t GroupNum, int64_t QQ, int32_t SetAdmin)
{
    return CQ_setGroupAdmin_Ptr(ac, GroupNum, QQ, SetAdmin);
}
int32_t CQ_setGroupSpecialTitle(int64_t GroupNum, int64_t QQ, char *Title, int64_t TimeOut)
{
    int32_t ret = CQ_setGroupSpecialTitle_Ptr(ac, GroupNum, QQ, Title, TimeOut);
    free(Title);
    return ret;
}
int32_t CQ_setGroupWholeBan(int64_t GroupNum, int32_t SetBan)
{
    return CQ_setGroupWholeBan_Ptr(ac, GroupNum, SetBan);
}
int32_t CQ_setGroupAnonymousBan(int64_t GroupNum, char *anonymous, int64_t BanTime)
{
    int32_t ret = CQ_setGroupAnonymousBan_Ptr(ac, GroupNum, anonymous, BanTime);
    free(anonymous);
    return ret;
}
int32_t CQ_setGroupAnonymous(int64_t GroupNum, int32_t anonymous)
{
    return CQ_setGroupAnonymous_Ptr(ac, GroupNum, anonymous);
}
int32_t CQ_setGroupCard(int64_t GroupNum, int64_t QQ, char *NewCard)
{
    int32_t ret = CQ_setGroupCard_Ptr(ac, GroupNum, QQ, NewCard);
    free(NewCard);
    return ret;
}
int32_t CQ_setGroupLeave(int64_t GroupNum, int32_t dissolve)
{
    return CQ_setGroupLeave_Ptr(ac, GroupNum, dissolve);
}
int32_t CQ_setDiscussLeave(int64_t DiscussNum)
{
    return CQ_setDiscussLeave_Ptr(ac, DiscussNum);
}
int32_t CQ_setFriendAddRequest(char *ReqFeedback, int32_t FbType, char *remark)
{
    int32_t ret = CQ_setFriendAddRequest_Ptr(ac, ReqFeedback, FbType, remark);
    free(ReqFeedback);
    free(remark);
    return ret;
}
int32_t CQ_setGroupAddRequest(char *ReqFeedback, int32_t ReqType, int32_t FbType)
{
    int32_t ret = CQ_setGroupAddRequest_Ptr(ac, ReqFeedback, ReqType, FbType);
    free(ReqFeedback);
    return ret;
}
int32_t CQ_setGroupAddRequestV2(char *ReqFeedback, int32_t ReqType, int32_t FbType, char *reason)
{
    int32_t ret = CQ_setGroupAddRequestV2_Ptr(ac, ReqFeedback, ReqType, FbType, reason);
    free(ReqFeedback);
    free(reason);
    return ret;
}
int32_t CQ_setFatal(char *errmsg)
{
    int32_t ret = CQ_setFatal_Ptr(ac, errmsg);
    free(errmsg);
    return ret;
}
char *CQ_getGroupMemberInfo(int64_t GroupNum, int64_t QQ)
{
    return CQ_getGroupMemberInfo_Ptr(ac, GroupNum, QQ);
}
char *CQ_getGroupMemberInfoV2(int64_t GroupNum, int64_t QQ, int32_t NoCatch)
{
    return CQ_getGroupMemberInfoV2_Ptr(ac, GroupNum, QQ, NoCatch);
}
char *CQ_getStrangerInfo(int64_t QQ, int32_t NoCatch)
{
    return CQ_getStrangerInfo_Ptr(ac, QQ, NoCatch);
}
char *CQ_getGroupMemberList(int64_t GroupNum)
{
    return CQ_getGroupMemberList_Ptr(ac, GroupNum);
}
char *CQ_getGroupList(int32_t AuthCode)
{
    return CQ_getGroupList_Ptr(ac, AuthCode);
}
int32_t CQ_deleteMsg(int64_t MsgID)
{
    return CQ_deleteMsg_Ptr(ac, MsgID);
}
