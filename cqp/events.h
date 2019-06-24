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