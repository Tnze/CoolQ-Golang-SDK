package cqp

//go:generate go run github.com/Tnze/CoolQ-Golang-SDK/v3/tools/cqgen $GOFILE
//cqapi CQ_sendPrivateMsg(int64,string)int32
//cqapi CQ_sendGroupMsg(int64,string)int32
//cqapi CQ_sendDiscussMsg(int64,string)int32
//cqapi CQ_deleteMsg(int64)int32
//cqapi CQ_sendLikeV2(int64,int32)int32
//cqapi CQ_setGroupKick(int64,int64,bool)int32
//cqapi CQ_setGroupBan(int64,int64,int64)int32
//cqapi CQ_setGroupAnonymousBan(int64,string,int64)int32
//cqapi CQ_setGroupWholeBan(int64,bool)int32
//cqapi CQ_setGroupAdmin(int64,int64,bool)int32
//cqapi CQ_setGroupAnonymous(int64,bool)int32
//cqapi CQ_setGroupCard(int64,int64,string)int32
//cqapi CQ_setGroupLeave(int64,bool)int32
//cqapi CQ_setGroupSpecialTitle(int64,int64,string,int64)int32
//cqapi CQ_setDiscussLeave(int64)int32
//cqapi CQ_setFriendAddRequest(string,int32,string)int32
//cqapi CQ_setGroupAddRequestV2(string,int32,int32,string)int32
//cqapi CQ_getLoginQQ()int64
//cqapi CQ_getLoginNick()string
//cqapi CQ_getStrangerInfo(int64,bool)string			getRawStrangerInfo
//cqapi CQ_getFriendList(bool)string					getRawFriendList
//cqapi CQ_getGroupList()string							getRawGroupList
//cqapi CQ_getGroupInfo(int64,bool)string				getRawGroupInfo
//cqapi CQ_getGroupMemberList(int64)string				getRawGroupMemberList
//cqapi CQ_getGroupMemberInfoV2(int64,int64,bool)string	getRawGroupMemberInfoV2
//cqapi CQ_getCookiesV2(string)string
//cqapi CQ_getCsrfToken()int32
//cqapi CQ_getAppDirectory()string
//cqapi CQ_getRecordV2(string,string)string
//cqapi CQ_getImage(string)string
//cqapi CQ_canSendImage()bool
//cqapi CQ_canSendRecord()bool
//cqapi CQ_addLog(int32,string,string)int32
