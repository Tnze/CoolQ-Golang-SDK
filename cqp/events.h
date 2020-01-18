// +build !websocket

#include <stdint.h>

#define CQEVENT extern int32_t __stdcall

//events
CQEVENT Initialize(int32_t p0);

extern char *_appinfo();
extern int32_t _on_enable();
extern int32_t _on_disable();
extern int32_t _on_start();
extern int32_t _on_exit();
extern int32_t _on_private_msg(int32_t subType, int32_t msgId, int64_t fromQQ,  char *msg, int32_t font);
extern int32_t _on_group_msg(int32_t subType, int32_t msgId, int64_t fromGroup, int64_t fromQQ, char *fromAnonymous, char *msg, int32_t font);
extern int32_t _on_discuss_msg(int32_t sub_type, int32_t msg_id, int64_t from_discuss, int64_t from_qq, char *msg, int32_t font);
extern int32_t _on_group_upload(int32_t sub_type, int32_t send_time, int64_t from_group, int64_t from_qq, char*file);
extern int32_t _on_group_admin(int32_t sub_type, int32_t send_time, int64_t from_group, int64_t being_operate_qq);
extern int32_t _on_group_member_decrease(int32_t sub_type, int32_t send_time, int64_t from_group, int64_t from_qq, int64_t being_operate_qq);
extern int32_t _on_group_member_increase(int32_t sub_type, int32_t send_time, int64_t from_group, int64_t from_qq, int64_t being_operate_qq);
extern int32_t _on_friend_add(int32_t sub_type, int32_t send_time, int64_t from_qq);
extern int32_t _on_friend_request(int32_t sub_type, int32_t send_time, int64_t from_qq, char* msg, char * response_flag);
extern int32_t _on_group_request(int32_t sub_type, int32_t send_time, int64_t from_group, int64_t from_qq, char *msg, char* response_flag);