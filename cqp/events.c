#include "events.h"

CQEVENT EVENT_ON_ENABLE() { return _on_enable(); }
CQEVENT EVENT_ON_DISABLE() { return _on_disable(); }
CQEVENT EVENT_ON_START() { return _on_start(); }
CQEVENT EVENT_ON_EXIT() { return _on_exit(); }
CQEVENT EVENT_ON_PRIVATE_MSG(int32_t subType, int32_t msgId, int64_t fromQQ, char *msg, int32_t font) { return _on_private_msg(subType, msgId, fromQQ, msg, font); }
CQEVENT EVENT_ON_GROUP_MSG(int32_t subType, int32_t msgId, int64_t fromGroup, int64_t fromQQ, char *fromAnonymous, char *msg, int32_t font) { return _on_group_msg(subType, msgId, fromGroup, fromQQ, fromAnonymous, msg, font); }
CQEVENT EVENT_ON_DISCUSS_MSG(int32_t sub_type, int32_t msg_id, int64_t from_discuss, int64_t from_qq, char *msg, int32_t font) { return _on_discuss_msg(sub_type, msg_id, from_discuss, from_qq, msg, font); }
CQEVENT EVENT_ON_GROUP_UPLOAD(int32_t sub_type, int32_t send_time, int64_t from_group, int64_t from_qq, char *file) { return _on_group_upload(sub_type, send_time, from_group, from_qq, file); }
CQEVENT EVENT_ON_GROUP_ADMIN(int32_t sub_type, int32_t send_time, int64_t from_group, int64_t being_operate_qq) { return _on_group_admin(sub_type, send_time, from_group, being_operate_qq); }
CQEVENT EVENT_ON_GROUP_MEMBER_DECREASE(int32_t sub_type, int32_t send_time, int64_t from_group, int64_t from_qq, int64_t being_operate_qq) { return _on_group_member_decrease(sub_type, send_time, from_group, from_qq, being_operate_qq); }
CQEVENT EVENT_ON_GROUP_MEMBER_INCREASE(int32_t sub_type, int32_t send_time, int64_t from_group, int64_t from_qq, int64_t being_operate_qq) { return _on_group_member_increase(sub_type, send_time, from_group, from_qq, being_operate_qq); }
CQEVENT EVENT_ON_FRIEND_ADD(int32_t sub_type, int32_t send_time, int64_t from_qq) { return _on_friend_add(sub_type, send_time, from_qq); }
CQEVENT EVENT_ON_FRIEND_REQUEST(int32_t sub_type, int32_t send_time, int64_t from_qq, char *msg, char *response_flag) { return _on_friend_request(sub_type, send_time, from_qq, msg, response_flag); }
CQEVENT EVENT_ON_GROUP_REQUEST(int32_t sub_type, int32_t send_time, int64_t from_group, int64_t from_qq, char *msg, char *response_flag) { return _on_group_request(sub_type, send_time, from_group, from_qq, msg, response_flag); }