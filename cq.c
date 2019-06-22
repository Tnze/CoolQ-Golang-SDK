#include <windows.h>
#include "app.h"
#include "cq.h"

extern char *__stdcall AppInfo()
{
    return _appinfo();
}

CQEVENT Initialize(int32_t access_code)
{
    ac = access_code;
    HMODULE hmod = LoadLibrary("CQP.dll");
    CQ_addLog_Ptr = GetProcAddress(hmod, "CQ_addLog");
    CQ_sendPrivateMsg_Ptr = GetProcAddress(hmod, "CQ_sendPrivateMsg");
    return 0;
}

CQEVENT EVENT_ON_ENABLE() { return _on_enable(); }

//apis

int32_t CQ_addLog(int32_t priority, char *type, char *reason)
{
    int32_t ret = CQ_addLog_Ptr(ac, priority, type, reason);
    free(type); //释放字符串内存
    free(reason);
    return ret;
}

int32_t CQ_sendPrivateMsg(int64_t qq, char *msg)
{
    int32_t ret = CQ_sendPrivateMsg_Ptr(ac, qq, msg);
    free(msg);
    return ret;
}

int32_t CQ_addLog(int32_t priority, char *type, char *content)
{
    int32_t ret = CQ_addLog_Ptr(priority, type, content);
    free(type);
    free(content);
    return ret;
}
int32_t CQ_sendPrivateMsg(int64_t QQ, char *msg)
{
    int32_t ret = CQ_sendPrivateMsg_Ptr(QQ, msg);
    free(msg);
    return ret;
}
int32_t CQ_sendGroupMsg(int64_t GroupNum, char *msg);
int32_t CQ_sendDiscussMsg(int64_t DiscussNum, char *msg);
int32_t CQ_sendLike(int64_t QQ);
int32_t CQ_sendLikeV2(int64_t QQ, int32_t times);
char *CQ_getCookies();
char *CQ_getRecord(char *file, char *outformat);
int32_t CQ_getCsrfToken();
char *CQ_getAppDirectory();
int64_t CQ_getLoginQQ();
char *CQ_getLoginNick();
int32_t CQ_setGroupKick(int64_t GroupNum, int64_t QQID, int32_t RejectNextTime);
int32_t CQ_setGroupBan(int64_t GroupNum, int64_t QQ, int64_t BanTime);
int32_t CQ_setGroupAdmin(int64_t GroupNum, int64_t QQID, int32_t SetAdmin);
int32_t CQ_setGroupSpecialTitle(int64_t GroupNum, int64_t QQID, char *Title, int64_t TimeOut);
int32_t CQ_setGroupWholeBan(int64_t GroupNum, int32_t SetBan);
int32_t CQ_setGroupAnonymousBan(int64_t GroupNum, char *匿名, int64_t BanTime);
int32_t CQ_setGroupAnonymous(int64_t GroupNum, int32_t 开启匿名);
int32_t CQ_setGroupCard(int64_t GroupNum, int64_t QQID, char *NewCard);
int32_t CQ_setGroupLeave(int64_t GroupNum, int32_t 是否解散);
int32_t CQ_setDiscussLeave(int64_t DiscussNum);
int32_t CQ_setFriendAddRequest(char *请求反馈标识, int32_t FbType, char *remark);
int32_t CQ_setGroupAddRequest(char *请求反馈标识, int32_t ReqType, int32_t FbType);
int32_t CQ_setGroupAddRequestV2(char *请求反馈标识, int32_t ReqType, int32_t FbType, char *reason);
int32_t CQ_setFatal(char *errmsg);
char *CQ_getGroupMemberInfo(int64_t GroupNum, int64_t QQID);
char *CQ_getGroupMemberInfoV2(int64_t GroupNum, int64_t QQID, int32_t NoCatch);
char *CQ_getStrangerInfo(int64_t QQID, int32_t NoCatch);
char *CQ_getGroupMemberList(int64_t GroupNum);
char *CQ_getGroupList(int32_t AuthCode);
int32_t CQ_deleteMsg(int64_t MsgId);
