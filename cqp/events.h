// +build windows,386,cgo

#include <stdint.h>

extern char *_appinfo();
extern int32_t _on_enable();
extern int32_t _on_disable();
extern int32_t _on_private_msg(int32_t subType, int32_t msgId, int64_t fromQQ, char *msg, int32_t font);