#include "events.h"

CQEVENT EVENT_ON_ENABLE() { return _on_enable(); }
CQEVENT EVENT_ON_DISABLE() { return _on_disable(); }
CQEVENT EVENT_ON_PRIVATE_MSG(int32_t subType, int32_t msgId, int64_t fromQQ, char *msg, int32_t font) { return _on_private_msg(subType, msgId, fromQQ, msg, font); }
CQEVENT EVENT_ON_GROUP_MSG(int32_t subType, int32_t msgId, int64_t fromGroup, int64_t fromQQ, char *fromAnonymous, char *msg, int32_t font) { return _on_group_msg(subType, msgId, fromGroup, fromQQ, fromAnonymous, msg, font); }